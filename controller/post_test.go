package controller

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreatePost(t *testing.T) {
	gin.SetMode(gin.TestMode)
	e := gin.Default()
	url := "/api/v1/post"
	e.POST(url, CreatePost)
	body := `{
		"community_id":1,
		"title":"giao",
		"content":"just a test"
	}`
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))
	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)
	fmt.Println(w.Code)
	fmt.Println(w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
}
