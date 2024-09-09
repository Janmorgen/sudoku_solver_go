package main 
import (
  "fmt"
)

func test_nested_loop(){
  for x := range 3 {
    for y := range 3 {
      fmt.Println("X: ",x-1," Y: ",y-1)
    }
  }
}

func main () {
  test_nested_loop()
}
