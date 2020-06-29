package test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	router2 "gotest/router"
	"net/http"
	"net/http/httptest"
	"testing"
)
var r *gin.Engine

func init()  {
	gin.SetMode(gin.TestMode)
	r = router2.RouterInit()
}

func TestIndexHtml(t *testing.T) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
