package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// JSONResponse 标准响应结构
type JSONResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// HandlerFunc 支持返回 error，方便统一处理
type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

// Router 是一个简单可扩展的路由器
type Router struct {
	mux *http.ServeMux
}

func NewRouter() *Router {
	return &Router{mux: http.NewServeMux()}
}

// Register 把 handler 注册到 path，并根据 method 做检查
func (rt *Router) Register(method, path string, h HandlerFunc) {
	wrapped := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			w.WriteHeader(http.StatusMethodNotAllowed)
			writeJSON(w, http.StatusMethodNotAllowed, JSONResponse{Success: false, Error: "method not allowed"})
			return
		}
		if err := h(w, r); err != nil {
			var status int
			// 可根据错误类型映射不同状态码，示例简单处理
			if errors.Is(err, context.Canceled) {
				status = http.StatusRequestTimeout
			} else {
				status = http.StatusInternalServerError
			}
			writeJSON(w, status, JSONResponse{Success: false, Error: err.Error()})
		}
	}
	rt.mux.HandleFunc(path, loggingMiddleware(jsonMiddleware(wrapped)))
}

func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rt.mux.ServeHTTP(w, r)
}

// loggingMiddleware 打印请求日志并记录耗时
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("%s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
		next(w, r)
		log.Printf("completed %s %s in %v", r.Method, r.URL.Path, time.Since(start))
	}
}

// jsonMiddleware 确保响应 Content-Type
func jsonMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next(w, r)
	}
}

// writeJSON 辅助写入 JSON
func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.WriteHeader(status)
	enc := json.NewEncoder(w)
	// 更好的控制可以设置缩进或其它选项
	if err := enc.Encode(v); err != nil {
		log.Printf("failed to write json: %v", err)
	}
}

func SocketHandler(w http.ResponseWriter, r *http.Request) error {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "world"
	}
	payload := map[string]string{"message": "hello " + name}
	writeJSON(w, http.StatusOK, JSONResponse{Success: true, Data: payload})
	return nil
}

// 简单示例：GET /hello
func helloHandler(w http.ResponseWriter, r *http.Request) error {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "world"
	}
	payload := map[string]string{"message": "hello " + name}
	writeJSON(w, http.StatusOK, JSONResponse{Success: true, Data: payload})
	return nil
}

// 简单示例：POST /echo 接受 JSON 并返回
func echoHandler(w http.ResponseWriter, r *http.Request) error {
	var body interface{}
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&body); err != nil {
		return err
	}
	writeJSON(w, http.StatusOK, JSONResponse{Success: true, Data: body})
	return nil
}

func main() {
	addr := ":8080"
	rt := NewRouter()

	// 注册路由，可按需添加更多
	rt.Register(http.MethodGet, "/QuerySockList", SocketHandler)
	rt.Register(http.MethodGet, "/hello", helloHandler)
	rt.Register(http.MethodPost, "/echo", echoHandler)

	srv := &http.Server{
		Addr:    addr,
		Handler: rt,
	}

	// 优雅关闭
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("server listening on %s", addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %v", err)
		}
	}()

	<-stop
	log.Println("shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown failed:%+v", err)
	}
	log.Println("server exited properly")
}
