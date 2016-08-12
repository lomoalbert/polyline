package polyline

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

type PolyLine struct {
	Image draw.Image
	Map map[image.Point]color.Color
}

type PointFoalt64 struct {
	X float64
	Y float64
}

func NewPolyLine(img image.Image)*PolyLine{
	newimage := image.NewRGBA(img.Bounds())
	draw.Draw(newimage,img.Bounds(),img,image.ZP,draw.Src)
	polyline := new(PolyLine)
	polyline.Image = newimage
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

	var LineImage = image.NewRGBA(img.Image.Bounds())
	for point, pointcolor := range img.Map{
		orgcolor := img.Image.At(point.X,point.Y)
		or,og,ob,oa:= orgcolor.RGBA()
		r,g,b,a := pointcolor.RGBA()
		nr := (r*a>>8+or*(255-a>>8))>>16
		ng := (g*a>>8+og*(255-a>>8))>>16
		nb := (b*a>>8+ob*(255-a>>8))>>16
		//fmt.Println("-",r,g,b,a)
		//fmt.Println("+",or,og,ob,oa)
		//fmt.Println("=",uint8(nr),uint8(ng),uint8(nb),MinUint32(255,a+oa),uint8(MinUint32(255,a+oa)))
		nowcolor := color.RGBA{uint8(nr),uint8(ng),uint8(nb),uint8(MinUint32(255,a+oa))}
		LineImage.Set(point.X,point.Y, nowcolor)
	}
	draw.Draw(img.Image,img.Image.Bounds(),LineImage,image.ZP,draw.Over)
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
	for x:= point.X-halfwidth-border;x <= point.X+halfwidth+border;x=x+0.1{
		for y:= point.Y-halfwidth-border;y <= point.Y+halfwidth+border;y=y+0.1{
			if ((x-point.X)*(x-point.X)+(y-point.Y)*(y-point.Y))>(halfwidth+border)*(halfwidth+border){
				continue
			}
			distance := (x-point.X)*(x-point.X)+(y-point.Y)*(y-point.Y)
			maxdistance := (halfwidth+border)*(halfwidth+border)
			mindistance := halfwidth*halfwidth
			pointa := uint8(float64(a>>8)*(maxdistance-distance)/(maxdistance-mindistance))
			img.AddPoint(image.Point{int(x),int(y)},color.RGBA{uint8(r),uint8(g),uint8(b),uint8(pointa)})
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

func isIn(this float64,start,end int)bool{

	if (float64(start) <= this && this <= float64(end)) || (float64(start) >= this && this >= float64(end)){
		return true
	}
	return false
}