package  main

import "fmt"

const MAX int = 999

type Participant struct {
	id       int
	name     string
	age      int
	interest string
	date     string
	active   bool
}

var participantData [MAX]Participant
var total int
var art, academic, technology int

func addParticipant() { //this function is used to add a new participant to the participantData array.
//  It will ask the user to input the participant's ID, name, age, interest category, and registration date.
//  The function will also update the total number of participants and count the number of participants in each category.
	var p Participant
	var category int
	var choice int

	fmt.Println("---- ADD PARTICIPANT ----")

	fmt.Print("ID Registration (6 digits) : ")
	fmt.Scan(&p.id)

	fmt.Print("Full Name : ")
	fmt.Scan(&p.name)

	fmt.Print("Age : ")
	fmt.Scan(&p.age)

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

		switch choice { 
			case 1 :
			p.interest = "Math"
			case 2 :
			p.interest = "Physics"
			case 3 :
			p.interest = "Biology"
			case 4 :
			p.interest = "Chemistry"
			case 5 :
			p.interest = "Economy"
			case 6 :
			p.interest = "History"
			case 7 :
			p.interest = "Sociology"
			case 8 :
			p.interest = "English"
		}
		

	} else if category == 2 {

		fmt.Println(" ~ Technology ~ ")
		fmt.Println("1. Coding")
		fmt.Println("2. Graphic Design")
		fmt.Println("3. Data & Office")
		fmt.Print("Choose Interest : ")
		fmt.Scan(&choice)

		switch choice {
			case 1 : //pakekan switch case (nested pun bisa)
			p.interest = "Coding"
			case 2 :
			p.interest = "Graphic Design"
			case 3 :
			p.interest = "Data & Office"
		}

		

	} else if category == 3 {

		fmt.Println(" ~ Art / Creativity ~ ")
		fmt.Println("1. Music")
		fmt.Println("2. Art")
		fmt.Println("3. Photography")
		fmt.Println("4. Cinematography")
		fmt.Print("Choose Interest : ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			p.interest = "Music"
		case 2:
			p.interest = "Art"
		case 3:
			p.interest = "Photography"
		case 4:
			p.interest = "Cinematography"
		}
		
		
	}

	fmt.Print("Registration Date (D/M/Y) : ")
	fmt.Scan(&p.date)

	p.active = true

	participantData[total] = p
	total = total + 1
	countCategory()
	fmt.Println("Participant successfully added")
}


func countCategory() { //masih belum siap (experimnent aja dulu)
	var art, academic, technology, choice int
	
	if  choice == 1 {
		academic = academic + 1
	}else if choice == 2 {
		technology = technology + 1
	}else if choice == 3 {
		art = art + 1
	}
}	//ini masalahnya karena variabel choice itu cuma ada di fungsi addParticipant, jadi dia ga bisa diakses di fungsi countCategory. Solusinya, kita bisa buat variabel global untuk menyimpan pilihan kategori, atau kita bisa langsung update countCategory di dalam fungsi addParticipant setelah kita menentukan kategori yang dipilih.


func viewParticipants() {
	//this function is used to display the data of all participants in the participantData array.
	var i int
	var enter string

	fmt.Println("---- PARTICIPANT DATA ----")

	if total == 0 {

		fmt.Println("No data available")

	} else {

		for i = 0; i < total; i++ {

			fmt.Println("-------------------")
			fmt.Println("ID       :", participantData[i].id)
			fmt.Println("Name     :", participantData[i].name)
			fmt.Println("Age      :", participantData[i].age)
			fmt.Println("Interest :", participantData[i].interest)
			fmt.Println("Registration Date :", participantData[i].date)
		}
	}

	fmt.Println()
	fmt.Print("Type EXIT to continue : ")
	fmt.Scan(&enter)
}

func updateParticipant() {
	//this function is used to update the data of a participant in the participantData array.
	var idSearch int
	var i int
	var found bool

	fmt.Println("---- UPDATE PARTICIPANT ----")

	fmt.Print("ID Registration : ")
	fmt.Scan(&idSearch)

	found = false
	i = 0
	for i < total {

		if participantData[i].id == idSearch {

			fmt.Print("New Age : ")
			fmt.Scan(&participantData[i].age)

			fmt.Print("New Interest : ")
			fmt.Scan(&participantData[i].interest)

			fmt.Print("New Date (D/M/Y) : ")
			fmt.Scan(&participantData[i].date)

			found = true
		}
		i++
	}

	if found == true {

		fmt.Println("Participant successfully updated")

	} else {

		fmt.Println("Participant not found")
	}
}

