package acmeproxy

import (
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/go-acme/lego/v4/challenge"
	"gopkg.in/urfave/cli.v1"
)

type Server struct {
	HttpServer	*http.Server
	Provider	challenge.Provider
	HtpasswdFile	string
	AllowedIPs	[]string
	AllowedDomains	[]string
	AccesslogFile	string
}

func NewServer(config *Config) (*Server, error) {
	return &Server{
		HttpServer:     config.HttpServer,
		Provider:       config.Provider,
		HtpasswdFile:   config.HtpasswdFile,
		AllowedIPs:     config.AllowedIPs,
		AllowedDomains: config.AllowedDomains,
		AccesslogFile:  config.AccesslogFile,
	}, nil
}

func RunServer(ctx *cli.Context, config *Config) {
	if config.HttpServer.TLSConfig != nil {
		log.WithFields(log.Fields{
			"endpoint": "https://" + ctx.GlobalString("interface") + ":" + strconv.Itoa(ctx.GlobalInt("port")),
			"addr":  config.HttpServer.Addr,
		}).Info("Starting acmeproxy")
		log.Fatal(config.HttpServer.ListenAndServeTLS("", ""))
	} else {
		log.WithFields(log.Fields{
			"endpoint": "http://" + ctx.GlobalString("interface") + ":" + strconv.Itoa(ctx.GlobalInt("port")),
			"addr":  config.HttpServer.Addr,
		}).Info("Starting acmeproxy")
		log.Fatal(config.HttpServer.ListenAndServe())
	}
}
