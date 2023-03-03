package qrlogo

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"

	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
	"github.com/skip2/go-qrcode"
)

const (
	maxLogoPercentage = 50
)

func GenerateQrLogo(qrContent string, qrSize int, logoUrl string, logoPercentage int) (image.Image, error) {

	if logoPercentage > maxLogoPercentage {
		logoPercentage = maxLogoPercentage
	}

	// Create QR
	qr, err := qrcode.New(qrContent, qrcode.Highest)
	if err != nil {
		return nil, err
	}
	qrImg := qr.Image(qrSize)

	// Download logo
	logo, err := downloadImage(logoUrl)
	if err != nil {
		return nil, err
	}
	logoBounds := logo.Bounds()
	if logoBounds.Dx() != 155 || logoBounds.Dy() != 155 {
		reImg := resizeImage(logo, 155)
		logo = reImg
	}

	// Convert logo into circular shape
	circularLogo := convertImageToCircularShape(logo)

	// Add logo into QR
	logoSize := float64(qrSize) * float64(logoPercentage) / 100
	qrLogo := addLogo(qrImg, circularLogo, int(logoSize))

	return qrLogo, nil
}

func EncodeToJPEG(img image.Image, imgName string) error {
	outFile, err := os.Create(imgName)
	if err != nil {
		return err
	}
	defer outFile.Close()
	jpeg.Encode(outFile, img, &jpeg.Options{Quality: 100})
	return nil
}

func EncodeToPNG(img image.Image, imgName string) error {
	outFile, err := os.Create(imgName)
	if err != nil {
		return err
	}
	defer outFile.Close()
	encoder := png.Encoder{CompressionLevel: png.DefaultCompression}
	err = encoder.Encode(outFile, img)
	return nil
}

func downloadImage(url string) (image.Image, error) {
	res, err := http.Get(url)
	if err != nil || res.StatusCode != 200 {
		return nil, err
	}
	defer res.Body.Close()
	m, _, err := image.Decode(res.Body)
	if err != nil {
		return nil, err
	}
	return m, err
}

func resizeImage(img image.Image, size uint) image.Image {
	newImg := resize.Resize(size, size, img, resize.Lanczos3)
	return newImg
}

func convertImageToCircularShape(img image.Image) *image.RGBA {
	newImg := image.NewRGBA(image.Rect(0, 0, 200, 200))
	dc := gg.NewContextForRGBA(newImg)
	dc.SetColor(color.White)
	dc.DrawCircle(100, 100, 95)
	dc.Fill()
	b := img.Bounds()
	x := 100 - b.Dx()/2
	y := 100 - b.Dy()/2
	draw.Draw(newImg, b.Add(image.Pt(x, y)), img, image.ZP, draw.Over)
	return newImg
}

func addLogo(srcImage image.Image, logo image.Image, size int) image.Image {
	logoImage := resizeImage(logo, uint(size))
	offset := image.Pt((srcImage.Bounds().Dx()-logoImage.Bounds().Dx())/2, (srcImage.Bounds().Dy()-logoImage.Bounds().Dy())/2)
	b := srcImage.Bounds()
	m := image.NewNRGBA(b)
	draw.Draw(m, b, srcImage, image.ZP, draw.Src)
	draw.Draw(m, logoImage.Bounds().Add(offset), logoImage, image.ZP, draw.Over)
	return m
}
