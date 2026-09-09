package auth

type GetPhoneCodeRequest struct {
	Phone string `json:"phone" validate:"required,len=11,numeric"`
}

type GetPhoneCodeResponse struct {
	SessionId string `json:"session_id" validate:"required"`
}

type VerifyPhoneCodeRequest struct {
	Code string `json:"code" validate:"required"`
	SessionId  string `json:"session_id" validate:"required"`
}

type VerifyPhoneCodeResponse struct {
	Token string `json:"token" validate:"required"`
}

var (
	ErrSessionIdAlreadyExists = "Session id already exists"
	ErrSessionIdDoesNotExist = "Session doesn't exist exists"
	ErrPhoneDoesNotMatchSession = "Phone number does not match session"
	ErrInvalidCode = "Invalid code"
)