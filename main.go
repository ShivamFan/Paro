package main

import (
	"os"

	consumerproducer "paro/consumer-producer"
	linuxtree "paro/linux-tree"
	topscript "paro/top-script"
)

func main() {
	argsWithProg := os.Args
	argSlice := argsWithProg[1:]
	if len(argSlice) > 0 {
		if argSlice[0] == "1" {
			consumerproducer.Ques1()
		} else if argSlice[0] == "2" {
			linuxtree.Ques2()
		} else if argSlice[0] == "3" {
			topscript.Ques3()
		}
	}
}
