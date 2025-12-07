package main

import "strings"

type advent struct{
	freshList string
	food string
}

func Advent_helper()advent{
	adventString := `

	`

	var adventList advent

	parts := strings.SplitN(adventString, " ", 2) 
	adventList.freshList = parts[0]
	adventList.food = parts[0]

	return adventList
}
