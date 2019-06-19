package cn.geekjb.strategy;

import cn.geekjb.Contract.DemoEat;
import cn.geekjb.Contract.DemoSleep;

public class Boy extends People {

    public Boy () {
        eatContract = new DemoEat();
        sleepContract = new DemoSleep();
    }

    @Override
    public void display() {
        System.out.println("Boy display...");
    }

}
