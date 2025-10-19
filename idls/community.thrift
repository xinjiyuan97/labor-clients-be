namespace go community

include "common.thrift"

// 获取帖子列表请求
struct GetPostListReq {
    1: common.PageReq page_req (api.body="page_req");
    2: string post_type (api.query="post_type");
    3: string sort (api.query="sort");
    4: string order (api.query="order");
}

// 获取帖子列表响应
struct GetPostListResp {
    1: common.BaseResp base (api.body="base");
    2: common.PageResp page_resp (api.body="page_resp");
    3: list<common.CommunityPostInfo> posts (api.body="posts");
}

// 获取帖子详情请求
struct GetPostDetailReq {
    1: i64 post_id (api.path="post_id", api.vd="$>0");
}

// 获取帖子详情响应
struct GetPostDetailResp {
    1: common.BaseResp base (api.body="base");
    2: common.CommunityPostInfo post (api.body="post");
}

// 发布帖子请求
struct CreatePostReq {
    1: string title (api.body="title", api.vd="len($)>0");
    2: string content (api.body="content", api.vd="len($)>0");
    3: string post_type (api.body="post_type");
}

// 发布帖子响应
struct CreatePostResp {
    1: common.BaseResp base (api.body="base");
    2: i64 post_id (api.body="post_id");
}

// 更新帖子请求
struct UpdatePostReq {
    1: i64 post_id (api.path="post_id", api.vd="$>0");
    2: string title (api.body="title");
    3: string content (api.body="content");
    4: string post_type (api.body="post_type");
}

// 更新帖子响应
struct UpdatePostResp {
    1: common.BaseResp base (api.body="base");
    2: common.CommunityPostInfo post (api.body="post");
}

// 删除帖子请求
struct DeletePostReq {
    1: i64 post_id (api.path="post_id", api.vd="$>0");
}

// 删除帖子响应
struct DeletePostResp {
    1: common.BaseResp base (api.body="base");
}

// 点赞帖子请求
struct LikePostReq {
    1: i64 post_id (api.path="post_id", api.vd="$>0");
}

// 点赞帖子响应
struct LikePostResp {
    1: common.BaseResp base (api.body="base");
    2: i32 like_count (api.body="like_count");
}

// 取消点赞请求
struct UnlikePostReq {
    1: i64 post_id (api.path="post_id", api.vd="$>0");
}

// 取消点赞响应
struct UnlikePostResp {
    1: common.BaseResp base (api.body="base");
    2: i32 like_count (api.body="like_count");
}

// 评论帖子请求
struct CommentPostReq {
    1: i64 post_id (api.path="post_id", api.vd="$>0");
    2: string content (api.body="content", api.vd="len($)>0");
    3: i64 parent_id (api.body="parent_id");
}

// 评论帖子响应
struct CommentPostResp {
    1: common.BaseResp base (api.body="base");
    2: i64 comment_id (api.body="comment_id");
}

// 获取评论列表请求
struct GetCommentListReq {
    1: i64 post_id (api.path="post_id", api.vd="$>0");
    2: common.PageReq page_req (api.body="page_req");
}

// 获取评论列表响应
struct GetCommentListResp {
    1: common.BaseResp base (api.body="base");
    2: common.PageResp page_resp (api.body="page_resp");
    3: list<common.MessageInfo> comments (api.body="comments");
}

// 获取入门课程请求
struct GetCoursesReq {
    1: common.PageReq page_req (api.body="page_req");
    2: string category (api.query="category");
}

// 获取入门课程响应
struct GetCoursesResp {
    1: common.BaseResp base (api.body="base");
    2: common.PageResp page_resp (api.body="page_resp");
    3: list<common.CommunityPostInfo> courses (api.body="courses");
}

service CommunityService {
    GetPostListResp GetPostList(1: GetPostListReq request) (api.get="/api/v1/community/posts");
    GetPostDetailResp GetPostDetail(1: GetPostDetailReq request) (api.get="/api/v1/community/posts/:post_id");
    CreatePostResp CreatePost(1: CreatePostReq request) (api.post="/api/v1/community/posts");
    UpdatePostResp UpdatePost(1: UpdatePostReq request) (api.put="/api/v1/community/posts/:post_id");
    DeletePostResp DeletePost(1: DeletePostReq request) (api.delete="/api/v1/community/posts/:post_id");
    LikePostResp LikePost(1: LikePostReq request) (api.post="/api/v1/community/posts/:post_id/like");
    UnlikePostResp UnlikePost(1: UnlikePostReq request) (api.delete="/api/v1/community/posts/:post_id/like");
    CommentPostResp CommentPost(1: CommentPostReq request) (api.post="/api/v1/community/posts/:post_id/comment");
    GetCommentListResp GetCommentList(1: GetCommentListReq request) (api.get="/api/v1/community/posts/:post_id/comments");
    GetCoursesResp GetCourses(1: GetCoursesReq request) (api.get="/api/v1/community/courses");
}
