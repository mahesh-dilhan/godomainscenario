package fund

type Fund struct {
	// balance is unexported (private), because it's lowercase
	balance int
}
