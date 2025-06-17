package src

type API struct{}

func New() *API {
	return &API{}
}

func (s *API) GetRank(user string) string {
	return "100"
}
