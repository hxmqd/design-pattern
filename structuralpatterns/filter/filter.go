// 过滤器模式,允许开发人员使用不同的标准来过滤一组对象，通过逻辑运算以解耦的方式把它们连接起来
package main

import "fmt"

// 人员
type Person struct {
	name string
	gender string
	maritalStatus string
}

func (p *Person) GetName() string{
	return p.name
}

func (p *Person) GetGender() string{
	return p.gender
}

func (p *Person) GetMaritalStatus() string{
	return p.maritalStatus
}

func NewPerson(name, gender, maritalStatus string) *Person{
	person := Person{
		name : name,
		gender : gender,
		maritalStatus: maritalStatus,
	}
	return &person
}

// 标准
type Criteria interface {
	MeetCriteria(persons []Person) []Person
}

type CriteriaMale struct {
}

func (c *CriteriaMale) MeetCriteria(persons []Person) []Person{
	malePersons := make([]Person, 0)
	for _, person := range persons {
		if person.GetGender() == "Male" {
			malePersons = append(malePersons, person)
		}
	}
	return malePersons
}

type CriteriaFemale struct {
}
func (c *CriteriaFemale) MeetCriteria(persons []Person) []Person{
	femalePersons := make([]Person, 0)
	for _, person := range persons {
		if person.GetGender() == "Female" {
			femalePersons = append(femalePersons, person)
		}
	}
	return femalePersons
}

type CriteriaSingle struct {
}

func (c *CriteriaSingle) MeetCriteria(persons []Person) []Person {
	singlePersons := make([]Person, 0)
	for _, person := range persons {
		if person.GetMaritalStatus() == "Single" {
			singlePersons = append(singlePersons, person)
		}
	}
	return singlePersons
}

func main(){
	persons := make([]Person, 0)
	persons = append(persons, *NewPerson("Robert","Male", "Single"),
		*NewPerson("John","Male", "Married"),
		*NewPerson("Laura","Female", "Married"),
		*NewPerson("Diana","Female", "Single"),
		)

	var male CriteriaMale
	fmt.Println("males:")
	fmt.Println(male.MeetCriteria(persons))

	var female CriteriaFemale
	fmt.Println("females:")
	fmt.Println(female.MeetCriteria(persons))

	var single CriteriaSingle
	fmt.Println("Single:")
	fmt.Println(single.MeetCriteria(persons))
}