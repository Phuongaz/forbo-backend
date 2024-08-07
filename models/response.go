package models

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type UserAuthResponse struct {
	Data struct {
		UserID string `json:"user_id"`
		Token  string `json:"token"`
	} `json:"data"`
	Message string `json:"message"`
}

type FeedResponse struct {
	Data    []Feed `json:"data"`
	Message string `json:"message"`
}
