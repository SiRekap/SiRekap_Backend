package main

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/aliyun/fc-runtime-go-sdk/fccontext"
)

func main() {
	fc.StartHttp(HandleHttpRequest)
}

// HandleHttpRequest ...
func HandleHttpRequest(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	lc, _ := fccontext.FromContext(ctx)
	fmt.Printf("context: %#v\n", lc)
	fmt.Printf("req.Headers: %#v\n", req.Header)
	fmt.Printf("req.URL: %#v\n", req.URL.String())
	w.Write([]byte("hello, world!"))
	io.WriteString(w, " 你好，世界!\n")
	return nil
}
