package main

//docker.io/library/testapp

import (
	"fmt"
	"io"
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful/v3"
)

func main() {
	//creating web service
	ws := new(restful.WebService)
	ws.Route(ws.GET("/").To(home))
	//creating GET route /hello
	ws.Route(ws.GET("/hello").To(hello))
	//creating GET route /healthz
	ws.Route(ws.GET("/healthz").To(handleHealthz))
	//creating POST route /break
	ws.Route(ws.POST("/break").To(handleBreak))
	//adding web service
	restful.Add(ws)
	//listen to port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}

var breakOccured bool = false

func home(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "This is home")
}

func handleBreak(req *restful.Request, resp *restful.Response) {
	fmt.Println("Break occured")
	//set break state
	breakOccured = true
}

func handleHealthz(req *restful.Request, resp *restful.Response) {
	//read break state
	if breakOccured {
		fmt.Println("Healthz status 500")
		resp.WriteHeader(500)
	} else {
		fmt.Println("Healthz status 200")
		resp.WriteHeader(200)
	}
}

func hello(req *restful.Request, resp *restful.Response) {
	namePar := req.QueryParameter("name")
	io.WriteString(resp, "Hello "+namePar)
}
