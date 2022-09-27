package student_card

import (
	ledgerapi "awesomeProject/src/StudentCard-Contract/ledger-api"
)

// StudentCardListInterface 定义StudentCardList应实现的方法
type StudentCardListInterface interface {
	AddStudentCard(card *StudentCard) error
	GetStudentCard(sno string) (*StudentCard, error)
	UpdateStudentCard(card *StudentCard) error
}
type List struct {
	stateList ledgerapi.StateListInterface
}

func (scl *List) AddStudentCard(card *StudentCard) error {
	return scl.stateList.AddState(card)
}
func (scl *List) GetStudentCard(sno string) (*StudentCard, error) {
	sc := new(StudentCard)
	err := scl.stateList.GetState(sno, sc)
	if err != nil {
		return nil, err
	} else {
		return sc, err
	}
}
func (scl *List) UpdateStudentCard(card *StudentCard) error {
	return scl.stateList.AddState(card)
}
func newList(ctx TransactionContextInterface) *List {
	stateList := new(ledgerapi.StateList)
	stateList.Ctx = ctx
	stateList.Name = "StudentCar-Contract.student-card"
	stateList.Deserialize = func(bytes []byte, state ledgerapi.StateInterface) error {
		return Deserialize(bytes, state.(*StudentCard))
	}

	list := new(List)
	list.stateList = stateList

	return list
}
