package main

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	Message string  // error message
	Field   string  // 不正なフィールド（高さ・幅・半径のいずれか）
	Value   float64 // 不正な値（図形の長さ）
}

func (e *ValidationError) Error() string {
	return e.Message
}

type Shape interface {
	Area() float64      // 面積を求める
	Perimeter() float64 // 周長を求める
}

type Rectangle struct {
	Width  float64
	Height float64
}

// 面積を求めるメソッド（長方形）
// Area()メソッドを実装することでShapeインターフェースを満たす
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return (r.Width + r.Height) * 2
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * 3.14
}

// 周長を求めるメソッド（円）
// Perimeter()メソッドを実装することでShapeインターフェースを満たす
func (c Circle) Perimeter() float64 {
	return c.Radius * 3.14
}

func validateShape(s Shape) error {
	switch shape := s.(type) {
	case Rectangle:
		if shape.Width <= 0 {
			return &ValidationError{
				Message: "幅は0より大きい値である必要があります",
				Field:   "Width",
				Value:   shape.Width,
			}
		}
		if shape.Height <= 0 {
			return &ValidationError{
				Message: "高さは0より大きい値である必要があります",
				Field:   "Height",
				Value:   shape.Height,
			}
		}
	case Circle:
		if shape.Radius <= 0 {
			return &ValidationError{
				Message: "半径は0より大きい値である必要があります",
				Field:   "Radius",
				Value:   shape.Radius,
			}
		}
	}
	return nil
}

func main() {
	shapes := []Shape{
		Rectangle{Width: 10.0, Height: 5.0},
		Circle{Radius: 3.0},
		Rectangle{Width: -2.0, Height: 5.0}, // 不正（幅が不正）
		Circle{Radius: -1.0},                // 不正（半径が不正）
		Rectangle{Width: 1.0, Height: 0.0},  // 不正（高さが不正）
	}

	// バリデーション後のShapeを格納するスライス
	var calcShapes []Shape

	//　スライスに対するバリデーション
	for _, v := range shapes {

		if err := validateShape(v); err != nil {
			var valErr *ValidationError
			if errors.As(err, &valErr) {
				fmt.Println(valErr.Message)
			} else {
				// その他のエラー
				fmt.Println("予期しないエラー:", err)
			}
		} else {
			calcShapes = append(calcShapes, v)
		}
	}

	for _, v := range calcShapes {
		area := v.Area()
		perimeter := v.Perimeter()
		fmt.Printf("図形: %T、面積は%f、周長は%fです。\n", v, area, perimeter)
	}

}
