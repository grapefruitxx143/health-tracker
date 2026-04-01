package reports

import (
	"fmt"
	"health-tracker/models"
	"strings"
)

func ShowingAcitve(u models.User) {

	fmt.Println("\n🏃 PHYSICAL ACTIVITY LEVEL")
	fmt.Println("Choose the option that fits you best:")
	fmt.Println("  [low]  - Little or no exercise")
	fmt.Println("  [mid]  - Moderate exercise (3-5 days/week)")
	fmt.Println("  [high] - Heavy exercise or physical job")

	BMR, ok := u.CalculateBMR()

	if !ok {
		fmt.Println("Last weight not founded.")
		return
	}

	var getActive string

	for {

		fmt.Print("\nYour selection: ")
		_, err4 := fmt.Scan(&getActive)
		getActive = strings.ToLower(getActive)

		if err4 != nil {
			fmt.Println("Please write a word.")

			var dump string
			fmt.Scanln(&dump)
			continue
		}
		if getActive != "low" && getActive != "mid" && getActive != "high" {
			fmt.Println("Please, write (low, mid, high) correct")
			continue
		}
		break
	}

	switch getActive {
	case "low":
		fmt.Printf("Your realistic BMR is %.2f kcal/day\n", BMR*1.2)
	case "mid":
		fmt.Printf("Your realistic BMR is %.2f kcal/day\n", BMR*1.5)
	case "high":
		fmt.Printf("Your realistic BMR is %.2f kcal/day\n", BMR*1.8)

	}

}

func UpdateWeight(u *models.User) {

	fmt.Println("\nDAILY WEIGHT CHECK-IN")
	fmt.Printf("----------------------------------------\n")
	fmt.Printf("Hello again, %s! \n", u.Name)
	CurrWeight, ok := u.CurrentWeight()
	if !ok {
		fmt.Println("Last weight not founded.")
	} else {
		fmt.Printf("Your last recorded weight was: %.1f kg\n", CurrWeight)
	}



	for {

		var CurrentWeight float64

		fmt.Print("\nHow much do you weigh today? (kg): ")

		_, err := fmt.Scan(&CurrentWeight)
		if err != nil {
			fmt.Println("Please, write a number.")

			var dump string
			fmt.Scanln(&dump)
			continue
		}

		if CurrentWeight <= 0 || CurrentWeight > 500 {
			fmt.Println("Invalid Weight. Please, enter realistic number.")
			continue
		}

		u.Weight = append(u.Weight, CurrentWeight)
		break
	}
}

func RegisterNewUser() models.User {
	var name, gender string
	var age int
	var height, weight float64
	var firstUser models.User

	fmt.Println("\nWELCOME TO GO-HEALTH TRACKER")
	fmt.Println("========================================")
	fmt.Println("Let's create your profile to get started.")

	for {

		fmt.Print("\nStep 1. Enter your Name and Age: ")
		_, err := fmt.Scan(&name, &age)

		if err != nil {
			fmt.Println("Please, enter word name AND number age!!!")

			var dump string
			fmt.Scanln(&dump)

			continue
		}
		if age <= 0 || age >= 120 {
			fmt.Println("Please, enter a realistic age.")
			continue
		}
		break
	}

	for {

		fmt.Print("Step 2. Gender (Male/Female): ")
		_, err1 := fmt.Scan(&gender)
		gender = strings.ToLower(gender)

		if err1 != nil {
			fmt.Println("Please, write a WORD.")
			var dump string
			fmt.Scanln(&dump)

			continue
		} else if gender != "male" && gender != "female" {
			fmt.Println("Please, write Male or Female correct.")
			continue
		}
		break
	}

	for {

		fmt.Print("Step 3. Enter your Height(cm) and Weight(kg): ")
		_, err2 := fmt.Scan(&height, &weight)

		if err2 != nil {
			fmt.Println("Please, write a NUMBERS.")

			var dump string
			fmt.Scanln(&dump)

			continue
		}
		break
	}

	fmt.Println("\nProfile created successfully!")
	firstUser = models.User{Name: name, Age: age, Gender: gender, Height: height, Weight: []float64{weight}}

	return firstUser
}

func ShowTotalProgress(u models.User) {

	fmt.Println("Doy you want to see your progress sincde first day? (yes/no): ")

	for {
		var CheckWord string
		_, errCheck := fmt.Scan(&CheckWord)
		CheckWord = strings.ToLower(CheckWord)

		if errCheck != nil {
			fmt.Println("Please, write a word 'yes' or 'no'.")
			var dump string
			fmt.Scanln(&dump)
			continue
		}

		if CheckWord != "yes" && CheckWord != "no" {
			fmt.Println("Invalid answer, please enter 'yes' or 'no'.")
			continue
		}

		if CheckWord == "yes" {

			if len(u.Weight) >= 2 {
				firstWeight := u.Weight[0]
				currentWeight := u.Weight[len(u.Weight)-1]
				if currentWeight > firstWeight {
					fmt.Println("You add since first day ", (currentWeight - firstWeight), " kg")
				} else if firstWeight > currentWeight {
					fmt.Println("You lost since first day ", (firstWeight - currentWeight), " kg")
				} else {
					fmt.Println("Weight since first day not changed.")
				}
			} else {
				fmt.Println("Not enough data to show.")
			}
		}
		break
	}
}
