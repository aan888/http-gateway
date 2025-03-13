#!/bin/bash

mkdir pb
# 清理旧的生成文件
rm -rf ./pb

# 遍历idls目录下的所有子目录
for dir in idls/*/; do
    # 获取目录名称（去掉idls/前缀和尾部的斜杠）
    service_name=$(basename "$dir")
    echo "Processing $service_name service..."

    # 确保输出目录存在
    mkdir -p "pb/$service_name"

    # 处理该目录下的所有proto文件
    for proto_file in "$dir"*.proto; do
        if [ -f "$proto_file" ]; then
            echo "  Generating code for $proto_file"

            # 使用protoc生成Go代码，输出到对应的服务目录
            protoc --proto_path=idls \
                --go_out=pb \
                --go_opt=paths=source_relative \
                --go-grpc_out=pb \
                --go-grpc_opt=paths=source_relative \
                "$proto_file"
        fi
    done
done

echo "Code generation completed!"