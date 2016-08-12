package polyline

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"math"
)

type PolyLine struct {
	Image draw.Image
	LineImage draw.Image
}

type PointFoalt64 struct {
	X float64
	Y float64
}

func NewPolyLine(img image.Image)*PolyLine{
	oldimage := image.NewRGBA(img.Bounds())
	draw.Draw(oldimage,img.Bounds(),img,image.ZP,draw.Src)
	newimage := image.NewRGBA(img.Bounds())
	polyline := new(PolyLine)
	polyline.Image = oldimage
	polyline.LineImage = newimage
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
	draw.Draw(img.Image,img.Image.Bounds(),img.LineImage,image.ZP,draw.Over)
}

func MinUint32(a,b uint32)uint32{
	if a>b{
		return b
	}
	return a
}

//AddPolyLine draws a line between (start.X, start.Y) and (end.X, end.Y)
func (img *PolyLine)AddLine(start, end image.Point, linecolor color.Color, width float64) {
	//fmt.Println("AddLine",start.X,start.Y,end.X,end.Y)
	point := PointFoalt64{float64(start.X),float64(start.Y)}
	for {
		if !isIn(point.X ,start.X ,end.X) || !isIn(point.Y,start.Y, end.Y){
			break
		}
		img.AddaroundPoint(PointFoalt64{point.X,point.Y},linecolor,width)
		if abs(start.X-end.X) >= abs(start.Y-end.Y){
			point.X += float64(sign(end.X-start.X))
			point.Y =float64(start.Y)+float64(end.Y-start.Y)/float64(end.X-start.X)*(point.X-float64(start.X))
		}else{
			point.Y += float64(sign(end.Y-start.Y))
			point.X =float64(start.X)+float64(end.X-start.X)/float64(end.Y-start.Y)*(point.Y-float64(start.Y))
		}
	}
}

func (img *PolyLine)AddaroundPoint(point PointFoalt64,pointcolor color.Color,width float64){
	//fmt.Println("AddaroundPoint",point.X,point.Y)
	halfwidth := width/2
	r,g,b,a := pointcolor.RGBA()
	//fmt.Println(pointcolor)
	border := 1.0
	for x:= point.X-halfwidth-border;x <= point.X+halfwidth+border;x=x+0.3{
		for y:= point.Y-halfwidth-border;y <= point.Y+halfwidth+border;y=y+0.3{
			if ((x-point.X)*(x-point.X)+(y-point.Y)*(y-point.Y))>(halfwidth+border)*(halfwidth+border){
				continue
			}
			distance := (x-point.X)*(x-point.X)+(y-point.Y)*(y-point.Y)
			maxdistance := (halfwidth+border)*(halfwidth+border)
			mindistance := halfwidth*halfwidth
			distance = math.Max(distance,mindistance)
			pointa := uint8(float64(a>>8)*(maxdistance-distance)/(maxdistance-mindistance))
			img.AddPoint(image.Point{int(x),int(y)},color.RGBA{uint8(r),uint8(g),uint8(b),uint8(pointa)})
		}
	}
}


func (img *PolyLine)AddPoint(point image.Point,pointcolor color.Color){
	col :=img.LineImage.At(point.X,point.Y)
	_,_,_,pta := col.RGBA()
	_,_,_,pointa := pointcolor.RGBA()
	if pointa<pta{
		return
	}
	img.LineImage.Set(point.X,point.Y,pointcolor)
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

func isIn(this float64,start,end int)bool{

	if (float64(start) <= this && this <= float64(end)) || (float64(start) >= this && this >= float64(end)){
		return true
	}
	return false
}