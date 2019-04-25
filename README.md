# 复制文件到另一个目录多份
要求go >= 1.11

# 克隆并运行
```
git clone https://github.com/zhaopei8948/copyfile
cd copyfile
go run copyfile.go 源文件 目标目录 复制几份
```

# windows 构建
```
cd copyfile
GO111MODULE=on GOOS=windows GOARCH=amd64 go build -o copyfile.exe
```

# linux 构建
```
cd copyfile
GO111MODULE=on GOOS=linux GOARCH=amd64 go build -o copyfile
```
