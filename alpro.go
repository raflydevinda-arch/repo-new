package main

import "fmt"

const MAX int = 999

type Participant struct {
	id       int
	name     string
	interest string
	date     string
	active   bool
}

var participantData [MAX]Participant
var total int

func addParticipant() {
	var p Participant
	var category int
	var choice int

	fmt.Println("---- ADD PARTICIPANT ----")

	fmt.Print("ID Registration : ")
	fmt.Scan(&p.id)

	fmt.Print("Full Name : ")
	fmt.Scan(&p.name)

	fmt.Println(" ~ Interest Category ~")
	fmt.Println("1. Academic")
	fmt.Println("2. Technology")
	fmt.Println("3. Art / Creativity")
	fmt.Print("Choose Category : ")
	fmt.Scan(&category)

	if category == 1 {

		fmt.Println(" ~ Academic ~ ")
		fmt.Println("1. Math")
		fmt.Println("2. Physics")
		fmt.Println("3. Biology")
		fmt.Println("4. Chemistry")
		fmt.Println("5. Economy")
		fmt.Println("6. History")
		fmt.Println("7. Sociology")
		fmt.Println("8. English")
		fmt.Print("Choose Interest : ")
		fmt.Scan(&choice)

		if choice == 1 {
			p.interest = "Math"
		} else if choice == 2 {
			p.interest = "Physics"
		} else if choice == 3 {
			p.interest = "Biology"
		} else if choice == 4 {
			p.interest = "Chemistry"
		} else if choice == 5 {
			p.interest = "Economy"
		} else if choice == 6 {
			p.interest = "History"
		} else if choice == 7 {
			p.interest = "Sociology"
		} else if choice == 8 {
			p.interest = "English"
		}

	} else if category == 2 {

		fmt.Println(" ~ Technology ~ ")
		fmt.Println("1. Coding")
		fmt.Println("2. Graphic Design")
		fmt.Println("3. Data & Office")
		fmt.Print("Choose Interest : ")
		fmt.Scan(&choice)

		if choice == 1 {
			p.interest = "Coding"
		} else if choice == 2 {
			p.interest = "Graphic Design"
		} else if choice == 3 {
			p.interest = "Data & Office"
		}

	} else if category == 3 {

		fmt.Println("=== Art / Creativity ===")
		fmt.Println("1. Music")
		fmt.Println("2. Art")
		fmt.Println("3. Photography")
		fmt.Println("4. Cinematography")
		fmt.Print("Choose Interest : ")
		fmt.Scan(&choice)

		if choice == 1 {
			p.interest = "Music"
		} else if choice == 2 {
			p.interest = "Art"
		} else if choice == 3 {
			p.interest = "Photography"
		} else if choice == 4 {
			p.interest = "Cinematography"
		}
	}

	fmt.Print("Registration Date (D/M/Y) : ")
	fmt.Scan(&p.date)

	p.active = true

	participantData[total] = p

	total = total + 1

	fmt.Println("Participant successfully added")
}

func viewParticipants() {
	var i int
	var enter string

	fmt.Println("---- Participant Data ----")

	if total == 0 {

		fmt.Println("No data available")

	} else {

		for i = 0; i < total; i++ {

			fmt.Println("-------------------")
			fmt.Println("ID       :", participantData[i].id)
			fmt.Println("Name     :", participantData[i].name)
			fmt.Println("Interest :", participantData[i].interest)
			fmt.Println("Registration Date (D/M/Y) :", participantData[i].date)
		}
	}

	fmt.Println()
	fmt.Print("Type EXIT to continue : ")
	fmt.Scan(&enter)
}

func changeParticipantStatus() {
	var idSearch int
	var i int
	var found bool
	fmt.Println("---- Change Participant Status ----")

	fmt.Print("ID Registration : ")
	fmt.Scan(&idSearch)

	found = false

	for i = 0; i < total; i++ {
		if participantData[i].id == idSearch {

			fmt.Print("New Name :")
			fmt.Scan(&participantData[i].name)

			fmt.Print("New Interest :")
			fmt.Scan(&participantData[i].interest)

			fmt.Print("New Date (D/M/Y) :")
			fmt.Scan(&participantData[i].date)

			found = true
		}
	}
	if found == true {
		fmt.Println("Participant not found")
	}
}

