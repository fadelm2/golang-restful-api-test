package model

type CreditcardEvent struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Type      string `json:"type" `
	Name      string `json:"name"`
	Number    string `json:"number" `
	Expired   string `json:"expired" `
	Cvv       string `json:"cvv"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

func (c *CreditcardEvent) GetId() string {
	return c.ID
}
