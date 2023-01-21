package response

type PresentResponse struct {
	ID    uint32       `json:"id" `
	User  UserResponse `json:"user"`
	Event EventRes     `json:"event"`
}
