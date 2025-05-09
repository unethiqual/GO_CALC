package orchestrator

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"sync"

	"github.com/unethiqual/GO_CALC/pkg/database"
)

const port = ":8080"

type Orchestrator struct {
}

func New() *Orchestrator {
	return &Orchestrator{}
}

var (
	base    = database.New()
	mu      sync.Mutex
	exprKey = contextKey{"expression"}
)

type contextKey struct {
	name string
}

type Expr struct {
	ID   int
	Expr *expression
}

type Request struct {
	Expression string `json:"expression"`
}

type Error struct {
	Res string `json:"error"`
}

type RespID struct {
	Id int `json:"id"`
}

func errorResponse(w http.ResponseWriter, err string, statusCode int) {
	w.WriteHeader(statusCode)
	e := Error{Res: err}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(e)
}

func checkId(id string) bool {
	pattern := "^[0-9]+$"
	r := regexp.MustCompile(pattern)
	return r.MatchString(id)
}

func (o *Orchestrator) Run() {
	StartManager()
	go runGRPC()

	mux := http.NewServeMux()

	expr := http.HandlerFunc(ExpressionHandler)
	GetData := http.HandlerFunc(GetDataHandler)

	mux.Handle("/api/v1/calculate", logsMiddleware(databaseMiddleware(expr)))
	mux.Handle("/api/v1/expressions/", logsMiddleware(GetData))

	log.Printf("Starting sevrer on port %s", port)
	log.Fatal(http.ListenAndServe(port, mux))

}
