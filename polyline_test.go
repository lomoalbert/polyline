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
	//AddLine(rgbaimage,image.Pt(5, 5), image.Pt(385, 385), colornames.Orange, 1)
	//AddLine(rgbaimage,image.Pt(385, 5), image.Pt(5, 385), colornames.Orange, 1)
	//AddLine(rgbaimage,image.Pt(5, 5), image.Pt(5, 385), colornames.Orange, 1)
	//AddLine(rgbaimage,image.Pt(5, 5), image.Pt(385, 5), colornames.Orange, 1)
	polyline := NewPolyLine(rgbaimage)
	polyline.AddPolyLine([]image.Point{image.Pt(15, 50), image.Pt(385, 15),image.Pt(15, 385),image.Pt(15, 15), image.Pt(385, 385)}, colornames.White, 15)
	//polyline.AddPolyLine([]image.Point{image.Pt(20, 20),  image.Pt(385, 385)}, colornames.White, 20)
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
