package polyline

import (
	"image"
	"testing"
	"os"
	"golang.org/x/image/colornames"
	"github.com/llgcode/draw2d/draw2dimg"
	"image/color"
	"image/draw"
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

func TestMinUint32(t *testing.T) {
	fi,err := os.Open("base.png")
	if err!=nil{
		return
	}
	rgbaimage,_,err := image.Decode(fi)
	if err!=nil{
		return
	}

	var LineImage = image.NewRGBA(rgbaimage.Bounds())
	draw.Draw(LineImage,rgbaimage.Bounds(),rgbaimage,image.ZP,draw.Src)
	gc := draw2dimg.NewGraphicContext(LineImage)
	gc.SetFillColor(color.RGBA{0x00, 0x00, 0x00, 0x00})
	gc.SetStrokeColor(color.RGBA{0xf1, 0x66, 0x00, 0xff})
	//gc.SetStrokeColor(color.RGBA{0xff, 0x00, 0x00, 0xff})
	gc.SetLineWidth(stroke)

	lastpoint := points[0]
	for _, point := range points[1:] {
			gc.MoveTo(float64(lastpoint.X),float64(lastpoint.Y)) // should always be called first for a new path
			gc.LineTo(float64(point.X),float64(point.Y))
			lastpoint = point
	}
	gc.Close()
	gc.FillStroke()

	draw2dimg.SaveToPngFile("./test1.png", LineImage)
}