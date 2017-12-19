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
		width, height = 1024, 1024
	)
	clarity,_:=strconv.ParseInt(os.Args[1], 10,8)
	t,_:= strconv.ParseInt(os.Args[2], 10,8)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
		    x := float64(px)/width*(xmax-xmin) + xmin
	            z := complex(x, y)
		    img.Set(px, py, newton(z,uint8(clarity),int(t)))
		}
	      }
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}
func min(x,y uint8) uint8{
	if x<y {return x}
	return y

}
func newton(z complex128, clarity uint8, t int) color.Color {
	const iterations = 200
	const contrast = 15

	for n := uint8(0); n < iterations; n++ {
	    z-=f(z,t)
	    if cmplx.Abs(f(z,t))<0.001 {
		sf:=min(clarity*n,255)
		return color.RGBA{sf,sf,0,255}
		}
	}
    	return color.Black
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
