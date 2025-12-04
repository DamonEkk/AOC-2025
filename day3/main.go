package main

import "fmt"
import "strings"
import "strconv"

func main(){


	adventInput := Get_input()
	_ = Get_battery(adventInput)
}


func Get_battery(adventInput string)int{
	
	adventInput = strings.ReplaceAll(adventInput, " ", "")
	adventArray := strings.Split(adventInput, "\n") 
	totalValue := 0
	
	for i := 0; i < len(adventArray); i++{
		totalValue += Get_line_num(adventArray[i])
		fmt.Printf("%s totalVal: %d \n", adventArray[i], totalValue)
	}

	return 0
}

func Get_line_num(line string)int{
	
	var firstValue byte = '0'
	firstPosition := 0
	var secondValue byte = '0'

	for i := 0; i < len(line) - 1; i++{
		if (line[i] > firstValue){
			firstValue = line[i]
			firstPosition = i
		}
	}

	for i := firstPosition + 1; i < len(line); i++{
		if (line[i] > secondValue){
			secondValue = line[i]
		}
	}

	combinedValue := string([]byte{firstValue, secondValue})

	finalVal, _ := strconv.Atoi(combinedValue)

	return finalVal
}
