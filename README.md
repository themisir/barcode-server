# Barcode Server

![GitHub](https://img.shields.io/github/license/themisir/barcode-server)
![Docker Pulls](https://img.shields.io/docker/pulls/themisir/barcode-server)
![Docker Image Size (tag)](https://img.shields.io/docker/image-size/themisir/barcode-server/latest)

Minimalist HTTP server application written in Go to generate Barcodes and QR
codes by sending http requests.

**GET** `/generate/{mode}/{size}?data=...`

|Name|Description                                                  |
|:---|:------------------------------------------------------------|
|Mode|Barcode encoding mode                                        |
|Size|Output image size (width and height separated by x character)|

### Supported Modes

* ean
* code39
* code93
* code128
* aztec
* qr

## Installation

You can use [`barcode-scanner`](https://hub.docker.com/r/themisir/barcode-server)
docker image to get started with the library.

```shell
docker run -p 8080:80 themisir/barcode-server
```

## Example

```
http://localhost:8080/generate/qr/300x300?data=https%3A%2F%2Fthemisir.com
```