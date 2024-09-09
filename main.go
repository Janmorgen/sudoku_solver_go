package main

import (
	"fmt"
	"math"
	"math/rand"
  "strings"
  "strconv"
  "github.com/fatih/color"
  "time"
)

// Helper function to remove a specific value from an array
func remove_value(array []int, value int) []int {
  for i,num := range array{
    if (num == value){
      array = append(array[:i],array[i+1:]...)
    }
  }
  return array
}

// First rule: Number cannot replace an existing number
// Second rule: Number cannot be repeated horizontally
// Third rule : Number cannot be repeated vertically
// Fourth rule : There cannot be more than 9 of a number
// Fith rule : There cannot be repeated numbers in the same 3x3 grid

// Finds the valid number options in a specific board position
func find_missing(board [9][9]int, pos [2]int) []int {
  // Define the total options available, 1-9
  possible_numbers := []int{1,2,3,4,5,6,7,8,9} 

  // If the board position is already occupied, there are no options available
  if board[pos[0]][pos[1]] != 0 {
    empty_slice := []int{}
    return empty_slice 
  }

  // Check the horizontal and vertical axis for numbers and remove them from possible_numbers
  for i := range 9 {
    v_num := board[pos[0]][i]
    h_num := board[i][pos[1]]
    if v_num != 0{
      possible_numbers= remove_value(possible_numbers,v_num)
    } 
    if h_num != 0{
      possible_numbers= remove_value(possible_numbers,h_num)
    }
  }

  // Check the 3x3 grid of the position and remove numbers which already exist from possible_numbers
  for _,i := range possible_numbers{
    if !check_grid(board, pos[0], pos[1], i){
      possible_numbers= remove_value(possible_numbers,i)
    }
  }
  return possible_numbers
  
}

// Deprecated function -- Replaced by find_missing
// Checks if the given number is valid in a specific board position
func check_insert(board [9][9]int, pos [2]int, number int) bool {
	// If the position is not empty return false
	if board[pos[0]][pos[1]] != 0 {
		// fmt.Println("Check revealed an existing number in that position")
		return false
	}
	// if the number exists in a horizontal or vertical line, return false
	for i := range 9 {
		if board[pos[0]][i] == number {
			// fmt.Println("Check revealed an existing number in the column: ", i)
			return false
		}
		if board[i][pos[1]] == number {
			// fmt.Println("Check revealed an existing number in the row: ", i)
			return false
		}
	}
  // Check the closest 3x3 grid and see if the number already exists within it
	return check_grid(board, pos[0], pos[1], number)

}

// Used to calculate the distance between two sets of board coordinates
func distance_between_points(xpos1 int, ypos1 int, xpos2 int, ypos2 int) float64 {
	return math.Sqrt(math.Pow(float64(xpos1)-float64(xpos2), 2) + math.Pow(float64(ypos1)-float64(ypos2), 2))
}

// Determines which grid the set of coordinates belong to, there are 9 3x3 grids in a 9x9 grid
func within_grid(xpos int, ypos int) int {
  // Define the center coordinate of each 3x3 grid
	var grid_centers = [9][2]int{
		{1, 1}, {1, 4}, {1, 7},
		{4, 1}, {4, 4}, {4, 7},
		{7, 1}, {7, 4}, {7, 7}}

  // Calculate the distance between {xpos,ypos} and each grid center coordinate,
  //  If the value is lower than 2, then the {xpos,ypos} coordinate belongs to that grid
  //    Return the index of the grid
	for index, pos := range grid_centers {
		if distance_between_points(pos[0], pos[1], xpos, ypos) < 2 {
			return index
		}
	}
  // This will never be returned, but a good thing to look out for upstream
	return -1
}

