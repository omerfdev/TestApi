# Network Devices API

Bu Go uygulaması, bağlı olduğunuz ağdaki cihazları listeleyen ve cihaz sayısını döndüren bir HTTP API'si sağlar.

## Gereksinimler

- Go 1.14 veya daha yeni bir sürüm

## Kurulum

1. Bu repoyu klonlayın:
    ```bash
    git clone https://github.com/username/network-devices-api.git
    cd network-devices-api
    ```

2. Gerekli bağımlılıkları yükleyin (Go modülünü kullanarak):
    ```bash
    go mod tidy
    ```

## Kullanım

1. API sunucusunu başlatın:
    ```bash
    go run main.go
    ```

2. Tarayıcıda veya terminalde bir HTTP GET isteği yaparak cihazları listeleyin:
    ```bash
    curl "http://localhost:8080/devices"
    ```

### Örnek Yanıt

API, bağlı olan cihazları ve cihaz sayısını JSON formatında döndürür. Örnek bir yanıt şu şekilde olacaktır:

```json
{
    "device_count": 5,
    "devices": [
        {"name": "192.168.1.1", "type": "Unknown", "status": "Connected"},
        {"name": "192.168.1.2", "type": "Unknown", "status": "Connected"},
        ...
    ]
}

