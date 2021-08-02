/*
	1.为【形状】创建trait
*/

pub trait IShape {
    fn draw(&self);
}

/*
	2.创建实现【形状】trait的实体类
*/
/// Square类
struct Square {}
impl IShape for Square {
    fn draw(&self) {
        println!("Draw a [Square].");
    }
}

/// Circle类
struct Circle {}
impl IShape for Circle {
    fn draw(&self) {
        println!("Draw a [Circle].")
    }
}

/// Rectangle类
struct Rectangle {}
impl IShape for Rectangle {
    fn draw(&self) {
        println!("Draw a [Rectangle].")
    }
}

/*
	3.为【颜色】创建接口
*/
pub trait IColor {
    fn fill(&self);
}

/*
	4.创建实现【颜色】接口的实体类
*/
/// Red类
struct Red {}
impl IColor for Red {
    fn fill(&self) {
        println!("Fill in [Red].");
    }
}

/// Green类
struct Green {}
impl IColor for Green {
    fn fill(&self) {
        println!("Fill in [Green].");
    }
} 

/// Blue类
struct Blue {}
impl IColor for Blue {
    fn fill(&self) {
        println!("Fill in [Blue].")
    }
}

/*
	5.抽象工厂trait，各种类型的工厂共享(这里指ShapeFactory和ColorFactory)
*/
/// 工厂类型枚举
enum FactoryType {
    ShapeFactory,
    ColorFactory
}

/// Shape类型枚举
pub enum ShapeType {
    Square,
    Circle,
    Rectangle
}

/// Color类型枚举
pub enum ColorType {
    Red,
    Green,
    Blue
}

pub trait IAbstractFactory {
    fn get_shape(&self, shape_type: &ShapeType) -> Option<Box<dyn IShape>>;
    fn get_color(&self, color_type: &ColorType) -> Option<Box<dyn IColor>>;
}

/*
	6.创建Shape工厂和Color工厂的实体类
*/
struct ShapeFactory {}
impl IAbstractFactory for ShapeFactory {
    // fn new() -> Self {
    //     ShapeFactory{}
    // }

    fn get_shape(&self, shape_type: &ShapeType) -> Option<Box<dyn IShape>> {
        match shape_type {
            ShapeType::Square => Some(Box::new(Square{})),
            ShapeType::Circle => Some(Box::new(Circle{})),
            ShapeType::Rectangle => Some(Box::new(Rectangle{})),
        }
    }

    // getColor
    fn get_color(&self, _color_type: &ColorType) -> Option<Box<dyn IColor>> {
        return None;
    }
}

struct ColorFactory {}
impl IAbstractFactory for ColorFactory {
    // fn new() -> Self {
    //     ColorFactory{}
    // }

    fn get_color(&self, color_type: &ColorType) -> Option<Box<dyn IColor>> {
        match color_type {
            ColorType::Red => Some(Box::new(Red{})),
            ColorType::Green => Some(Box::new(Green{})),
            ColorType::Blue => Some(Box::new(Blue{})),
        }
    }

    // getShape
    fn get_shape(&self, _shape_type: &ShapeType) -> Option<Box<dyn IShape>> {
        return None;
    }
}

/*
	7.建立超级工厂类，用于获取工厂实例
*/
struct FactoryProducer {}
impl FactoryProducer {
    fn new() -> Self {
        return FactoryProducer{};
    }

    fn get_factory(&self, factory_type: &FactoryType) -> Box<dyn IAbstractFactory> {
        match factory_type {
            FactoryType::ShapeFactory => Box::new(ShapeFactory{}),
            FactoryType::ColorFactory => Box::new(ColorFactory{}),
        }
    }
}

#[test]
fn it_works() {
    let factory_producer = FactoryProducer::new();

    factory_producer.get_factory(&FactoryType::ShapeFactory).get_shape(&ShapeType::Square).unwrap().draw();
    factory_producer.get_factory(&FactoryType::ShapeFactory).get_shape(&ShapeType::Rectangle).unwrap().draw();
    factory_producer.get_factory(&FactoryType::ShapeFactory).get_shape(&ShapeType::Circle).unwrap().draw();
    factory_producer.get_factory(&FactoryType::ColorFactory).get_color(&ColorType::Red).unwrap().fill();
    factory_producer.get_factory(&FactoryType::ColorFactory).get_color(&ColorType::Blue).unwrap().fill();
    factory_producer.get_factory(&FactoryType::ColorFactory).get_color(&ColorType::Green).unwrap().fill();
}