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
    2: string image_url (api.body="image_url");
    3: string file_name (api.body="file_name");
    4: i64 file_size (api.body="file_size");
}

// 上传文件请求
struct UploadFileReq {
    1: string file (api.form="file", api.vd="len($)>0");
    2: string upload_type (api.form="upload_type");
}

// 上传文件响应
struct UploadFileResp {
    1: common.BaseResp base (api.body="base");
    2: string file_url (api.body="file_url");
    3: string file_name (api.body="file_name");
    4: i64 file_size (api.body="file_size");
    5: string file_type (api.body="file_type");
}

// 上传认证文件请求
struct UploadCertFileReq {
    1: string cert_file (api.form="cert_file", api.vd="len($)>0");
    2: string cert_type (api.form="cert_type", api.vd="len($)>0");
}

// 上传认证文件响应
struct UploadCertFileResp {
    1: common.BaseResp base (api.body="base");
    2: string file_url (api.body="file_url");
    3: string cert_type (api.body="cert_type");
    4: string file_name (api.body="file_name");
    5: i64 file_size (api.body="file_size");
}

service UploadService {
    UploadImageResp UploadImage(1: UploadImageReq request) (api.post="/api/v1/upload/image");
    UploadFileResp UploadFile(1: UploadFileReq request) (api.post="/api/v1/upload/file");
    UploadCertFileResp UploadCertFile(1: UploadCertFileReq request) (api.post="/api/v1/upload/cert");
}
