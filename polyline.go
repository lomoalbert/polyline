package polyline

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"fmt"
)

type PolyLine struct {
	Image draw.Image
	Map map[image.Point]color.Color
}

func NewPolyLine(img draw.Image)*PolyLine{
	polyline := new(PolyLine)
	polyline.Image = img
	polyline.Map = make(map[image.Point]color.Color)
	return polyline
}

func (img *PolyLine)AddPolyLine(points []image.Point, linecolor color.Color, width float64) {
	if len(points) < 2{
		return
	}
	var startpoint = points[0]
	for _,point := range points[1:]{
		img.AddLine(startpoint,point,linecolor,width)
		startpoint = point
	}
}

func (img *PolyLine)Draw(){
	for point, pointcolor := range img.Map{
		_,_,_,a := pointcolor.RGBA()
		fmt.Println("rgba:A=",a)
		if a == 65535{
			img.Image.Set(point.X,point.Y, pointcolor)
			continue
		}
		orgcolor := img.Image.At(point.X,point.Y)
		or,og,ob,_:= orgcolor.RGBA()
		r,g,b,a := pointcolor.RGBA()
		nr := (r>>8*(65535-a)>>8+or>>8*a>>8)
		ng := (g>>8*(65535-a)>>8+og>>8*a>>8)
		nb := (b>>8*(65535-a)>>8+ob>>8*a>>8)
		nowcolor := color.RGBA{uint8(nr),uint8(ng),uint8(nb),uint8(255)}
		img.Image.Set(point.X,point.Y, nowcolor)
	}
}

//AddPolyLine draws a line between (start.X, start.Y) and (end.X, end.Y)
func (img *PolyLine)AddLine(start, end image.Point, linecolor color.Color, width float64) {
	//fmt.Println("AddLine",start.X,start.Y,end.X,end.Y)
	point := start
	for {
		if !isIn(point.X ,start.X ,end.X) || !isIn(point.Y,start.Y, end.Y){
			break
		}
		img.AddaroundPoint(point,linecolor,width)
		if abs(start.X-end.X) >= abs(start.Y-end.Y){
			point.X += sign(end.X-start.X)
			point.Y =start.Y+int(float64(end.Y-start.Y)/float64(end.X-start.X)*float64(point.X-start.X))
		}else{
			point.Y += sign(end.Y-start.Y)
			point.X =start.X+int(float64(end.X-start.X)/float64(end.Y-start.Y)*float64(point.Y-start.Y))
		}
	}
}

func (img *PolyLine)AddaroundPoint(point image.Point,pointcolor color.Color,width float64){
	//fmt.Println("AddaroundPoint",point.X,point.Y)
	halfwidth := width/2
	r,g,b,a := pointcolor.RGBA()
	//fmt.Println(pointcolor)
	border := 1
	for x:= point.X-int(halfwidth)-border;x <= point.X+int(halfwidth)+border;x++{
		for y:= point.Y-int(halfwidth)-border;y <= point.Y+int(halfwidth)+border;y++{
			var ptcolor color.RGBA
			if ((x-point.X)*(x-point.X)+(y-point.Y)*(y-point.Y))>int((halfwidth+1)*(halfwidth+1)){
				continue
			}else if ((x-point.X)*(x-point.X)+(y-point.Y)*(y-point.Y))>int(halfwidth+1)*int(halfwidth+1){
				ptcolor = color.RGBA{uint8(r>>8/4),uint8(g>>8/4),uint8(b>>8/4),uint8(a>>8/4)}
			}else if ((x-point.X)*(x-point.X)+(y-point.Y)*(y-point.Y))>int((halfwidth)*(halfwidth)){
				ptcolor = color.RGBA{uint8(r>>8/2),uint8(g>>8/2),uint8(b>>8/2),uint8(a>>8/2)}
			}else if ((x-point.X)*(x-point.X)+(y-point.Y)*(y-point.Y))>(int(halfwidth)*int(halfwidth)){
				ptcolor = color.RGBA{uint8(r>>8/4*3),uint8(g>>8/4*3),uint8(b>>8/4*3),uint8(a>>8/4*3)}
				//fmt.Println(ptcolor)
			}else{
				ptcolor = color.RGBA{uint8(r),uint8(g),uint8(b),uint8(a)}
			}
			//fmt.Println("AddPoint",x,y,ptcolor)
			img.AddPoint(image.Pt(x,y),ptcolor)
		}
	}
}

func (img *PolyLine)AddPoint(point image.Point,pointcolor color.Color){
	pt,ok := img.Map[point]
	if ok{
		_,_,_,pta := pt.RGBA()
		_,_,_,pointa := pointcolor.RGBA()
		if pointa<pta{
			return
		}
	}
	img.Map[point]=pointcolor
}

func (img *PolyLine)SaveToPngFile(imagename string){
	fi,err := os.Create(imagename)
	if err != nil{
		panic(err)
	}
	defer func(){
		fi.Close()
	}()
	err = png.Encode(fi, img.Image)
	if err != nil{
		panic(err)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sign(x int)int{
	if x >=0{
		return 1
	}
	return -1
}

func isIn(this,start,end int)bool{
	if (start <= this && this <= end) || (start >= this && this >= end){
		return true
	}
	return false
}