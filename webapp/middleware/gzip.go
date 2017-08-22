package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

type GzipMiddleware struct {
	Next http.Handler // a pointer to the next handler in the chain
}

func (gm *GzipMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) { //responseWriter is an interface
	if gm.Next == nil {
		gm.Next = http.DefaultServeMux //if last piece of middleware populate w/DefaultServeMux
	}

	encodings := r.Header.Get("Accept-Encoding") //look for header
	if !strings.Contains(encodings, "gzip") {    //looks in string to see if it contains gzip
		gm.Next.ServeHTTP(w, r) //if it doesn't pass it on to the next handler
		return
	}

	w.Header().Add("Content-Encoding", "gzip") //write to the headers before we write to the body
	gzipwriter := gzip.NewWriter(w)            //writer that wraps
	defer gzipwriter.Close()
	grw := gzipResponseWriter{
		ResponseWriter: w,          //pass in writer i recieved
		Writer:         gzipwriter, //created above
	}
	gm.Next.ServeHTTP(grw, r)
}

type gzipResponseWriter struct { //not expoerted (lowercase)
	http.ResponseWriter
	io.Writer
}

func (grw gzipResponseWriter) Write(data []byte) (int, error) { //override write method
	return grw.Writer.Write(data)
}
