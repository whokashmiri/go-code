package middleware

import (
	"log"
	"net/http"
	"time"
)

type statusRecorder struct{
    http.ResponseWriter
    status int
    bytes int
}

func (rec *statusRecorder) WriteHeader(code int){
    if rec.status != 0{
    return
    }
    rec.status = code
    rec.ResponseWriter.WriteHeader(code)
}

func(rec *statusRecorder) Write(data []byte (int , error){
if rec.status == 0{
 rec.WriteHeader(http.StatusOK)
 }
n , err:= rec.ResponseWriter.Write(data)
rec.bytes+=n
return e,err
}



func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
		start :=  time.Now()

		recorder := &statusRecorder{
		    ResponseWriter:w,
		    }
		next.ServeHTTP(w, r)

		if recorder.status == 0{
		recorder.status = http.StatusOK
		}

		duration :=  time.Since(start)
			log.Printf(
				  "%s %s status=%d bytes=%d duration=%s",
                   r.Method,
                   r.URL.Path,
                    recorder.status,
                   recorder.bytes,
                    duration,
			)


		},
	)
}