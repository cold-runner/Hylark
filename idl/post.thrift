namespace go post

include "common.thrift"

struct CreatePostRequest {
    1: optional string token
    2: optional string user_id
    3: optional string category_id
    4: optional string title
    5: optional string summary
    6: optional string content
    7: optional binary picture
    8: optional string link_url
    9: optional set<string> tag_id
}

struct CreatePostResponse {
}