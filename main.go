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

func chooseWhoFirst(number1 string, number2 string) (string, string) {
	fmt.Println("Who first:", number1, "or", number2, "?")
	var firstMoveName string
	var secondMoveName string
	fmt.Scanf("%s\n", &firstMoveName) // Выбираем первое имя
	fmt.Println("Good,", firstMoveName)
	if firstMoveName == number1 { // Второе имя относительно выбранного первого
		secondMoveName = number2
	} else {
		secondMoveName = number1
	}
	return firstMoveName, secondMoveName
}

func game() int {
	var object1 string
	var object2 string
	fmt.Println("You can write: stone, paper or scissors")
	fmt.Println("Move1")
	fmt.Scanf("%s\n", &object1)
	fmt.Println("Move2")
	fmt.Scanf("%s\n", &object2)

	if (object1 == "stone" && object2 == "scissors") || (object1 == "paper" && object2 == "stone") || (object1 == "scissors" && object2 == "paper") {
		return 1
	} else if (object1 == "scissors" && object2 == "stone") || (object1 == "stone" && object2 == "paper") || (object1 == "paper" && object2 == "scissors") {
		return 2
	} else {
		return 3
	}
}

func main() {
	name1, name2 := start() // Оставляем, чтобы начислять победы по именам
	first, second := chooseWhoFirst(name1, name2)
	numberScenaries := game()
	if numberScenaries == 1 {
		fmt.Println(first, ", win")
	} else if numberScenaries == 2 {
		fmt.Println(second, ", win")
	} else {
		fmt.Println("Draw")
	}
}
