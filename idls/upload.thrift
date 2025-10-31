namespace go upload

include "common.thrift"

// 上传图片请求
struct UploadImageReq {
    1: string image_file (api.form="image_file", api.vd="len($)>0");
    2: string upload_type (api.form="upload_type");
}

// 上传图片响应
struct UploadImageResp {
    1: common.BaseResp base (api.body="base");
    2: string image_url (api.body="image_url");          // 原始地址（用于存储）
    3: string display_url (api.body="display_url");      // 签名后的展示地址
    4: string file_name (api.body="file_name");
    5: i64 file_size (api.body="file_size" go.tag="json:\"file_size,string\"");
}

// 上传文件请求
struct UploadFileReq {
    1: string file (api.form="file", api.vd="len($)>0");
    2: string upload_type (api.form="upload_type");
}

// 上传文件响应
struct UploadFileResp {
    1: common.BaseResp base (api.body="base");
    2: string file_url (api.body="file_url");            // 原始地址（用于存储）
    3: string display_url (api.body="display_url");      // 签名后的展示地址
    4: string file_name (api.body="file_name");
    5: i64 file_size (api.body="file_size" go.tag="json:\"file_size,string\"");
    6: string file_type (api.body="file_type");
}

// 上传认证文件请求
struct UploadCertFileReq {
    1: string cert_file (api.form="cert_file", api.vd="len($)>0");
    2: string cert_type (api.form="cert_type", api.vd="len($)>0");
}

// 上传认证文件响应
struct UploadCertFileResp {
    1: common.BaseResp base (api.body="base");
    2: string file_url (api.body="file_url");            // 原始地址（用于存储）
    3: string display_url (api.body="display_url");      // 签名后的展示地址
    4: string cert_type (api.body="cert_type");
    5: string file_name (api.body="file_name");
    6: i64 file_size (api.body="file_size" go.tag="json:\"file_size,string\"");
}

// 获取签名URL请求
struct GetSignedURLReq {
    1: string file_url (api.query="file_url", api.vd="len($)>0");
    2: i64 expire_seconds (api.query="expire_seconds" go.tag="json:\"expire_seconds,string\"");
}

// 获取签名URL响应
struct GetSignedURLResp {
    1: common.BaseResp base (api.body="base");
    2: string signed_url (api.body="signed_url");
    3: i64 expire_seconds (api.body="expire_seconds" go.tag="json:\"expire_seconds,string\"");
    4: string expire_time (api.body="expire_time");
}

service UploadService {
    UploadImageResp UploadImage(1: UploadImageReq request) (api.post="/api/v1/upload/image");
    UploadFileResp UploadFile(1: UploadFileReq request) (api.post="/api/v1/upload/file");
    UploadCertFileResp UploadCertFile(1: UploadCertFileReq request) (api.post="/api/v1/upload/cert");
    GetSignedURLResp GetSignedURL(1: GetSignedURLReq request) (api.get="/api/v1/upload/signed-url");
}
