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
						huh.NewOption("Exit", "exit"),
					).
					Value(&mainChoice),
			),
		)

		// Run the form
		err := form.Run()
		if err != nil {
			log.Fatal(err)
		}
		//test commit
		// Handle menu selection
		switch mainChoice {
		case "create":
			createDate()
		case "view":
			viewDateIdeas()
		case "randomExpensive":
			randomExpensive()
		case "exit":
			fmt.Println("Exiting... Have a great date! ❤️")
			return
		}
	}
}

func randomExpensive() {
	var expensiveCategory bool
	randomExpensiveForm := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().Title("Random Expensive Date").
				Value(&expensiveCategory),
		),
	)
	err := randomExpensiveForm.Run()
	if err != nil {
		log.Fatal(err)
	}

	if expensiveCategory {
		if len(ExpensiveDates) == 0 {
			fmt.Println("\n⚠️ No expensive date ideas saved yet! Try creating one first!")
			fmt.Println("Press Enter to return to the main menu...")
			fmt.Scanln()
			return
		}

		// Ensure randomness
		rand.Seed(time.Now().UnixNano())

		randomIndex := rand.Intn(len(ExpensiveDates))
		selectedDate := ExpensiveDates[randomIndex]

		// Display the chosen date
		fmt.Printf("\n💎 Your random expensive date idea: %s\n", selectedDate.Name)
		fmt.Println("\nPress Enter to return to the main menu...")
		fmt.Scanln()
	}
}

// Function to create a new date
func createDate() {
	var dateName string
	var category string

	// Ask for the date name
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

	// Ask for the category
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

	// Save the date idea to the slice
	newDate := DateIdea{Name: dateName, Category: category}
	if category == "expensive" {
		ExpensiveDates = append(ExpensiveDates, newDate)
	} else if category == "cheap" {
		CheapestDates = append(CheapestDates, newDate)
	} else if category == "average" {
		MediumDates = append(MediumDates, newDate)
	}

	fmt.Println("\n✅ Date idea saved successfully!")

	// Ask if the user wants to create another
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

// Function to view all saved date ideas
func viewDateIdeas() {
	if len(ExpensiveDates) == 0 && len(MediumDates) == 0 && len(CheapestDates) == 0 {
		fmt.Println("\n⚠️ No date ideas saved yet! Try creating one first!")
		fmt.Println("Press Enter to return to the main menu...")
		fmt.Scanln()
		return
	}

	fmt.Println("\n=== All Date Ideas ===")

	// Print Expensive Dates
	if len(ExpensiveDates) > 0 {
		fmt.Println("\n💎 Expensive Dates:")
		for i, date := range ExpensiveDates {
			fmt.Printf("%d. %s\n", i+1, date.Name)
		}
	}

	// Print Medium Dates
	if len(MediumDates) > 0 {
		fmt.Println("\n🎯 Medium Dates:")
		for i, date := range MediumDates {
			fmt.Printf("%d. %s\n", i+1, date.Name)
		}
	}

	// Print Cheap Dates
	if len(CheapestDates) > 0 {
		fmt.Println("\n💰 Cheap Dates:")
		for i, date := range CheapestDates {
			fmt.Printf("%d. %s\n", i+1, date.Name)
		}
	}

	fmt.Println("\nPress Enter to return to the main menu...")
	fmt.Scanln()
}
