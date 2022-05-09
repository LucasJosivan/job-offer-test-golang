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
	lines := strings.Split(csvData, "\n")
	names := strings.Split(lines[0], ",")
	firstValue := strings.Split(lines[1], ",")
	secondValue := strings.Split(lines[2], ",")
	var pList = People{}
	for i := 0; i < len(names); i++ {
		pList = append(pList, Person{
			Name:        names[i],
			FistValue:   firstValue[i],
			SecondValue: secondValue[i],
		})
	}
	sort.Sort(People(pList))
	result := fmt.Sprintf(strings.Join(pList.rangeName(), ",") + "\n")
	result += fmt.Sprintf(strings.Join(pList.rangeFirstValue(), ",") + "\n")
	result += fmt.Sprint(strings.Join(pList.rangeSecondValue(), ","))
	return result
}

func (p People) rangeName() (names []string) {
	for _, person := range p {
		names = append(names, person.Name)
	}
	return
}

func (p People) rangeFirstValue() (values []string) {
	for _, person := range p {
		values = append(values, person.FistValue)
	}
	return
}

func (p People) rangeSecondValue() (values []string) {
	for _, person := range p {
		values = append(values, person.SecondValue)
	}
	return
}
