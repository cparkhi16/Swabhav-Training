package logger

type BankError struct {
	Err string
}

func (b BankError) Error() string {
	return b.Err
}
