package src

type GetUserSummaryResponse struct {
	User struct {
		ID string `json:"id"`
	} `json:"user"`
}

type GetUserLeaderboardResponse struct {
	Runs []struct {
		ID         string  `json:"id"`
		GameID     string  `json:"gameId"`
		CategoryID string  `json:"categoryId"`
		Obsolete   bool    `json:"obsolete"`
		Place      int     `json:"place,omitempty"`
		Igt        float64 `json:"igt"`
	} `json:"runs"`
}
