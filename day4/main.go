package main


import "fmt"
import "time"
import "strings"

func main(){
	start := time.Now()
	advent := Advent_input()




	fmt.Printf("%s\n\n", advent)
	


	Check_paper(advent)	

	elapsed := time.Since(start)
	fmt.Printf("Finished %d ms\n", elapsed.Milliseconds())

}

func Check_paper(advent string){
	matrix := Create_matrix(advent)
	count := 0
	total := 0
	changes := 0
	
	for {
		changes = 0
		count = 0
		for i := 0; i < len(matrix); i++{
			for k := 0; k < len(matrix[i]); k++{
					
				if (Check_border(matrix, k, i) == false && matrix[i][k] == '@'){
					count++
					changes++
					fmt.Printf("x")
					matrix[i][k] = '.'
				} else{
					fmt.Printf("%s", string(matrix[i][k]))
				}

			}
			fmt.Printf("\n")
		}

		total += count

		fmt.Println("")
		fmt.Printf("Count: %d\n", total)
		
		// Look at the @ and check if they are the same as old matrix. 
		if (changes == 0){
			break
		}

	}
}


func Create_matrix(advent string)[][]byte{
	line := strings.Split(advent, "\n")
	matrix := make([][]byte, len(line))

	for i:= 0; i < len(line); i++{
		matrix[i] = make([]byte, len(line[i]))
		for k:= 0; k < len(line[i]); k++ {
			matrix[i][k] = line[i][k]
		}
	}  

	return matrix
}

func Check_border(matrix [][]byte, x int, y int)bool{
	var count int = 0

	for iterator := y-1; iterator < y+2; iterator++{
		for kterator := x-1; kterator < x+2; kterator++{
			if (iterator < 0){
				continue
			}
			if (iterator >= len(matrix)){
				continue
			}
			if (kterator < 0){
				continue
			}
			if (kterator >= len(matrix[iterator])){
				continue
			}

			if (iterator == y && kterator == x){
				continue
			}

			if (matrix[iterator][kterator] == '@'){
				count++
			}

		}
	} 

	if (count < 4){
		return false
	}
	return true
} 
