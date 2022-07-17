package main

type Rectangle struct {
	x, y int
}

func (r Rectangle) CalculateRect() (int, int) {
	area := r.x * r.y
	circum := (r.x + r.y) * 2

	return area, circum
}

func (r Rectangle) SetRect(x int, y int) {
	r.x = x
	r.y = y
}

func SetAndReturnRect(x int, y int) (Rectangle, error) {
	var r Rectangle
	var err error = nil
	if x < 0 || y < 0 {
		return r, err.New("negative parameter")
	}
	r.SetRect(x, y)
	return r, err
}

func main() {

}
