package chainOfResponsibility

import (
	"fmt"
	"net/http"
	"testing"
)

func Test_ChainOfResponsibility(t *testing.T) {
	r := &Engine{}
	r.Use(LogMiddleware)
	r.Use(AuthMiddleware)

	fmt.Println("web server on :8080 ...")
	http.ListenAndServe(":8080", r)
}
