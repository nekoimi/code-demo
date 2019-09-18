let bool: boolean = true
let num: number = 233
let str: string = "string"

// 数组
let arr1: number[] = [1, 2, 3]
let arr2: string[] = ['a', 'b', 'c']
let arr3: Array<number> = arr1
let arr4: Array<number | boolean> = [1, 2, true]

// 元组
let tuple: [string, boolean] = ["string", true]
// tuple.push("aaa")

// 函数
let add = (x: number, y: number): number => x + y
let func: (x: number, y: number) => number
// exec
func = (a, b) => a + b

// 对象
// FIXME 错误的用法
// let obj: object = {
//     x: 1,
//     y: 2
// }
// obj.x = 100
let obj2: {
    x: number,
    y: number,
    name: string
} = {
    x: 1,
    y: 2,
    name: "zhangsan"
}

obj2.x = 100

let sym: symbol = Symbol()

// undefined, null
let un: undefined = undefined
let nu: null = null
let num3: number | undefined = un

// void
let noReturn = () => {}

// any
let anyVar: any = 1
anyVar = "var"

// never
// 永远不会有返回指
// Error
// Loop


// Enum

// 数字枚举
enum Role{
    Super,
    Master,
    Admin,
    Member
}
console.log(Role)
console.log(Role.Super)

// 字符串枚举
enum Status {
    Success = "Success !",
    Fail = "Fail !"
}
console.log(Status)
console.log(Status.Success)

// 常量枚举
// fixme 常量枚举编译之后会被删除 不会保留在代码中
const enum ListAbc {
    A,
    B,
    C
}
