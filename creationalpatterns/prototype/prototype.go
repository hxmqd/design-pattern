// 原型模式：用原型实例指定创建对象的种类，并且通过拷贝这些原型创建新的对象

package main

import "fmt"

// 个人信息
type PersonalInfo struct {
	name string
	sex string
	age string
}

// 工作经历
type WorkExperience struct {
	timeArea string
	company string
}

// 简历
type Resume struct {
	PersonalInfo
	WorkExperience
}

func (r *Resume) SetPersonalInfo(name string, sex string, age string) {
	r.name = name
	r.sex = sex
	r.age = age
}

func (r *Resume) SetWorkExperience(timeArea string, company string) {
	r.timeArea = timeArea
	r.company = company
}

func (r *Resume) Display(){
	fmt.Println("个人信息:",r.name, r.sex, r.age)
	fmt.Println("工作经历:", r.timeArea, r.company)
}

func (r *Resume) clone() *Resume{
	resume := *r
	return &resume
}

func main(){
	r1 := &Resume{}
	r1.SetPersonalInfo("tim", "男", "23")
	r1.SetWorkExperience("2012-2015", "北京嘀嘀无限科技发展有限公司")

	r2 := r1.clone()
	r2.SetWorkExperience("2012-2015", "北京百度科技有限公司")

	r3 := r1.clone()
	r3.SetPersonalInfo("jack", "女", "25")

	r1.Display()
	r2.Display()
	r3.Display()
	}



