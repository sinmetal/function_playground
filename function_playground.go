package function_playground

import (
	"fmt"
	"net/http"
	"time"
)

func HelloFunction(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Printf("%s:%v\n", k, v)
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Cache-Control", "public, max-age=300, s-maxage=600")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hello Cloud Function %s", time.Now().String())))
}
