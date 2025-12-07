package main

import "fmt"
import "strconv"


func main(){
	

}


func Get_ingrediant_map(advent string){

	food := Advent_helper()
	m := make(map[string]string) // Homemade set

	rangeA := ""
	rangeB := ""
	compareString := ""

	mode := 0
	iterator := 0
	maxLength := len(food.freshList)
	count := 0

	for {

		if (iterator == maxLength){
			break
		}

		if (food.freshList[iterator] == '-'){
			mode = 1
			iterator++
			continue
		}

		if (food.freshList[iterator] == '\n'){

			if (rangeB == "" || rangeA == ""){
				mode = 0
				continue
			}

			a, _ := strconv.ParseInt(rangeA, 10, 64)
			b, _ := strconv.ParseInt(rangeB, 10, 64)
			
			for ; a <= b; a++{

				_, ok := m[strconv.FormatInt(a, 10)]

				if (ok){
					continue
				}
				m[strconv.FormatInt(a, 10)] = "Fresh"	
			}
			
			mode = 0
			continue
		}

		if (mode == 0){
			rangeA += string(food.freshList[iterator])
		} else {
			rangeB += string(food.freshList[iterator])
		}

		iterator++
	}

	for i := 0; i < len(food.food); i++{
		
		if (food.food[i] == '\n'){
			compareString = ""
		} else{
			compareString += food.food[i]
		}

		_, ok := m[compareString]
		if (ok){
			count++	
		}

	}

	
	// Do till end of string. Get the first digit and the second digit.
}
