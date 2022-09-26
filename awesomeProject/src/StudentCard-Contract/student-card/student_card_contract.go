package student_card

import (
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type StudentCardContract struct {
	contractapi.Contract
}

func (c *StudentCardContract) Instantiate() {
	fmt.Println("StudentContract is Instantiated")
}
func (sc *StudentCardContract) AddStudentCard(ctx TransactionContextInterface, sno string, name string, age int) error {
	card := StudentCard{
		Sno:  sno,
		Name: name,
		Age:  age,
	}
	err := ctx.GetStudentCardList().AddStudentCard(&card)
	if err != nil {
		fmt.Println("AddStudentCard failed")
	} else {
		fmt.Println("AddStudentCard successed")
	}
	return err
}
func (sc *StudentCardContract) GetStudentCard(ctx TransactionContextInterface, sno string) (*StudentCard, error) {
	card, err := ctx.GetStudentCardList().GetStudentCard(sno)
	if err != nil {
		fmt.Println("GetStudentCard failed")
	} else {
		fmt.Println("GetStudentCard successed")
	}
	return card, err
}
func (sc *StudentCardContract) UpdateStudentCard(ctx TransactionContextInterface, sno string, name string, age int) error {
	card := StudentCard{
		Sno:  sno,
		Name: name,
		Age:  age,
	}
	err := ctx.GetStudentCardList().UpdateStudentCard(&card)
	if err != nil {
		fmt.Println("UpdateStudentCard failed")
	} else {
		fmt.Println("UpdateStudentCard successed")
	}
	return err
}
