package qrlogo

import (
	"log"
	"testing"
)

func TestQrLogo(t *testing.T) {

	qrContent := "https://github.com/nsrvel"
	qrSize := 2000
	logoUrl := "https://github.blog/wp-content/uploads/2013/04/0cf7be70-a5e3-11e2-8943-6ac7a953f26d.jpg?resize=1234%2C631"
	logoPercentage := 30

	jpegName := "example/qrsample.jpeg"
	pngName := "example/qrsample.png"

	qrLogo, err := GenerateQrLogo(qrContent, qrSize, logoUrl, logoPercentage)
	if err != nil {
		log.Fatal(err)
	}

	err = EncodeToJPEG(qrLogo, jpegName)
	if err != nil {
		log.Fatal(err)
	}

	err = EncodeToPNG(qrLogo, pngName)
	if err != nil {
		log.Fatal(err)
	}
}
