package model

type UserResponse struct {
	ID        string `json:"id",omitempty`
	Name      string `json:"name",omitempty`
	Address   string `json:"address",omitempty`
	Photos    string `json:"photos",omitempty`
	Token     string `json:"token",omitempty`
	CreatedAt int64  `json:"createdAt",omitempty`
	UpdatedAt int64  `json:"updatedAt",omitempty`
}

type VerifyUserRequest struct {
	Token string `validate:"required,max=100"`
}

type GetUserRequest struct {
	ID string `json:"id" validate:"required,max=100"`
}

type RegisterUserRequest struct {
	ID       string `json:"id" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
	Name     string `json:"name" validate:"required,max=100"`
	Address  string `json:"address,omitempty" validate:"max=100"`
	Photos   string `json:"photos,omitempty" validate:"max=100"`
}

type UpdateUserRequest struct {
	ID       string `json:"-" validate:"required,max=100"`
	Password string `json:"password,omitempty" validate:"max=100"`
	Name     string `json:"name,omitempty" validate:"max=100"`
	Address  string `json:"address,omitempty" validate:"max=100"`
	Photos   string `json:"photos,omitempty" validate:"max=100"`
}
type LoginUserRequest struct {
	ID       string `json:"id" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
}

type LogoutUserRequest struct {
	ID string `json:"id" validate:"required,max=100"`
}
