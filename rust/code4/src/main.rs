fn main() {
    println!("Hello, world!");


    let v = {
        let a = 1;
        let b = 2;
        a + b
    };
    println!("v = {}", v);

    let add_func = add_2;
    let result = add_func(1);
    println!("result is {}", result);
}


fn add_2(a: i32) -> i32 {
    if a < 10 {
        return a + add_2(a + 1);
    }

    a + 2
}
