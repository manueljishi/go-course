package main

import "fmt"

func main(){
	// colors := make(map[string]string)
	colors := map[string]string{
		"red": "#ff0000",
		"green": "#00ff00",
		"blue": "#0000ff",
		"white": "ffffff",
		"black": "000000",
	}
	fmt.Println(colors["red"])

	delete(colors,"red")
	printMap(colors)
}

func printMap(c map[string]string){
	for color, hex := range c{
		fmt.Println("Hex code for ", color, " is ", hex)
	}
}