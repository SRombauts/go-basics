package multi

import "testing"

func TestReturn(t *testing.T) {
	const ina, inb = 0, 0
	const outx, outy = 0, 0
	x, y := Return(ina, inb)
	if (x != outx) || (y != outy) {
		t.Errorf("Return(%v, %v) = (%v, %v), want (%v, %v)", ina, inb, x, y, outx, outy)
	}
}
