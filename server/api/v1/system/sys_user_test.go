package system

import (
	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBaseApi_GetUserInfo(t *testing.T) {
	router := gin.Default()
	b := &BaseApi{}
	router.GET("/", b.GetUserInfo)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "", w.Body.String())

}
