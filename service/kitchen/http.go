package main 
import (
	"context"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/mrinalxdev/go-res/services/common/genproto/orders"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()
	
	
}