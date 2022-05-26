package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/delivery-club/qrlogo"
)

var (
	input  = flag.String("i", "assets/logo.jpeg", "Logo to be placed over QR code")
	output = flag.String("o", "assets/qr.png", "Output filename")
	size   = flag.Int("size", 512, "Image size in pixels")
)

const defaultExt = "png"

func main() {
	flag.Usage = Usage
	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
		return
	}

	text := flag.Arg(0)

	file, err := os.Open(*input)
	if err != nil {
		log.Println("Failed to open logo:", err)
		return
	}
	defer file.Close()

	logo, _, err := image.Decode(file)
	if err != nil {
		log.Println("Failed to decode PNG with logo:", err)
		return
	}

	ext := strings.ReplaceAll(filepath.Ext(*output), ".", "")
	if ext == "" {
		ext = defaultExt
	}

	qr, err := qrlogo.Encode(text, logo, *size, ext)
	if err != nil {
		log.Println("Failed to encode QR:", err)
		return
	}

	out, err := os.Create(*output)
	if err != nil {
		log.Println("on create output image:", err)
		return
	}
	defer func() {
		if err := out.Close(); err != nil {
			log.Println("on close result file:", err)
		}
	}()

	if _, err = io.Copy(out, qr); err != nil {
		log.Println("on copy data to destination:", err)
		return
	}

	fmt.Println("Result:", *output)
}

// Usage overloads flag.Usage.
func Usage() {
	fmt.Println("Usage: qrlogo [options] text")
	flag.PrintDefaults()
}
