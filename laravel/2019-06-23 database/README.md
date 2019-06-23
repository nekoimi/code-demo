# 数据库 - illuminate/database

你可以在一个空项目中使用 `composer` 来安装这个包

```bash
composer require illuminate/database -vvv
```

### 基本

`Laravel` 使用的 `IOC容器` 来管理依赖

想来之前很不理解这样做的目的

但是将一个类放到一个全局的数组中，需要这个类的实例的时候直接从数组中拿出来用确实比到处用 `new` 来获取类的实例优雅的多

> 这里的全局数组就充当了 `IOC容器`，所有的类、类的方法都可以用这种方式来管理

##### 提问: 在使用 Laravel 的时候，经常提到自动注入，这个是怎么实现的呢?
>
>
> 实现这个其实也很简单，基于 PHP 的 `反射` 可以在代码运行时为所欲为
>
> 可以拿到类的属性和方法(包括私有方法)
>
> 可以拿到类、方法的注释(这个就是现在很多注解框架的实现原理)
>
> 可以拿到方法M所需要的参数 (如果这些参数的类型是某个其他的类A，通过反射获取类A的参数然后实例化，将类A的实例放到方法M的参数里面，这样是不是就实现了自动注入参数了呢 ? )
>
> 

有兴趣可以去了解 PHP 的反射

- 以 Reflection 开头的类

### 正题

Laravel 中大多使用使用 ServiceProvider 来向容器注入一个对象

都继承自 `Illuminate\Support\ServiceProvider` 这个类

该类及其子类，只需要关注 `boot`、`register` 两个方法

##### 源码 `DatabaseServiceProvider`

- boot 方法

```php
Model::setConnectionResolver($this->app['db']);

Model::setEventDispatcher($this->app['events']);
```

这个暂且不看，不过这里的`Model`很明显是`Eloquent`

- register 方法

![20190623214327.png](https://i.loli.net/2019/06/23/5d0f82481508e71548.png)

在这里初始化了Models列表、向容器注册了链接和Eloquent

- registerConnectionServices 注册数据库链接服务

```php
// The connection factory is used to create the actual connection instances on
// the database. We will inject the factory into the manager so that it may
// make the connections while they are actually needed and not of before.
$this->app->singleton('db.factory', function ($app) {
    return new ConnectionFactory($app);
});

// The database manager is used to resolve various connections, since multiple
// connections might be managed. It also implements the connection resolver
// interface which may be used by other components requiring connections.
$this->app->singleton('db', function ($app) {
    return new DatabaseManager($app, $app['db.factory']);
});

$this->app->bind('db.connection', function ($app) {
    return $app['db']->connection();
});
```

哈！看到没有

常常使用的 `app('db')->xxx` 就是在这里注册进容器的

这个链接工厂类中定义各种方法

关于详细的数据库链接可以查看 `Illuminate\Database\Connectors\MySqlConnector` 这个类

链接使用的是 `PDO`

所有的 connection 类都继承自 `Illuminate\Database\Connectors\Connector`

还有其他的数据库类型支持: Postgres、SqlServer、SQLite

详细请自行查看源代码
