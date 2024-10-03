package main

import (
	"fmt"
	"strconv"
	"strings"
)
func  main() {

	addStudent(Student{id: 1, name: "Nguyen Dang", age: 34, sex: "Male", sub: "GoLang"})
	addStudent(Student{id: 2, name: "Thanh Thư", age: 62, sex: "Female", sub: "English"})
	addStudent(Student{id: 3, name: "Thanh Nga", age: 64, sex: "Female", sub: "English"})
						
	 showListStudent()
	 line()
	 findStudent(1)
	 line()
	 editStudent(2 , Student{id: 2, name: "Thanh Thơi", age: 64, sex: "Nữ", sub: "English"})
	 line()
	 deleteStudent(2)
	 line()
	 showListStudent()
	 line()
	 findStudentBySex("Female")
	 line()
	 
	 // Interger to Roman
	num := 1987
	roman := intToRoman(num)
	fmt.Printf("The Roman numeral for %d is %s\n", num, roman)

	// kiểm tra số đối xứng
	r := isPalindrome(121) 
	fmt.Println(r)

	// kiểm tra tiền tố dài nhất
	strs := []string{"flower", "flow", "flight"}
	fmt.Println("Longest Common Prefix:", longestCommonPrefix(strs))
}
type Student struct{
	id int
	name string
	age int
	sex string 
	sub string
}
var listStudent =make(map[int]Student)
func addStudent(std Student)  {
	listStudent [std.id] = std
	fmt.Println("Add Successfull")
}
func showListStudent()  {
	fmt.Println("List Student")
	for _, st := range listStudent {
		fmt.Printf(" ID: %d\n Name: %s\n Age: %d\n Sex: %s\n Subject: %s\n", st.id, st.name, st.age, st.sex, st.sub)
	
	}
}
func findStudent(id int)  {
	if st , found := listStudent[id] ; found{
		fmt.Printf("Thông tin sinh viên: %+v\n", st)
	}
}
func editStudent(id int ,st Student )  {
	if _, found := listStudent[id] ; found{
		listStudent[id] = st
		fmt.Printf("Edit Successfull\n ID: %d\n Name: %s\n Age: %d\n Sex: %s\n Subject: %s\n", st.id, st.name, st.age, st.sex, st.sub )
		}else{
		fmt.Println("Can not find student")
	}
}
func deleteStudent(id int)  {
	if _, found := listStudent[id] ; found{
		delete(listStudent,id)
		fmt.Println("Delete Successfull!")
	} else {
		fmt.Println("Can not find student")
	}
}

func findStudentBySex(sex string )  {
	fmt.Printf("List student by sex %s:\n", sex)
	for _, st := range listStudent {
		if strings.ToLower(st.sex) == strings.ToLower(sex){
			fmt.Printf("%+v\n", st)
			
		}
	}
}
func intToRoman(num int) string {
	// Các cặp giá trị và ký hiệu La Mã
	vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var result strings.Builder

	for i := 0; num > 0; i++ {
		// Tìm ký hiệu La Mã phù hợp với giá trị hiện tại
		for num >= vals[i] {
			result.WriteString(romans[i])
			num -= vals[i]
		}
	}

	return result.String()
}

func isPalindrome(num int) bool  {
	if num < 0 {
		return false
	}
	s := strconv.Itoa(num)
	n := len(s)
	for i:=0 ; i < n/2 ; i++ {
		if s[i] != s[n-1-i]{
			return false
		}
	}
	return true
}

func longestCommonPrefix(strs []string) string {
    if len(strs)==0 {
		return ""
	}
	prefix := strs[0]
	for _, str := range strs[1:]{
		for !strings.HasPrefix(str , prefix){
			prefix = prefix[:len(prefix)-1]
			if len(prefix)==0{
				return ""
			}
		}
	}
	return prefix
}
func hasPrefix(str, prefix string) bool {
	return len(str) >= len(prefix) && str[:len(prefix)] == prefix
}

func line()  {
	fmt.Println("------------------")
}
