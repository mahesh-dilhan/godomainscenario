package fund

type Fund struct {
	// balance is unexported (private), because it's lowercase
	balance int
}

// A regular function returning a pointer to a fund
func NewFund(initialBalance int) *Fund {
	// We can return a pointer to a new struct without worrying about
	// whether it's on the stack or heap: Go figures that out for us.
	return &Fund{
		balance: initialBalance,
	}
}
