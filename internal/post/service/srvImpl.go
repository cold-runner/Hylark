package service

import (
	"bytes"
	"context"
	"sync"

	"github.com/cold-runner/Hylark/gorm_gen/post_srv"
	"github.com/cold-runner/Hylark/gorm_gen/post_srv/model"
	"github.com/cold-runner/Hylark/internal/pkg"
	"github.com/cold-runner/Hylark/internal/pkg/response"
	"github.com/cold-runner/Hylark/kitex_gen/post"

	"github.com/bytedance/gopkg/util/gopool"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (s *Srv) CreatePost(ctx context.Context, req *post.CreatePostRequest) (r *post.CreatePostResponse, err error) {
	// 参数校验
	if err = func(context.Context, *post.CreatePostRequest) error {
		// 图片类型校验
		ft, err := pkg.FileTypeFromBinary(req.GetPicture())
		if err != nil {
			return response.BizErr(response.ErrInternal)
		}
		if ft != "jpg" && ft != "png" {
			return response.BizErrWithExtra(response.ErrBadRequest, response.BizErrExtra{"msg": "file type is illegal."})
		}

		// 查询是否存在category_id
		cid, err := uuid.Parse(req.GetCategoryId())
		if err != nil {
			return response.BizErr(response.ErrInternal)
		}
		_, err = s.Repository.Category().Get(ctx, post_srv.Q.Category.ID.Eq(cid))
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return response.BizErr(response.ErrCategoryNotExist)
		case err != nil:
			return response.BizErr(response.ErrInternal)
		}

		// 查询tags是否合法
		// TODO 抽象成工具函数：并发执行若干任务，当有任何一个任务失败时返回，否则直到所有任务完成（成功）后返回
		reqTagId := req.GetTagId()
		if len(reqTagId) > 20 {
			return response.BizErrWithExtra(response.ErrBadRequest, response.BizErrExtra{"msg": "tag too many."})
		}
		start, end := make(chan struct{}), make(chan struct{})
		sign := make(chan error)
		var wg sync.WaitGroup
		for _, tag := range reqTagId {
			t := tag
			wg.Add(1)
			gopool.Go(func() {
				<-start
				tagId, err := uuid.Parse(t)
				if err != nil {
					sign <- response.BizErr(response.ErrInternal)
					return
				}
				_, err = s.Repository.Tag().Get(ctx, post_srv.Q.Tag.ID.Eq(tagId))
				switch {
				case errors.Is(err, gorm.ErrRecordNotFound):
					sign <- response.BizErr(response.ErrTagNotExist)
				case err != nil:
					sign <- response.BizErr(response.ErrInternal)
				}

				wg.Done()
			})
		}
		gopool.Go(func() {
			wg.Wait()
			end <- struct{}{}
		})
		close(start)

		select {
		case <-sign:
			return err
		case <-end:
			return nil
		}
	}(ctx, req); err != nil {
		return nil, err
	}

	uid := uuid.New()

	// 上传图片
	url, err := s.Oss.Upload(ctx, bytes.NewReader(req.GetPicture()), "cover."+pkg.MustFileType(req.GetPicture()), "/post/"+uid.String())
	if err != nil {
		return nil, response.BizErr(response.ErrInternal)
	}

	p := &model.Post{
		ID:           uid,
		Title:        req.GetTitle(),
		CoverImage:   url,
		UserID:       req.GetUserId(),
		Summary:      req.GetSummary(),
		Content:      req.GetContent(),
		CategoryID:   req.GetCategoryId(),
		Temperature:  0,
		LikeCount:    0,
		ViewCount:    0,
		StarCount:    0,
		CommentCount: 0,
		ShareCount:   0,
		State:        0,
		LinkURL:      req.GetLinkUrl(),
	}

	err = s.Repository.Post().Persist(ctx, p)
	if err != nil {
		return nil, response.BizErr(response.ErrInternal)
	}

	return post.NewCreatePostResponse(), nil
}
