package handlers

import (
    "net/http"
    "net"
    "bytes"
    "fmt"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)

    var health_check_msg bytes.Buffer
    health_check_msg.WriteString("Starbucks User API V1.0 - Bo - ")

    list, err := net.Interfaces()
    if err != nil {
        panic(err)
    }

    for _, iface := range list {
        // fmt.Printf("%d name=%s %v\n", i, iface.Name, iface)
        if iface.Name == "en0" {
            addrs, err := iface.Addrs()
            if err != nil {
                panic(err)
            }
            for j, ip_address := range addrs {
                fmt.Printf(" %d %v\n", j, ip_address.String())
                if j == 1 {
                    health_check_msg.WriteString(ip_address.String())
                }
            }
        }
    }

    fmt.Fprintf(w, health_check_msg.String())
}
