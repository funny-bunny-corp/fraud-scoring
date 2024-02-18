package errors

type AverageTransactionsNotFound struct {
	Err error
}

func (atn AverageTransactionsNotFound) Error() string {
	return "fail to retrieve avg transactions: " + atn.Err.Error()
}
