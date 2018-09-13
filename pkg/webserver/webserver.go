// 'webserver.go'

// by David Skinner
// on September 12, 2018
// for Davsk℠ Universe 4.0 project web-server

// Copyright © 2018 Davsk℠. All rights reserved.
// Use of this source code is governed by an ISC License (ISC)
// that can be found in the LICENSE file.

// Simple Web Server
package webserver

import "net/http"

func Start() (err error) {
	http.Handle("/", http.FileServer(http.Dir("./")))
	http.HandleFunc("/ping", ping)
	err = http.ListenAndServe(":8080", nil)
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
