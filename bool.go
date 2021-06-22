package ecrypto

func Trues(bs ...bool) int {
	var r int
	for _, b := range bs {
		if b {
			r++
		}
	}
	return r
}
