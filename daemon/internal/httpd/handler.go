package httpd

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"rvhc/daemon/internal/vmgr"
)

func handler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		host := r.Host

		mu.RLock()
		vmID, exists := hosts[host]
		mu.RUnlock()

		if !exists {
			http.Error(w, "unknown host", http.StatusNotFound)
			return
		}

		portMap, err := vmgr.GetPorts(vmID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		target := fmt.Sprintf("http://127.0.0.1:%d", portMap[80])
		targetURL, err := url.Parse(target)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		proxy := &httputil.ReverseProxy{
			Rewrite: func(pr *httputil.ProxyRequest) {
				pr.SetURL(targetURL)
				pr.SetXForwarded()
			},
		}

		proxy.ServeHTTP(w, r)
	})
}
