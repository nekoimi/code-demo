# Symfony Components —— Asset 组件

> Asset组件管理URL的生成以及web assets的版本，诸如CSS样式表，JavaScripts以及图片文件等。

### 问题

- html 模板中静态文件链接

- 没有版本控制

- 静态资源的地址是死的

- 不能使用不同的存储服务

### 安装

```bash
composer require symfony/asset -vvv
```

### Master

> Asset组件通过包（packages）来管理资源（assets）。一个包，集合了全部资源，这些资源共享着相同属性：版本策略，基本路径（base path），CDN主机等。

### PackageInterface

`Package`类扩展至这个接口

```php
interface PackageInterface
{
    /**
     * Returns the asset version for an asset.
     *
     * @param string $path A path
     *
     * @return string The version string
     */
    public function getVersion($path);

    /**
     * Returns an absolute or root-relative public path.
     *
     * @param string $path A path
     *
     * @return string The public path
     */
    public function getUrl($path);
}
```

#### Package

基本类库，继承自 `PackageInterface`，提供基本的静态资源管理

- 基础使用

```php
require_once __DIR__ . '/../vendor/autoload.php';

use Symfony\Component\Asset\Package;
use Symfony\Component\Asset\VersionStrategy\EmptyVersionStrategy;

$asset = new Package(
    new EmptyVersionStrategy()
);

echo $asset->getUrl('style.css');
// style.css
```

- 版本化管理

```php
require_once __DIR__ . '/../vendor/autoload.php';

use Symfony\Component\Asset\Package;
use Symfony\Component\Asset\VersionStrategy\StaticVersionStrategy;

$asset = new Package(
    new StaticVersionStrategy('v1', '%2$s/%1$s')
);

echo $asset->getUrl('style.css');
// v1/style.css
```

#### UrlPackage

继承上面的`Package` 类

构造函数里面多了一个 `baseUrls` 的参数

```php
public function __construct($baseUrls, VersionStrategyInterface $versionStrategy, ContextInterface $context = null)
```

#### PathPackage

继承上面的`Package` 类

构造函数里面多了一个 `basePath` 的参数

```php
public function __construct($basePath, VersionStrategyInterface $versionStrategy, ContextInterface $context = null)
```

### Eg

- 使用CDN

```php
require_once __DIR__ . '/../vendor/autoload.php';
use Symfony\Component\Asset\UrlPackage;
use Symfony\Component\Asset\VersionStrategy\StaticVersionStrategy;

$asset = new UrlPackage(
    '//www.baidu.com',
    new StaticVersionStrategy('v1', '%2$s/%1$s')
);

echo $asset->getUrl('style.css');
# //www.baidu.com/v1/style.css
```

### 参考

[文档 CN](http://www.symfonychina.com/doc/current/components/asset.html)
[文档 EN](https://symfony.com/doc/current/components/asset.html)

