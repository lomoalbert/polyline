package polyline

import (
	"image"
	"testing"
	"os"
	"golang.org/x/image/colornames"
)

var points = []image.Point{image.Pt(15, 15), image.Pt(385, 15),image.Pt(15, 300),image.Pt(15, 15), image.Pt(50, 385)}
var stroke = 1.0
var col = colornames.Orange

func TestAddPolyLine(t *testing.T) {
	fi,err := os.Open("base.png")
	if err!=nil{
		return
	}
	rgbaimage,_,err := image.Decode(fi)
	if err!=nil{
		return
	}
	polyline := NewPolyLine(rgbaimage)
	polyline.AddPolyLine(points, col, stroke)
	polyline.Draw()
	polyline.SaveToPngFile("./test.png")
}
