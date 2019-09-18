console.log("interface")

interface List {
    id: number,
    name: string,
    age?: number
}

interface Result {
    data: Array<List>
}

function render(result: Result) {
    result.data.forEach(item => {
        console.log(item.id, item.name, item.age)
    })
}

let result = {
    data: [
        {id: 1, name: 'A', sex: 'man'},
        {id: 2, name: 'B', age: 18},
        {id: 3, name: 'C'}
    ]
}
render(result)

// FIXME 当后端返回的数据有多个字段的时候 这种传参报错
// render({
//     data: [
//         {id: 1, name: 'A', sex: 'man'},
//         {id: 2, name: 'B'},
//         {id: 3, name: 'C'}
//     ]
// })
// TODO 解决办法一 => 想上面那样定义一个变量传递
// TODO 解决办法二 => 类型断言
render({
    data: [
        {id: 1, name: 'A', sex: 'man'},
        {id: 2, name: 'B'},
        {id: 3, name: 'C'}
    ]
} as Result )
// render(<Result>{
//     data: [
//         {id: 1, name: 'A', sex: 'man'},
//         {id: 2, name: 'B'},
//         {id: 3, name: 'C'}
//     ]
// })
// TODO 解决办法三 => 字符串索引类型签名
// interface List {
//     id: number,
//     name: string,
//     [x: string]: any  // 用任意的字符串可以索引任意值
// }



/**
 * 函数接口
 */

// interface Div {
//     (a: number, b: number) : number
// }
// let d: Div = (a, b) => { return a - b;}

type Div = (a: number, b: number) => number

interface LibInterface {
    (): void,
    version: string
    doHello(): string
}

function getLib(): LibInterface {
    let lib: LibInterface = (() => {}) as LibInterface
    lib.version = "v1"
    lib.doHello = () => {
        return "Hello World"
    }
    return lib
}
