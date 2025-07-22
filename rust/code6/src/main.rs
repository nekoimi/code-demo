fn main() {
    println!("Hello, world!");

    let tup = (500, 6.4, 1);
    let (x, y, z) = tup;
    println!("x = {}, y = {}, z = {}", x, y, z);

    // 元组可以直接用.来访问
    println!("x = {}, y = {}, z = {}", tup.0, tup.1, tup.2);

    let check_str = String::from("hello");
    let (s, len) = check_length(check_str);
    println!("s = {}, len = {}", s, len);
}

fn check_length(input: String) -> (String, usize) {
    let length = input.len();

    (input, length)
}
