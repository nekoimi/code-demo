package cn.geekjb.strategy;

import cn.geekjb.Contract.EatContract;
import cn.geekjb.Contract.SleepContract;

public abstract class People {

    EatContract eatContract;
    SleepContract sleepContract;

    public People(){
    }

    public void SayHello () {
        System.out.println("Hello ~");
    }

    public abstract void display();

    public void swim() {
        System.out.println("swim ~");
    }

    public void eat() {
        eatContract.eat();
    }

    public void sleep() {
        sleepContract.sleep();
    }

}
