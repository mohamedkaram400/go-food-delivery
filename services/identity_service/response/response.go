package response

type RegisterError struct {
    Fields map[string]string
}

func (e RegisterError) Error() string {
    return "registration failed"
}