---
title: "一款好用的golang初始化工具"
date: 2023-03-19T20:03:04+08:00
draft: false
image: "img/GI4wXur48BY.jpg"
categories: 
  - golang
tag:
---


## 1.简介

用过java spring的同学应该对依赖注入这个概念不陌生，通过把类和对象托管给容器，在使用时只需要使用注解申明即可，这可以避免写很多的`A a = new A()`这种代码。

wire就是一个实现依赖注入的工具，不同之处在于: wire不是在运行时通过反射动态地进行依赖注入，而是在编译期通过**静态代码生成**的方式达成这个目的。

我们只需要在一个`.go`文件中声明各个组件的依赖关系，wire就能够自动完成实例的创建和注入，这能够极大地降低我们开发的工作量。

[地址](https://github.com/google/wire)

## 2.两个概念

wire中有两个很重要的概念: Provider和Injector。

### 2.1 Provider

从名称上可大概猜出来，Provider的主要功能应该是提供某些对象。实际上，在golang中Provider可以简单地理解为生成对象的工厂函数或者构造函数(类似于`NewXXX()`这种)。

使用时遵守一些简单的规则即可:

1. Provider函数必须是导出函数(即首字母大写，因为其他地方会使用到)
2. Provider通过参数来表示各个组件的依赖关系
3. Provider可以返回error对象
4. Provider通过进行分组，例如通过`wire.NewSet()`将几个一起使用的provider划分到一个组
5. **任意两个Provider不能返回相同类型的值**

例如:

```golang
package main

type TypeA struct {
  B *TypeB
}
type TypeB struct {}

func NewTypeA(b *TypeB) *TypeA{
  return &TypeA{B: b,}
}

func NewTypeB() *TypeB{
  return &TypeB{}
}

```

### 2.2 Injector

上述Provider只是一个简单的生成对象的函数，它告诉应用程序在需要某些对象时应该如何获取。但为了实现依赖注入，程序还需要知道何时使用这些Provider函数以及各个Provider函数之间的调用顺序，这就需要Injector来发挥作用。

Injector实际就是一个能够根据各个组件的依赖关系按顺序调用Provider的函数。下面是一个简单的Injector声明。

```golang
// +build wireinject
package main

func Inject() *TypeA {
	wire.Build(NewTypeA, NewTypeB)
	return &TypeA{}
}
```

需要注意的点:

1. 文件首行通过`// +build wireinject`声明
2. Inject函数体调用`wire.Build`，函数参数为对应的Provider函数
3. 运行前通过`wire`命令生成代码
4. Inject函数的返回值并不重要，其仅仅作为函数签名，只需要保证返回值类型正确即可(例如上面的Inject函数，可以返回&TypeA{}，也可以返回nil)

下面是上述案例生成的Inject函数体:

```golang
func Inject() *TypeA {
	typeB := NewTypeB()
	typeA := NewTypeA(typeB)
	return typeA
}
```

## 3.高级特性

### 3.1 接口绑定

有时候为了便于拓展，我们会希望基于接口编程。比如将参数声明为某个接口类型，但是提供具体类型的构造函数，wire对这种场景提供了支持。

使用时需要在`wire.NewSet`中使用`wire.Bind`对接口和具体实现进行绑定。

```golang
var set = wire.NewSet(NewTypeA, 
NewTypeB, 
wire.Bind(new(ITypeA), new(*TypeA)))

func Inject() *TypeB {
	wire.Build(set)
	return &TypeB{}
}
```

### 3.2 构造结构体

wire提供了`wire.Struct`函数构造结构体对象，通过该函数可以指定哪些字段需要被注入。例如下面的案例中，wire.Struct的第一个参数指定了待注入的结构体实例指针，后面的参数指定该结构的哪些字段需要被注入(MyFoo和MyBar)。

```golang
type Foo int
type Bar int

func ProvideFoo() Foo {
	return Foo(0)
}

func ProvideBar() Bar {
	return Bar(0)
}

type FooBar struct {
	MyFoo Foo
	MyBar Bar
}

var Set = wire.NewSet(
	ProvideFoo,
	ProvideBar,
	wire.Struct(new(FooBar), "MyFoo", "MyBar"))

func Inject() FooBar {
	wire.Build(Set)
	return FooBar{}
}
```

> **可以指定为`*`，表示注入所有字段，如`wire.Struct(new(FooBar), "*"))`，对于一些不想要被注入的字段，可以在结构体tag中加入`wire:"-"`**

### 3.3 值绑定

有时候想要将基本类型的值绑定到结构体中，而不想要提供对应的Provider函数(因为是一次性使用的，没有必要)，wire提供了`wire.Value`来完成此类绑定关系。

```golang
type Foo struct {
	X int
}

func Inject() Foo {
	wire.Build(wire.Value(Foo{X: 42}))
	return Foo{}
}
```

### 3.4 获取结构体的值

有时候我们希望能够将结构体的某些字段提供出去。一般情况下，为了达到这个目的，可以先获取到该结构体，然后调用对应的GetXXX方法(或者直接访问字段)，但是这样我们得实现Get方法。

```golang
type Foo struct {
	S string
	N int
	F float64
}

func getS(foo Foo) string {
	// Bad! Use wire.FieldsOf instead.
	return foo.S
}

func provideFoo() Foo {
	return Foo{S: "Hello, World!", N: 1, F: 3.14}
}

func injectedMessage() string {
	wire.Build(
		provideFoo,
		getS)
	return ""
}
```

wire提供了`wire.FieldsOf`函数，能够达成同样的目的。

```golang
type Foo struct {
	S string
	N int
	F float64
}

func provideFoo() Foo {
	return Foo{S: "Hello, World!", N: 1, F: 3.14}
}

func injectedMessage() string {
	wire.Build(
		provideFoo,
		wire.FieldsOf(new(Foo), "S"))
	return ""
}
```

### 3.5 一个使用Injector的小技巧

在使用Injector时，需要返回正确的类型，这样做很麻烦，尤其是返回多个值时(这些值通常是没有意义的)，可以简单地在`wire.Build`后调用panic。这样不会影响wire生成的代码。

## 4. 总结

wire通过代码生成的方式实现依赖注入，这依赖于两个很重要的概念，即Provider和Injector。

可以理解为: 

1. Provider用于提供依赖注入过程中需要的对象，它将所有的对象放到一个大池子里面，供Injector使用；
2. Injector能够根据各个对象的依赖关系，将池子中的对象组合起来，**根据Inject函数(即调用wire.Build函数的函数)的返回值决定将池子中哪个对象(或对象的字段)返回出来**


