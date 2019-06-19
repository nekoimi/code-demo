package cn.geekjb.Contract;

public class DemoEat implements EatContract {
    @Override
    public void eat() {
        System.out.println("Eat...");
    }
}
