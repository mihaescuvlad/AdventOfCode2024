package solver

func IAbs[V int | uint](x V) V {
	if x < 0 {
		return -x
	}
	return x
}
