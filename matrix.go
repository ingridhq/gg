package gg

import "math"

type Matrix struct {
	XX, YX, XY, YY, X0, Y0 float64
}

func Identity() Matrix {
	return Matrix{
		1, 0,
		0, 1,
		0, 0,
	}
}

func Translate(x, y float64) Matrix {
	return Matrix{
		1, 0,
		0, 1,
		x, y,
	}
}

func Scale(x, y float64) Matrix {
	return Matrix{
		x, 0,
		0, y,
		0, 0,
	}
}

func Rotate(angle float64) Matrix {
	c := math.Cos(angle)
	s := math.Sin(angle)
	return Matrix{
		c, s,
		-s, c,
		0, 0,
	}
}

func Shear(x, y float64) Matrix {
	return Matrix{
		1, y,
		x, 1,
		0, 0,
	}
}

func (a Matrix) Multiply(b Matrix) Matrix {
	return Matrix{
		XX: float64(a.XX*b.XX) + float64(a.YX*b.XY),
		YX: float64(a.XX*b.YX) + float64(a.YX*b.YY),
		XY: float64(a.XY*b.XX) + float64(a.YY*b.XY),
		YY: float64(a.XY*b.YX) + float64(a.YY*b.YY),
		X0: float64(a.X0*b.XX) + float64(a.Y0*b.XY) + b.X0,
		Y0: float64(a.X0*b.YX) + float64(a.Y0*b.YY) + b.Y0,
	}
}

func (a Matrix) TransformVector(x, y float64) (tx, ty float64) {
	tx = float64(a.XX*x) + float64(a.XY*y)
	ty = float64(a.YX*x) + float64(a.YY*y)
	return
}

func (a Matrix) TransformPoint(x, y float64) (tx, ty float64) {
	tx = float64(a.XX*x) + float64(a.XY*y) + a.X0
	ty = float64(a.YX*x) + float64(a.YY*y) + a.Y0
	return
}

func (a Matrix) Translate(x, y float64) Matrix {
	return Translate(x, y).Multiply(a)
}

func (a Matrix) Scale(x, y float64) Matrix {
	return Scale(x, y).Multiply(a)
}

func (a Matrix) Rotate(angle float64) Matrix {
	return Rotate(angle).Multiply(a)
}

func (a Matrix) Shear(x, y float64) Matrix {
	return Shear(x, y).Multiply(a)
}
