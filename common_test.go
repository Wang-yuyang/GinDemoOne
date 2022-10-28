package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var tmpArticleList []article

func TestMain(m *testing.M) {
	// 设置为测试模式下的Gin
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func getRouter(withTemplates bool) *gin.Engine {
	// 创建类main 返回路由器
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("templates/*")
	}
	return r
}

func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request,
	f func(w *httptest.ResponseRecorder) bool) {
	// 执行传入的函数，返回一个bool
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if !f(w) {
		t.Fail()
	}
}

func saveLists() {
	// 将原始文章列表报错在一个临时变量中
	tmpArticleList = articleList
}

func restoreLists() {
	// 在执行单元测试结束后将列表恢复为初始状态
	articleList = tmpArticleList
}
