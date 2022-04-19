package afrikpay

// Mode is used to determine the payment mode
type Mode string

const (
	// ModeCash payment from agent partner account
	ModeCash = Mode("cash")
	// ModeMoney send payment request
	ModeMoney = Mode("money")
	// ModeAccount make payment operation from afrikpay client
	ModeAccount = Mode("account")
)

// String converts the Mode to a string
func (mode Mode) String() string {
	return string(mode)
}
