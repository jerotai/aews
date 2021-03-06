package api

import (
	"aews/api/router"
	"github.com/labstack/echo"
	"github.com/facebookgo/grace/gracehttp"
)

type APIService struct {
}

func NewApiService() *APIService {
	return &APIService{}
}

/**
 * 啟動API服務
 */
func (s *APIService) Start() {
	echo := echo.New()
	router.InitRouting(echo)
	echo.Server.Addr = ":8080"
	gracehttp.Serve(echo.Server)
	echo.Close()
}
