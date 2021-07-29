/// factory.go

package factory

import (
	"fmt"
)
/*
	1.创建接口
*/
type Shape interface {
	Draw()
}
/************/

/*
	2.创建实现接口的实体类
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
/************/

/*
	3.创建工厂，生成实现接口的实体类的对象
*/

type ShapeFactory struct {

}
// 使用getShape方法获取形状类型的对象
func (sf *ShapeFactory) getShape(shapeType string) Shape {
	switch shapeType {
	case "SQUARE":
		return &Square{}
	case "CIRCLE":
		return &Circle{}
	case "RECTANGLE":
		return &Rectangle{}
	}
	return nil
}