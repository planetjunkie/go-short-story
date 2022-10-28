package main


import (
	"fmt"
)

var hero string = ""
var age int = 0
	

func main() {
	

	fmt.Println("Input your hero's name")
	fmt.Scan(&hero)
	fmt.Println("Input your hero's age")
	fmt.Scan(&age)


	fmt.Println("Your hero's name is:", hero, "\nYour hero's age is:", age)
	fmt.Println("------")
	fmt.Println("Let the journy begin")
	fmt.Println("------")

	storypt1()
}

func storypt1() {
	fmt.Println("Our hero", hero, "Was wandering around the forest, but he encounters a big goblin blocking the cave, What should he do?")
	fmt.Println("1) Ask the Goblin \n2) Slay the Goblin")

	var choice int = 0

	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("He asked the goblin to move out the way, but the goblin got angry due to him not understanding what ", hero, " Said and he ate Our hero Up")
		fmt.Println("Game over")
	} else if choice == 2 {
		fmt.Println("Our Hero slayed the goblin with ease and was able to Move on to into the cave, in there, he found the teasure and took the artifact he needed")
		fmt.Println("The end")
	}
}