package result

type Result struct {
	value any
	err   error
}

func Ok(value any) Result {
	return Result{
		value,
		nil,
	}
}

func Err(err error) Result {
	return Result{
		nil,
		err,
	}
}

func (r Result) Bind(fn func(any) Result) Result {
	if r.err != nil {
		return r
	}
	return fn(r.value)
}

func (r Result) BindWithWrap(fn func(any) (any, error)) Result {
	if r.err != nil {
		return r
	}
	var (
		val any
		err error
	)
	if val, err = fn(r.value); err != nil {
		return Err(err)
	}

	return Ok(val)
}

func (r Result) Unwrap() (any, error) {
	return r.value, r.err
}

func (r Result) Error() error {
	return r.err
}
