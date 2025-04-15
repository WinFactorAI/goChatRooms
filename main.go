// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// 获取 room 参数
	roomID := r.URL.Query().Get("room")
	log.Println("Room ID:", roomID) // 示例：记录参数值

	// 可选：检查参数是否为空
	if roomID == "" {
		http.Error(w, "Room parameter is required", http.StatusBadRequest)
		return
	}

	http.ServeFile(w, r, "home.html")
}

func main() {
	flag.Parse()
	// hub := newHub()
	// go hub.run()

	mainHubs := NewMainHub()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		// serveWs(hub, w, r)
		serveWsRoom(mainHubs, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
