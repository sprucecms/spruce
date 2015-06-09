package spruce

import (
	"bytes"
	"strconv"
)

// ax7fds/c4x3/s250x250.jpg
// <hash>/c<width_ratio>x<height_ratio>/s<width>x<height>.<format>
//
type Image struct {
	Key      string
	OrigBase string
	OrigExte string
	Width    int
	Height   int
	Crops    []CroppedImage
}

//
//
type CroppedImage struct {
	Key    string
	Width  int
	Height int
	X      int
	Y      int
}

//
//
type SizedImage struct {
	Key    string
	Width  int
	Height int
}

func (image SizedImage) GetPath() string {
	buffer := bytes.NewBufferString(image.Key)
	buffer.WriteString("/s")
	buffer.WriteString(strconv.Itoa(image.Width))
	buffer.WriteString("x")
	buffer.WriteString(strconv.Itoa(image.Height))
	return buffer.String()
}
