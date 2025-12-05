package main


import "fmt"
import "time"
import "strings"

func main(){
	start := time.Now()
	advent := Advent_input()

	Check_paper(advent)	

	elapsed := time.Since(start)
	fmt.Printf("Finished %d ms\n", elapsed.Milliseconds())



}

func Check_paper(advent string){
	matrix := Create_matrix(advent)
	count := 0

	for i := 0; i < len()  

}


func Create_matrix(advent string)[][]byte{
	line := strings.Split(advent, "\n")
	matrix := make([][]byte, len(line))

	for i:= 0; i < len(line); i++{
		for k:= 0; k < len(line[i]); k++ {
			matrix[i][k] = line[i][k]
		}
	}  

	return matrix
}
