package methods

import (
	"fmt"
	"io"
	"strings"
	"image"
	"image/color"
)

type MyImage struct {
	Height int
	Width int
}


func (img MyImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (img MyImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.Height, img.Width)
}

func (img MyImage) At(x, y int) color.Color {
	c := uint8(x ^ y)
	return color.RGBA{c, c, 255, 255}
}

// func (T) Read(b []byte) (n int, err error)
func ReaderErrorTest () {
	r := strings.NewReader("Hello, Reader")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n ,err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

func Images() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}




//func ShowImage(m image.Image) {
//	var buf bytes.Buffer
//	err := png.Encode(&buf, m)
//	if err != nil {
//		panic(err)
//	}
//	enc := base64.StdEncoding.EncodeToString(buf.Bytes())
//	fmt.Println("IMAGE:" + enc)
//}

