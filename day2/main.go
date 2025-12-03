package main

import "fmt"
import "strings"
import "strconv"


func main(){
	inputString := get_input()	
	stringParts := strings.Split(inputString, ",")
	var totalSum int64 = 0
	
	for i:= 0; i < len(stringParts); i++{
		point := strings.Split(stringParts[i], "-")
		first, _ := strconv.ParseInt(point[0], 10, 64)
		end, _ := strconv.ParseInt(point[1], 10, 64)

		for k := first; k <= end; k++{
		
			stringK := strconv.FormatInt(k, 10)
			sizeK := len(stringK) / 2 
			part0 := stringK[0:sizeK]
			part1 := stringK[sizeK:len(stringK)]
			if (part0 == part1){
				totalSum += k	
			}

			
		}
	}	

	fmt.Printf("Sum = %d", totalSum)

}
