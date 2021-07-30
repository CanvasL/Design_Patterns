/// abstract_factory_test.go

package abstract_factory

import "testing"

func TestAbstract_Factory(t *testing.T) {
	var iFactory IAbstractFactory

	iFactory = new(ShapeFactory)
	iFactory.getShape(RECTANGLE).Draw() 
	iFactory.getShape(CIRCLE).Draw() 
	iFactory.getShape(SQUARE).Draw()
	
	iFactory = new(ColorFactory)
	iFactory.getColor(BLUE).FILL()
	iFactory.getColor(RED).FILL()
	iFactory.getColor(GREEN).FILL()
}
