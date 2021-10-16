package sort

type Person struct {
	Year int
	Name string
}

// CountingSort sorts Persons by year
func CountingSort(persons []Person) []Person {
	count := make([][]Person, 200) // from 1900 to 2099

	for _, person := range persons {
		count[person.Year-1900] = append(count[person.Year-1900], person)
	}

	i := 0
	for _, personsPerYear := range count {
		for _, person := range personsPerYear {
			persons[i] = person
			i++
		}
	}

	return persons
}
