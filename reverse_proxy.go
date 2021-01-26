package main

import(
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var count = 0

const(
	ser1 = "https://www.google.com/"
	ser2 = "https://www.facebook.com/"
	ser3 = "https://www.youtube.com/"
)

func serveReverseProxy(link string, res http.ResponseWriter, req *http.Request){

	// parse url
	url,_ := url.Parse(link)

	// create the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(url)

	// serveHttp is non blocking and uses a go routine under the hood
	proxy.ServeHTTP(res, req)
}

// print the server 
func logRequestPayload(url string)  {
	log.Printf("Proxy URL : %s\n", url)
}

// choose the server
func getProxyURL() string{
	var servers = [] string {ser1,ser2,ser3}

	server := servers[count]
	count++

	if count >= len(servers){
		count = 0
	}
	return server
}

func handleRequest(res http.ResponseWriter, req *http.Request){
	
	url := getProxyURL()
	logRequestPayload(url)
	serveReverseProxy(url,res,req)
}

func main(){
	http.HandleFunc("/",handleRequest)
	log.Fatal(http.ListenAndServe(":8080",nil))
}