package main
import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
  var res = make([][]uint8, dy)
  for x := range res {
    res[x] = make([]uint8, dx)
    for y := range res[x] {
      res[x][y] = uint8(x*y)
    }
  }
  return res 
}

func main() {
    pic.Show(Pic)
}
