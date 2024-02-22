package errors

type LastOrderNotFound struct {
	Err error
}

func (lon LastOrderNotFound) Error() string {
	return "fail to retrieve last order: " + lon.Err.Error()
}
