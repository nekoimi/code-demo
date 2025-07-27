enum Status {
    Alive,
    Dead,
    ActionA(String),
    ActionB(u8)
}

fn main() {
    println!("Hello, world!");

    // Java 的枚举是“有点聪明的常量”，而 Rust 的枚举是“类型系统中的变体表达工具”，强大到可以代替许多类层级结构。

    let _some_number = Some(5);
    let _some_string = Some("a string");
    let absent_number: Option<i32> = None;

    match absent_number {
        Some(value) => println!("value: {}", value),
        None => println!("None")
    }
}
