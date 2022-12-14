package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

var serverCount = 0

// Константы для серверов бэкэнда, понимаю что хардкодить адреса плохо,
// правельнее считывать настройки из файла, но работу с файлом показал в функционале логирования.
const (
	SERVER1 = "http://localhost:8080"
	SERVER2 = "http://192.168.1.82:8080"
	SERVER3 = "http://192.168.1.81:8080"
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

func loadBalacer(res http.ResponseWriter, req *http.Request) {
	url := getProxyURL()
	ip := ReadUserIP(req)
	logRequestPayload(url, ip)
	serveReverseProxy(url, res, req)
}
func getProxyURL() string {
	var servers = []string{SERVER1, SERVER2, SERVER3}
	server := servers[serverCount]
	serverCount++
	if serverCount >= len(servers) {
		serverCount = 0
	}
	return server
}

func logRequestPayload(proxyURL, userIP string) {
	log.Infof("proxy_url: %s, userIP: %s", proxyURL, userIP)
}

func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	url, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ServeHTTP(res, req)
}
