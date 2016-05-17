package main

import (
    "net"
    "net/http"
    "net/http/fcgi"
)

type FastCGIServer struct{}

func (s FastCGIServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
    resp.Write([]byte("Hello World!!"))
    resp.Write([]byte("<form> Your name: <input type=\"text\" name=\"fname\"> </form>"))
    if req.Body != nil {
        yourname := req.FormValue("fname")
        resp.Write([]byte("Hello " + yourname))
    }
}

func main() {
    listener, _ := net.Listen("tcp", "127.0.0.1:9001")
    srv := new(FastCGIServer)
    fcgi.Serve(listener, srv)
}
