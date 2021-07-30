/// abstract_factory.go

package abstract_factory

import "fmt"

/*
	1.为【形状】创建接口
*/
type IShape interface {
	Draw()
}

/*
	2.创建实现【形状】接口的实体类
*/
// Square类
type Square struct {
}

func (s *Square) Draw() {
	fmt.Println("Draw a [Square].")
}

// Circle类
type Circle struct {
}

func (c *Circle) Draw() {
	fmt.Println("Draw a [Circle].")
}

// Rectangle类
type Rectangle struct {
}

func (r *Rectangle) Draw() {
	fmt.Println("Draw a [Rectangle].")
}

/*
	3.为【颜色】创建接口
*/
type IColor interface {
	Fill()
}

/*
	4.创建实现【颜色】接口的实体类
*/
// Red类
type Red struct {
}

func (r *Red) Fill() {
	fmt.Println("Fill in [Red] color.")
}
// Green类
type Green struct {
}

func (g *Green) Fill() {
	fmt.Println("Fill in [Green] color.")
}
// Blue类
type Blue struct {
}

func (g *Blue) Fill() {
	fmt.Println("Fill in [Blue] color.")
}

/*
	5.抽象工厂接口，能够获取工厂Shape和工厂Color
*/
type IAbstractFactory interface {
	getShape() IShape
	getColor() IColor
}

/*
	6.抽象工厂类
*/
type AbstractFactory struct {

}

func (af *AbstractFactory) getShape(shapeType )


const (
	Shape = iota
	SQUARE 
	CIRCLE
	RECTANGLE
	Color
	RED
	GREEN
	BLUE
)

type ShapeFactory struct {

}

func (sf *ShapeFactory) getShape(shapeType int) IShape {
	switch shapeType {
	case SQUARE:
		return &Square{}
	case CIRCLE:
		return &Circle{}
	case RECTANGLE:
		return &Rectangle{}
	}
	return nil
}

type ColorFactory struct {

}

func (cf *ColorFactory) getColor(colorType int) IColor {
	switch colorType {
	case RED:
		return &Red{}
	case GREEN:
		return &Green{}
	case BLUE:
		return &Blue{}
	}
	return nil
}