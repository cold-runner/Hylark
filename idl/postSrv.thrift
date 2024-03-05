namespace go post

include "post.thrift"

service srv {
   post.CreatePostResponse CreatePost(post.CreatePostRequest req)
}
