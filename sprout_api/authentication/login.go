package authentication

type LoginSuccessResponse struct {
	Message  string `json:"message"`
	Redirect string `json:"redirect"`
	Results  int64  `json:"results"`
	Status   string `json:"status"`
}
