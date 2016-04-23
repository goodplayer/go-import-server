package main // "import.moetang.info/go/tool/go-import-server"

import (
	"flag"
	"log"

	"github.com/gin-gonic/gin"
)

import (
	"import.moetang.info/go/lib/gin-startup"
)

var (
	fastcgiAddress string
	httpAddress    string
)

func init() {
	flag.StringVar(&fastcgiAddress, "bind", "tcp://127.0.0.1:15001", "fastcgi bind address")
	flag.StringVar(&httpAddress, "httpbind", "tcp://127.0.0.1:15002", "http bind address")
	flag.Parse()
}

func main() {
	entries, err := ReadEntries()
	if err != nil {
		log.Println(err)
		return
	}

	g := gin_startup.NewGinStartup()
	g.EnableFastCgi(fastcgiAddress)
	g.EnableHttp(httpAddress)
	g.Custom(func(r *gin.Engine) {
		r.Use(gin.Recovery(), gin.Logger())

		for _, v := range entries {
			RegisterAction(v, r)
		}
	})
	g.Start()
	c := make(chan bool)
	<-c
}
