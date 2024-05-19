package iface

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DefaultHttpSPI interface {
	CommonSPI
	GetServer() *http.Server
	GetEngine() *gin.Engine
}
