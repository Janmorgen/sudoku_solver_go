package main 
import (
  "testing"
  // "regexp"
  "fmt"
)

func test_nested_loop(){
  for x := range 3 {
    for y := range 3 {
      fmt.Println("X: ",x-1," Y: ",y-1)
    }
  }
}
func is_board_valid(board [9][9]int) bool {
  for x := range 9{
    for y := range 9{
      test_pos := [2]int{x,y}
      if (board[x][y]==0){
        options := find_missing(board,test_pos)
        if len(options)==0{
          return false
        }
      }
    }
  }
  return true

}
func TestBoards(t *testing.T){

  for difficulty_value := 17; difficulty_value<57; difficulty_value+=10{
    failure_count:=0
    for test_board_count := 0; test_board_count < 10 ; test_board_count++{
      board := create_sudoku_board(difficulty_value)
      if !is_board_valid(board){
        failure_count ++
      }
    }

    if (failure_count != 0){
      t.Fatalf("[%d] Expected test success of 10, but got %d",difficulty_value,(10-failure_count))
    }

    // fmt.Printf("Difficulty: {difficulty_value} -- Success rate {(10 - failure_count)/10}%")

  }
}
