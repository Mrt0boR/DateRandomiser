package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/charmbracelet/huh"
)

type DateIdea struct {
	Name     string
	Category string
	Checked  bool
}

// Slice to store saved date ideas
var ExpensiveDates []DateIdea
var MediumDates []DateIdea
var CheapestDates []DateIdea

func main() {
	for {
		var mainChoice string
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Title("Date Randomizer").
					Options(
						huh.NewOption("Create a New Date", "create"),
						huh.NewOption("View All Date Ideas", "view"),
						huh.NewOption("Random Expensive Date", "randomExpensive"),
						huh.NewOption("Random Medium Date", "randomMedium"),
						huh.NewOption("Random Cheap Date", "randomCheap"),
						huh.NewOption("Exit", "exit"),
					).
					Value(&mainChoice),
			),
		)

		err := form.Run()
		if err != nil {
			log.Fatal(err)
		}

		switch mainChoice {
		case "create":
			createDate()
		case "view":
			viewAndCheckDateIdeas()
		case "randomExpensive":
			randomDate(&ExpensiveDates, "Expensive")
		case "randomMedium":
			randomDate(&MediumDates, "Medium")
		case "randomCheap":
			randomDate(&CheapestDates, "Cheap")
		case "exit":
			fmt.Println("Exiting... Have a great date! ❤️")
			return
		}
	}
}

func viewAndCheckDateIdeas() {
	if len(ExpensiveDates) == 0 && len(MediumDates) == 0 && len(CheapestDates) == 0 {
		fmt.Println("\n⚠️ No date ideas saved yet! Try creating one first!")
		fmt.Println("Press Enter to return to the main menu...")
		fmt.Scanln()
		return
	}

	fmt.Println("\n=== All Date Ideas ===")

	checkAndPrintDates("💎 Expensive Dates:", &ExpensiveDates)
	checkAndPrintDates("🎯 Medium Dates:", &MediumDates)
	checkAndPrintDates("💰 Cheap Dates:", &CheapestDates)

	fmt.Println("\nPress Enter to return to the main menu...")
	fmt.Scanln()
}

func checkAndPrintDates(title string, dateList *[]DateIdea) {
	if len(*dateList) > 0 {
		fmt.Println("\n" + title)
		for i, date := range *dateList {
			checkedMark := ""
			if date.Checked {
				checkedMark = " ✅"
			}
			fmt.Printf("%d. %s%s\n", i+1, date.Name, checkedMark)
		}

		var indexToCheck int
		fmt.Println("\nEnter the number of the date you want to check off, or 0 to skip:")
		fmt.Scanln(&indexToCheck)

		if indexToCheck > 0 && indexToCheck <= len(*dateList) {
			(*dateList)[indexToCheck-1].Checked = true
			fmt.Println("\n✅ Date checked off!")
		}
	}
}

func randomDate(dateList *[]DateIdea, category string) {
	var confirm bool
	randomForm := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().Title("Random " + category + " Date").
				Value(&confirm),
		),
	)
	err := randomForm.Run()
	if err != nil {
		log.Fatal(err)
	}

	if confirm {
		if len(*dateList) == 0 {
			fmt.Printf("\n⚠️ No %s date ideas saved yet! Try creating one first!\n", category)
			fmt.Println("Press Enter to return to the main menu...")
			fmt.Scanln()
			return
		}

		rand.Seed(time.Now().UnixNano())
		randomIndex := rand.Intn(len(*dateList))
		selectedDate := (*dateList)[randomIndex]

		fmt.Printf("\n🎉 Your random %s date idea: %s\n", category, selectedDate.Name)
		fmt.Println("\nPress Enter to return to the main menu...")
		fmt.Scanln()
	}
}

func createDate() {
	var dateName string
	var category string

	nameForm := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Enter the date").
				Value(&dateName),
		),
	)
	err := nameForm.Run()
	if err != nil {
		log.Fatal(err)
	}

	categoryForm := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select a category").
				Options(
					huh.NewOption("Cheap", "cheap"),
					huh.NewOption("Average", "average"),
					huh.NewOption("Expensive", "expensive"),
				).
				Value(&category),
		),
	)

	err = categoryForm.Run()
	if err != nil {
		log.Fatal(err)
	}

	newDate := DateIdea{Name: dateName, Category: category}
	if category == "expensive" {
		ExpensiveDates = append(ExpensiveDates, newDate)
	} else if category == "cheap" {
		CheapestDates = append(CheapestDates, newDate)
	} else if category == "average" {
		MediumDates = append(MediumDates, newDate)
	}

	fmt.Println("\n✅ Date idea saved successfully!")

	var createAnother bool
	confirmForm := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Would you like to create another date?").
				Value(&createAnother),
		),
	)

	err = confirmForm.Run()
	if err != nil {
		log.Fatal(err)
	}

	if createAnother {
		createDate()
	}
}
