package polyline

import (
	"github.com/golang/image/colornames"
	"image"
	"testing"
	"image/png"
	"os"
	"fmt"
)

func TestAddPolyLine(t *testing.T) {
	rgbaimage := image.NewRGBA(image.Rect(0, 0, 400, 400))
	//AddLine(rgbaimage,image.Pt(5, 5), image.Pt(395, 395), colornames.Orange, 1)
	//AddLine(rgbaimage,image.Pt(395, 5), image.Pt(5, 395), colornames.Orange, 1)
	//AddLine(rgbaimage,image.Pt(5, 5), image.Pt(5, 395), colornames.Orange, 1)
	//AddLine(rgbaimage,image.Pt(5, 5), image.Pt(395, 5), colornames.Orange, 1)
	polyline := NewPolyLine(rgbaimage)
	polyline.AddPolyLine([]image.Point{image.Pt(5, 5), image.Pt(395, 5),image.Pt(5, 395),image.Pt(5, 5), image.Pt(395, 395)}, colornames.Orange, 1)
	polyline.Draw()
	fi,err := os.Create("./test.png")
	if err != nil{
		fmt.Println(err.Error())
		return
	}else{
		fmt.Println("created")
	}
	defer func(){
		fi.Close()
	}()
	png.Encode(fi, polyline.Image)
}
