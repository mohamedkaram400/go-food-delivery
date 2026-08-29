package requests


type LoginRequest struct {
	Email 		string
	Password	string
}


type RegisterRequest struct {
	Name 		string
	Email 		string
	Password	string
}