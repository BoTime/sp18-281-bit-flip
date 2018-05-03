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

    for i, iface := range list {
        fmt.Printf("%d name=%s %v\n", i, iface.Name, iface)
        if iface.Name == "eth0" {
            addrs, err := iface.Addrs()
            if err != nil {
                panic(err)
            }
            for j, ip_address := range addrs {
                fmt.Printf(" %d %v\n", j, ip_address.String())
                if j == 0 {
                    health_check_msg.WriteString(ip_address.String())
                }
            }
        }
    }

    fmt.Fprintf(w, health_check_msg.String())
}
