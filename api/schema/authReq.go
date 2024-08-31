package schema

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResetPasswordReq struct {
	Email string `json:"email"`
}

type UpdatePasswordReq struct {
	Password string `json:"password"`
}

type UpdateEmailReq struct {
	Email string `json:"email"`
}

type ChangeEmailReq struct {
	UserId string `json:"userId"`
}

type SendForgetPasswordMailReq struct {
	Email string `json:"email"`
}

type CheckAuthenticationCodeForResetPasswordReq struct {
	Email              string `json:"email"`
	AuthenticationCode string `json:"authenticationCode"`
}
