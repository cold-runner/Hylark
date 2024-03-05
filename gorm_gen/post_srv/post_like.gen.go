// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package post_srv

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/cold-runner/Hylark/gorm_gen/post_srv/model"
)

func newPostLike(db *gorm.DB, opts ...gen.DOOption) postLike {
	_postLike := postLike{}

	_postLike.postLikeDo.UseDB(db, opts...)
	_postLike.postLikeDo.UseModel(&model.PostLike{})

	tableName := _postLike.postLikeDo.TableName()
	_postLike.ALL = field.NewAsterisk(tableName)
	_postLike.ID = field.NewField(tableName, "id")
	_postLike.CreatedAt = field.NewTime(tableName, "created_at")
	_postLike.DeletedAt = field.NewField(tableName, "deleted_at")
	_postLike.UpdatedAt = field.NewTime(tableName, "updated_at")
	_postLike.CommentID = field.NewString(tableName, "comment_id")
	_postLike.UserID = field.NewString(tableName, "user_id")

	_postLike.fillFieldMap()

	return _postLike
}

// postLike 评论-点赞表
type postLike struct {
	postLikeDo

	ALL       field.Asterisk
	ID        field.Field  // 自然主键
	CreatedAt field.Time   // 创建时间
	DeletedAt field.Field  // 删除时间（软删除）
	UpdatedAt field.Time   // 更新时间
	CommentID field.String // 评论id（考虑性能，不加外键）
	UserID    field.String // 点赞人（考虑性能，不加外键）

	fieldMap map[string]field.Expr
}

func (p postLike) Table(newTableName string) *postLike {
	p.postLikeDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p postLike) As(alias string) *postLike {
	p.postLikeDo.DO = *(p.postLikeDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *postLike) updateTableName(table string) *postLike {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewField(table, "id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.DeletedAt = field.NewField(table, "deleted_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.CommentID = field.NewString(table, "comment_id")
	p.UserID = field.NewString(table, "user_id")

	p.fillFieldMap()

	return p
}

func (p *postLike) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *postLike) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 6)
	p.fieldMap["id"] = p.ID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["comment_id"] = p.CommentID
	p.fieldMap["user_id"] = p.UserID
}

func (p postLike) clone(db *gorm.DB) postLike {
	p.postLikeDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p postLike) replaceDB(db *gorm.DB) postLike {
	p.postLikeDo.ReplaceDB(db)
	return p
}

type postLikeDo struct{ gen.DO }

type IPostLikeDo interface {
	gen.SubQuery
	Debug() IPostLikeDo
	WithContext(ctx context.Context) IPostLikeDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPostLikeDo
	WriteDB() IPostLikeDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPostLikeDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPostLikeDo
	Not(conds ...gen.Condition) IPostLikeDo
	Or(conds ...gen.Condition) IPostLikeDo
	Select(conds ...field.Expr) IPostLikeDo
	Where(conds ...gen.Condition) IPostLikeDo
	Order(conds ...field.Expr) IPostLikeDo
	Distinct(cols ...field.Expr) IPostLikeDo
	Omit(cols ...field.Expr) IPostLikeDo
	Join(table schema.Tabler, on ...field.Expr) IPostLikeDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPostLikeDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPostLikeDo
	Group(cols ...field.Expr) IPostLikeDo
	Having(conds ...gen.Condition) IPostLikeDo
	Limit(limit int) IPostLikeDo
	Offset(offset int) IPostLikeDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPostLikeDo
	Unscoped() IPostLikeDo
	Create(values ...*model.PostLike) error
	CreateInBatches(values []*model.PostLike, batchSize int) error
	Save(values ...*model.PostLike) error
	First() (*model.PostLike, error)
	Take() (*model.PostLike, error)
	Last() (*model.PostLike, error)
	Find() ([]*model.PostLike, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PostLike, err error)
	FindInBatches(result *[]*model.PostLike, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.PostLike) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPostLikeDo
	Assign(attrs ...field.AssignExpr) IPostLikeDo
	Joins(fields ...field.RelationField) IPostLikeDo
	Preload(fields ...field.RelationField) IPostLikeDo
	FirstOrInit() (*model.PostLike, error)
	FirstOrCreate() (*model.PostLike, error)
	FindByPage(offset int, limit int) (result []*model.PostLike, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPostLikeDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p postLikeDo) Debug() IPostLikeDo {
	return p.withDO(p.DO.Debug())
}

func (p postLikeDo) WithContext(ctx context.Context) IPostLikeDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p postLikeDo) ReadDB() IPostLikeDo {
	return p.Clauses(dbresolver.Read)
}

