package main

import "fmt"

func start() (string, string) {
	var (
		Player1 string
		Player2 string
	)
	fmt.Println("Insert your names")
	fmt.Scanf("%s\n", &Player1)
	fmt.Scanf("%s\n", &Player2)
	fmt.Println("All right,", Player1, "and", Player2)
	return Player1, Player2
}

func game(chel1 string, chel2 string) {
	var object1 string
	var object2 string
	fmt.Println("You can write: stone, paper or scissors")
	fmt.Println(chel1, "Throw")
	fmt.Scanf("%s\n", &object1)
	fmt.Println(chel2, "Throw")
	fmt.Scanf("%s\n", &object2)

	if (object1 == "stone" && object2 == "scissors") || (object1 == "paper" && object2 == "stone") || (object1 == "scissors" && object2 == "paper") {
		fmt.Println(chel1, ",wins.")
	} else if (object1 == "scissors" && object2 == "stone") || (object1 == "stone" && object2 == "paper") || (object1 == "paper" && object2 == "scissors") {
		fmt.Println(chel2, ",wins.")
	} else {
		fmt.Println("Draw")
	}
}

func main() {
	name1, name2 := start() // Оставляем, чтобы начислять победы по именам
	game(name1, name2)
}
