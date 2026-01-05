# Go 本地 JSON HTTP Server 示例

这是一个简洁、可扩展的 Go HTTP 服务器示例。当用户向指定 URL 发出 GET 或 POST 请求时，会调用对应的处理函数并返回 JSON 响应。

默认监听端口: `:8080`

示例路由:
- `GET /hello?name=alice` -> 返回 {"success": true, "data": {"message": "hello alice"}}
- `POST /echo` (JSON body) -> 返回接收到的 JSON

运行:

```bash
go run main.go
```

示例请求:

GET:

```bash
curl "http://localhost:8080/hello?name=yin"
```

POST:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"foo": "bar"}' http://localhost:8080/echo
```

可扩展性说明:

## ReadBTFandGetItsMember (Linux-only) ✅

- 文件: `ReadBTFandGetItsMember.go` (package `main`)
- 功能: 读取 `./.cache/btf.json`，查找与 `sk_buff` 相关（深度 ≤ 5）的 types，筛选出参数中包含这些 type 的 `FUNC` 项并写入 `./.cache/relatedFuncD5.json`。
- 导出函数: `ReadBTFandGetItsMember()`，返回 `([]map[string]interface{}, error)`。
- 注意: 仅在 **Linux** 上编译（文件包含 `//go:build linux`）。实现对 JSON 结构的断言做了防御性处理，尽量兼容原 Python 脚本的行为。

### 示例用法

```go
func main() {
    funcs, err := ReadBTFandGetItsMember()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("found %d related funcs\n", len(funcs))
}
```

- 使用 `Router.Register(method, path, handler)` 注册新的路由。
- 处理函数签名为 `func(w http.ResponseWriter, r *http.Request) error`，返回 `error` 可统一处理各种错误并返回 JSON。
- 中间件（如日志、JSON header）集中注册，易于插拔。

---

## TranslateJSON (Linux-only) ✅

- 文件: `translateJSON.go` (package `main`)
- 功能: 读取 `./.cache/relatedFuncD5.json`，将函数列表按 `id` 重建为字典并写入 `./.cache/FuncIDMap.json`，同时根据函数名规则生成 BPF C 源文件 `./.cache/kProberFunc.c`。
- 导出函数: `TranslateJSON()`，返回 `error`。
- 注意: 仅在 **Linux** 上编译（文件包含 `//go:build linux`）。此 Go 实现行为与原 `translateJSON.py` 等价，尽量保留原脚本的选择逻辑和模板。

使用示例：

```go
if err := TranslateJSON(); err != nil {
    log.Fatalf("translate json failed: %v", err)
}
```

