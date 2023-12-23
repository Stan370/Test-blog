package main

import (
	"fmt"
	"net/http"

	"github.com/Stan370/Test-blog/db"
)

func main() {
	dbObject := db.InitDatabase()

	// 创建一个路由并将处理器函数与路径关联
	http.HandleFunc("/api", corsHandler)

	// 启动 HTTP 服务器并监听端口
	port := ":8080"
	fmt.Printf("Server is running on port %s\n", port)
	http.ListenAndServe(port, nil)
}
