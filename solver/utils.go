package solver

func IAbs[V int | uint](x V) V {
	if x < 0 {
		return -x
	}
	return x
}

type Coord2D struct {
	x int
	y int
}

// func (self *Coord2D) AddLvalue(data *Coord2D) {
// 	self.x += data.x
// 	self.y += data.y
// }

// func (self *Coord2D) MultiplyByPrvalue(data int) {
// 	self.x *= data
// 	self.y *= data
// }
