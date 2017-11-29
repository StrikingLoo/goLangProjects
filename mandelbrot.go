package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"runtime"
	"sync"
	"fmt"
)

type tile struct {
	x1, x2, y1, y2 int
}

func mandelbrot(w, h, i int, z float32, seed int64) *image.RGBA {

	
	var wg = new(sync.WaitGroup)
	work := make(chan tile)

	zoom := 1 / z

	m := image.NewRGBA(image.Rect(0, 0, w, h))
	fmt.Println(runtime.NumCPU())
	for t := 0; t < runtime.NumCPU(); t++ {
		go func() {
			for tile := range work {
				for x := tile.x1; x < tile.x2; x++ {
					for y := tile.y1; y < tile.y2; y++ {
						setColor(m, x, y, i, zoom)
					}
				}
				wg.Done()
			}
		}()
	}

	tx := w/2
	ty := h

	wg.Add(tx * ty)

	go func() {
		for x := 0; x < w; x += w / tx {
			for y := 0; y < h; y += h / ty {
				work <- tile{x, x + w/tx, y, y + h/ty}
			}
		}

		close(work)
	}()

	wg.Wait()
	return m

}


func setColor(m *image.RGBA, px, py, maxi int, zoom float32) {

	x0 := zoom * (3.5*float32(px)/float32(m.Bounds().Size().X) - 2.5)
	y0 := zoom * (2*float32(py)/float32(m.Bounds().Size().Y) - 1.0)
	x := float32(0)
	y := float32(0)

	i := 0

	for x*x+y*y < 2*2 && i < maxi {
		xtemp := x*x - y*y + x0
		y = 2*x*y + y0
		x = xtemp

		i++
	}

	m.Set(px, py, color.RGBA{min( uint8(510*i/maxi),255),0,0,255})
}
func min(x,y uint8) uint8{
	if x<y {return x}
	return y

}
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	m := mandelbrot(10000, 5000, 3000, 1.0, 50)

	w, _ := os.Create("mandelbrot.png")
	defer w.Close()
	png.Encode(w, m)
}