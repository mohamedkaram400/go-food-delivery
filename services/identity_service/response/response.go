package response

type RegisterError struct {
    Fields map[string]string
}

func (e RegisterError) Error() string {
    return "registration failed"
}

type InvalidCredentialsError struct {
    Message string
}

func (e InvalidCredentialsError) Error() string {
    return e.Message
}

