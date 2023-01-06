package param

type UserUpdate struct {
	ID       uint32 `json:"id" `
	Name     string `json:"name"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"required,min=6"`
}
