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
	
	
	// I need to find the highest digit with x minimal spots. Could reverse search and cut it down limit being how many digits we need. / up to previous digit.

	previousIndex := -1
	finalNum := ""

	for remaindingDigits := 12; remaindingDigits > 0; remaindingDigits--{
		tempIndex := 0
		var tempNum byte = '0'
		for k:= len(line) - remaindingDigits; k > previousIndex; k--{
			if (tempNum <= line[k]){
				tempNum = line[k]
				tempIndex = k
			}	
		}
		previousIndex = tempIndex
		finalNum += string(tempNum)
	}

	convertNum, _ := strconv.Atoi(finalNum)

	return convertNum

}
