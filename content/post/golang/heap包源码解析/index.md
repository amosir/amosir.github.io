---
title: "Heap包源码解析"
date: 2022-05-02T20:12:12+08:00
draft: false
image: "img/YpUFf0kOWQ0.jpg"
categories: 
  - golang
  - 源码学习
tag:
---


## 1.1 概述

这个包定义了实现一个堆所需的结构和堆的操作方法

heap 包中提供了一个 Interface 接口，只要实现了这个接口，就能将该结构当做堆来使用。

```Go
type Interface interface {
  sort.Interface
  Push(x any) // add x as element Len()
  Pop() any   // remove and return element Len() - 1.
}
```

其中 Push 和 Pop 方法分别用于向堆中添加元素和移除堆顶元素。

由于堆在构建和维护过程需要对节点进行排序和移动位置，因此还需要实现 sort.Interface 接口

```Go
package sort

type Interface interface {

  Len() int

  Less(i, j int) bool

  Swap(i, j int)
}
```

总的来说，定义一个堆结构需要实现以下方法

- `Len() int`
- `Less(i, j int) bool`
- `Swap(i, j int)`
- `Push(x any)`
- `Pop() any`

## 1.2 非导出函数

首先需要说明的是，尽管大多数时候堆的底层是用数组实现的，但是在逻辑上还是将其视为一棵树。

heap 包中只有两个非导出函数，即:

- down: 将节点下沉
- up: 将节点上浮

这两个函数是堆的大多数导出方法的基础。

### 1.2.1 down 函数

该方法定义如下，主要作用是将当前节点下沉。
比如在小根堆中，删除了堆顶节点后，会将堆的底层数组的最后一个元素移动到堆顶。这种情况下，整个堆就不是一个小根堆了(但是堆顶节点的左右子树各自是小根堆)，需要将堆顶元素下沉。

```Go
func down(h Interface, i0, n int) bool {
  i := i0  // 当前位置
  for {
    j1 := 2*i + 1    // 当前位置节点的左子节点的位置
    if j1 >= n || j1 < 0 {   // 左子节点索引超过堆长度或溢出成负数
      break
    }
    j := j1 // left child
    if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {   // 右子节点比左子节点小
      j = j2 // = 2*i + 2  // right child
    }
    if !h.Less(j, i) {      // j表示左右子节点中更小的，如果当前节点i比左右子节点中最小的j还小，那么已经是小根堆了，就不用下沉了
      break
    }
    h.Swap(i, j)   // 否则，将当前节点i与j交换，将较小的节点交换到当前位置，当前节点下沉
    i = j    // 以左子节点(或右子节点)为当前节点，继续下沉
  }
  return i > i0   // 如果没有执行下沉操作，则返回false
}
```

### 1.2.2 up 函数

方法定义如下，主要作用就是将当前节点上浮。

```Go
func up(h Interface, j int) {
  for {
    i := (j - 1) / 2 // 父节点
    if i == j || !h.Less(j, i) {    // 当前节点比父节点大，不用上浮
      break
    }
    h.Swap(i, j)   // 当前节点比父节点小，将当前节点和父节点交换位置
    j = i    // 以父节点为当前节点继续上浮
  }
}
```

## 1.3 导出函数

### 1.3.1 Init 函数

Init 函数用于初始化一个堆。

```Go
func Init(h Interface) {
  // heapify
  n := h.Len()
  for i := n/2 - 1; i >= 0; i-- {
    down(h, i, n)
  }
}
```

### 1.3.2 Push 和 Pop 函数

这两个函数分别用于向堆中添加一个元素和移除堆顶元素(小根堆中是最小元素，大根堆中是最大元素)

底层实现依赖于 up 和 down 函数

