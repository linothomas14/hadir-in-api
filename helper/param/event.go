package param

type CreateEvent struct {
	Title        string `json:"title"`
	Date         string `json:"date"`
	ExpiredToken string `json:"expired_token"`
}
