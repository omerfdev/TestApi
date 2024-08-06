package main

import (
    "encoding/json"
    "fmt"
    "net"
    "net/http"
)

type Device struct {
    Name   string `json:"name"`
    Type   string `json:"type"`
    Status string `json:"status"`
}

type Response struct {
    DeviceCount int      `json:"device_count"`
    Devices     []Device `json:"devices"`
}

func GetDevices(w http.ResponseWriter, r *http.Request) {
    devices := []Device{}
    interfaces, err := net.Interfaces()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    for _, iface := range interfaces {
        addrs, err := iface.Addrs()
        if err != nil {
            continue
        }
        for _, addr := range addrs {
            devices = append(devices, Device{Name: addr.String(), Type: "Unknown", Status: "Connected"})
        }
    }

    response := Response{
        DeviceCount: len(devices),
        Devices:     devices,
    }

    // JSON olarak cevap verisi hazırlama
    responseData, err := json.Marshal(response)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Content-Type başlığı ayarlama ve JSON yanıtını yazma
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(responseData)
}

func main() {
    http.HandleFunc("/devices", GetDevices)

    // API'yi 8080 portunda başlatma
    fmt.Println("API server listening on port 8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println(err)
    }
}
