abstract class Dog {
    constructor(name: string) {
        this.name = name
    }
    public name: string
    private index: number = 233
    protected pro: number = 233
    readonly log: string = "DOG"
    static staconst: string = "DDOG"
    run() {}
    abstract say(): void
}

class Hasky extends Dog {

    constructor(name: string, color: string) {
        super(name)
        this.color = color
    }
    private color: string

    say(): void {
        console.log("wang!wang!")
    }

}

let hasky = new Hasky("hhh", "2")
hasky.say()