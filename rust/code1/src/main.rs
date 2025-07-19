fn main() {
    println!("Hello, world!");

    let mut x = 5;
    println!("x value: {}", x);
    x = 6;
    println!("x value: {}", x);
    println!("Hello, world!");


    // 下划线开头的变量忽略未使用警告
    let _y1 = 5;
    let _y2 = 5;


    // 变量结构
    let (a, mut b) = (true, false);
    println!("a: {}, b: {}", a, b);
    b = true;
    println!("a: {}, b: {}", a, b);
}
