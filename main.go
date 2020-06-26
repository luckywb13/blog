package main

import (
	"flag"
	"github.com/luckywb13/blog/conf"
	"github.com/luckywb13/blog/models"
	"github.com/luckywb13/blog/pkg/log"
	"github.com/luckywb13/blog/server"
)


func main() {
	flag.Parse()
	var err error
	var c *conf.Config

	if c, err = conf.Init(); err != nil {
		panic(err)
	}


	log.InitLog(c.LogConf)

	if err = models.Init(c.DBConf); err != nil {
		panic(err)
	}

	defer func() {
		models.Close()
		log.DestroyLogger()
	}()

	if err = server.NewServer(c.ServerConf).Run(); err != nil {
		panic(err)
	}
}

