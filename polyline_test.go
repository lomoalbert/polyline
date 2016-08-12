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
	//rgbaimage := image.NewRGBA(image.Rect(0, 0, 400, 400))
	fi,err := os.Open("base.png")
	if err!=nil{
		return
	}
	rgbaimage,_,err := image.Decode(fi)
	if err!=nil{
		return
	}
	//AddLine(rgbaimage,image.Pt(5, 5), image.Pt(385, 385), colornames.Orange, 1)
	//AddLine(rgbaimage,image.Pt(385, 5), image.Pt(5, 385), colornames.Orange, 1)
	//AddLine(rgbaimage,image.Pt(5, 5), image.Pt(5, 385), colornames.Orange, 1)
	//AddLine(rgbaimage,image.Pt(5, 5), image.Pt(385, 5), colornames.Orange, 1)
	polyline := NewPolyLine(rgbaimage)
	polyline.AddPolyLine(points, col, stroke)
	//polyline.AddPolyLine([]image.Point{image.Pt(20, 20),  image.Pt(385, 385)}, colornames.White, 20)
	polyline.Draw()
	polyline.SaveToPngFile("./test.png")
}
