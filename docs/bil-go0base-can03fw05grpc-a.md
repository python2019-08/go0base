# 2025年最新Golang保姆级公开课教程-零基础也可学！（完整版）
2024-12-31 14:17:39
https://www.bilibili.com/video/BV1Y26GYhEGq?spm_id_from=333.788.videopod.episodes&vd_source=4212b105520112daf65694a1e5944e23&p=26
# 5.【golang框架】如何使用rpc框架grpc开发微服务项目

# 1.Golang 与 Protobuf 的结合使用
 
## 1.1什么是 Protobuf？
Protobuf（Protocol Buffers）是谷歌开发的一种数据序列化格式，特点是高效、跨平台、语言中立，常用于微服务通信、数据存储等场景。相比 JSON/XML，它序列化后体积更小、解析速度更快。

## 1.2在 Golang 中使用 Protobuf 的步骤

### 1.2.1环境准备
### 1.2.2安装 Protobuf 编译器
(1) 从 [官网](https://github.com/protocolbuffers/protobuf/releases) 下载对应系统的二进制包（如 protoc-xxx.zip），解压后将 protoc 加入环境变量。

(2)安装 Golang 插件
```bash
go get -u github.com/golang/protobuf/protoc-gen-go
```

### 1.2.2 定义 .proto 文件
在项目中创建 .proto 文件（如 user.proto），示例：
```
syntax = "proto3";  // 指定版本

// 包名（对应 Golang 中的包路径）
package pb;

// 定义消息结构
message User {
  int64 id = 1;
  string name = 2;
  int32 age = 3;
}

// 定义服务接口（可选，用于 gRPC）
service UserService {
  rpc GetUser(UserRequest) returns (User);
}

message UserRequest {
  int64 id = 1;
}
```

#### 5.1.2.3、生成 Golang 代码
在 .proto 文件所在目录执行：
```sh
protoc --go_out=plugins=grpc:. user.proto
```
> 1. --go_out：指定生成 Golang 代码，plugins=grpc 表示同时生成 gRPC 相关代码
> 2. 生成的文件会包含结构体、序列化 / 反序列化方法及 gRPC 接口实现

#### 5.1.2.4、在 Golang 中使用生成的代码
(1)导入包
import "your-project/pb"  // 对应 .proto 中的 package 路径

(2)创建和序列化消息
```go
// 创建消息实例
user := &pb.User{
    Id:   1,
    Name: "张三",
    Age:  30,
}

// 序列化（二进制格式）
data, err := user.Marshal()
if err != nil {
    // 错误处理
}

// 反序列化
newUser := &pb.User{}
err = newUser.Unmarshal(data)
```

(3)结合 gRPC 使用（可选）
服务端：实现 .proto 定义的服务接口
客户端：通过生成的客户端代码调用远程服务

#### 5.1.2.5、进阶用法

(1)字段编号（Field Number）
字段编号在序列化后是固定的，修改 .proto 时需注意：
* 新增字段可分配新编号，不影响旧数据
* 删除字段需谨慎（避免与后续新增字段冲突），建议标记为 deprecated

(2)默认值与空值处理
Protobuf 中未赋值的字段会使用默认值（如 int32 为 0，string 为空字符串），需在代码中注意空值判断。

(3)与其他工具结合
配合 protoc-gen-go-grpc 生成更优的 gRPC 代码
使用 protobuf-json 实现 Protobuf 与 JSON 的转换

#### 5.1.2.6、最佳实践
* 按功能模块组织 .proto 文件，避免过度集中
* 版本迭代时遵循兼容性原则（新增字段、不修改已有字段编号）
* 对重要消息添加 json_name 选项，便于与 JSON 交互：
``` 
  string name = 2 [json_name = "user_name"];
```
通过以上步骤，即可在 Golang 项目中高效使用 Protobuf 进行数据序列化和服务通信，尤其适合微服务、API 接口等对性能要求较高的场景。

  
# 2 举一个完整的包含service结构中rpc方法的proto文件的例子
下面是一个完整的 `.proto` 文件示例，包含了 `service` 结构、多种 `rpc` 方法（普通 RPC、客户端流式、服务器流式、双向流式）以及相关的 `message` 定义。这个例子模拟了一个简单的博客服务：
```blog.proto 
syntax = "proto3";

// 包名，避免命名冲突
package blog;

// 导入 google/protobuf 中的空消息类型
import "google/protobuf/empty.proto";
import "google/protobuf/timestamp.proto";

// 博客文章消息
message Post {
  string id = 1;
  string title = 2;
  string content = 3;
  string author_id = 4;
  google.protobuf.Timestamp created_at = 5;
  google.protobuf.Timestamp updated_at = 6;
  repeated string tags = 7;
}

// 创建博客文章请求
message CreatePostRequest {
  Post post = 1;
}

// 创建博客文章响应
message CreatePostResponse {
  Post post = 1;
}

// 获取博客文章请求
message GetPostRequest {
  string post_id = 1;
}

// 获取博客文章列表请求
message ListPostsRequest {
  int32 page_size = 1;
  string page_token = 2;
  string author_id = 3;  // 可选：按作者筛选
  repeated string tags = 4;  // 可选：按标签筛选
}

// 获取博客文章列表响应
message ListPostsResponse {
  repeated Post posts = 1;
  string next_page_token = 2;
}

// 更新博客文章请求
message UpdatePostRequest {
  Post post = 1;
}

// 删除博客文章请求
message DeletePostRequest {
  string post_id = 1;
}

// 批量操作结果
message BatchOperationResult {
  int32 success_count = 1;
  int32 failure_count = 2;
  map<string, string> failures = 3;  // key: post_id, value: error message
}

// 评论消息
message Comment {
  string id = 1;
  string post_id = 2;
  string user_id = 3;
  string content = 4;
  google.protobuf.Timestamp created_at = 5;
}

// 博客服务定义
service BlogService {
  // 1. 普通 RPC：创建单篇博客
  rpc CreatePost(CreatePostRequest) returns (CreatePostResponse) {}
  
  // 2. 普通 RPC：获取单篇博客
  rpc GetPost(GetPostRequest) returns (Post) {}
  
  // 3. 普通 RPC：更新单篇博客
  rpc UpdatePost(UpdatePostRequest) returns (Post) {}
  
  // 4. 普通 RPC：删除单篇博客
  rpc DeletePost(DeletePostRequest) returns (google.protobuf.Empty) {}
  
  // 5. 普通 RPC：获取博客列表（分页）
  rpc ListPosts(ListPostsRequest) returns (ListPostsResponse) {}
  
  // 6. 客户端流式 RPC：批量创建博客（客户端发送多篇文章，服务器返回汇总结果）
  rpc BatchCreatePosts(stream Post) returns (BatchOperationResult) {}
  
  // 7. 服务器流式 RPC：订阅特定作者的新博客（客户端发送作者 ID，服务器实时推送新文章）
  rpc SubscribeToAuthorPosts(GetPostRequest) returns (stream Post) {}
  
  // 8. 双向流式 RPC：实时评论系统（客户端发送评论，服务器广播给所有订阅者）
  rpc LiveComments(stream Comment) returns (stream Comment) {}
  
  // 9. 普通 RPC：获取博客统计信息
  rpc GetBlogStats(google.protobuf.Empty) returns (BlogStats) {}
}

// 博客统计信息
message BlogStats {
  int32 total_posts = 1;
  int32 total_comments = 2;
  map<string, int32> author_post_count = 3;  // key: author_id, value: post_count
  google.protobuf.Timestamp last_updated = 4;
}
```


### 关键点说明

1. **消息定义**：
   - 定义了 `Post`、`Comment` 等核心实体消息
   - 使用了 `google.protobuf.Timestamp` 处理时间戳
   - 包含了分页查询、批量操作等场景的请求/响应消息

2. **服务定义**：
   - 实现了 CRUD 基本操作（CreatePost、GetPost、UpdatePost、DeletePost）
   - 展示了三种流式 RPC 的用法：
     - 客户端流式：`BatchCreatePosts`（客户端批量发送数据）
     - 服务器流式：`SubscribeToAuthorPosts`（服务器实时推送数据）
     - 双向流式：`LiveComments`（客户端和服务器双向实时通信）
   - 包含了复杂参数的设计（如分页、过滤条件）

3. **最佳实践**：
   - 使用 `google.protobuf.Empty` 表示无返回值
   - 为列表请求添加分页支持（page_size、page_token）
   - 为批量操作设计统一的结果消息结构
   - 合理使用字段编号（1-15 用于高频字段）

4. **扩展性考虑**：
   - 为 `Post` 消息预留了可扩展字段（tags）
   - 设计了统计信息接口 `GetBlogStats`

这个例子涵盖了 gRPC 服务的常见场景，你可以根据实际需求扩展更多功能，如权限验证、搜索功能等。

# 6 protocol error: received DATA after END_STREAM
### gRPC 协议错误 "received DATA after END_STREAM" 解析与解决方案

这个错误表示 gRPC 在通信过程中违反了协议规范：**在接收到流结束标记（END_STREAM）后，又收到了数据（DATA）**。这通常发生在流式 RPC 场景中，是客户端或服务器端代码逻辑错误导致的。


### 一、错误根源分析
gRPC 流式 RPC 有严格的消息顺序规则：
1. **客户端流式 RPC**：客户端发送多个 DATA 后，必须发送 END_STREAM 表示请求结束，服务器处理后返回一个响应。
2. **服务器流式 RPC**：服务器发送多个 DATA，最后用 END_STREAM 表示响应结束。
3. **双向流式 RPC**：双方可交替发送 DATA，但一旦一方发送 END_STREAM，另一方不能再发送 DATA。

**错误场景示例**：
- 客户端在流式请求中发送了 END_STREAM 后，又继续发送 DATA。
- 服务器在流式响应中发送完所有 DATA 并标记 END_STREAM 后，又意外发送了新 DATA。


### 二、常见代码错误场景

#### 1. 客户端错误（以 Go 为例）
```go
// ❌ 错误示例：客户端流式 RPC 中先发送 END_STREAM，再发送数据
stream, err := client.StreamData(ctx)
if err != nil { /* 处理错误 */ }

// 发送前两条数据
if err := stream.Send(&Data{Value: "1"}); err != nil { /* 处理错误 */ }
if err := stream.Send(&Data{Value: "2"}); err != nil { /* 处理错误 */ }

// 错误：先调用 CloseSend() 发送 END_STREAM
if err := stream.CloseSend(); err != nil { /* 处理错误 */ }

// 之后又尝试发送数据（导致协议错误）
if err := stream.Send(&Data{Value: "3"}); err != nil { 
    // 这里会触发 "received DATA after END_STREAM"
}
```

#### 2. 服务器错误（以 Java 为例）
```java
// ❌ 错误示例：服务器流式 RPC 中先发送 END_STREAM，再发送数据
@Override
public void listData(StreamRequest request, StreamObserver<Data> responseObserver) {
    // 发送前两条数据
    responseObserver.onNext(Data.newBuilder().setValue("1").build());
    responseObserver.onNext(Data.newBuilder().setValue("2").build());
    
    // 错误：先调用 onCompleted() 发送 END_STREAM
    responseObserver.onCompleted();
    
    // 之后又尝试发送数据（导致协议错误）
    responseObserver.onNext(Data.newBuilder().setValue("3").build());
}
```


### 三、解决方案

#### 1. 检查流式 RPC 的消息顺序
- **客户端流式 RPC**：
  - 确保在所有 `Send()` 调用后，仅调用一次 `CloseSend()`（或类似结束方法），且之后不再调用 `Send()`。
  - **Go 正确示例**：
    ```go
    stream, err := client.StreamData(ctx)
    if err != nil { /* 处理错误 */ }
    
    for i := 0; i < 3; i++ {
        if err := stream.Send(&Data{Value: strconv.Itoa(i)}); err != nil {
            break
        }
    }
    // 仅在所有发送完成后调用 CloseSend()
    if err := stream.CloseSend(); err != nil { /* 处理错误 */ }
    ```

- **服务器流式 RPC**：
  - 确保在所有 `onNext()` 调用后，仅调用一次 `onCompleted()`（或类似结束方法），且之后不再调用 `onNext()`。
  - **Java 正确示例**：
    ```java
    @Override
    public void listData(StreamRequest request, StreamObserver<Data> responseObserver) {
        for (int i = 0; i < 3; i++) {
            responseObserver.onNext(Data.newBuilder().setValue(String.valueOf(i)).build());
        }
        // 仅在所有响应发送完成后调用 onCompleted()
        responseObserver.onCompleted();
    }
    ```

#### 2. 处理异步流式 RPC 的并发问题
- 在异步编程模型中，确保流操作的线程安全性：
  - **JavaScript 示例（gRPC-Web）**：
    ```javascript
    const stream = client.streamData(request, metadata);
    let count = 0;
    
    // 监听响应
    stream.on('data', (response) => {
        count++;
        if (count === 3) {
            stream.cancel(); // 避免后续数据
        }
    });
    
    stream.on('end', () => {
        console.log('流已结束');
    });
    ```

#### 3. 排查中间件或拦截器的影响
- 检查是否有拦截器意外关闭了流或发送了额外数据：
  - **Go 拦截器示例（正确关闭流）**：
    ```go
    func streamInterceptor(ctx context.Context, desc *grpc.StreamDesc, 
                          cc *grpc.ClientConn, method string,
                          streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
        stream, err := streamer(ctx, desc, cc, method, opts...)
        if err != nil {
            return nil, err
        }
        // 包装 Stream 以确保正确关闭
        return &wrappedStream{stream}, nil
    }
    
    type wrappedStream struct {
        grpc.ClientStream
    }
    
    // 确保 CloseSend() 只调用一次
    func (w *wrappedStream) CloseSend() error {
        // 可添加日志或校验逻辑
        return w.ClientStream.CloseSend()
    }
    ```


### 四、调试与日志定位
1. **启用 gRPC 详细日志**：
   - **Go**：设置环境变量 `GRPC_GO_LOG_SEVERITY_LEVEL=info` 和 `GRPC_GO_LOG_VERBOSITY_LEVEL=99`。
   - **Java**：配置 `io.grpc` 日志级别为 `FINE`。

2. **捕获错误时的详细信息**：
   - 在错误处理中打印完整的错误堆栈和元数据：
   ```go
   if err != nil {
       if status, ok := status.FromError(err); ok {
           log.Printf("gRPC 错误: code=%s, message=%s, details=%v", 
                     status.Code(), status.Message(), status.Details())
       } else {
           log.Printf("普通错误: %v", err)
       }
   }
   ```

3. **使用网络抓包工具**：
   - 通过 Wireshark 或 tcpdump 分析 gRPC 流量（需解码 HTTP/2 协议），确认 DATA 和 END_STREAM 的顺序。


### 五、协议层面的补充说明
gRPC 基于 HTTP/2 协议，流式 RPC 的消息顺序由 HTTP/2 流控制规范定义：
- **END_STREAM**：表示流的结束，后续不能再发送同方向的 DATA 帧。
- 该错误本质是 HTTP/2 协议违规，gRPC 层将其转换为 `protocol error`。


通过以上方法，可准确定位并修复流式 RPC 中的消息顺序错误，确保 gRPC 通信符合协议规范。

# 7 go mod tidy 时 protocol error: received DATA after END_STREAM
```
go: downloading go.opentelemetry.io/otel/trace v1.34.0
protocol error: received DATA after END_STREAM
```


### 下载 OpenTelemetry 包时的 gRPC 协议错误分析

你遇到的错误是在下载 Go 依赖包 `go.opentelemetry.io/otel/trace v1.34.0` 时发生的，这表明问题可能与 Go 模块下载机制（基于 gRPC）或网络连接有关。


### 一、错误原因分析

这个 `protocol error: received DATA after END_STREAM` 错误通常在以下场景出现：

1. **网络不稳定**：
   - 代理、防火墙或 VPN 中断了正常的 gRPC 通信。
   - DNS 解析异常导致连接到不可靠的服务器。

2. **Go 模块代理问题**：
   - 默认代理 `proxy.golang.org` 在某些地区可能不稳定。
   - 自定义代理服务器配置错误。

3. **本地缓存损坏**：
   - `$GOPATH/pkg/mod` 中的缓存文件可能损坏。

4. **服务器端问题**：
   - OpenTelemetry 的官方仓库或 CDN 暂时不可用。


### 二、解决方案

#### 1. **切换 Go 模块代理**
临时使用阿里云或其他国内代理：
```bash
go env -w GOPROXY=https://goproxy.cn,direct
```

#### 2. **清除本地缓存**
删除损坏的缓存文件：
```bash
go clean -modcache
```

#### 3. **重试下载**
使用 `-v` 选项获取详细日志：
```bash
go get -v go.opentelemetry.io/otel/trace@v1.34.0
```

#### 4. **检查网络连接**
- 关闭代理、VPN 或防火墙后重试：
  ```bash
  # 临时禁用代理
  unset http_proxy https_proxy all_proxy
  ```

#### 5. **手动下载依赖**
先单独下载 OpenTelemetry 核心包：
```bash
go get -v go.opentelemetry.io/otel@v1.34.0
```

#### 6. **使用 GOPROXY=direct 绕过代理**
直接从源仓库下载（不推荐生产环境使用）：
```bash
GOPROXY=direct go get -v go.opentelemetry.io/otel/trace@v1.34.0
```

#### 7. **升级 Go 版本**
确保使用最新稳定版 Go（如 1.21.x）：
```bash
# 检查当前版本
go version

# 升级到最新版
go install golang.org/dl/go1.21.8@latest
go1.21.8 download
```


### 三、其他可能的解决方案

#### 1. **设置 GOSUMDB**
禁用或切换校验和数据库：
```bash
# 禁用校验和验证（不推荐生产环境）
go env -w GOSUMDB=off

# 或使用国内替代
go env -w GOSUMDB=https://sum.golang.google.cn
```

#### 2. **使用 go.work 工作区（Go 1.18+）**
如果项目包含多个模块：
```bash
# 创建工作区文件
go work init .
go work use ./path/to/module
```

#### 3. **检查文件权限**
确保 GOPATH 目录可读写：
```bash
chmod -R 755 $GOPATH
```


### 四、示例命令组合
```bash
# 1. 清除缓存
go clean -modcache

# 2. 设置国内代理
go env -w GOPROXY=https://goproxy.cn,direct

# 3. 单独下载 OpenTelemetry 核心包
go get -v go.opentelemetry.io/otel@v1.34.0

# 4. 再次尝试完整下载
go mod tidy
```


### 五、预防措施

1. **定期更新依赖**：
   ```bash
   go get -u
   go mod tidy
   ```

2. **使用 Go Modules 缓存**：
   ```bash
   # 缓存模块下载结果
   go env -w GOMODCACHE=$HOME/.cache/go-mod
   ```

3. **固定依赖版本**：
   在 `go.mod` 中明确指定版本，避免使用 `@latest`。


通过以上步骤，应该能解决下载 OpenTelemetry 包时的 gRPC 协议错误。如果问题仍然存在，可能需要联系 OpenTelemetry 社区或检查网络基础设施。

 