---
title: "List包源码解析"
date: 2022-05-01T15:38:18+08:00
draft: false
image: "img/mpwF3Mv2UaU.jpg"
categories: 
  - golang
  - 源码学习
tag:
---

## 1.1 概览

这个包是对双向链表的实现，该链表的特点是:

- 存在一个固定的头结点
- 链表首尾相连

下面是该结构的示意图

## 1.2 结构定义

下面是链表结构

```Go
type List struct {
  root Element // sentinel list element, only &root, root.prev, and root.next are used
  len  int     // current list length excluding (this) sentinel element
}
```

整个链表包括一个头节点对象 root 和链表的长度，从这里可知**获取 list 长度并不需要从头开始遍历**。

链表节点的定义如下:

```Go
type Element struct {
  // Next and previous pointers in the doubly-linked list of elements.
  // To simplify the implementation, internally a list l is implemented
  // as a ring, such that &l.root is both the next element of the last
  // list element (l.Back()) and the previous element of the first list
  // element (l.Front()).
  next, prev *Element

  // The list to which this element belongs.
  list *List

  // The value stored with this element.
  Value any
}
```

除了必须的节点左右指针和节点值外，还多了一个指定当前链表的指针，这样从每个节点都能够获取关于 list 的信息。

## 1.3 函数

### 1.3.1 构造器

调用 list.New 函数创建一个双向链表，在创建时会进行链表的初始化

```Go
func New() *List { return new(List).Init() }

```

### 1.3.2 Init 方法

Init 函数可用于**初始化链表或者清空链表**

```Go
func (l *List) Init() *List {
  l.root.next = &l.root
  l.root.prev = &l.root
  l.len = 0
  return l
}
```

从这里也能看出，**双向链表结构中存在一个占位的头结点，并且头结点不计入总长度**

### 1.3.3 Len 方法

由于 List 结构中有单独的字段存储了当前链表的长度，所以直接返回该字段即可，不需要从头开始遍历。

```Go
func (l *List) Len() int { return l.len }

```

### 1.3.4 Front 和 Back 方法

分别用于获取链表的头结点和尾结点

```Go
func (l *List) Front() *Element {
  if l.len == 0 {
    return nil
  }
  return l.root.next
}

// Back returns the last element of list l or nil if the list is empty.
func (l *List) Back() *Element {
  if l.len == 0 {
    return nil
  }
  return l.root.prev
}
```

### 1.3.5 InsertBefore 和 InsertAfter 方法

分别用于在指定的节点前和节点后插入新的节点

```Go
func (l *List) InsertBefore(v any, mark *Element) *Element {
  if mark.list != l {
    return nil
  }
  // see comment in List.Remove about initialization of l
  return l.insertValue(v, mark.prev)
}

// InsertAfter inserts a new element e with value v immediately after mark and returns e.
// If mark is not an element of l, the list is not modified.
// The mark must not be nil.
func (l *List) InsertAfter(v any, mark *Element) *Element {
  if mark.list != l {
    return nil
  }
  // see comment in List.Remove about initialization of l
  return l.insertValue(v, mark)
}
```

底层都是用 insertValue 方法实现的，这个方法在其他导出函数中也有用到。

基本原理就是先构造一个新的节点，然后将新节点插入到指定位置，最后更新链表长度。

从 insert 的实现可知: **整个链表是一个首尾相连的环形结构**

```Go
// insert inserts e after at, increments l.len, and returns e.
func (l *List) insert(e, at *Element) *Element {
  e.prev = at
  e.next = at.next
  e.prev.next = e
  e.next.prev = e
  e.list = l
  l.len++
  return e
}

// insertValue is a convenience wrapper for insert(&Element{Value: v}, at).
func (l *List) insertValue(v any, at *Element) *Element {
  return l.insert(&Element{Value: v}, at)
}
```

### 1.3.6 PushFront 和 PushBack 方法

分别用于在链表头和链表尾部插入新节点

```Go
// PushFront inserts a new element e with value v at the front of list l and returns e.
func (l *List) PushFront(v any) *Element {
  l.lazyInit()
  return l.insertValue(v, &l.root)
}

// PushBack inserts a new element e with value v at the back of list l and returns e.
func (l *List) PushBack(v any) *Element {
  l.lazyInit()
  return l.insertValue(v, l.root.prev)
}
```

### 1.3.7 MoveToBack 和 MoveToFront 方法

这两个函数内部实际都是调用了 move 方法，下面是该方法定义

move 方法可以将节点 e 移动到 at 节点后面

```Go
func (l *List) move(e, at *Element) {
  if e == at {
    return
  }
  // 将节点e摘掉
  e.prev.next = e.next
  e.next.prev = e.prev

  // 将e连接到at后面
  e.prev = at
  e.next = at.next
  e.prev.next = e
  e.next.prev = e
}
```

下面是 MoveToBack 和 MoveToFront 方法的定义。

```Go
// MoveToFront moves element e to the front of list l.
// If e is not an element of l, the list is not modified.
// The element must not be nil.
func (l *List) MoveToFront(e *Element) {
  if e.list != l || l.root.next == e {
    return
  }
  // see comment in List.Remove about initialization of l
  l.move(e, &l.root)
}

// MoveToBack moves element e to the back of list l.
// If e is not an element of l, the list is not modified.
// The element must not be nil.
func (l *List) MoveToBack(e *Element) {
  if e.list != l || l.root.prev == e {
    return
  }
  // see comment in List.Remove about initialization of l
  l.move(e, l.root.prev)
}
```

### 1.3.8 MoveBefore 和 MoveAfter

这两个函数分别用于将节点移动到指定节点前或指定节点后

```Go
// MoveBefore moves element e to its new position before mark.
// If e or mark is not an element of l, or e == mark, the list is not modified.
// The element and mark must not be nil.
func (l *List) MoveBefore(e, mark *Element) {
  if e.list != l || e == mark || mark.list != l {
    return
  }
  l.move(e, mark.prev)
}

// MoveAfter moves element e to its new position after mark.
// If e or mark is not an element of l, or e == mark, the list is not modified.
// The element and mark must not be nil.
func (l *List) MoveAfter(e, mark *Element) {
  if e.list != l || e == mark || mark.list != l {
    return
  }
  l.move(e, mark)
}
```

### 1.3.9 PushBackList 和 PushFrontList 方法

**PushBackList**方法用于将 other 链表从头结点开始用尾插法一个一个插入到原始链表的尾部

```Go
// PushBackList inserts a copy of another list at the back of list l.
// The lists l and other may be the same. They must not be nil.
func (l *List) PushBackList(other *List) {
  l.lazyInit()
  for i, e := other.Len(), other.Front(); i > 0; i, e = i-1, e.Next() {
    l.insertValue(e.Value, l.root.prev)
  }
}

```

**PushFrontList**方法用于将 other 链表从尾部一个一个用头插法插入到原始链表的头部

```Go
// PushFrontList inserts a copy of another list at the front of list l.
// The lists l and other may be the same. They must not be nil.
func (l *List) PushFrontList(other *List) {
  l.lazyInit()
  for i, e := other.Len(), other.Back(); i > 0; i, e = i-1, e.Prev() {
    l.insertValue(e.Value, &l.root)
  }
}
```

假定原始链表为: A1 - A2 - A3，other 链表为: B1 - B2 - B3

则 PushFrontList 的结果是: B1 - B2 - B3 - A1 - A2 - A3

PushBackList 的结果是: A1 - A2 - A3 - B1 - B2 - B3

