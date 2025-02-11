package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/charmbracelet/huh"
)

type DateIdea struct {
	Name     string `json:"name"`
	Category string `json:"category"`
}

type DateStorage struct {
	ExpensiveDates []DateIdea `json:"expensive_dates"`
	MediumDates    []DateIdea `json:"medium_dates"`
	CheapestDates  []DateIdea `json:"cheapest_dates"`
}

var (
	ExpensiveDates []DateIdea
	MediumDates    []DateIdea
	CheapestDates  []DateIdea
)

const dateFile = "dates.json"

func main() {
	loadDates()
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
			saveDates()
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

	saveDates()
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

func loadDates() {
	if _, err := os.Stat(dateFile); os.IsNotExist(err) {
		return
	}

	jsonData, err := ioutil.ReadFile(dateFile)
	if err != nil {
		log.Fatal("Error reading JSON file:", err)
	}

	var data DateStorage
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Fatal("Error unmarshalling JSON:", err)
	}

	ExpensiveDates = data.ExpensiveDates
	MediumDates = data.MediumDates
	CheapestDates = data.CheapestDates
}

func saveDates() {
	data := DateStorage{
		ExpensiveDates: ExpensiveDates,
		MediumDates:    MediumDates,
		CheapestDates:  CheapestDates,
	}
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal("Error marshalling JSON:", err)
	}
	err = ioutil.WriteFile(dateFile, jsonData, 0644)
	if err != nil {
		log.Fatal("Error writing JSON file:", err)
	}
}

func viewDateIdeas() {
	if _, err := os.Stat(dateFile); os.IsNotExist(err) {
		fmt.Println("\n⚠️ No date ideas saved yet! Try creating one first!")
		fmt.Println("Press Enter to return to the main menu...")
		fmt.Scanln()
		return
	}

	jsonData, err := ioutil.ReadFile(dateFile)
	if err != nil {
		log.Fatal("Error reading JSON file:", err)
	}

	var data DateStorage
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Fatal("Error unmarshalling JSON:", err)
	}

	fmt.Println("\n=== All Date Ideas ===")
	printDates("💎 Expensive Dates:", data.ExpensiveDates)
	printDates("🎯 Medium Dates:", data.MediumDates)
	printDates("💰 Cheap Dates:", data.CheapestDates)

	fmt.Println("\nPress Enter to return to the main menu...")
	fmt.Scanln()
}

func printDates(title string, dates []DateIdea) {
	if len(dates) == 0 {
		fmt.Println(title, "No date ideas available.")
		return
	}

	fmt.Println(title)
	for _, date := range dates {
		fmt.Println(" -", date.Name)
	}
}
