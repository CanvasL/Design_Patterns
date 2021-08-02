/// abstract_factory_test.go

package abstract_factory

import "testing"

func TestMain(t *testing.T) {
	var factory_producer = new(FactoryProducer)

	shape_factory := factory_producer.getFactory(SHAPE_FACTORY)
	shape_factory.getShape(RECTANGLE).Draw()
	shape_factory.getShape(CIRCLE).Draw()
	shape_factory.getShape(SQUARE).Draw()

	color_factory := factory_producer.getFactory(COLOR_FACTORY)
	color_factory.getColor(BLUE).Fill()
	color_factory.getColor(RED).Fill()
	color_factory.getColor(GREEN).Fill()
}
