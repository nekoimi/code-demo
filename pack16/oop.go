/**
 * # ------------Oooo---
 * # -----------(----)---
 * # ------------)--/----
 * # ------------(_/-
 * # ----oooO----
 * # ----(---)----
 * # -----\--(--
 * # ------\_)-
 * # ----
 * #     author : Yprisoner <yyprisoner@gmail.com>
 * #     19-5-26 下午9:32
 * #                            ------
 * #    「 涙の雨が頬をたたくたびに美しく 」
 **/
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
