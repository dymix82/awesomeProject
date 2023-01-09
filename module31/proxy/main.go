package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

var serverCount = 0

// These constant is used to define all backend servers
const (
	SERVER1 = "http://localhost:8080"
	SERVER2 = "http://localhost:8080"
	SERVER3 = "http://localhost:8080"
	PORT    = "1338"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)
	f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(f)
}

func main() {
	// start server
	http.HandleFunc("/", loadBalacer)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
func ReadUserIP(req *http.Request) string {
	IPAddress := req.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = req.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = req.RemoteAddr
	}
	return IPAddress
}

// Given a request send it to the appropriate url
func loadBalacer(res http.ResponseWriter, req *http.Request) {
	// Get address of one backend server on which we forward request
	url := getProxyURL()
	// Log the request
	ip := ReadUserIP(req)
	logRequestPayload(url, ip)
	// Forward request to original request
	serveReverseProxy(url, res, req)
}
func getProxyURL() string {
	var servers = []string{SERVER1, SERVER2, SERVER3}
	server := servers[serverCount]
	serverCount++
	// reset the counter and start from the beginning
	if serverCount >= len(servers) {
		serverCount = 0
	}
	return server
}

// Log the redirect url
func logRequestPayload(proxyURL, userIP string) {
	log.Infof("proxy_url: %s, userIP: %s", proxyURL, userIP)
}

// Serve a reverse proxy for a given url
func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	// parse the url
	url, _ := url.Parse(target)
	// create the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(url)
	// Note that ServeHttp is non blocking & uses a go routine under the hood
	proxy.ServeHTTP(res, req)
}
