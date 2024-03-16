package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

var (
	Version     = "0.0.1"
	ServerName  = "backend"
	DeafultPort = "4112"
)

func main() {
	// read first argument as a postfix for the server name
	if len(os.Args) > 1 {
		ServerName = ServerName + "-" + os.Args[1]
	}

	http.HandleFunc("/", handler)
	if os.Getenv("SERVERPORT") == "" {
		os.Setenv("SERVERPORT", DeafultPort)
	}
	log.Println("Server listening on port " + os.Getenv("SERVERPORT") + "...")

	log.Fatal(http.ListenAndServe(":"+os.Getenv("SERVERPORT"), nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server Name: %s\n", ServerName)
	fmt.Fprintf(w, "Version: %s\n", Version)
	fmt.Fprintf(w, "Machine IP(s): %s\n", getIPs())
	fmt.Fprintf(w, "Date: %s\n", time.Now().Format("2006-01-02 15:04:05.000000"))
	fmt.Fprintf(w, "URI: %s\n", r.URL.Path)

}

// getIPs gets the unicast IP addresses of the machine
func getIPs() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}
	var IPs string

	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			log.Println(err)
			continue
		}

		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)

			// To filter out loopback addresses and IPv6 if desired
			if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
				IPs += fmt.Sprintf("(Interface Name: %s, IP Address: %s)\n", iface.Name, ipNet.IP.String())
			}
		}
	}
	return IPs
}
