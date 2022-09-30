package main

import (
	"fmt"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a *Animal) Eat() {
	fmt.Println(a.food)
}
func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}
func (a *Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	var name, action string
	var vre Animal
	var cow = Animal{
		food:       "grass",
		locomotion: "walk",
		noise:      "moo",
	}
	var bird = Animal{
		food:       "worms",
		locomotion: "fly",
		noise:      "peep",
	}
	var snake = Animal{
		food:       "mice",
		locomotion: "slither",
		noise:      "hsss",
	}

	for {
		fmt.Println("Animal are: cow , bird , snake\nThey can eat, move, speak")
		for {
			fmt.Println("enter name")
			fmt.Scanf("%s", &name)
			nm := strings.ToLower(name)

			switch nm {
			case "cow":
				vre = cow
			case "bird":
				vre = bird
			case "snake":
				vre = snake
			default:
				fmt.Println("cant find entered option")
				continue
			}
			break
		}

		for {

			fmt.Println("enter action")
			fmt.Scanf("%s", &action)
			act := strings.ToLower(action)

			switch act {
			case "eat":
				vre.Eat()
			case "speak":
				vre.Speak()
			case "move":
				vre.Move()
			default:
				fmt.Println("cant find entered option")
				continue
			}
			break
		}
	}
}
