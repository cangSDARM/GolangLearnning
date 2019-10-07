package web

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// V1 ---------------------------------------------------------------------------
func webV1() {
	//设置路由
	http.HandleFunc("/", helloWorld) //HandleFunc是一个type, 但经过类型转换了

	//监听端口
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err) //打印错误
	}
}
func helloWorld(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world, this is version 1.")
}

// V2 ---------------------------------------------------------------------------
func webV2() {
	mux := http.NewServeMux()

	mux.Handle("/", &muxHandle{})
	mux.HandleFunc("/hello", helloWorld2)

	//简易静态文件服务器. 不包含在版本区别中
	wd, err := os.Getwd() //获取当前工作目录 work direction
	if err != nil {
		log.Fatal(err)
	}
	mux.Handle("/static/", http.StringPrefix("/static/", http.FileServer(http.Dir(wd))))
	//---------------------------------------------------

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err) //打印错误
	}
}

type muxHandle struct{}

func (*muxHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "URL: "+r.URL.String())
}
func helloWorld2(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world, this is version 2.")
}

// V3 ---------------------------------------------------------------------------
var mux map[string]func(http.ResponseWriter, *http.Request)

func webV3() {
	server := &http.Server{
		Addr:        ":8080",
		Handler:     &muxHandle2{},
		ReadTimeout: 5 * time.Second,
	}

	mux = make([string]func(http.ResponseWriter, *http.Request))
	mux["/"] = helloWorld3

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

type muxHandle2 struct{}

func (*muxHandle2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, or := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}

	io.WriteString(w, "URL: "+r.URL.String())
}

func helloWorld3(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world, this is version 2.")
}
