package main

import (
    "fmt"
    "net/http"
)

// 处理根路径请求
func HandleMainRoute(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Main Route")
}

// 处理/api路径请求
func HandleApiRoute(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Second Route")
}
