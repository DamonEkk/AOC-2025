package main

import "fmt"
import "strconv"

type password struct{
	currentPosition int
	zeroCount int
	code string
	i int
	direction int; // 0 is L  && 1 is R
}

func main(){
	pass := password{}
	pass.currentPosition = 50
	pass.zeroCount = 0
	pass.i = 0
	Get_code(&pass)
	

	for {
		temp:= 0
		if (len(pass.code) == pass.i){
			fmt.Println(pass.zeroCount)
			break;
		}
		if (pass.code[pass.i] == 'L'){
			pass.direction = 0
			temp = 0
			pass.i++

		} else if (pass.code[pass.i] == 'R'){
			pass.direction = 1
			temp = 0
			pass.i++
		} else{
			temp = Get_substring(&pass)
			Move(&pass, temp)
		}

	}
	fmt.Printf("final zeroCount: %d",pass.zeroCount)
}

func Get_substring(pass *password)int{
	subString := []byte("")
	for pass.i < len(pass.code){
		if (pass.code[pass.i] != 'L' && pass.code[pass.i] != 'R'){
			subString = append(subString, pass.code[pass.i])
			pass.i++
		} else{
			break
		}
	}
	temp := string(subString)
	num, _ := strconv.Atoi(temp)
	return num
}


func Move(pass *password, value int) {
	direc := ""
	if pass.direction == 0 {
		direc = "LEFT"
	} else {
		direc = "RIGHT"
	}

	fmt.Printf("The direction is %s Value: %d Initial location is: %d ", direc, value, pass.currentPosition)

	for i := 0; i < value; i++{
		
		// LEFT
		if (pass.direction == 0){
			pass.currentPosition--
			if (pass.currentPosition == 0){
				pass.zeroCount++
			} else if (pass.currentPosition == -1){
				pass.currentPosition = 99
			}
		// RIGHT
		} else if (pass.direction == 1){
			pass.currentPosition++
			if (pass.currentPosition == 100){
				pass.zeroCount++
				pass.currentPosition = 0
			}
		} 
	}

	fmt.Printf("New Position: %d ZeroCount: %d\n", pass.currentPosition, pass.zeroCount)
}



