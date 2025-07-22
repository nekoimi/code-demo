fn main() {
    println!("Hello, world!");

    // greet("a");

    let s = String::from("hello world");

    let hello = &s[0..5];
    let world = &s[6..11];
    println!("Hello, {}!", hello);
    println!("Hello, {}!", world);

    let s2 = s;
    println!("s2: {}", s2);


    // 不可变引用
    let s3: &str = "hello s3";
    println!("s3: {}", s3);

    // &str -> 字符串切片 不可变字符串
    // String -> 标准库  可变字符串

    // String -> &str
    let s4: &str = "hello s4";
    println!("s4: {}", s4);
    let s5: String = s4.to_string();
    println!("s5: {}", s5);

    // &str -> String
    let mut s6: String = String::from("hello s6");
    println!("s6: {}", s6);
    let s7: &str = &s6[6..];
    println!("s7: {}", s7);

    s6.push_str(", world!");
    println!("s6: {}", s6);

    // 操作中文
    for char in "中文汉字".chars() {
        println!("{}", char);
    }
}

fn greet(name: String) {
    println!("Hello, {}!", name);
}
