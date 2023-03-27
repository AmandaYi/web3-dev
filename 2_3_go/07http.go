package main

import "net/http"

func main() {
	/**
	HTTP协议

	请求行 GET / HTTP/1.1
	请求头 k:v，如Content-Type:
	空行 \r\n ，不管是什么系统都是这个\r\n
	请求包体:携带数据
	*/

	http.HandleFunc("/1.jpg", func(w http.ResponseWriter, r *http.Request) {

	})
	//The handler is typically nil, in which case the DefaultServeMux is used.
	http.ListenAndServe("127.0.0.1:9000", nil)
}
