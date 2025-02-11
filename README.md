# Date Randomizer

## Description

A CLI-based **Date Idea Randomizer** built with Go and the [Charmbracelet Huh](https://github.com/charmbracelet/huh) library. This program allows users to create, view, delete, and randomly select date ideas based on cost categories (**Expensive, Medium, Cheap**). Data is stored in a `dates.json` file for persistence. I made this for me and my partner to solve the issue of great date ideas, but poor recollection and indecisiveness for what dates to go on!

## Features

-   📌 **Create** a new date idea and categorize it.

-   📜 **View** all saved date ideas.

-   ❌ **Delete** date ideas with confirmation.

-   🎲 **Randomly select** a date idea from a specific category.

-   💾 **Persistent storage** via JSON file.

-   🛑 **Exit** and save progress automatically.


## Installation & Usage

### Prerequisites

-   Go 1.18+


### Clone the Repository

```
git clone https://github.com/yourusername/date-randomizer.git
cd date-randomizer
```

### Install Dependencies

```
go mod tidy
```

### Run the Program

```
go run main.go
```

## Menu Options

1.  **Create a New Date** - Enter a date idea and select a cost category.

2.  **View All Date Ideas** - Display all saved date ideas by category.

3.  **Delete a Date** - Choose and confirm deletion of a saved date idea.

4.  **Random Expensive Date** - Get a surprise from your expensive list.

5.  **Random Medium Date** - Get a random mid-range date.

6.  **Random Cheap Date** - Find a budget-friendly date idea.

7.  **Exit** - Save changes and close the program.


## File Structure

```
📂 date-randomizer
 ├── 📄 main.go        # Main application logic
 ├── 📄 dates.json     # Persistent storage for date ideas
 ├── 📄 go.mod         # Go module file
 ├── 📄 README.md      # Project documentation
```

