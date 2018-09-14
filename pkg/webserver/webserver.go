// 'webserver.go'

// by David Skinner
// on September 12, 2018
// for Davsk℠ Universe 4.0 project web-server

// Copyright © 2018 Davsk℠. All rights reserved.
// Use of this source code is governed by an ISC License (ISC)
// that can be found in the LICENSE file.

// Simple Web Server
package webserver

import (
	"github.com/goinggo/tracelog"
	"net/http"
)

// const Title for tracelog
const traceTitle = "webserver"

func Start() (err error) {
	// const FunctionName for tracelog.
	const traceFunctionName = "Start"
	tracelog.Started(traceTitle, traceFunctionName)

	http.Handle("/", http.FileServer(http.Dir("./")))
	http.HandleFunc("/ping", ping)
	err = http.ListenAndServe(":8080", nil)

	// Completed.
	if err == nil {
		tracelog.Completed(traceTitle, traceFunctionName)
	} else {
		tracelog.CompletedError(err, traceTitle, traceFunctionName)
	}

	return err
}

func MustStart() {
	if err := Start(); err != nil {
		panic(err)
	}
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
