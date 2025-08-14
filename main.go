package main

import (
	"github.com/boomthdev/wld_check_bk/config"
	"github.com/boomthdev/wld_check_bk/server"
)

func main() {
	conf := config.ConfigGetting()
	server := server.NewFiberServer(conf)
	server.Start()
}
