fn main() {
    println!("Hello, world!");

    // 变量遮蔽：重复定义
    let x = 5;
    let x = x + 1;

    {
        let x = x * 2;
        println!("The value of x in the inner scope is: {}", x);
    }
    println!("The value of x is: {}", x);

    let a = [1, 2, 3, 4, 5];
    println!("a length is: {}", a.len());

    {
        let a = [7, 8, 9];
        println!("a length is: {}", a.len());
        println!("The value of a index0 is: {}", a[0]);
    }

    println!("The value of a index0 is: {}", a[0]);
}
