package src

type GetUserSummaryResponse struct {
	User struct {
		ID string `json:"id"`
	} `json:"user"`
}
