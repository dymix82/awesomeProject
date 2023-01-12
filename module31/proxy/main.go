package main

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

var serverCount = 0
var Cfg1 Config
var Con1 *Config

type Config struct {
	Server1 string `yaml:"server1"`
	Server2 string `yaml:"server2"`
	Server3 string `yaml:"server3"`
	Port    string `yaml:"port"`
}

func GetConf2(file string, cnf interface{}) error {
	yamlFile, err := ioutil.ReadFile(file)
	if err == nil {
		err = yaml.Unmarshal(yamlFile, cnf)
	}
	return err
}
func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)
	f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(f)
	if err := GetConf2("proxy_conf.yml", &Cfg1); err != nil {
		log.Panicln(err)
	}
	Con1 = &Cfg1
}

func main() {
	http.HandleFunc("/", loadBalacer)
	log.Fatal(http.ListenAndServe(":"+Con1.Port, nil))
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
	var servers = []string{Con1.Server1, Con1.Server2, Con1.Server3}
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
