package student_card

import "github.com/hyperledger/fabric-contract-api-go/contractapi"

type TransactionContextInterface interface {
	contractapi.TransactionContextInterface
	GetStudentCardList() StudentCardListInterface
}

type TransactionContext struct {
	contractapi.TransactionContext
	studentCardList *List
}

func (ctx *TransactionContext) GetStudentCardList() StudentCardListInterface {
	if ctx.studentCardList == nil {
		ctx.studentCardList = newList(ctx)
	}
	return ctx.studentCardList
}
