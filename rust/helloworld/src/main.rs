fn greet_world() {
    let str = "hello world";
    let str2 = "hi, rust";
    let strs = [str, str2];
    for s in strs {
        println!("{}", &s)
    }
}

fn main() {
    greet_world();

    println!("Hello, world!");

    let text = "\
    Rust 这门语言对于 Haskell 和 Java 开发者来说，
    可能会觉得很熟悉，因为它们在高阶表达方面都很优秀。
    简而言之，就是可以很简洁的写出原本需要一大堆代码才能表达的含义。
    但是，Rust 又有所不同：它的性能是底层语言级别的性能，可以跟 C/C++ 相媲美。
    ";

    let lines = text.lines();
    for (i, s) in lines.enumerate() {
        println!("{} -> {}", i, &s);

        let coll: Vec<_> = s.split('，')
            .map(|x| {
                x.trim()
            })
            .collect();

        println!("Size: {}", coll.len());
    }
}
