---
title: "BTC06 挖矿难度"
date: 2023-05-28T13:46:55+08:00
draft: false
image: ""
categories: 
tag:
---

> 整理自 [北京大学肖臻老师《区块链技术与应用》公开课](https://www.bilibili.com/video/BV1Vt411X7JF?from=search&seid=14488407572640514229)

### 什么是调整挖矿难度？

挖矿就是不断尝试nonce，使整个block header中的哈希值小于等于给定的**目标阈值 target**。H(block header) <= target。

target越小，挖矿难度越大。调整挖矿难度，就是调整目标空间在整个输出空间中所占的比例。比特币中的hash算法是SHA-256，整个输出空间是2^256个可能取值。通俗的说，调整目标空间占整个输出空间的比例，就是哈希值前面要多少个0，但不是完全准确。

挖矿难度 difficulty与目标阈值 target成反比。$difficulty = \frac{ difficulty1\\_target(挖矿难度为1时所对应的阈值) }{ target(当前的目标阈值) }$，difficulty最小就是1。

### 为什么调整挖矿难度？

**系统总算力越来越强，如果难度不变，出块时间将会越来越短。**一个区块传播给大多数节点所需要的时间大约是几十秒，**如果出块时间缩短，如几秒钟一个，那么会很容易出现分叉。分叉过多，系统不易达成共识，威胁系统安全性。**

假设大部分算力掌握在诚实的矿工手里。系统中总算力越强，安全性越好，因为要发动51%攻击所需要的算力也就越大。但是如果出现过多分叉，系统中的总算力就会被分散，恶意节点集中算力扩展自己的分叉，很快就可以成为最长链，此时可能不再需要51%算力才能发动攻击，攻击成本大大降低。

![img](https://gitee.com//tiansir-wg/blogimg/raw/master/imgs/20200608121637.png)

综上，出块时间不是越小越好，因此要通过调整挖矿难度进而调整出块时间

### 如何调整挖矿难度？

**每2016个区块调一下阈值 target。**

$\frac{2016 \times 10min}{60min \times 24h}= 14天$，也就是每14天调整一下阈值 target。
$$
target = target \times \frac{actual\\_time}{expected\\_time}
$$

$$
difficulty = difficulty \times \frac{expected\\_time}{actual\\_time}
$$



其中`actual_time`就是指的系统中产生的最近的2016个区块实际花费的时间；`expected_time`就是2016*10min，也就是14天。

如果最近2016个区块产生的时间超过14天，说明平均每一个区块的出块时间超过10min，这时候挖矿难度应该降低。如果实际时间小于14天，说明出块太快，这时候应该提高挖矿难度。实际代码中，上调或下调都有4倍的限制。比如实际时间非常长，超过8周，那么算的时候也按照8周来算，最多增大4倍。相反如果实际时间很短，不到半周，那么算的时候按照半周来算，最小也是1/4。

### 如何让所有矿工同时调整目标阈值？

计算target的方法是写在比特币系统的代码中，每挖到2016个区块，会自动进行调整。如果有的恶意节点故意不调，那么发布的区块，诚实矿工不会认同。block header中不直接存储target，因为它是256位的，nBits域在header中只有32位，可以认为是target的压缩编码。如果有恶意矿工不调难度，那么在检查区块合法性(nBits域)时就不会通过。

```cpp
class CBlockHeader
{
public:
    // header
    int32_t nVersion;
    uint256 hashPrevBlock; //前一个区块的hash
    uint256 hashMerkleRoot;
    uint32_t nTime;
    uint32_t nBits; //难度值
    uint32_t nNonce; //随机数
```
