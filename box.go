package golang_united_school_homework

import "fmt"

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) == b.shapesCapacity {
		return fmt.Errorf("it goes out of the shapesCapacity range")
	}
	b.shapes = append(b.shapes, shape)

	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i < 0 || i >= len(b.shapes) {
		return nil, fmt.Errorf("index went out of the range")
	}

	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if i < 0 || i >= len(b.shapes) {
		return nil, fmt.Errorf("index went out of the range")
	}

	res := b.shapes[i]
	b.shapes = append(b.shapes[0:i], b.shapes[i+1:]...)

	return res, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if i < 0 || i >= len(b.shapes) {
		return nil, fmt.Errorf("index went out of the range")
	}

	res := b.shapes[i]
	b.shapes[i] = shape

	return res, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	res := float64(0)
	for indx1, indx2 := 0, len(b.shapes)-1; indx1 < indx2; {
		res = res + b.shapes[indx1].CalcPerimeter() + b.shapes[indx2].CalcPerimeter()
		indx1++
		indx2--

		if indx1 == indx2 {
			res = res + b.shapes[indx1].CalcPerimeter()
			break
		}
	}

	return res
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	res := float64(0)
	for indx1, indx2 := 0, len(b.shapes)-1; indx1 < indx2; {
		res = res + b.shapes[indx1].CalcArea() + b.shapes[indx2].CalcArea()
		indx1++
		indx2--

		if indx1 == indx2 {
			res = res + b.shapes[indx1].CalcArea()
			break
		}
	}

	return res
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	var i Shape
	var hasCircle bool
	newShapes := make([]Shape, 0, b.shapesCapacity)

	for indx, e := range b.shapes {
		i = e
		_, ok := i.(*Circle)

		if ok {
			hasCircle = true
		} else {
			newShapes = append(newShapes, b.shapes[indx])
		}
	}

	b.shapes = newShapes

	if !hasCircle {
		return fmt.Errorf("circles are not exist")
	}
	return nil
}
