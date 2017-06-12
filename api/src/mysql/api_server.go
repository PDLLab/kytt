package main

import (
    "github.com/julienschmidt/httprouter"
    "net/http"
    "encoding/json"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
    "time"
    "os"
    "io"
    "io/ioutil"
)

const (
    // username:password@protocol(address)/dbname?param=value 
    cConnectString = "root:kytt207@tcp(120.24.177.49:3306)/kaoyantoutiao?charset=utf8"
    cMaxConnectionCount = 128
    cLoggerFile = "server.log"
    cPostUserHeadlines = "INSERT INTO kytt_user_headline(user_id, user_nickname, title, content, post_date, like_count, comment_count, forward_count, view_count, tag, title_image, summary) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)" 
    cPostOfficialHeadlines = "INSERT INTO kytt_official_headline(user_id, user_nickname, title, content, post_date, like_count, comment_count, forward_count, view_count, tag, title_image, summary) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)" 
    cPostUsers = "INSERT INTO kytt_user(nickname, telephone, email, type, signup_date, last_signin_date, active_time, is_auth, user_state, follower_count, following_count, answer_count, headline_count) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
    cPostQuestions = "INSERT INTO kytt_user_question() VALUES()"
    cPostAnswers = "INSERT INTO kytt_user_answer() VALUES()"
)


var Logger *log.Logger

type Server struct {
    mRouter             *httprouter.Router
    mDB                 *sql.DB
    mLogFile            *os.File
    mStmp               Stmp
}

type RecommendData struct {
    Id              int                 `json:"id"`
    UserId          int                 `json:"userId"`
    UserNickname    string              `json:"userNickname"`
    Title           string              `json:"title"`
    PostDate        string              `json:"postDate"`
    LikeCount       int                 `json:"likeCount"`
    CommentCount    int                 `json:"commentCount"`
    ForwardCount    int                 `json:"forwardCount"`
    ViewCount       int                 `json:"ViewCount"`
    Tag             string              `json:"tag"`
    Summary         string              `json:"summary"`
    QuestionId      int                 `json:"questionId"`
}

type Recommend struct {
    Type            string              `json:"type"`
    IsOfficial      bool                `json:"isOfficial"`
    Data            RecommendData       `json:"data"`
}

type Content struct {
    Content         string              `json:"content"`
}

type BaseHeadline struct {
    Id              uint                `json:"id"`
    UserId          uint                `json:"userId"`
    UserNickname    string              `json:"userNickname"`
    Title           string              `json:"title"`
    Content         string              `json:"content"`
    PostDate        string              `json:"postDate"`
    LikeCount       int                 `json:"likeCount"`
    CommentCount    int                 `json:"commentCount"`
    ForwardCount    int                 `json:"forwardCount"`
    ViewCount       int                 `json:"viewCount"`
    Tag             string              `json:"tag"`
    TitleImage      string              `json:"titleImage"`
    Summary         string              `json:"summary"`
}

type Headline struct {
    BaseHeadline
    IsOfficial      bool                `json:"isOfficial"`
}

type User struct {
    Nickname        string              `json:"nickname"`    
    Realname        string              `json:"realname"`
    Gender          uint8               `json:"gender"`
    Birthday        string              `json:"birthday"`
    AvaterUrl       string              `json:"avaterUrl"` 
    Telephone       string              `json:"telephone"`
    Email           string              `json:"email"`
    Type            uint8               `json:"type"`
    SignupDate      string              `json:"signupDate"`
    LastSigninDate  string              `json:"lastSigninDate"`
    LastSigninLocation  string          `json:"lastSigninLocation"`
    LastSigninIp    string              `json:"lastSigninIp"`
    ActiveTime      int                 `json:"activeTime"`
    IsAuth          uint8               `json:"isAuth"`
    UserState       uint8               `json:"userState"`
    FollowerCount   int                 `json:followerCount"`
    FollowingCount  int                 `json:followingCount"`
    AnswerCount     int                 `json:answerCount"`
    HeadlineCount   int                 `json:headlineCount"`
}

