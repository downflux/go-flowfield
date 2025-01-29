package vector

type V [2]int

func (v V) X() int { return v[0] }
func (v V) Y() int { return v[1] }

func (v V) Add(u V) {
	v[0] += u[0]
	v[1] += u[1]
}

func (v V) Sub(u V) {
	v[0] -= u[0]
	v[1] -= u[1]
}

func Sub(v, u V) V { return [2]int{v.X() - u.X(), v.Y() - u.Y()} }
func Add(v, u V) V { return [2]int{v.X() + u.X(), v.Y() + u.Y()} }

type B [2]V

func (b B) P() V { return b[0] }
func (b B) D() V { return b[1] }

// In checks for membership within the box.
//
// We are checking the half-open interval [min, max).
func (b B) In(v V) bool {
	return v.X() >= b.P().X() && v.X() < b.P().X()+b.D().X() && v.Y() >= b.P().Y() && v.Y() < b.P().Y()+b.D().Y()
}
