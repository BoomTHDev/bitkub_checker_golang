package handler

import (
	"net/http"
	"sync"

	"github.com/boomthdev/wld_check_bk/config"
	"github.com/boomthdev/wld_check_bk/server"
)

var (
	once sync.Once
	conf *config.Config
)

func init() {
	once.Do(func() {
		conf = config.ConfigGetting()
	})
}

func Handler(w http.ResponseWriter, r *http.Request) {
	r.RequestURI = r.URL.String()
	server.BuildVercelHandler(conf).ServeHTTP(w, r)
}
