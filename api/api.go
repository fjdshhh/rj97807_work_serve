package main

import (
	"flag"
	"fmt"
	"net/http"
	"rj97807_work_serve/api/internal/config"
	"rj97807_work_serve/api/internal/handler"
	"rj97807_work_serve/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/api-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	c.MaxBytes = 62914560 //gozero默认最大传输byte为8MB 此处修改为60MB
	c.Timeout = 3000000   //gozero默认超时时间为3s 此处修改为30s
	server := rest.MustNewServer(c.RestConf, rest.WithCustomCors(nil, CorsAllow, "*"))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()

}
func CorsAllow(w http.ResponseWriter) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
}
