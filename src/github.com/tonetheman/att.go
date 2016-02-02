package main
import ( "math"
"fmt"
"os"
"encoding/binary" )

// saw this
// http://paulbourke.net/fractals/clifford/

func writePPMHeader(outf *os.File,
  width,height,depth int) {
  fheader := []byte("P6\n")
  binary.Write(outf,binary.LittleEndian,fheader)
  ts := fmt.Sprintf("%d %d\n",width,height)
  binary.Write(outf,binary.LittleEndian,[]byte(ts))
  ts2 := fmt.Sprintf("%d\n",depth)
  binary.Write(outf,binary.LittleEndian,[]byte(ts2))
}

func writePPM(filename string, board * [800][800]int) {
    outf,_ := os.Create(filename)
    writePPMHeader(outf,800,800,255)
    for x := 0;x<800;x++ {
      for y := 0;y<800;y++ {

          var v = board[x][y]
          if v>0 {
            binary.Write(outf, binary.LittleEndian,uint8(255-v))
      			binary.Write(outf, binary.LittleEndian,uint8(0))
      			binary.Write(outf, binary.LittleEndian,uint8(0))
          } else {
            binary.Write(outf, binary.LittleEndian,uint8(255))
      			binary.Write(outf, binary.LittleEndian,uint8(255))
      			binary.Write(outf, binary.LittleEndian,uint8(255))
          }
      }
    }
    outf.Close()
}

func nextPoint(x,y float64) (float64,float64) {
  var a float64 = -1.4;
  var b float64 = 1.6;
  var c float64 = 1.0;
  var d float64 = 0.7


  // this ranges from 1.97 to -1.4 X
  // this ranges from 1.53 to -1.3 Y
  var newx float64 = math.Sin(a * y) +
    c * math.Cos(a * x)
  var newy float64 = math.Sin(b * x) +
    d * math.Cos(b * y)

  return newx,newy;
}

func plot(board *[800][800]int, x,y float64) {
    var x_scaled = (x * 400)
    var y_scaled = (y * 400)
    if int(x_scaled) > 0 && int(x_scaled) < 800 {
      if int(y_scaled) > 0 && int(y_scaled) < 800 {
        board[int(x_scaled)][int(y_scaled)] += 1
      }
    }
    //fmt.Println("point(",x_scaled,",",y_scaled,")")
//
    // now plot the points

}

func main() {
  const TotalPoints int = 1000000
  var board [800][800]int
  var x,y float64
  x = 0.0
  y = 0.0
  var maxX,maxY,minX,minY float64

  var counter int = 0
  for counter < TotalPoints {
    x1,y1 := nextPoint(x,y)
    //fmt.Println(x,y,x1,y1)
    if x1<minX {
      minX = x1
    }
    if y1 <minY {
      minY = y1
    }
    if x1>maxX {
      maxX = x1
    }
    if y1>maxY {
      maxY = y1
    }
    x = x1
    y = y1

    plot(&board,x,y)
    counter ++
  }

  writePPM("junk.ppm",&board)

  //fmt.Println("Max x and y", maxX, maxY)
  //fmt.Println("Min x and y", minX, minY)

}
