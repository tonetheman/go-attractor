package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"
)

const Width = 800
const Height = 800

// saw this
// http://paulbourke.net/fractals/clifford/

func dbg_write(board *[800][800]int) {
	outf, _ := os.Create("dbg.ppm")
	fheader := []byte("P3\n")
	binary.Write(outf, binary.LittleEndian, fheader)
	ts := fmt.Sprintf("%d %d\n", Width, Height)
	binary.Write(outf, binary.LittleEndian, []byte(ts))
	binary.Write(outf, binary.LittleEndian, []byte("255\n"))

	for x := 0; x < Width; x++ {
		for y := 0; y < Height; y++ {

			var v = board[x][y]
			//var buf [3]uint8
			Empty := []byte(fmt.Sprintf("255 255 255\n"))
			if v > 0 {
				binary.Write(outf, binary.LittleEndian,
					[]byte(fmt.Sprintf("%d 255 255\n", v)))
				//buf[0] = uint8(255 - v)
				//binary.Write(outf, binary.LittleEndian, buf)
				//binary.Write(outf, binary.LittleEndian, uint8(255-v))
				//binary.Write(outf, binary.LittleEndian, uint8(0))
				//binary.Write(outf, binary.LittleEndian, uint8(0))
			} else {
				binary.Write(outf, binary.LittleEndian, Empty)
				//buf[0] = 255
				//buf[1] = 255
				//buf[2] = 255
				//binary.Write(outf, binary.LittleEndian, buf)
				//binary.Write(outf, binary.LittleEndian, uint8(255))
				//binary.Write(outf, binary.LittleEndian, uint8(255))
				//binary.Write(outf, binary.LittleEndian, uint8(255))
			}
		}
	}
	outf.Close()

}

func writePPMHeader(outf *os.File,
	width, height, depth int) {
	fheader := []byte("P6\n")
	binary.Write(outf, binary.LittleEndian, fheader)
	ts := fmt.Sprintf("%d %d\n", width, height)
	binary.Write(outf, binary.LittleEndian, []byte(ts))
	ts2 := fmt.Sprintf("%d\n", depth)
	binary.Write(outf, binary.LittleEndian, []byte(ts2))
}

func writePPM(filename string, board *[Width][Height]int) {
	outf, _ := os.Create(filename)
	writePPMHeader(outf, Width, Height, 255)
	for x := 0; x < Width; x++ {
		for y := 0; y < Height; y++ {

			var v = board[x][y]
			var buf [3]uint8
			if v > 0 {
				buf[0] = uint8(255 - v)
				binary.Write(outf, binary.LittleEndian, buf)
				//binary.Write(outf, binary.LittleEndian, uint8(255-v))
				//binary.Write(outf, binary.LittleEndian, uint8(0))
				//binary.Write(outf, binary.LittleEndian, uint8(0))
			} else {
				buf[0] = 255
				buf[1] = 255
				buf[2] = 255
				binary.Write(outf, binary.LittleEndian, buf)
				//binary.Write(outf, binary.LittleEndian, uint8(255))
				//binary.Write(outf, binary.LittleEndian, uint8(255))
				//binary.Write(outf, binary.LittleEndian, uint8(255))
			}
		}
	}
	outf.Close()
}

func nextPoint(x, y float64) (float64, float64) {
	const a float64 = -1.4
	const b float64 = 1.6
	const c float64 = 1.0
	const d float64 = 0.7

	// this ranges from 1.97 to -1.4 X
	// this ranges from 1.53 to -1.3 Y
	//var newx float64 = math.Sin(a*y) +
	//	c*math.Cos(a*x)
	//var newy float64 = math.Sin(b*x) +
	//	d*math.Cos(b*y)

	return math.Sin(a*y) + c*math.Cos(a*x), math.Sin(b*x) + d*math.Cos(b*y)
}

func plot(board *[Width][Height]int, x, y float64) {
	var x_scaled int = int(x * 400)
	var y_scaled int = int(y * 300)
	if x_scaled > 0 && x_scaled < Width {
		if y_scaled > 0 && y_scaled < Height {
			board[x_scaled][y_scaled] += 1
		}
	}
	//fmt.Println("point(",x_scaled,",",y_scaled,")")
	//
	// now plot the points

}

func main() {
	const TotalPoints int = 1000000 * 2
	var board [Width][Height]int
	var x, y float64
	x = 0.0
	y = 0.0
	var maxX, maxY, minX, minY float64

	var counter int = 0
	for counter < TotalPoints {
		x1, y1 := nextPoint(x, y)
		//fmt.Println(x,y,x1,y1)
		if x1 < minX {
			minX = x1
		}
		if y1 < minY {
			minY = y1
		}
		if x1 > maxX {
			maxX = x1
		}
		if y1 > maxY {
			maxY = y1
		}
		x = x1
		y = y1

		plot(&board, x, y)
		counter++
	}

	writePPM("junk.ppm", &board)
	dbg_write(&board)
	//fmt.Println("Max x and y", maxX, maxY)
	//fmt.Println("Min x and y", minX, minY)

}
