@echo off

:: 创建输出目录
mkdir pb 2>nul

:: 清理旧的生成文件
if exist pb rmdir /s /q pb
mkdir pb

:: 遍历idls目录下的所有子目录
for /d %%d in (idls\*) do (
    set "service_name=%%~nxd"
    echo Processing !service_name! service...

    :: 确保输出目录存在
    mkdir "pb\!service_name!" 2>nul

    :: 处理该目录下的所有proto文件
    for %%f in ("%%d\*.proto") do (
        echo   Generating code for %%f

        :: 使用protoc生成Go代码，输出到对应的服务目录
        protoc --proto_path=idls ^
            --go_out=pb ^
            --go_opt=paths=source_relative ^
            --go-grpc_out=pb ^
            --go-grpc_opt=paths=source_relative ^
            "%%f"
    )
)

echo Code generation completed!