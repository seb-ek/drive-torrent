package main

import (
	"github.com/ppkavinda/drive-torrent/server"
	"fmt"
)

func main() {

	port := os.Getenv("Port");

	s := server.Server{
		Port:       port,
		ConfigPath: "drive-torrent.json",
	}
	fmt.Printf("Server is up at port %v and Listning for a connection!!\n",port)
	s.Start()
}
