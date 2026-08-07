package httpd

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"sync"

	"github.com/caddyserver/certmagic"
)

var hostsFile = "httpdhosts.json"
var hosts = map[string]string{}
var mu sync.RWMutex

func Init() {
	if err := LoadHosts(); err != nil {
		log.Fatalln(err)
	}

	certmagic.DefaultACME.Agreed = true
	certmagic.DefaultACME.Email = "runoneall@nodeloc.com"

	cfg := certmagic.NewDefault()
	cfg.OnDemand = &certmagic.OnDemandConfig{
		DecisionFunc: func(ctx context.Context, name string) error {
			mu.RLock()
			_, exists := hosts[name]
			mu.RUnlock()

			if !exists {
				return fmt.Errorf("domain %s not allowed", name)
			}

			return nil
		},
	}

	server := &http.Server{
		Addr:      ":443",
		Handler:   handler(),
		TLSConfig: cfg.TLSConfig(),
	}

	log.Println("start httpd")
	go func() {
		if err := server.ListenAndServeTLS("", ""); err != nil {
			slog.Error(err.Error())
		}
	}()
}
