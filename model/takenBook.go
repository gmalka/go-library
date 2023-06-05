package model

type TakenBook struct {
	UserWhoTake   User    	`json:"user"`
	TakenBook Book 			`json:"book"`
}
