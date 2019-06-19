### 策略模式

> 
> 通过面向对象的角度设计超类和扩展类
>

- 超类一般是一个抽象类

- 在子类中实现这些抽象方法

在普通的继承关系中出现的问题:

- 对类的局部改动，尤其是对超类的改动会影响到子类的实现

- 影响存在溢出效应

实例

父类:

```java
package cn.geekjb.strategy;

public abstract class People {
    
    public People(){
    }
    
    public void SayHello () {
        System.out.println("Hello ~");
    }
    
    public abstract void display();

    public void swim() {
        System.out.println("swim ~");
    }

}
```


子类:

```java
package cn.geekjb.strategy;

public class Boy extends People {

    @Override
    public void display() {
        System.out.println("Boy display...");
    }

}
```

子类:

```java
package cn.geekjb.strategy;

public class Girl extends People {

    @Override
    public void display() {
        System.out.println("Girl display...");
    }

}
```

#### 分析

- 根据需求，分析不变与变的部分  =>  抽象 + 就扣实现

#### 演变

策略模式:

>
> 封装行为为接口，实现算法族，超类里面固定行为接口对象
>
> 在子类具体设置行为对象
>
>

原则:

> 分离变化部分，封装接口，基于接口编程。
