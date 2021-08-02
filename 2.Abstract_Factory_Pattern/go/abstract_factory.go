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
	fmt.Println("Fill in [Red].")
}

// Green类
type Green struct {
}

func (g *Green) Fill() {
	fmt.Println("Fill in [Green].")
}

// Blue类
type Blue struct {
}

func (g *Blue) Fill() {
	fmt.Println("Fill in [Blue].")
}

/*
	5.抽象工厂接口，各种类型的工厂共享(这里指ShapeFactory和ColorFactory)
*/
type IAbstractFactory interface {
	getShape(int) IShape
	getColor(int) IColor
}

/*
	6.创建Shape工厂和Color工厂的实体类
*/

const (
	SHAPE_FACTORY = iota
	SQUARE
	CIRCLE
	RECTANGLE
	COLOR_FACTORY
	RED
	GREEN
	BLUE
)

/// ShapeFactory类
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

// 无实义，仅用来保证ShapeFactory实现了IAbstractFactory接口
func (sf *ShapeFactory) getColor(shapeType int) IColor {
	return nil
}

/// ColorFactory类
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

// 无实义，仅用来保证ColorFactory实现了IAbstractFactory接口
func (cf *ColorFactory) getShape(colorType int) IShape {
	return nil
}

/*
	7.建立超级工厂类，用于获取工厂实例
*/
type FactoryProducer struct {
}

func (af *FactoryProducer) getFactory(factoryType int) IAbstractFactory {
	switch factoryType {
	case SHAPE_FACTORY:
		defer fmt.Println("[ShapeFactory] has been built...")
		return &ShapeFactory{}
	case COLOR_FACTORY:
		defer fmt.Println("[ColorFactory] has been built...")
		return &ColorFactory{}
	}
	return nil
}