```Go
func Push(h Interface, x any) {
  h.Push(x)    // 向底层数组的最后添加一个元素
  up(h, h.Len()-1)   // 添加完后可能会破坏堆结构，因此需要对新加的元素执行上浮操作
}

func Pop(h Interface) any {
  n := h.Len() - 1
  h.Swap(0, n)  // 将当前堆顶元素(底层数组第一个元素)和底层数组最后一个元素交换位置
  down(h, 0, n)    // 由于进行了交换，当前堆顶可能不再是最小元素了(小根堆中)，因此堆顶元素需进行下沉
  return h.Pop()  // 从底层数组返回并移除最后一个元素
}
```

### 1.3.3 Remove 函数

用于移除位置 i 处的元素。

该函数的主要思路和 Pop 类似，首先将当前位置和底层数组的最后一个位置进行交换，并通过上浮和下沉操作弥补交换导致的结构破坏，最后将底层数组的最后一个位置移除。

定义如下:

```Go
func Remove(h Interface, i int) any {
  n := h.Len() - 1
  if n != i {   // 如果要移除的元素不是最后一个元素
    h.Swap(i, n)  // 将位置i处的元素和底层数组最后一个元素交换
    // 从当前位置进行下沉
    // 1. 返回false，表示未执行下沉操作，这意味着交换后以当前节点为根的树构成小根堆。但是有可能会破坏其父节点及其祖先节点所在的树，所以还要进行上浮
    // 2. 返回true，表示执行了下沉操作，这意味着当前位置节点并不是最小的节点，最小的节点来自于原本的堆中，必然不可能大于当前位置父节点以前的节点，因此无需上浮
    if !down(h, i, n) {
      up(h, i)
    }
  }
  return h.Pop()
}
```

### 1.3.4 Fix 函数

用于在节点值发生变化后修补因此导致的堆结构破坏。

```Go
func Fix(h Interface, i int) {
  if !down(h, i, h.Len()) {
    up(h, i)
  }
}
```

## 1.4 使用案例

下面构造一个堆结构存储学生信息，根据年龄对信息进行排序，最后按年龄升序输出

```Go
package main

import (
  "container/heap"
  "fmt"
  "math/rand"
  "strconv"
)

type Stu struct {
  Name string
  Age int
}

// StuHeap is a heap of type stu
type StuHeap []Stu

// Len returns the length of heap
func (sh *StuHeap) Len() int{
  return len(*sh)
}

// Less return true if element i is less than j in heap
func (sh *StuHeap) Less(i, j int) bool{
  return (*sh)[i].Age < (*sh)[j].Age
}

// Swap swaps element in i and j of the head
func (sh *StuHeap) Swap(i, j int) {
  (*sh)[i], (*sh)[j] = (*sh)[j], (*sh)[i]
}

// Push pushes x to heap
func (sh *StuHeap) Push(x interface{}) {
  *sh = append(*sh, x.(Stu))
}

// Pop pops element from the tail of the heap
func (sh *StuHeap) Pop() interface{} {
  s := (*sh)[sh.Len() - 1]
  *sh = (*sh)[:sh.Len() - 1]
  return s
}

func main() {
  h := new(StuHeap)
  for i := 0;i < 10; i++ {
    t := rand.Intn(20)
    s := Stu{
      Name: "stu_" + strconv.Itoa(i),
      Age: t,
    }
    h.Push(s)
    fmt.Println(s)
  }
  fmt.Println("---")
  heap.Init(h)
  for i := h.Len();i > 0;i-- {
    fmt.Println(heap.Pop(h))
  }

}
```

## 1.5 总结

1. 定义堆结构需要实现 heap.Interface 中的所有方法，包括 Len、Swap、Less、Push 和 Pop，其中 Push 和 Pop 只能在底层数组尾部进行操作
2. 对堆的所有操作，包括 Push、Pop、Remove、Init、Fix 都是 heap 包的函数，**不要调用堆对象的方法**
3. 根据 Less 方法的结果确定是小根堆还是大根堆，返回 true 表示是小根堆
