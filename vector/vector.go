package vector

type V [2]int

func (v V) X() int { return v[0] }
func (v V) Y() int { return v[1] }

func Sub(v, u, buf V) V {
	buf[0] = v[0] - u[0]
	buf[1] = v[1] - u[1]

	return buf
}

func Add(v, u, buf V) V {
	buf[0] = v[0] + u[0]
	buf[1] = v[1] + u[1]

	return buf
}
