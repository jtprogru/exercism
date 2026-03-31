package secret

var signals = []string{"wink", "double blink", "close your eyes", "jump"}

func Handshake(code uint) []string {
	var res []string

	switch {
	case code < 1:
	case code&16 == 0:
		for _, signal := range signals {
			if code&1 != 0 {
				res = append(res, signal)
			}
			code >>= 1
		}
	default:
		for idx := 3; idx >= 0; idx-- {
			if code&8 != 0 {
				res = append(res, signals[idx])
			}
			code <<= 1
		}
	}
	return res
}
