package model

type TakenBook struct {
	UserWhoTake  	User    				`json:"user"`
	TakenBook 		BookWithAuthor 			`json:"book"`
}
