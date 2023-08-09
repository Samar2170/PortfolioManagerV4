package internal

type SignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type BankAccountRequest struct {
	Bank      string `json:"bank"`
	AccountNo string `json:"account_no"`
}

type DematAccountRequest struct {
	AccountCode string `json:"account_code"`
	Broker      string `json:"broker"`
}
