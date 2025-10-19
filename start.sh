#!/bin/bash

# 零工APP后端服务启动脚本

set -e

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 项目信息
PROJECT_NAME="零工APP后端服务"
VERSION="1.0.0"

# 打印带颜色的消息
print_message() {
    local color=$1
    local message=$2
    echo -e "${color}${message}${NC}"
}

# 显示帮助信息
show_help() {
    echo "用法: $0 [选项]"
    echo ""
    echo "选项:"
    echo "  server [env]     启动服务器 (env: example|prod, 默认: example)"
    echo "  migrate [env]    执行数据库迁移 (env: example|prod, 默认: example)"
    echo "  build           编译项目"
    echo "  test            运行测试"
    echo "  clean           清理构建文件"
    echo "  help            显示帮助信息"
    echo "  version         显示版本信息"
    echo ""
    echo "示例:"
    echo "  $0 server example    # 启动开发环境服务器"
    echo "  $0 server prod       # 启动生产环境服务器"
    echo "  $0 migrate example   # 执行开发环境数据库迁移"
    echo "  $0 migrate prod      # 执行生产环境数据库迁移"
    echo ""
}

# 显示版本信息
show_version() {
    print_message $BLUE "$PROJECT_NAME"
    print_message $BLUE "版本: $VERSION"
    print_message $BLUE "Go版本: $(go version | cut -d' ' -f3)"
    print_message $BLUE "构建时间: $(date +"%Y-%m-%d %H:%M:%S")"
}

# 检查Go环境
check_go() {
    if ! command -v go &> /dev/null; then
        print_message $RED "错误: Go 未安装或不在 PATH 中"
        exit 1
    fi
}

# 检查配置文件
check_config() {
    local env=$1
    local config_file="conf/config.${env}.yaml"
    
    if [ ! -f "$config_file" ]; then
        print_message $YELLOW "警告: 配置文件 $config_file 不存在"
        if [ -f "conf/config.yaml" ]; then
            print_message $YELLOW "将使用默认配置文件: config/config.yaml"
        else
            print_message $RED "错误: 没有找到任何配置文件"
            exit 1
        fi
    fi
}

# 创建必要目录
create_dirs() {
    mkdir -p logs
    mkdir -p output/bin
}

# 启动服务器
start_server() {
    local env=${1:-example}
    
    print_message $GREEN "启动服务器模式 (环境: $env)..."
    check_go
    check_config $env
    create_dirs
    
    go run . -mode server -env $env
}

# 执行数据库迁移
run_migrate() {
    local env=${1:-example}
    
    print_message $GREEN "执行数据库迁移 (环境: $env)..."
    check_go
    check_config $env
    create_dirs
    
    go run . -mode migrate -env $env
}

# 编译项目
build_project() {
    print_message $GREEN "编译项目..."
    check_go
    create_dirs
    
    go build -ldflags "-X main.Version=$VERSION -X main.BuildTime=$(date +"%Y-%m-%d %H:%M:%S")" -o output/bin/labor-clients-be .
    print_message $GREEN "编译完成: output/bin/labor-clients-be"
}

# 运行测试
run_tests() {
    print_message $GREEN "运行测试..."
    check_go
    
    go test -v ./...
}

# 清理构建文件
clean_project() {
    print_message $GREEN "清理构建文件..."
    rm -rf output/bin/*
    rm -rf logs/*.log
    go clean
    print_message $GREEN "清理完成"
}

# 主函数
main() {
    case "${1:-help}" in
        "server")
            start_server $2
            ;;
        "migrate")
            run_migrate $2
            ;;
        "build")
            build_project
            ;;
        "test")
            run_tests
            ;;
        "clean")
            clean_project
            ;;
        "help"|"-h"|"--help")
            show_help
            ;;
        "version"|"-v"|"--version")
            show_version
            ;;
        *)
            print_message $RED "错误: 未知选项 '$1'"
            echo ""
            show_help
            exit 1
            ;;
    esac
}

# 执行主函数
main "$@"
