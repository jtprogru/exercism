package erratum

func Use(opener ResourceOpener, input string) (err error) {
	var res Resource
	for {
		res, err = opener()
		if err == nil {
			break
		}
		if _, ok := err.(TransientError); !ok {
			return err
		}
	}
	defer res.Close()
	defer func() {
		if x := recover(); x != nil {
			if frobErr, ok := x.(FrobError); ok {
				res.Defrob(frobErr.defrobTag)
			}
			err = x.(error)
		}
	}()
	res.Frob(input)
	if err != nil { // Set in recover handler
		return err
	}
	return nil
}
