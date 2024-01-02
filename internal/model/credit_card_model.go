package model

type CreditcardResponse struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	Name      string `json:"name"`
	Number    string `json:"number"`
	Expired   string `json:"expired"`
	Cvv       string `json:"cvv"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type CreateCreditcardRequest struct {
	UserId  string `json:"-" validate:"required"`
	Type    string `json:"type" validate:"required,max=100"`
	Name    string `json:"name" validate:"required,max=100"`
	Number  string `json:"number" validate:"required,max=32"`
	Expired string `json:"expired" validate:"required,max=100"`
	Cvv     string `json:"cvv" validate:"required,max=100"`
}

type UpdateCreditcardRequest struct {
	UserId  string `json:"-" validate:"required"`
	ID      string `json:"-" validate:"required,max=100,uuid"`
	Type    string `json:"type" validate:"required,max=100"`
	Name    string `json:"name" validate:"required,max=100"`
	Number  string `json:"number" validate:"required,max=32"`
	Expired string `json:"expired" validate:"required,max=100"`
	Cvv     string `json:"cvv" validate:"required,max=100"`
}

type SearchCreditcardRequest struct {
	UserId string `json:"-" validate:"required"`
	Name   string `json:"name" validate:"max=100"`
	Number string `json:"number" validate:"max=32"`
	Page   int    `json:"page" validate:"min=1"`
	Size   int    `json:"size" validate:"min=1,max=100"`
}
type GetCreditcardRequest struct {
	UserId string `json:"-" validate:"required"`
	ID     string `json:"-" validate:"required,max=100,uuid"`
}

type DeleteCreditcardRequest struct {
	UserId string `json:"-" validate:"required"`
	ID     string `json:"-" validate:"required,max=100,uuid"`
}
