package models

type UserResponse struct {
	Data struct {
		UserID string `json:"user_id"`
		Token  string `json:"token"`
	} `json:"user"`
	Message string `json:"message"`
}
