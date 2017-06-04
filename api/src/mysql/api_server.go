package main

import (
    "github.com/julienschmidt/httprouter"
    "net/http"
    "encoding/json"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

const (
    // username:password@protocol(address)/dbname?param=value 
    cConnectString = "root:Kytt@207@tcp(localhost:3306)/kytt?charset=utf8"
    cMaxConnectionCount = 128
    cLoggerFile = "server.log"
)

var Logger *log.Logger

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
    mDB                 *sql.DB
    mLoggerFile         *os.File
}

func (this *Server) StmtPrepare(operation string) *sql.Stmt {
    stmt, err := db.Prepare(operation)
    if err != nil {
        Logger.Panic("[StmtPrepare] err = ", err)
    }
    return stmt
}

func (this *Server) InitRouter() {
    this.mRouter.POST("/v1/signup", this.PostSignup)
    this.mRouter.POST("/v1/signin", this.PostSignin)

    this.mRouter.POST("/v1/headlines", this.PostHeadlines)
    this.mRouter.GET("/v1/headlines", this.GetHeadlines)
    this.mRouter.GET("/v1/headlines/:headlineId", this.GetHeadline)
    this.mRouter.POST("/v1/headlines/:headlineId/comments", this.PostHeadlineComments)
    this.mRouter.POST("/v1/headlines/:headlineId/likes", this.PostHeadlineLikes)
    
    this.mRouter.POST("/v1/comment/:commentId/comments", this.PostCommentComments)
    this.mRouter.POST("/v1/comment/:commentId/likes", this.PostCommentLikes)

    this.mRouter.GET("/v1/users/:userId/followings", this.PostUserFollowings)
    this.mRouter.GET("/v1/headlines/:headlineId/followings", this.PostHeadlineFollowings)
}

func (this *Server) PostSignup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func (this *Server) PostSignin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func (this *Server) PostHeadlines(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func (this *Server) GetHeadlines(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func (this *Server) GetHeadline(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func (this *Server) PostHeadlineComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func (this *Server) PostHeadlineLikes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func (this *Server) PostCommentComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func (this *Server) PostCommentLikes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func (this *Server) PostUserFollowings(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func (this *Server) PostHeadlineFollowings(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func (this *Server) InitLogger() {
    logFile, err := os.Create(cLoggerFile)
    this.mLogFile = logFile 
    if err != nil {
        os.Exit(1)
    }
    writers := []io.Writer {
        logFile,
        os.Stdout,
    }
    fileAndStdoutWriter := io.MultiWriter(writers...)
    Logger = log.New(fileAndStdoutWriter, "[Debug] ", logLstdFlags | log.Lshortfile)
}

func (this *Server) InitDB() {
    db, err := sql.Open("mysql", cConnectString)
    if err != nil {
        Logger.Panic("[InitDB]" err = ", err)
    }
    this.mDB = db
    err = this.mDB.Ping()
    if err != nil {
        Logger.Panic("[InitDB] err = ", err)
    }
}

func (this *Server) Init() {
    this.InitLogger()
    Logger.Println("[Init Begin]")
    router := httprouter.New()
    this.mRouter = router
    this.InitRouter()
    Logger.Println("[Init End]")
}

func (this *Server) Start() {
    http.ListenAndServe(":80", this.mRouter)
}
   
func (this *Server) Stop() {
    this.mLogFile.Close()
    this.mDB.Close()
}
func main() {
    var server Server
    server.Init()
    server.Start()
    defer server.Stop()
}
