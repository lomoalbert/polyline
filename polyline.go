package polyline

import (
	"image"
	"image/color"
	"image/draw"
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
	for point,color := range img.Map{
		img.Image.Set(point.X,point.Y,color)
	}
}

//AddPolyLine draws a line between (start.X, start.Y) and (end.X, end.Y)
func (img *PolyLine)AddLine(start, end image.Point, linecolor color.Color, width float64) {
	point := start
	for {
		img.AddaroundPoint(point,linecolor,width)
		if point.X == end.X && point.Y == end.Y{
			break
		}
		if abs(start.X-end.X) >= abs(start.Y-end.Y){
			point.Y += (end.Y-start.Y)/abs(end.Y-start.Y)
			point.X += int(float64(end.X-start.X)/float64(end.Y-start.Y)*float64(point.Y))
		}else{
			point.X += (end.X-start.X)/abs(end.X-start.X)
			point.Y += int(float64(end.Y-start.Y)/float64(end.X-start.X)*float64(point.X))
		}
	}
}

func (img *PolyLine)AddaroundPoint(point image.Point,pointcolor color.Color,width float64){
	halfwidth := width/2
	r,g,b,_ := pointcolor.RGBA()
	for x:= point.X-int(halfwidth)-1;x <= point.X+int(halfwidth)+1;x++{
		for y:= point.Y-int(halfwidth)-1;y <= point.Y+int(halfwidth)+1;y++{
			var a uint8= 255
			if float64(abs(x-point.X))>halfwidth{
				a = a/2
			}
			if float64(abs(y-point.Y))>halfwidth{
				a = a/2
			}
			var ptcolor = color.RGBA{uint8(r),uint8(g),uint8(b),a}
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}