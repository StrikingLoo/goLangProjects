package main 
import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"strconv"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height = 4096, 4096
	)
	contrast,_:=strconv.ParseInt(os.Args[1], 10,8)
	t,_:= strconv.ParseInt(os.Args[2], 10,8)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		go setLine(ymin,ymax,xmin,xmax,width,height,contrast,t,py,img)
	      }
	png.Encode(os.Stdout, img) 
}
func setLine(ymin,ymax,xmin,xmax,width,height float64,contrast,t int64,py int,  img *image.RGBA) {
	y := getC(ymin,ymax,height,int(py))
	for px := 0; px < int(width); px++ {
	    x := getC(xmin,xmax,width,px)
            z := complex(x, y)
	    img.Set(px, int(py), newton(z,uint8(contrast),int(t)))
	}
}
func getC(min,max,sizeC float64,coor int) float64{
	return float64(coor)/sizeC*(max-min) + min

}

func min(x,y uint8) uint8{
	if x<y {return x}
	return y

}
func newton(z complex128, contrast uint8, t int) color.Color {
	const iterations = 200
	for n := uint8(0); n < iterations; n++ {
	    z-=f(z,t)
	    if cmplx.Abs(f(z,t))<0.01 {
		sf:=min(contrast*n,255)
		return color.RGBA{255-sf,255-sf/2,0,255}
		}
	}
    	return color.RGBA{255,255,0,255}
}
func f(v complex128,t int) complex128 {
	w := v
	for i:=0;i<t-2;i++{
	    w*=v		
	}
	dw:=w*complex(float64(t),0.)
	w*=v
	return (w-1)/dw
}
