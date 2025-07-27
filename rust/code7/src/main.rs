use std::fmt::{Display, Formatter};

// #[derive(Debug)]
struct User {
    username: String,
    email: String,
    sign_in_count: u64,
    active: bool,
}

impl Display for User {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        f.write_str(&self.username)
    }
}

fn create_user(username: String, email: String) -> User {
    User {
        username,
        email,
        sign_in_count: 0,
        active: false,
    }
}

struct Point2D(f32, f32);
struct Run;

fn main() {
    println!("Hello, world!");

    // 结构体
    let current_user = User {
        username: String::from("张三"),
        email: String::from("aaa@aaa.com"),
        sign_in_count: 1,
        active: true,
    };
    println!("{}", current_user);

    let user2 = create_user(String::from("李四"), String::from("lisi@qqq.com"));
    println!("{}", user2);

    // 结构体合并更新
    let user3 = User {
        username: String::from("王五"),
        ..user2
    };
    println!("{}", user3);

    // 把结构体中具有所有权的字段转移出去后，将无法再访问该字段，但是可以正常访问其它的字段。
    println!("user3 email address: {}", user3.email);
    println!("user2 username: {}", user2.username);
    // println!("user2 email address: {}", user2.email);  // error, value move to user3
}
