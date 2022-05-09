package sort_string_csv

import (
	"fmt"
	"sort"
	"strings"
)

type People []Person

type Person struct {
	Name        string
	FistValue   string
	SecondValue string
}

func (p People) Len() int {
	return len(p)
}

func (p People) Less(i, j int) bool {
	return p[i].Name < p[j].Name
}

func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func SortCsvColumns(csvData string) string {
	people := strings.Split(csvData, "\n")
	var names []string
	var firstValue []string
	var secondValue []string
	var pList = People{}
	for i := 0; i < len(people); i++ {
		names = strings.Split(people[0], ",")
		firstValue = strings.Split(people[1], ",")
		secondValue = strings.Split(people[2], ",")
	}
	for i := 0; i < len(names); i++ {
		pList = append(pList, Person{
			Name:        names[i],
			FistValue:   firstValue[i],
			SecondValue: secondValue[i],
		})
	}
	sort.Sort(People(pList))
	result := pList.rangeName()
	result += pList.rangeFirstValue()
	result += pList.rangeSecondValue()
	return result
}

func (p People) rangeName() (names string) {
	for n, person := range p {
		if (n + 1) == len(p) {
			names += fmt.Sprintf("%v\n", person.Name)
			return
		}
		names += fmt.Sprintf("%v,", person.Name)
	}
	return
}

func (p People) rangeFirstValue() (values string) {
	for n, person := range p {
		if (n + 1) == len(p) {
			values += fmt.Sprintf("%v\n", person.FistValue)
			return
		}
		values += fmt.Sprintf("%v,", person.FistValue)
	}
	return
}

func (p People) rangeSecondValue() (values string) {
	for n, person := range p {
		if (n + 1) == len(p) {
			values += fmt.Sprintf("%v", person.SecondValue)
			return
		}
		values += fmt.Sprintf("%v,", person.SecondValue)
	}
	return
}
