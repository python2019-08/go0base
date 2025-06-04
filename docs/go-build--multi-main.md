# 1.一个go工程下,有超过1个main函数,如何编译

在 Go 工程中，如果有多个 `main` 函数（即多个入口点），通常表示你有多个可执行程序。Go 编译器默认每个包只能有一个 `main` 函数，因此需要为每个可执行程序单独编译。以下是几种常见的解决方案：


### 一、多入口文件结构
假设你的项目结构如下：
```
project/
├── cmd/
│   ├── server/
│   │   └── main.go    # 包名为 main，包含 server 的 main 函数
│   └── client/
│       └── main.go    # 包名为 main，包含 client 的 main 函数
├── internal/         # 内部包
└── pkg/              # 公共包
```


### 二、编译方法

#### 1. **使用 `go build` 编译单个入口**
指定入口文件的路径：
```bash
# 编译 server
go build -o bin/server ./cmd/server

# 编译 client
go build -o bin/client ./cmd/client
```
- `-o bin/server`：指定输出可执行文件的路径和名称。
- `./cmd/server`：入口文件所在的目录路径。

#### 2. **使用 `go install` 安装到 `$GOPATH/bin`**
```bash
# 安装 server
go install ./cmd/server

# 安装 client
go install ./cmd/client
```
- 安装后的可执行文件会被放在 `$GOPATH/bin` 目录下，可直接在命令行调用。

#### 3. **使用 Makefile 批量编译**
创建 `Makefile` 简化操作：
```makefile
.PHONY: build clean

# 输出目录
OUTPUT_DIR := bin

# 所有可执行文件
BINS := server client

# 编译所有目标
build: $(BINS)

# 编译单个目标
$(BINS): %: $(OUTPUT_DIR)/%

# 生成输出目录
$(OUTPUT_DIR):
    mkdir -p $(OUTPUT_DIR)

# 编译每个目标
$(OUTPUT_DIR)/%: cmd/%/main.go | $(OUTPUT_DIR)
    go build -o $@ ./cmd/$*

# 清理
clean:
    rm -rf $(OUTPUT_DIR)
```
使用方法：
```bash
# 编译所有可执行文件
make build

# 编译单个目标
make server
```


### 三、注意事项

1. **包名必须为 `main`**  
   每个入口文件的包声明必须是 `package main`，且包含 `main()` 函数。

2. **避免重复依赖**  
   不同入口可能依赖相同的包，Go 会自动处理依赖关系，但需确保依赖版本一致。

3. **使用 Go Modules**  
   推荐使用 Go Modules 管理依赖：
   ```bash
   go mod init example.com/project
   go mod tidy  # 自动整理依赖
   ```

4. **交叉编译**  
   如果需要为不同平台编译：
   ```bash
   # 编译 Linux 平台的 server
   GOOS=linux GOARCH=amd64 go build -o bin/server-linux ./cmd/server
   ```


### 四、更复杂的场景

#### 1. **条件编译**  
   如果多个入口共享部分代码，但需要根据条件编译：
   ```go
   // +build server

   // 仅在编译 server 时包含的代码
   ```
编译时使用标签：
```bash
go build -tags "server" -o bin/server ./cmd/server
```

#### 2. **使用 `go.work` 工作区（Go 1.18+）**  
   如果项目包含多个模块，可使用工作区管理：
```bash
# 创建 go.work 文件
go work init ./cmd/server ./cmd/client
```
编译时自动识别工作区配置。


### 五、示例命令
编译后，可执行文件会出现在指定目录（如 `bin/`），直接运行即可：
```bash
./bin/server  # 启动服务器
./bin/client  # 启动客户端
```


通过合理组织目录结构和使用编译命令，可以轻松管理 Go 项目中的多个入口点。