package main

import (
	"beyond/pkg/consul"
	"beyond/pkg/xcode"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"

	"beyond/application/article/api/internal/config"
	"beyond/application/article/api/internal/handler"
	"beyond/application/article/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "application/article/api/etc/article-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 自定义错误处理方法
	httpx.SetErrorHandler(xcode.ErrHandler)

	err := consul.Register(c.Consul, fmt.Sprintf("%s:%d", c.ServiceConf.Prometheus.Host, c.ServiceConf.Prometheus.Port))
	if err != nil {
		fmt.Printf("register consul error: %v\n", err)
	}

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
