package main

import "fmt"

func GenDisplaceFn(s0, v0, a float64) func(float64) float64 {
	return func(t float64) float64 {
		return s0 + v0*t + 0.5*a*t*t
	}
}

func REadValue(message string) float64 {
	var value float64
	for {

		fmt.Println(message)
		_, err := fmt.Scanf("%f", &value)
		if err != nil {
			var discard string
			fmt.Scanln(&discard)
			continue

		}

		break
	}
	return value
}

func main() {
	var displ, vel, acsel, tiime float64
	acsel = REadValue("Enter value for acceleration: ")
	vel = REadValue("Enter value for velocity: ")
	displ = REadValue("Enter value for displacement: ")

	fn := GenDisplaceFn(displ, vel, acsel)

	for {
		var cont string
		tiime = REadValue("Enter the value of time:")
		fmt.Println(fn(tiime), "\n to exit enter x or anything else to continue")
		fmt.Scan(&cont)
		if cont == "x" {
			break
		}
	}

}
