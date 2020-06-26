package server

import (
	"github.com/gin-gonic/gin"
	"github.com/luckywb13/blog/conf"
	"github.com/luckywb13/blog/routers"
)

type Server struct {
	engine *gin.Engine
	c *conf.Server
}

func NewServer(c *conf.Server) *Server {
	if !conf.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	s := &Server{
		engine: routers.InitRouter(),
		c:c,
	}

	return s
}

func (s *Server)Run() (err error)  {
	if err = s.engine.Run(s.c.Address); err != nil {
		return
	}
	return
}

