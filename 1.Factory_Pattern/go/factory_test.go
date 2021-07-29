/// factory_test.go
package factory

import "testing"

func TestDraw(t *testing.T) {
	sf := new(ShapeFactory)
	sf.getShape("CIRCLE").Draw()
	sf.getShape("RECTANGLE").Draw()
	sf.getShape("SQUARE").Draw()
}