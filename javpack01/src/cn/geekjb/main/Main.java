package cn.geekjb.main;

import cn.geekjb.strategy.Boy;
import cn.geekjb.strategy.Girl;

public class Main {

    public static void main(String[] args) {

        Boy boy = new Boy();
        Girl girl = new Girl();

        boy.SayHello();
        boy.swim();
        boy.display();
        boy.eat();
        boy.sleep();

        girl.SayHello();
        girl.swim();
        girl.display();

    }

}
