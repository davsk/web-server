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
	"github.com/goinggo/tracelog"
)

func main() {
	tracelog.Start(tracelog.LevelTrace)
	webserver.MustStart()
	nothing.MustDo()
	tracelog.Stop()
}