func (p postLikeDo) WriteDB() IPostLikeDo {
	return p.Clauses(dbresolver.Write)
}

func (p postLikeDo) Session(config *gorm.Session) IPostLikeDo {
	return p.withDO(p.DO.Session(config))
}

func (p postLikeDo) Clauses(conds ...clause.Expression) IPostLikeDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p postLikeDo) Returning(value interface{}, columns ...string) IPostLikeDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p postLikeDo) Not(conds ...gen.Condition) IPostLikeDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p postLikeDo) Or(conds ...gen.Condition) IPostLikeDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p postLikeDo) Select(conds ...field.Expr) IPostLikeDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p postLikeDo) Where(conds ...gen.Condition) IPostLikeDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p postLikeDo) Order(conds ...field.Expr) IPostLikeDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p postLikeDo) Distinct(cols ...field.Expr) IPostLikeDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p postLikeDo) Omit(cols ...field.Expr) IPostLikeDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p postLikeDo) Join(table schema.Tabler, on ...field.Expr) IPostLikeDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p postLikeDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPostLikeDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p postLikeDo) RightJoin(table schema.Tabler, on ...field.Expr) IPostLikeDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p postLikeDo) Group(cols ...field.Expr) IPostLikeDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p postLikeDo) Having(conds ...gen.Condition) IPostLikeDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p postLikeDo) Limit(limit int) IPostLikeDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p postLikeDo) Offset(offset int) IPostLikeDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p postLikeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPostLikeDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p postLikeDo) Unscoped() IPostLikeDo {
	return p.withDO(p.DO.Unscoped())
}

func (p postLikeDo) Create(values ...*model.PostLike) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p postLikeDo) CreateInBatches(values []*model.PostLike, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p postLikeDo) Save(values ...*model.PostLike) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p postLikeDo) First() (*model.PostLike, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PostLike), nil
	}
}

func (p postLikeDo) Take() (*model.PostLike, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PostLike), nil
	}
}

func (p postLikeDo) Last() (*model.PostLike, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PostLike), nil
	}
}

func (p postLikeDo) Find() ([]*model.PostLike, error) {
	result, err := p.DO.Find()
	return result.([]*model.PostLike), err
}

func (p postLikeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PostLike, err error) {
	buf := make([]*model.PostLike, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p postLikeDo) FindInBatches(result *[]*model.PostLike, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p postLikeDo) Attrs(attrs ...field.AssignExpr) IPostLikeDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p postLikeDo) Assign(attrs ...field.AssignExpr) IPostLikeDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p postLikeDo) Joins(fields ...field.RelationField) IPostLikeDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p postLikeDo) Preload(fields ...field.RelationField) IPostLikeDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p postLikeDo) FirstOrInit() (*model.PostLike, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PostLike), nil
	}
}

func (p postLikeDo) FirstOrCreate() (*model.PostLike, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PostLike), nil
	}
}

func (p postLikeDo) FindByPage(offset int, limit int) (result []*model.PostLike, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p postLikeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p postLikeDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p postLikeDo) Delete(models ...*model.PostLike) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *postLikeDo) withDO(do gen.Dao) *postLikeDo {
	p.DO = *do.(*gen.DO)
	return p
}
