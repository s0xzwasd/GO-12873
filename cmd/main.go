package main

import (
	"github.com/s0xzwasd/GO-12873/internal/proxy"
	log "github.com/sirupsen/logrus"
)

func main() {
	proxy.Bar()
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}