// Check whether the supplied value at a specific coordinate already exists in its 3x3 grid
func check_grid(board [9][9]int, xpos int, ypos int, value int) bool {
  // Define the center coordinate of each 3x3 grid
	var grid_centers = [9][2]int{
		{1, 1}, {1, 4}, {1, 7},
		{4, 1}, {4, 4}, {4, 7},
		{7, 1}, {7, 4}, {7, 7}}
  // Get the index of the 3x3 grid {xpos, ypos} belongs to
	var index = within_grid(xpos, ypos)
	var center_x = grid_centers[index][0]
	var center_y = grid_centers[index][1]

  // for each index within the 3x3 grid, test if its value is equal to the value supplied
	for x := range 3 {
		for y := range 3 {
			var test_x = center_x + x - 1
			var test_y = center_y + y - 1
			if board[test_x][test_y] == value {
				// fmt.Println("Check revealed an existing number in grid: ", index, " at position X: ", test_x, ", Y: ", test_y)
				return false
			}
		}

	}
	return true

}


// Print the 9x9 board, include color coding for coordinates supplied through highlights
func print_board(board [9][9]int, highlights ...[2]int) {
	for y := range 9 {
		for x := range 9 {
      // Only highlight colors which match the first coordinate of highlights
      if len(highlights) == 1 && highlights[0][0]==x && highlights[0][1]==y{

        // If the existing value on the board is 0 print the number as Green 
        //  If the existing value isn't 0 print the number as Red
        if board[x][y] != 0{
        fmt.Print(" ",color.RedString(strconv.Itoa(board[x][y]))," ")
        }else {
          fmt.Print(" ",color.GreenString(strconv.Itoa(board[x][y]))," ")
        }
      }else{

			fmt.Print(" ", board[x][y], " ")
      }
		}
		fmt.Println()
	}
  fmt.Println(strings.Repeat(" - ",9))
}

// The number passed to this function determines how many tiles get revealed
func create_sudoku_board(difficultiy int) [9][9]int {
  // Define the board
	var board = [9][9]int{}


	for i := 0; i < difficultiy; i++ {
		var x = rand.Intn(9)
		var y = rand.Intn(9)
		insert_pos := []int{x, y}

    // Find the valid number options for the random tile
    options := find_missing(board, [2]int(insert_pos)) 

    // Ensure that empty options do not affect the choice of difficulty
    if len(options)!=0{
      var number = options[rand.Intn(len(options))]
      board[x][y] = number
    }else{
      i --
    }
		

	}
	return board

}
func main() {

	// example_board := [9][9]int{
	// 	{0, 0, 6, 9, 1, 2, 4, 8, 0},
	// 	{0, 1, 0, 3, 0, 0, 7, 6, 0},
	// 	{3, 8, 0, 0, 0, 0, 0, 0, 2},
	// 	{8, 0, 1, 0, 7, 3, 0, 0, 4},
	// 	{0, 0, 0, 0, 8, 0, 1, 7, 0},
	// 	{5, 0, 7, 0, 6, 0, 0, 0, 8},
	// 	{0, 3, 0, 0, 0, 1, 2, 4, 0},
	// 	{0, 9, 4, 0, 0, 7, 6, 0, 5},
	// 	{2, 0, 0, 6, 0, 4, 0, 9, 0}}
  example_board := create_sudoku_board(30)
	print_board(example_board)
  // check_pos := [2]int{1,0}
  for tries := 0 ; tries < 4 ; tries ++ {
  for x := range 9{

    for y := range 9{

    fmt.Print("\033[0;0H")
    check_pos:=[2]int{x,y}
    missing_values := find_missing(example_board,check_pos)
    fmt.Println("Missing Values: ",missing_values)
    if len(missing_values) == 1{
      example_board[x][y]=missing_values[0]
    }

    print_board(example_board, check_pos)
    time.Sleep(500000000)

fmt.Print("\033[0;0H")
    for i :=0; i<30; i++ {
      fmt.Println(strings.Repeat(" ", 50))
    }
    }
  }

}
fmt.Print("\033[0;0H")
fmt.Println("Final board")
print_board(example_board)
	// board := create_sudoku_board(30, example_board)
	// print_board(board)
}