func deleteParticipant() {
	//this function is used to delete a participant from the participantData array.
	var idSearch int
	var i int
	var j int
	var found bool

	fmt.Println("---- DELETE PARTICIPANT ----")

	fmt.Print("ID Registration : ")
	fmt.Scan(&idSearch)

	found = false
	i = 0
	for i < total {

		if participantData[i].id == idSearch {

			for j = i; j < total-1; j++ {

				participantData[j] = participantData[j+1]
			}

			total = total - 1
			found = true
		}
		i++
	}

	if found == true {

		fmt.Println("Participant successfully deleted")

	} else {

		fmt.Println("ID not found")
	}
}

func sequentialSearch() {
	//this function is used to search for a participant in the participantData array using the sequential search algorithm.
	var idSearch int
	var i int
	var found bool

	fmt.Println("---- SEQUENTIAL SEARCH ----")

	fmt.Print("ID Registration : ")
	fmt.Scan(&idSearch)

	found = false

	for i = 0; i < total; i++ {

		if participantData[i].id == idSearch {

			fmt.Println("Data Found!")
			fmt.Println("ID       :", participantData[i].id)
			fmt.Println("Name     :", participantData[i].name)
			fmt.Println("Age      :", participantData[i].age)
			fmt.Println("Interest :", participantData[i].interest)

			found = true
		}
	}

	if found == false {

		fmt.Println("ID not found")
	}
}

func binarySearch() {
	//this function is used to search for a participant in the participantData array using the binary search algorithm.
	var nameSearch string
	var left int
	var right int
	var mid int
	var found bool

	insertionSortName()

	fmt.Println("---- BINARY SEARCH ----")

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
			fmt.Println("Age      :", participantData[mid].age)
			fmt.Println("Interest :", participantData[mid].interest)

			found = true
			left = right + 1

		} else if participantData[mid].name > nameSearch {

			left = mid + 1

		} else {

			right = mid - 1
		}
	}

	if found == false {

		fmt.Println("Data not found")
	}
}

func selectionSortID() {
	//this function is used to sort the participants in the participantData array by their ID using the selection sort algorithm
	//  in increasing order.
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

func statistics() {
	//this function is used to display the statistics of the participants in the participantData array,
	//  such as the number of participants in each category and the total number of active participants.
	var i int
	var totalActive int

	for i = 0; i < total; i++ {

		
		if participantData[i].active == true {

			totalActive = totalActive + 1
		}
	}
	
	fmt.Println("---- STATISTICS ----")
	fmt.Println("Academic    :", academic)
	fmt.Println("Technology  :", technology)
	fmt.Println("Arts/Creativity :", art)
	fmt.Println("Total Active:", totalActive) 
}
func searchMenu() {
	//this function is used to display the search menu and call the appropriate search function based on the user's choice.
	var choose int
	var enter string

	fmt.Println("---- SEARCH DATA ----")
	fmt.Println("1. Sequential Search")
	fmt.Println("2. Binary Search")
	fmt.Print("Choose : ")
	fmt.Scan(&choose)

	if choose == 1 {

		sequentialSearch()

	} else if choose == 2 {

		binarySearch()

	} else {

		fmt.Println("Menu not found")
	}

	fmt.Println()
	fmt.Print("Type EXIT to continue : ")
	fmt.Scan(&enter)
}

func sortMenu() { 
	//this function is used to display the sort menu and call the appropriate sort function based on the user's choice.
	var choose int
	var enter string

	fmt.Println("---- SORT DATA ----")
	fmt.Println("1. Selection Sort ID")
	fmt.Println("2. Insertion Sort Name")
	fmt.Print("Choose : ")
	fmt.Scan(&choose)

	if choose == 1 {

		selectionSortID()

	} else if choose == 2 {

		insertionSortName()

	} else {

		fmt.Println("Menu not found")
	}

	fmt.Println()
	fmt.Print("Type EXIT to continue : ")
	fmt.Scan(&enter)
}

func menu() {

	fmt.Println()
	fmt.Println("===== KURSUS IN =====")
	fmt.Println("1. Add Participant")
	fmt.Println("2. View Participants")
	fmt.Println("3. Update Participant")
	fmt.Println("4. Delete Participant")
	fmt.Println("5. Search Data")
	fmt.Println("6. Sort Data")
	fmt.Println("7. Statistics")
	fmt.Println("0. Exit")
}

func main() {
	var choose int

	choose = -1

	for choose != 0 {

		menu()

		fmt.Print("Choose Menu : ")
		fmt.Scan(&choose)

		switch choose {
		case 1:

			addParticipant()
		case 2:

			viewParticipants()
		case 3:

			updateParticipant()

		case 4:
			
			deleteParticipant()

		case 5:

			searchMenu()

		case 6:

			sortMenu()

		case 7:

			statistics()

		case 0:

			fmt.Println("Bye-byeee. See you soonn~~")

		default:
			fmt.Println("Menu not found")
		}
	}
}