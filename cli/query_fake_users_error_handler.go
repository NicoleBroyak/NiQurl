package cli

func QFUErrHandler(errs [4]error) error {
	for _, v := range errs {
		if v != nil {
			return v
		}
	}
	return nil

}
