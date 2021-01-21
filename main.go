package main

import (
	"fmt"
	"image/png"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/aztec"
	"github.com/boombuler/barcode/code128"
	"github.com/boombuler/barcode/code39"
	"github.com/boombuler/barcode/code93"
	"github.com/boombuler/barcode/ean"
	"github.com/boombuler/barcode/qr"

	"github.com/julienschmidt/httprouter"
)

func logError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Index(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "text/plain")
	w.WriteHeader(200)
	fmt.Fprintf(w, "unnamed library to generate barcodes & qrcodes using http requests\n\n"+
		"GET /generate/<mode>/<value>[?scale=<scale>]\n\n"+
		"mode  - barcode mode (one of: ean, code39, code93, code128, aztec, qr)\n"+
		"value - data to encode\n"+
		"scale - output image scale")
}

func Health(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "OK")
}

func Generate(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	query := req.URL.Query()
	value := ps.ByName("value")
	mode := ps.ByName("name")

	var code barcode.Barcode
	var err error

	width := 150
	height := 100

	switch mode {
	case "ean":
		code, err = ean.Encode(value)
		break
	case "code39":
		code, err = code39.Encode(value, true, true)
		break
	case "code93":
		code, err = code93.Encode(value, true, true)
		break
	case "code128":
		code, err = code128.Encode(value)
		break
	case "aztec":
		width = 200
		height = 200
		code, err = aztec.Encode([]byte(value), aztec.DEFAULT_EC_PERCENT, aztec.DEFAULT_LAYERS)
		break
	case "qr":
		width = 200
		height = 200
		code, err = qr.Encode(value, qr.M, qr.Auto)
		break

	default:
		w.WriteHeader(404)
		return
	}

	scaleStr := query.Get("scale")
	scale := 1.0

	if len(scaleStr) > 0 {
		if strings.HasPrefix(scaleStr, "x") {
			scaleStr = scaleStr[1:]
		}

		scale, err = strconv.ParseFloat(scaleStr, 64)
	}

	width = int(math.Round(scale * float64(width)))
	height = int(math.Round(scale * float64(height)))

	if err == nil {
		code, err = barcode.Scale(code, width, height)
	}

	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Header().Add("content-type", "image/png")
	w.WriteHeader(200)

	logError(png.Encode(w, code))
}



func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/health", Health)
	router.GET("/generate/:name/:value", Generate)

	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "8080"
	}
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	logError(http.ListenAndServe(port, router))
}
