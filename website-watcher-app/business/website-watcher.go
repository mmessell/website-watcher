package business

type MyResponse struct {
	Message string `json:"Answer:"`
}

func Run() (MyResponse, error) {
	return MyResponse{Message: "Test"}, nil
}
