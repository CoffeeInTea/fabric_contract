package student_card

import (
	"fmt"
	"testing"
)

func TestDeserialize(t *testing.T) {
	sc := &StudentCard{
		Sno:  "123",
		Name: "zhangsan",
		Age:  18,
	}
	data, err := sc.Serialize()
	fmt.Println(data)
	//测试过程中发现调用newSC的函数需要初始化，对比study/json.go
	newSC := &StudentCard{}
	err = newSC.UnmarshalJson(data)
	if err != nil {
		t.Fatalf("Deserialize failed")
	} else {
		t.Logf("Deserialize successed")
		fmt.Println(*newSC)
	}
}

func TestStudentCard_MarshalJson(t *testing.T) {
	sc := &StudentCard{
		Sno:  "123",
		Name: "zhangsan",
		Age:  18,
	}
	if data, err := sc.MarshalJson(); err != nil {
		t.Fatalf("MarshalJson failed")
	} else {
		t.Logf("MarshalJson successed")
		fmt.Println(data)
		fmt.Println(string(data))
	}
}

func TestStudentCard_Serialize(t *testing.T) {
	sc := &StudentCard{
		Sno:  "123",
		Name: "zhangsan",
		Age:  18,
	}
	if data, err := sc.Serialize(); err != nil {
		t.Fatalf("Serialize failed")
	} else {
		t.Logf("Serialize successed")
		fmt.Println(data)
	}
}

func TestStudentCard_UnmarshalJson(t *testing.T) {
	sc := &StudentCard{
		Sno:  "123",
		Name: "zhangsan",
		Age:  18,
	}
	data, err := sc.Serialize()
	if err != nil {
		t.Fatalf("Deserialize failed")
	} else {
		t.Logf("Deserialize successed")
		fmt.Println(data)
		fmt.Println(string(data))
	}
}
