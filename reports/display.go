package reports

import (
	"fmt"
	"health-tracker/models"
)


func ShowWeightChange(u models.User) {
	if len(u.Weight) >= 2 {
		current := u.Weight[len(u.Weight)-1]
		prev := u.Weight[len(u.Weight)-2]

		if current > prev {
			fmt.Println("You gained", (current - prev), "kg. Keep focused on your goal!")
		} else if prev > current {
			fmt.Println("Great job! You lost", (prev - current), "kg since your last check-in!")
		} else {
			fmt.Println("Your weight is stable since last time.")
		}
	} else {
		fmt.Println("There's no data.")
	}

}

func ShouldAddLose(u models.User) {
	result, ok := u.CalculateBMI()

	if !ok {
		fmt.Println("Last weight not founded.")
		return
	}

	MinWeight, MaxWeight := u.GetWeightRange()

	if result < 18.5 {
		needToAdd := MinWeight - u.Weight[len(u.Weight)-1]
		fmt.Printf("You should add %.2f kg to reach healthy weight. \n", needToAdd)
	} else if result > 24.9 {
		needToLose := u.Weight[len(u.Weight)-1] - MaxWeight
		fmt.Printf("You should lose %.2f kg to reach healthy weight. \n", needToLose)
	}

	BMR, ok1 := u.CalculateBMR()

	if !ok1 {
		fmt.Println("Last weight not founded.")
		return
	}

	fmt.Printf("Your Basal Metabolic Rate (BMR) is %.2f kcal/day\n", BMR)
}

func ShowHealthReport(u models.User) {
	result, ok := u.CalculateBMI()

	if !ok {
		fmt.Println("Last weight not founded.")
		return
	}

	fmt.Println("\n📊 YOUR HEALTH REPORT")
	fmt.Println("========================================")
	fmt.Printf("Current BMI: %.2f\n", result)

	var minBMI, midBMI, maxBMI float64
	minBMI = 18.5
	midBMI = 25
	maxBMI = 30

	switch {
	case result < minBMI:

		fmt.Println("🔵 UNDERWEIGHT")
	case result >= minBMI && result < midBMI:
		fmt.Println("🟢 HEALTHY WEIGHT")
	case result >= midBMI && result < maxBMI:
		fmt.Println("🟡 OVERWEIGHT")
	default:
		fmt.Println("🔴 OBESITY")
	}
	fmt.Println("----------------------------------------")
}
