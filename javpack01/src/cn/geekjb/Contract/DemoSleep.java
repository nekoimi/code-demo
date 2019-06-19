package cn.geekjb.Contract;

public class DemoSleep implements SleepContract {
    @Override
    public void sleep() {
        System.out.println("Sleep...");
    }
}
