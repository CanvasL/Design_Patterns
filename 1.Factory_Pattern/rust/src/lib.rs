/*
    1.创建trait
*/
pub trait IShape {
    fn draw(&self);
}

/*
    2.创建实现trait的实体类
*/
struct Square {}
impl IShape for Square {
    fn draw(&self) {
        println!("Draw a [Square].");
    }
}

struct Circle {}
impl IShape for Circle {
    fn draw(&self) {
        println!("Draw a [Circle].");
    }
}

struct Rectangle {}
impl IShape for Rectangle {
    fn draw(&self) {
        println!("Draw a [Rectangle].");
    }
}

/*
    3.创建工厂
*/
enum ShapeType {
    Square,
    Circle,
    Rectangle
}

struct ShapeFactory {}

impl ShapeFactory {
    fn new() -> Self{
        ShapeFactory{}
    }

    fn get_shape(&self, shape_type: &ShapeType) -> Box<dyn IShape> {
        match shape_type {
            ShapeType::Square => Box::new(Square {}),
            ShapeType::Circle => Box::new(Circle {}),
            ShapeType::Rectangle => Box::new(Rectangle {}),
        }
    }
}

#[test]
fn it_works() {
    let sf = ShapeFactory::new();
    sf.get_shape(&ShapeType::Square).draw();
    sf.get_shape(&ShapeType::Circle).draw();
    sf.get_shape(&ShapeType::Rectangle).draw();
}