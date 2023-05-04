/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"main/cmd"
)

func main() {
	cmd.Execute()
	Jwt()
}
func Jwt() {
	fmt.Println("jwt")
}
