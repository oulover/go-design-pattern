# 简单工厂模式

`go` 语言没有构造函数一说，所以一般会定义 `NewXXX` 函数来初始化相关类。
`NewXXX` 函数返回接口时就是简单工厂模式，也就是说 `Golang` 的一般推荐做法就是简单工厂。
在这个 `simplefactory` 包中只有Car 接口和 `NewCar` 函数为包外可见，封装了实现细节。

The `golanguage` dose not have a constructor,so the `NewXXX` function is 
generally defined to initialize related classes, it's  a simple factory 
pattern when the `NewXXX` function returns an interface, that is to say, 
the general recommended practice of `Golang` is a simple factory.
In this `simplefactory` package only the Car interface and the `NewCar` 
function are visible outside the package, encapsulating the implementation 
details.

