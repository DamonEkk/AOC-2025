package main

import "fmt"
import "strings"
import "strconv"


func main(){
	inputString := get_input()
	//inputString = "212118-212124"
	//inputString = "100000-200000"
	//inputString = `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,
	//1698522-1698528,446443-446449,38593856-38593862,565653-565659,
	//824824821-824824827,2121212118-2121212124`
	//inputString = "38593856-38593862"
	//inputString = "1188511880-1188511890"
	stringParts := strings.Split(inputString, ",")
	var totalSum int64 = 0
	
	for i:= 0; i < len(stringParts); i++{
		point := strings.Split(stringParts[i], "-")
		first, _ := strconv.ParseInt(point[0], 10, 64)
		end, _ := strconv.ParseInt(point[1], 10, 64)

		for k := first; k <= end; k++{
			stringK := strconv.FormatInt(k, 10)

			if (Check_pattern(stringK) == true){
				totalSum += k
			}

			

		}
	}	

	fmt.Printf("Sum = %d", totalSum)

}


func Check_pattern(tempString string)bool{

	pattern := Get_pattern(tempString)
	patternFlag := 0
	

	if (tempString == "1188511885"){
		fmt.Println(pattern)
	}

	if (len(pattern) == 0){
		return false
	}

	// for loop that checks whole string for pattern
	for i := 0; i < len(tempString); i++{
		if (strings.Count(tempString, pattern) != len(tempString) / len(pattern)){
			return false
		}
		patternFlag = 1
		if (i == len(tempString) - 1){
			fmt.Printf("Pattern = %s and input = %s\n", pattern, tempString)
		}
	}

	return patternFlag == 1	

}


func Get_pattern(tempString string)string{
	var dividedVals[] int
	pattern := ""

	
	for i := 1; i <= len(tempString) / 2; i++{
		if (len(tempString) % i == 0){
			dividedVals = append(dividedVals, i)
		}
	}
	
	// Check for patterns 
	for i := 0; i < len(dividedVals); i++{
		tempPattern := tempString[:dividedVals[i]]
		if (strings.Count(tempString, tempPattern) == len(tempString) / len(tempPattern)){
			pattern = tempPattern
			break
		}
	}

	return pattern	
}
