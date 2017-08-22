package middleware

import (
	"context"
	"net/http"
	"time"
)

type TimeoutMiddleware struct {
	Next http.Handler
}

func (tm TimeoutMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if tm.Next == nil { // look to see if next field is nil
		tm.Next = http.DefaultServeMux //if yes uses SSM
	}

	ctx := r.Context()                               //pull current context off of req obj
	ctx, _ = context.WithTimeout(ctx, 3*time.Second) //with to modify context to recieve a signal on its done channel 3 sec after req starts being processed
	r.WithContext(ctx)                               //replace centext on the req with our new one
	ch := make(chan struct{})                        // make channel thats going to recieve a singal if req finishes processing normally
	go func() {
		tm.Next.ServeHTTP(w, r) //call Next handler using goRoutine
		ch <- struct{}{}        //if that returns send signal in to ch channel, will tell us everything processed normally
	}() //envoke routine
	select {
	case <-ch:
		return // evrything processed normally, get singal on ch channel and return out of mw
	case <-ctx.Done():
		w.WriteHeader(http.StatusRequestTimeout) // assume req has timed out, writeHeader back to the req'tr
	}
}