type Question struct {
    UserId          uint                `json:"userId"`
    UserNickname    string              `json:"userNickname"`
    Title           string              `json:"title"`
    Content         string              `json:"content"`
    PostDate        string              `json:"postDate"`
    LikeCount       int                 `json:"likeCount"`
    CommentCount    int                 `json:"commentCount"`
    ForwardCount    int                 `json:"forwardCount"`
    ViewCount       int                 `json:"ViewCount"`
    Tag             string              `json:"tag"`
}

type Stmp struct {
    PostUserHeadlines       *sql.Stmt
    PostOfficialHeadlines   *sql.Stmt
    PostUsers               *sql.Stmt
    PostUserQuestions       *sql.Stmt
    PostUserAnswers         *sql.Stmt
}

func (this *Server) InitStmp() {
    this.mStmp.PostUserHeadlines = this.StmtPrepare(cPostUserHeadlines)
    this.mStmp.PostOfficialHeadlines = this.StmtPrepare(cPostOfficialHeadlines)
    this.mStmp.PostUserQuestions = this.StmpPrepare(cPostUserQuestions)
    this.mStmp.PostUserAnswers = this.StmpPrepare(cPostUserAnswers)
    this.mStmp.PostUsers = this.StmtPrepare(cPostUsers)
}

func (this *Server) StmtPrepare(operation string) *sql.Stmt {
    stmt, err := this.mDB.Prepare(operation)
    if err != nil {
        Logger.Panic("[StmtPrepare] err = ", err)
    }
    return stmt
}

func GetCurrentTime() string {
    return time.Now().Format("2006-01-02 15:04:05")
}

func (this *Server) InitRouter() {
    this.mRouter.GET("/v1/home/recommend/:userId", this.GetUserRecommend)
    this.mRouter.GET("/v1/headlines/content/:headlineId", this.GetHeadlineContent)
    this.mRouter.GET("/v1/answer/content/:answerId", this.GetAnswerContent)
    this.mRouter.POST("/v1/headlines", this.PostHeadlines)
    this.mRouter.POST("/v1/questions", this.PostQuestions)
    this.mRouter.POST("/v1/answers", this.PostAnswers)

    this.mRouter.POST("/v1/users", this.PostUsers)
    this.mRouter.POST("/v1/signin", this.PostSignin)

    this.mRouter.GET("/v1/headlines", this.GetHeadlines)
//    this.mRouter.GET("/v1/headlines/:headlineId", this.GetHeadline)
//    this.mRouter.POST("/v1/headlines/:headlineId/comments", this.PostHeadlineComments)
//    this.mRouter.POST("/v1/headlines/:headlineId/likes", this.PostHeadlineLikes)
    
    this.mRouter.POST("/v1/comment/:commentId/comments", this.PostCommentComments)
    this.mRouter.POST("/v1/comment/:commentId/likes", this.PostCommentLikes)

    this.mRouter.GET("/v1/users/:userId/followings", this.PostUserFollowings)
 //   this.mRouter.GET("/v1/headlines/:headlineId/followings", this.PostHeadlineFollowings)
}

func (this *Server) GetUserRecommend(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    userId := ps.ByName("userId")
    Logger.Println("[GetUserRecommend] userId = ", userId)
    var recommend []Recommend
    var mockRecommend Recommend
    mockRecommend.Type = "headline"
    mockRecommend.IsOfficial = true
    mockRecommend.Data.Id = 12832401734
    mockRecommend.Data.UserId = 238978232
    mockRecommend.Data.UserNickname = "迷糊的小蘑菇"
    mockRecommend.Data.Title = "以梦为马——选择比努力重要（跨专业没你想象的那么难）"
    mockRecommend.Data.PostDate = "2017-06-07 09:27:23"
    mockRecommend.Data.LikeCount = 306
    mockRecommend.Data.CommentCount = 46
    mockRecommend.Data.ForwardCount = 2
    mockRecommend.Data.ViewCount = 283232
    mockRecommend.Data.Tag = "考研,会计,经验"
    mockRecommend.Data.Summary = "　这篇是写给所有想跨专业但在犹豫的考生的强心剂，也是想考现当代文学的考生的经验帖，其实去年就想把自己艰难的考研之路记录下来，但觉得全是教训，负能量满满，不想提笔正视当时的自己，但如今已经确认考上自己喜欢的学校，喜欢的专业，并且以初试400+的成绩跨入自己只用不到七个月时间准备的专业，所以真心想用自己的亲身经历给想跨专业的战友们打打鸡血，也让希望走上研究生道路的小汪们避免一些弯路。"

    recommend = append(recommend, mockRecommend)
    content, err := json.Marshal(recommend)
    if err != nil {
        Logger.Println("[GetUserRecommend] = err", err)
    }
    
    w.Write(content)
    w.WriteHeader(http.StatusOK)
}

