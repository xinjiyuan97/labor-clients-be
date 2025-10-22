#!/usr/bin/env python3
"""
根据 Thrift IDL 文件生成 OpenAPI 3.0 规范文件
"""

import os
import re
import json
import argparse
from typing import Dict, List, Any

class ThriftToOpenAPIConverter:
    def __init__(self):
        self.openapi_spec = {
            "openapi": "3.0.0",
            "info": {
                "title": "零工APP API",
                "description": "零工APP后端API接口文档",
                "version": "1.0.0",
                "contact": {
                    "name": "API Support",
                    "email": "support@labor-clients.com"
                }
            },
            "servers": [
                {
                    "url": "https://api.labor-clients.com/v1",
                    "description": "生产环境"
                },
                {
                    "url": "http://localhost:8080/v1",
                    "description": "开发环境"
                }
            ],
            "tags": [],
            "paths": {},
            "components": {
                "schemas": {},
                "securitySchemes": {
                    "BearerAuth": {
                        "type": "http",
                        "scheme": "bearer",
                        "bearerFormat": "JWT"
                    }
                }
            },
            "security": [
                {"BearerAuth": []}
            ]
        }
        
        # Thrift 类型到 OpenAPI 类型的映射
        self.type_mapping = {
            "i32": "integer",
            "i64": "integer",
            "double": "number",
            "string": "string",
            "bool": "boolean",
            "list": "array",
            "map": "object"
        }
        
        # 模块信息
        self.modules = {
            "auth": {"name": "auth", "description": "用户注册、登录、认证相关接口"},
            "schedule": {"name": "schedule", "description": "日程创建、查询、管理相关接口"},
            "job": {"name": "job", "description": "岗位列表、详情、推荐相关接口"},
            "job_application": {"name": "job_application", "description": "岗位申请、状态管理相关接口"},
            "message": {"name": "message", "description": "消息发送、接收、管理相关接口"},
            "user": {"name": "user", "description": "用户信息、收藏、收入相关接口"},
            "community": {"name": "community", "description": "社区帖子、互动相关接口"},
            "attendance": {"name": "attendance", "description": "打卡、考勤记录相关接口"},
            "review": {"name": "review", "description": "双向评价、评分管理相关接口"},
            "payment": {"name": "payment", "description": "收入统计、提现相关接口"},
            "system": {"name": "system", "description": "系统配置、公告相关接口"},
            "upload": {"name": "upload", "description": "文件、图片上传相关接口"},
            "admin": {"name": "admin", "description": "管理员管理品牌方等信息"},
        }

    def parse_thrift_file(self, file_path: str) -> Dict[str, Any]:
        """解析 Thrift 文件"""
        with open(file_path, 'r', encoding='utf-8') as f:
            content = f.read()
        
        # 提取模块名
        namespace_match = re.search(r'namespace go (\w+)', content)
        module_name = namespace_match.group(1) if namespace_match else "unknown"
        
         # 提取 include    
        includes = re.findall(r'include "([^"]+)"', content)

        # 提取结构体定义
        structs = self._extract_structs(content)
        
        # 提取服务定义
        services = self._extract_services(content)
        
        return {
            "module_name": module_name,
            "includes": includes,
            "structs": structs,
            "services": services
        }

    def _extract_structs(self, content: str) -> Dict[str, Any]:
        """提取结构体定义"""
        structs = {}
        
        # 匹配结构体定义
        struct_pattern = r'struct\s+(\w+)\s*\{([^}]+)\}'
        struct_matches = re.findall(struct_pattern, content, re.DOTALL)
        
        for struct_name, struct_body in struct_matches:
            fields = []
            field_lines = struct_body.strip().split('\n')
            
            for line in field_lines:
                line = line.strip()
                if line and not line.startswith('//'):
                    # 解析字段定义
                    field_match = re.match(r'(\d+):\s*([^;]+);', line)
                    if field_match:
                        field_num, field_def = field_match.groups()
                        
                        # 解析字段类型、名称和注解
                        field_parts = field_def.strip().split()
                        if len(field_parts) >= 2:
                            field_type = field_parts[0]
                            field_name = field_parts[1]
                            
                            # 提取注解
                            annotations = {}
                            if '(' in field_def and ')' in field_def:
                                ann_start = field_def.find('(')
                                ann_end = field_def.rfind(')')
                                ann_text = field_def[ann_start+1:ann_end]
                                
                                # 解析注解
                                ann_matches = re.findall(r'api\.(\w+)="([^"]*)"', ann_text)
                                for ann_key, ann_value in ann_matches:
                                    annotations[ann_key] = ann_value
                            
                            fields.append({
                                "number": int(field_num),
                                "type": field_type,
                                "name": field_name,
                                "annotations": annotations
                            })
            
            structs[struct_name] = {
                "fields": fields
            }
        
        return structs

    def _extract_services(self, content: str) -> Dict[str, Any]:
        """提取服务定义"""
        services = {}
        
        # 匹配服务定义
        service_pattern = r'service\s+(\w+)\s*\{([^}]+)\}'
        service_matches = re.findall(service_pattern, content, re.DOTALL)
        
        for service_name, service_body in service_matches:
            methods = []
            method_lines = service_body.strip().split('\n')
            
            for line in method_lines:
                line = line.strip()
                if line and not line.startswith('//'):
                    # 解析方法定义
                    method_match = re.match(r'(\w+Resp)\s+(\w+)\(1:\s*(\w+Req)\s+request\)\s*\(([^)]+)\);', line)
                    if method_match:
                        resp_type, method_name, req_type, annotations = method_match.groups()
                        
                        # 解析注解
                        api_matches = re.findall(r'api\.(\w+)="([^"]*)"', annotations)
                        api_info = dict(api_matches)
                        
                        methods.append({
                            "name": method_name,
                            "request_type": req_type,
                            "response_type": resp_type,
                            "api_info": api_info
                        })
            
            services[service_name] = {
                "methods": methods
            }
        
        return services

    def thrift_type_to_openapi(self, thrift_type: str) -> Dict[str, Any]:
        """将 Thrift 类型转换为 OpenAPI 类型"""
        if thrift_type in self.type_mapping:
            return {"type": self.type_mapping[thrift_type]}
        elif thrift_type.startswith("list<"):
            inner_type = thrift_type[5:-1]
            return {
                "type": "array",
                "items": self.thrift_type_to_openapi(inner_type)
            }
        elif thrift_type.startswith("map<"):
            # 简化处理，map 转换为 object
            return {"type": "object"}
        else:
            # 自定义类型，作为引用处理
            # 去掉命名空间前缀（如 common.BaseResp -> BaseResp）
            schema_name = thrift_type.split('.')[-1]
            return {"$ref": f"#/components/schemas/{schema_name}"}

    def convert_struct_to_schema(self, struct_name: str, struct_info: Dict[str, Any]) -> Dict[str, Any]:
        """将 Thrift 结构体转换为 OpenAPI Schema"""
        properties = {}
        required_fields = []
        
        for field in struct_info["fields"]:
            field_name = field["name"]
            field_type = field["type"]
            
            # 检查是否必需字段
            if "vd" in field["annotations"]:
                required_fields.append(field_name)
            
            properties[field_name] = self.thrift_type_to_openapi(field_type)
            
            # 添加描述
            if field_name in ["code", "message"]:
                properties[field_name]["description"] = "状态码" if field_name == "code" else "消息"
        
        schema = {
            "type": "object",
            "properties": properties
        }
        
        if required_fields:
            schema["required"] = required_fields
        
        return schema

    def convert_method_to_operation(self, method_info: Dict[str, Any], module_name: str) -> Dict[str, Any]:
        """将 Thrift 方法转换为 OpenAPI 操作"""
        method_name = method_info["name"]
        api_info = method_info["api_info"]
        
        if "get" in api_info:
            http_method = "get"
            path = api_info["get"]
        elif "post" in api_info:
            http_method = "post"
            path = api_info["post"]
        elif "put" in api_info:
            http_method = "put"
            path = api_info["put"]
        elif "delete" in api_info:
            http_method = "delete"
            path = api_info["delete"]
        else:
            return None
        
        # 转换路径参数：将 :xxx 格式转换为 {xxx} 格式
        # 例如：/api/v1/admin/admins/:admin_id -> /api/v1/admin/admins/{admin_id}
        path = re.sub(r':([a-zA-Z_][a-zA-Z0-9_]*)', r'{\1}', path)
        
        operation = {
            "tags": [self.modules.get(module_name, {}).get("name", module_name)],
            "summary": method_name,
            "description": f"{method_name} 接口",
            "operationId": f"{module_name}_{method_name}",
        }
        
        # 添加请求体（对于 POST/PUT 请求）
        if http_method in ["post", "put"]:
            operation["requestBody"] = {
                "required": True,
                "content": {
                    "application/json": {
                        "schema": {"$ref": f"#/components/schemas/{method_info['request_type']}"}
                    }
                }
            }
        
        # 添加参数（对于 GET 请求或路径参数）
        parameters = []
        if http_method == "get":
            # 为 GET 请求添加查询参数
            parameters.append({
                "name": "page",
                "in": "query",
                "description": "页码",
                "schema": {"type": "integer", "minimum": 1, "default": 1}
            })
            parameters.append({
                "name": "limit",
                "in": "query",
                "description": "每页数量",
                "schema": {"type": "integer", "minimum": 1, "maximum": 100, "default": 20}
            })
        
        # 检查路径参数
        path_params = re.findall(r'\{([^}]+)\}', path)
        for param in path_params:
            parameters.append({
                "name": param,
                "in": "path",
                "required": True,
                "description": f"{param} 参数",
                "schema": {"type": "integer"}
            })
        
        if parameters:
            operation["parameters"] = parameters
        
        # 添加响应
        operation["responses"] = {
            "200": {
                "description": "成功",
                "content": {
                    "application/json": {
                        "schema": {"$ref": f"#/components/schemas/{method_info['response_type']}"}
                    }
                }
            },
            "400": {
                "description": "请求参数错误",
                "content": {
                    "application/json": {
                        "schema": {"$ref": "#/components/schemas/BaseResp"}
                    }
                }
            },
            "401": {
                "description": "未授权",
                "content": {
                    "application/json": {
                        "schema": {"$ref": "#/components/schemas/BaseResp"}
                    }
                }
            },
            "500": {
                "description": "服务器内部错误",
                "content": {
                    "application/json": {
                        "schema": {"$ref": "#/components/schemas/BaseResp"}
                    }
                }
            }
        }
        
        return http_method, path, operation

    def generate_openapi_spec(self, idl_dir: str) -> Dict[str, Any]:
        """生成完整的 OpenAPI 规范"""
        # 首先处理 common.thrift
        common_file = os.path.join(idl_dir, "common.thrift")
        if os.path.exists(common_file):
            common_data = self.parse_thrift_file(common_file)
            for struct_name, struct_info in common_data["structs"].items():
                self.openapi_spec["components"]["schemas"][struct_name] = self.convert_struct_to_schema(struct_name, struct_info)
        
        # 添加标签
        for module_key, module_info in self.modules.items():
            self.openapi_spec["tags"].append({
                "name": module_info["name"],
                "description": module_info["description"]
            })
        
        # 处理其他模块文件
        for filename in os.listdir(idl_dir):
            if filename.endswith('.thrift') and filename != 'common.thrift' and filename != 'main.thrift':
                file_path = os.path.join(idl_dir, filename)
                module_data = self.parse_thrift_file(file_path)
                module_name = module_data["module_name"]
                
                # 添加结构体定义
                for struct_name, struct_info in module_data["structs"].items():
                    self.openapi_spec["components"]["schemas"][struct_name] = self.convert_struct_to_schema(struct_name, struct_info)
                
                # 添加服务方法
                for service_name, service_info in module_data["services"].items():
                    for method_info in service_info["methods"]:
                        operation_result = self.convert_method_to_operation(method_info, module_name)
                        if operation_result:
                            http_method, path, operation = operation_result
                            
                            if path not in self.openapi_spec["paths"]:
                                self.openapi_spec["paths"][path] = {}
                            
                            self.openapi_spec["paths"][path][http_method] = operation
        
        return self.openapi_spec

    def save_openapi_spec(self, output_file: str):
        """保存 OpenAPI 规范到文件"""
        with open(output_file, 'w', encoding='utf-8') as f:
            json.dump(self.openapi_spec, f, ensure_ascii=False, indent=2)

def main():
    """主函数"""
    parser = argparse.ArgumentParser(description="根据 Thrift IDL 文件生成 OpenAPI 3.0 规范文件")
    parser.add_argument("idl_dir", help="IDL 文件所在目录")
    parser.add_argument("--output-dir", "-o", dest="output_dir", help="OpenAPI 输出目录，默认当前工作目录", default=None)
    args = parser.parse_args()

    idl_dir = os.path.abspath(args.idl_dir)
    output_dir = os.path.abspath(args.output_dir) if args.output_dir else os.getcwd()
    output_file = os.path.join(output_dir, "openapi.json")
    
    converter = ThriftToOpenAPIConverter()
    openapi_spec = converter.generate_openapi_spec(idl_dir)
    converter.save_openapi_spec(output_file)
    
    print(f"OpenAPI 规范已生成: {output_file}")
    print(f"总共包含 {len(openapi_spec['paths'])} 个接口")
    print(f"总共包含 {len(openapi_spec['components']['schemas'])} 个数据结构")

if __name__ == "__main__":
    main()