func deleteParticipant() {
	var idSearch int
	var i int
	var j int
	var found bool

	fmt.Println("---- Delete Participant ----")

	fmt.Print("ID Registration : ")
	fmt.Scan(&idSearch)

	found = false

	for i = 0; i < total; i++ {
		if participantData[i].id == idSearch {
			for j = i; j < total-1; j++ {
				participantData[j] = participantData[j+1]
			}
			total = total - 1
			found = true
		}
	}

	if found == true {
		fmt.Println("Participant successfully deleted")
	} else {
		fmt.Println("ID not found")
	}
}

func sequentialSearch() {
	var idSearch int
	var i int
	var found bool

	fmt.Println("=== Search Participant ===")

	fmt.Print("ID Registration : ")
	fmt.Scan(&idSearch)

	found = false

	for i = 0; i < total; i++ {
		if participantData[i].id == idSearch {
			fmt.Println("Data Found!")
			fmt.Println("ID       :", participantData[i].id)
			fmt.Println("Name     :", participantData[i].name)
			fmt.Println("Interest :", participantData[i].interest)

			found = true
		}
	}

	if found == false {
		fmt.Println("ID not found")
	}
}

func selectionSortID() {
	var i int
	var j int
	var minIndex int
	var temp Participant

	for i = 0; i < total-1; i++ {
		minIndex = i
		for j = i + 1; j < total; j++ {
			if participantData[j].id < participantData[minIndex].id {
				minIndex = j
			}
		}
		temp = participantData[i]
		participantData[i] = participantData[minIndex]
		participantData[minIndex] = temp
	}

	fmt.Println("Participants sorted by ID")
}

func insertionSortName() {
	var i int
	var j int
	var key Participant

	for i = 1; i < total; i++ {
		key = participantData[i]
		j = i - 1
		for j >= 0 && participantData[j].name > key.name {
			participantData[j+1] = participantData[j]
			j = j - 1
		}
		participantData[j+1] = key
	}
	fmt.Println("Participants sorted by Name")
}

func binarySearch() {
	var nameSearch string
	var left int
	var right int
	var mid int
	var found bool

	insertionSortName()

	fmt.Println("=== Binary Search ===")
	fmt.Print("Search Name : ")
	fmt.Scan(&nameSearch)

	left = 0
	right = total - 1
	found = false

	for left <= right {
		mid = (left + right) / 2
		if participantData[mid].name == nameSearch {
			fmt.Println("Data Found!")
			fmt.Println("ID       :", participantData[mid].id)
			fmt.Println("Name     :", participantData[mid].name)
			fmt.Println("Interest :", participantData[mid].interest)

			found = true
			left = right + 1
		} else if participantData[mid].name < nameSearch {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if found == false {
		fmt.Println("Data not found")
	}
}

func statistics() {
	var i int
	var design int
	var coding int
	var business int
	var totalActive int

	for i = 0; i < total; i++ {

		if participantData[i].interest == "Desain" {

			design = design + 1
		}

		if participantData[i].interest == "Coding" {

			coding = coding + 1
		}

		if participantData[i].interest == "Bisnis" {

			business = business + 1
		}

		if participantData[i].active == true {

			totalActive = totalActive + 1
		}
	}

	fmt.Println("=== Statistik ===")
	fmt.Println("Desain      :", design)
	fmt.Println("Coding      :", coding)
	fmt.Println("Bisnis      :", business)
	fmt.Println("Total Aktif :", totalActive)
}

func menu() {

	fmt.Println()
	fmt.Println("===== KURSUS IN =====")
	fmt.Println("1. Add Participant")
	fmt.Println("2. View Participants")
	fmt.Println("3. Update Participant")
	fmt.Println("4. Delete Participant")
	fmt.Println("5. Sequential Search")
	fmt.Println("6. Binary Search")
	fmt.Println("7. Selection Sort ID")
	fmt.Println("8. Insertion Sort Name")
	fmt.Println("9. Statistics")
	fmt.Println("0. Exit")
}

func main() {
	var choose int

	choose = -1

	for choose != 0 {

		menu()

		fmt.Print("Choose Menu : ")
		fmt.Scan(&choose)

		if choose == 1 {

			addParticipant()

		} else if choose == 2 {

			viewParticipants()

		} else if choose == 3 {

			changeParticipantStatus()

		} else if choose == 4 {

			deleteParticipant()

		} else if choose == 5 {

			sequentialSearch()

		} else if choose == 6 {

			binarySearch()

		} else if choose == 7 {

			selectionSortID()

		} else if choose == 8 {

			insertionSortName()

		} else if choose == 9 {

			statistics()

		} else if choose == 0 {

			fmt.Println("DONE")

		} else {

			fmt.Println("Menu not found")
		}
	}
}
