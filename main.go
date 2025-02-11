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
}

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
			viewDateIdeas()
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

func createDate() {
	var dateName string
	var category string

	for {
		nameForm := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("Enter the date idea (cannot be empty)").
					Value(&dateName),
			),
		)

		err := nameForm.Run()
		if err != nil {
			log.Fatal(err)
		}
		if dateName != "" {
			break
		}
	}

	categoryForm := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select a category").
				Options(
					huh.NewOption("Cheap", "cheap"),
					huh.NewOption("Medium", "medium"),
					huh.NewOption("Expensive", "expensive"),
				).
				Value(&category),
		),
	)

	err := categoryForm.Run()
	if err != nil {
		log.Fatal(err)
	}

	newDate := DateIdea{Name: dateName, Category: category}
	switch category {
	case "expensive":
		ExpensiveDates = append(ExpensiveDates, newDate)
	case "medium":
		MediumDates = append(MediumDates, newDate)
	case "cheap":
		CheapestDates = append(CheapestDates, newDate)
	}

	fmt.Println("\n✅ Date idea saved successfully!")
}

func randomDate(dateList *[]DateIdea, category string) {
	if len(*dateList) == 0 {
		fmt.Printf("\n⚠️ No %s date ideas saved yet! Try creating one first!\n", category)
		fmt.Println("Press Enter to return to the main menu...")
		fmt.Scanln()
		return
	}

	rand.Seed(time.Now().UnixNano())
	selectedDate := (*dateList)[rand.Intn(len(*dateList))]

	fmt.Printf("\n🎉 Your random %s date idea: %s\n", category, selectedDate.Name)
	fmt.Println("\nPress Enter to return to the main menu...")
	fmt.Scanln()
}

func viewDateIdeas() {
	if len(ExpensiveDates) == 0 && len(MediumDates) == 0 && len(CheapestDates) == 0 {
		fmt.Println("\n⚠️ No date ideas saved yet! Try creating one first!")
		fmt.Println("Press Enter to return to the main menu...")
		fmt.Scanln()
		return
	}

	fmt.Println("\n=== All Date Ideas ===")

	printDates("💎 Expensive Dates:", &ExpensiveDates)
	printDates("🎯 Medium Dates:", &MediumDates)
	printDates("💰 Cheap Dates:", &CheapestDates)

	fmt.Println("\nPress Enter to return to the main menu...")
	fmt.Scanln()
}

func printDates(title string, dateList *[]DateIdea) {
	if len(*dateList) == 0 {
		return
	}

	fmt.Println("\n" + title)
	for i, date := range *dateList {
		fmt.Printf("%d. %s\n", i+1, date.Name)
	}
}
