package main

import (
	"flag"
	"github.com/MenciusCheng/superman/conf"
	"github.com/MenciusCheng/superman/server/http"
	"github.com/MenciusCheng/superman/service"
	"github.com/MenciusCheng/superman/util/dragons"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	configS := flag.String("config", "config/config.toml", "Configuration file")
	flag.Parse()

	dragons.Init(
		dragons.ConfigPath(*configS),
	)

}

func main() {

	defer dragons.Shutdown()

	// init local config
	cfg, err := conf.Init()
	if err != nil {
		log.Fatalf("service config init error %s", err)
	}

	// create a service instance
	srv := service.New(cfg)

	// init and start http server
	http.Init(srv, cfg)

	defer http.Shutdown()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-sigChan
		log.Printf("get a signal %s\n", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Println("demo server exit now...")
			return
		case syscall.SIGHUP:
		default:
		}
	}
}
