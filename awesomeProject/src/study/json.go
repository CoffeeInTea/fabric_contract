package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string `json:"stu_name"`
	Age  int    `json:"stu_age"`
	Sex  string `json:"stu_sex"`
}

func main() {

	//将结构体转换未json
	student := Student{
		Name: "zhangsan",
		Age:  18,
		Sex:  "boy",
	}
	//返回的结构类型为[]byte
	jsonStu, err := json.Marshal(student)
	if err == nil {
		fmt.Println(string(jsonStu))
	}

	//将json字符串转换为结构体

	jsonStu2 := []byte(`{"stu_name":"zhangsan","stu_age":18,"stu_sex":"boy"}`)
	fmt.Println("jsonStu=")
	fmt.Println(jsonStu)
	fmt.Println("jsonStu2=")
	fmt.Println(jsonStu2)

	var student2 Student
	err = json.Unmarshal(jsonStu2, &student2)
	fmt.Println(student2)
}
