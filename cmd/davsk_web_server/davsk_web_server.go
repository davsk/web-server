// /////////////////////////////////////////////////////////////
// 'davsk_web_srv.go'                                          /
//                                                             /
// Copyright (c) 2018 Davsk℠. All Rights Reserved.             /
// Use of this source code is governed by an ISC License (ISC) /
// that can be found in the LICENSE file.                      /
//                                                             /
// by David Skinner                                            /
// on July 20, 2018                                            /
// for Davsk℠ Universe 4.0 project gbase                       /
//                                                             /
// /////////////////////////////////////////////////////////////

// Command davsk_web_srv is an http server.
package main

import (
	"davsk.net/gbase/pkg/nothing"
	"davsk.net/web-server/pkg/webserver"
	"fmt"
	"github.com/goinggo/tracelog"
	"runtime"
)

func main() {
	tracelog.Start(tracelog.LevelTrace)
	fmt.Println("davsk_web_server v0.1.1-Alpha ", runtime.GOOS,
		runtime.GOARCH, runtime.Compiler, runtime.Version(),
		"CPU", runtime.NumCPU(), ".")
	fmt.Println("Part of the Davsk℠ Universe 4.0 App Dev Suite.")
	fmt.Println("Copyright (c) 2018 Davsk℠. All Rights Reserved.")
	fmt.Println("Released under the MIT license.")
	fmt.Println("davsk_web_server is charityware. Please buy me a beer.")
	fmt.Println("Browse to:  http://localhost:8080 .")
	fmt.Println("Press Ctrl-C to Stop server.")
	fmt.Println()

	webserver.MustStart()
	nothing.MustDo()
	tracelog.Stop()
}
