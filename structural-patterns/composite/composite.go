// 组合模式（Composite Pattern），又叫部分整体模式，是用于把一组相似的对象当作一个单一的对象。
// 组合模式依据树形结构来组合对象，用来表示部分以及整体层次

package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type Employee struct {
	name string
	dept string
	salary int
	subordinates []*Employee
}

func (e *Employee) Add(sub *Employee){
	e.subordinates = append(e.subordinates, sub)
}

func (e *Employee) GetSubordinates() []*Employee{
	return e.subordinates
}

func (e *Employee) String() string{
	var buf bytes.Buffer
	buf.WriteString("Employee:[ name:")
	buf.WriteString(e.name+", ")
	buf.WriteString("dept:")
	buf.WriteString(e.dept+", ")
	buf.WriteString("salary:")
	buf.WriteString(strconv.Itoa(e.salary)+" ]")
	return buf.String()
}

func main(){
	ceo := &Employee{"John","CEO", 3000, make([]*Employee, 0)}

	cto := &Employee{"Robert","CTO",0000, make([]*Employee, 0)}

	coo := &Employee{"Michel","COO", 200000, make([]*Employee, 0)}

	rd1 := &Employee{"Laura","RD", 18000, nil}
	rd2 := &Employee{"Bob","RD", 17000, nil}

	op1 := &Employee{"tim","OP", 18000, nil}
	op2 := &Employee{"jack","OP", 18000, nil}

	ceo.Add(cto)
	ceo.Add(coo)

	cto.Add(rd1)
	cto.Add(rd2)

	coo.Add(op1)
	coo.Add(op2)

	//打印该公司所有员工
	fmt.Println(ceo.String())
	for _, headEmployee := range ceo.GetSubordinates(){
		fmt.Println(headEmployee.String())
		for _, employee := range  headEmployee.GetSubordinates(){
			fmt.Println(employee.String())
		}
	}

}



