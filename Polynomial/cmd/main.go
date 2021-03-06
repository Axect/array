package main

import (
	"fmt"

	"github.com/Axect/Numeric/Polynomial"
)

// Vector is type alias
type Vector = []float64

func main() {
	var A, B Polynomial.Polynomial
	A.Coeff = Vector{1, -2, 3, 4}
	B.Coeff = Vector{2, 1}
	C := Polynomial.Add(A, B)
	D := Polynomial.Sub(A, B)
	C.Show()
	D.Show()
	E := Polynomial.Diff(D)
	E.Show()
	T := Polynomial.Integrate(E, 3)
	T.Show()
	fmt.Println()

	var P, Q Polynomial.Polynomial
	P.Coeff = Vector{1, 3, 3, 1}
	Q.Coeff = Vector{1, 1}
	R, r := Polynomial.Honor(P, Q)
	R.Show()
	fmt.Println(r)
	r1, r2 := Polynomial.RealQuad(R)
	fmt.Println(r1, r2)
	fmt.Println()

	var S Polynomial.Polynomial
	S.Coeff = Vector{1, 2, 3, 4}
	F := S.PolyFunc()
	S.Show()
	fmt.Println(F(1))
	fmt.Println(F(2))
}
