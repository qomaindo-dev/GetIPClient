package main

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gorilla/mux"
)

func last(s string, b byte) int {
	ip := len(s)
	for ip--; ip >= 0; ip-- {
		if s[ip] == b {
			break
		}
	}
	return ip
}

func splitZone(s string) (host, zone string) {
	if i := last(s, '%'); i > 0 {
		host, zone = s[:i], s[i+1:]
		return
	}
	host = s
	return
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****************************")

	// GET Server IP
	adr := r.Context().Value(http.LocalAddrContextKey)
	serverAddr, serverPort, err := net.SplitHostPort(
		fmt.Sprintf("%v", adr))
	if err != nil {
		fmt.Println(err)
		return
	}
	serverIP, serverZone := splitZone(serverAddr)
	fmt.Println("server ip:", serverIP, "port:", serverPort)
	fmt.Println("zone:", serverZone)
	fmt.Println(net.JoinHostPort(serverIP, serverPort))

	fmt.Println("*****************************")

	// GET Client IP
	clientAddr, clientPort, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	clientIP, clientZone := splitZone(clientAddr)
	fmt.Println("Client IP:", clientIP, "Port:", clientPort)
	fmt.Println("-------------------------------")
	fmt.Println("Zone:", clientZone)
	fmt.Println(net.JoinHostPort(clientIP, clientPort))
}

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", handleMain).Methods(http.MethodGet)

	err := http.ListenAndServe(`:101010`, r)
	if err != nil {
		fmt.Println(err)
	}
}
