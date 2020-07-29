package handlers

import (
	"chitchat/models"
	"net/http"
)

// 论坛首页路由处理器方法
func Index(w http.ResponseWriter, r *http.Request) {
	threads, err := models.Threads();
	if err == nil {
		generateHTML(w, threads, "layout", "navbar", "index")
	}
}