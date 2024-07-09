package filter

import (
	"image/jpeg"
	"os"

	"github.com/disintegration/imaging"
)

type Filter interface {
	//method
	Process(srcPath, dstPath string) error
}

type Grayscale struct {}
 
func (g Grayscale) Process(srcPath, dstPath string) error {
	// OPEN IMAGE
	src, err := imaging.Open(srcPath)
	if err != nil {
		return err
	}

	// apply filter into image
	dst := imaging.Grayscale(src)

	// creat destination file
	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	opts := jpeg.Options{Quality: 90}
	return jpeg.Encode(dstFile, dst, &opts)
}

type Blur struct {}

 
func (b Blur) Process(srcPath, dstPath string) error {
	// OPEN IMAGE
	src, err := imaging.Open(srcPath)
	if err != nil {
		return err
	}
	
	// apply filter into image
	dst := imaging.Blur(src, 3.5)

	// creat destination file
	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	opts := jpeg.Options{Quality: 90}
	return jpeg.Encode(dstFile, dst, &opts)
}