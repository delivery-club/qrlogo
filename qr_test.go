package qrlogo

import (
	"crypto/md5"
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"os"
	"testing"
)

func TestEncodePNG(t *testing.T) {
	testEncode(t, "png")
}

func TestEncodeJPEG(t *testing.T) {
	testEncode(t, "jpeg")
}

func testEncode(t *testing.T, outputFormat string) {
	qr, err := os.Open("./assets/qr." + outputFormat)
	if err != nil {
		t.Fatalf("on open file:%s", err)
	}
	defer qr.Close()

	b, err := ioutil.ReadAll(qr)
	if err != nil {
		t.Fatalf("on read body of existed file:%s", err)
	}
	oldSum := fmt.Sprintf("%x", md5.Sum(b))

	logo, err := os.Open("./assets/logo.jpeg")
	if err != nil {
		t.Fatalf("on open file:%s", err)
	}
	img, err := jpeg.Decode(logo)

	if err != nil {
		t.Fatalf("on decode jpeg:%s", err)
	}

	buf, err := Encode("https://delivery-club.ru", img, 512, outputFormat)
	if err != nil {
		t.Fatalf("on encode: qr-code:%s", err)
	}

	newSum := fmt.Sprintf("%x", md5.Sum(buf.Bytes()))

	if oldSum != newSum {
		t.Fatalf("sums are not equal")
	}
}
