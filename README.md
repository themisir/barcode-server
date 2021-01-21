# Barcode Server

Minimalist HTTP server application written in Go to generate Barcodes and QR
codes by sending http requests.

```
GET /generate/{mode}/{value}
```

### Supported Modes

* EAN
* Code 39 
* Code 93
* Code 128
* Aztec
* QR

## Example

```
http://localhost:8080/generate/qr/https%3A%2F%2Fthemisir.com
```