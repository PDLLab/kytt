package main

import (
    "github.com/julienschmidt/httprouter"
    "net/http"
    "encoding/json"
)

type R struct {
    Uid     string      `json:"uid"`
    Response string     `json:"response"`
}

func (this *Server) testing(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    var t R
    t.Uid = "testing"
    t.Response = "Hello World"
    content, _ := json.Marshal(t)
    w.WriteHeader(http.StatusOK)
    w.Write(content)
}

type Server struct {
    mRouter             *httprouter.Router
}

func (this *Server) InitRouter() {
    this.mRouter.POST("/v1/singnup", this.testing)
    this.mRouter.POST("/v1/signin", this.testing)

    this.mRouter.POST("/v1/headlines", this.tetsing)
    this.mRouter.GET("/v1/headlines", this.tetsing)
    this.mRouter.GET("/v1/headlines:headlineId", this.tetsing)
    this.mRouter.POST("/v1/headlines:headlineId/comments", this.tetsing)
    this.mRouter.POST("/v1/headlines:headlineId/likes", this.tetsing)
    
    this.mRouter.POST("/v1/comment/:commentId/likes", this.tetsing)
    this.mRouter.POST("/v1/comment/:commentId/likes", this.tetsing)

    this.mRouter.POST("/v1/users/:userId/following", this.tetsing)
    this.mRouter.POST("/v1/headlines/:userId/followings", this.tetsing)
}

func (this *Server) Init() {
    router := httprouter.New()
    this.mRouter = router
    this.InitRouter()
}

func (this *Server) Start() {
    http.ListenAndServe(":80", this.mRouter)
}
    
func main() {
    var server Server
    server.Init()
    server.Start()
}
