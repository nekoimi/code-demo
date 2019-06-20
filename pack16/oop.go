package pack16

import "fmt"

type Point struct {
	px float32
	py float32
}

func (p *Point) setXY (px, py float32) {
	p.px = px
	p.py = py
}

func (p *Point) getXY () (x, y float32) {
	return p.px, p.py
}

func Pack16Main () {
	p := Point{}
	p.setXY(1.2,2.3)
	fmt.Println(p.getXY())
}
