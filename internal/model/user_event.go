package model

type UserEvent struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Address   string `json:"address,omitempty"`
	Photos    string `json:"photos,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}

func (u *UserEvent) GetId() string {
	return u.ID
}