func (this *Server) GetHeadlineContent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    headlineId := ps.ByName("headlineId")
    Logger.Println("[GetHeadlineContent] headlineId = ", headlineId)
    var content Content
    content.Content = `"<p>加油</p>"`
    c, err := json.Marshal(content)
    if err != nil {
        Logger.Println("[GetHeadlineContent] err = ", err)
    }
    w.Write(c)
    w.WriteHeader(http.StatusOK)
}

func (this *Server) GetAnswerContent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    answerId := ps.ByName("answerId")
    Logger.Println("[GetAnswerContent] answerId = ", answerId)
}

func (this *Server) PostQuestion(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}


func (this *Server) PostSignin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func (this *Server) PostUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    content, err := ioutil.ReadAll(r.Body)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        Logger.Println("[PostUsers] err = ", err)
        return
    }
    var user User
    err = json.Unmarshal(content, &user)
    if err != nil {
        Logger.Println("[PostUsers] err = ", err)
        w.WriteHeader(http.StatusBadRequest)
        return
    } 
    user.SignupDate = GetCurrentTime()
    user.LastSigninDate = GetCurrentTime()
    Logger.Println("user = ", user)
    _, err = this.mStmp.PostUsers.Exec(user.Nickname, user.Telephone, user.Email, user.Type, user.SignupDate, user.LastSigninDate, user.ActiveTime, user.IsAuth, user.UserState, user.FollowerCount, user.FollowingCount, user.AnswerCount, user.HeadlineCount)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        Logger.Println("[PostUsers] err = ", err)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

func (this *Server) PostHeadlines(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    content, err := ioutil.ReadAll(r.Body)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        Logger.Println("[PostHeadlines] err = ", err)
        return
    }
    var headline Headline
    err = json.Unmarshal(content, &headline)
    if err != nil {
        Logger.Println("[PostHeadlines] err = ", err)
        w.WriteHeader(http.StatusBadRequest)
        return
    } 
    headline.PostDate = GetCurrentTime()
    headline.Summary = "waiting to coding"
    Logger.Println("headline = ", headline)
    if headline.IsOfficial {
        _, err = this.mStmp.PostOfficialHeadlines.Exec(headline.UserId, headline.UserNickname, headline.Title, headline.Content, headline.PostDate, headline.LikeCount, headline.CommentCount, headline.ForwardCount, headline.ViewCount, headline.Tag, headline.TitleImage, headline.Summary)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            Logger.Println("[PostHeadlines] err = ", err)
            return
        }
    } else {
        _, err = this.mStmp.PostUserHeadlines.Exec(headline.UserId, headline.UserNickname, headline.Title, headline.Content, headline.PostDate, headline.LikeCount, headline.CommentCount, headline.ForwardCount, headline.ViewCount, headline.Tag, headline.TitleImage, headline.Summary)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            Logger.Println("[PostHeadlines] err = ", err)
            return
        }
    }
    w.WriteHeader(http.StatusCreated)
}

func (this *Server) PostQuestions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func (this *Server) PostAnswers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
    Logger = log.New(fileAndStdoutWriter, "[Debug] ", log.LstdFlags | log.Lshortfile)
}

func (this *Server) InitDB() {
    db, err := sql.Open("mysql", cConnectString)
    if err != nil {
        Logger.Panic("[InitDB] err = ", err)
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
    this.InitDB()
    this.InitStmp()
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
