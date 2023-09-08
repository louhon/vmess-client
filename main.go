package main

import (
	"log"

	"github.com/louhon/vmess.cli/vmesshub"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Ltime | log.Lshortfile)
	socks5ListenUri := "127.0.0.1:9000"

	log.Println("Local Socks5 :", socks5ListenUri)
	uuid := "8f38b509-14fc-47d6-d357-6e2483e987e5"
	remoteHost := "38.6.175.85:9099"
	securityType := "none"
	alterId := 0

	vh, err := vmesshub.CreateVmessHub(socks5ListenUri, remoteHost, uuid, securityType, alterId)

	if err != nil {
		log.Println(err)
		return
	}

	vh.StartSocks5Listen()
}
