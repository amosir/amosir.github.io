---
title: "ETH17 以太坊的挖矿算法"
date: 2023-05-28T13:51:40+08:00
draft: true
image: ""
categories: 
tag:
---


> 整理自 [北京大学肖臻老师《区块链技术与应用》公开课](https://www.bilibili.com/video/BV1Vt411X7JF?from=search&seid=14488407572640514229)

### 设计mining puzzle的目标

对于基于工作量证明的区块链系统来说，挖矿是保障区块链安全的一个重要手段。总的来说，比特币的挖矿算法是比较成功的，目前为止没有发现大的漏洞。但是比特币的挖矿算法也有一些需要改进的地方，其中一个问题就是挖矿设备的专业化: 用普通的计算机挖不了矿，只能用专业的asic芯片来挖矿。有人认为这种做法与去中心化的理念不符，也和比特币的设计初衷相违背，中本聪最早的比特币论文提出了一个说法叫 **one cpu,one vote**。理想状况下，应该让普通人也能够参与挖矿过程，用自己的电脑手机来挖矿，这样也更安全，因为算力分散之后，恶意节点想要聚集到51%的算力发动攻击难度很大，所以比特币之后出现的加密货币包括以太坊设计mining puzzle的时候一个目标就是**要做到ASIC resistance**。

### 具体方案

想要设计一个ASIC resistance的mining puzzle，常用做法就是增加mining puzzle对内存访问的需求，也就是所谓的**memory hard mining puzzle**。因为ASIC芯片相对于普通计算机而言主要优势是算力强，但是在内存访问的性能上没有什么大的优势。所以如果能够设计一个对内存要求高的puzzle，就能起到遏制ASIC芯片的作用。这方面早期例子就是LiteCoin(莱特币)，它的puzzle是基于scrypt，这是一个对内存要求很高的哈希函数，它的设计思想是开设一个很大的数组，然后按照顺序填充一些伪随机数。有一个种子节点seed，把seed的值通过运算算出一个值，填在第一个位置，然后后面每个位置都是前一个位置的值取哈希得到的。之所以叫做伪随机数，因为哈希取完之后算出来的值看上去好像是随机数，但实际上不可能真的用随机数，否则没有办法验证。

<img src="https://gitee.com//tiansir-wg/blogimg/raw/master/imgs/20200613123644.png" style="zoom:50%;" />

数组填充的一个特点是里面的取值是有前后关系的，是从第一个数依次算出来的，然后需要求解puzzle的时候，按照伪随机的顺序从数组当中读取一些数，每次读取的位置也是和前一个数相关的。比如要解puzzle，一开始读取A位置的数，把A位置的数读取出来之后，根据取值进行一些运算算出下一次要读取的位置(假设为B)，然后把B位置的数读取出来， 再经过一些运算算出下一个要读取的位置(C位置)。这也是一种伪随机的顺序，因为经过哈希运算之后得到下一个读取的位置，这样做的好处是如果数组开的足够大，那么对于挖矿的矿工来说就是memory hard，因为如果不保存数组，那么挖矿的计算复杂度会大幅度上升。所以要想高效的挖矿这个内存区域是需要保存的。



### 优点和缺点

 这样设计的**好处**是对矿工来说挖矿的时候是memory hard，不好的地方是对于轻节点来说也是memory hard。设计puzzle的原则是**difficult to solve，but easy to verify**。该设计的问题就在于验证puzzle需要的内存区域跟求解puzzle需要的内存区域几乎是一样大的。轻节点要验证也需要保存数组，不然计算复杂度也是大幅度提高。这样造成的结果是莱特币设计的时候数组不敢设计的太大， 实际莱特币在使用的时候数组只有128kb，非常小，就是为了照顾轻节点。实验证明，莱特币要求的128K内存不足以对ASIC芯片的生产造成障碍，从这一点来说莱特币的目标没有达到。

### 改进

以太坊使用的也是memory hard的mining puzzle，但是在设计上与莱特币有很大的不同，以太坊用的是两个数据集，一大一小。小的是一个16M的cache，大的数据集是一个1G的dataset(DAG)， 1G的数据集是从16M的cache生成出来的。之所以设计成一大一小的两个数据集，就是为了便于验证，轻节点只要保存16M的cache就好，只有需要挖矿的矿工才需要保存1G的大数据集。基本思想是: 小的数据集cache的生成方式和前面数组的生成方式是类似的，首先从种子节点经过一些运算算出数组第一个元素，然后依次取哈希，把整个数组填充伪随机数，就得到一个cache。以太坊生成一个更大的数组，比小数组大得多。而且小的cache和大的dataset都是定期增长的，每隔一段时间就要增大，因为计算机的内存容量也是定期增长的。大数据集每个元素都是从小cache里按照伪随机的顺序读取一些元素，方法和莱特币求解puzzle的过程类似， 比如第一次实际上是读取A位置的元素，对当前哈希值进行更新迭代，算出下一个要读取的位置B，然后是C......从cache中来回读256次，读256个数，最后算出来一个数，放到大dataset中的第一个元素，第二个元素生成采用一样的方式......然后求解puzzle的时候用的是大数据集中的数，cache是不用的，按照伪随机的顺序从dataset中读取128个数。

![image-20200613133216382](https://gitee.com//tiansir-wg/blogimg/raw/master/imgs/20200613133216.png)

到目前为止以太坊的挖矿是以GPU为主，用ASIC矿机的很少，从这一点来说要比莱特币成功，起到了ASIC resistance的作用，这个和以太坊挖矿算法需要的大内存是很有关系的(算法叫做ethash)。以太坊没有出现ASIC矿机还有一个原因就是以太坊很早就计划要从工作量证明转向权益证明。这个对ASIC芯片厂商来说是很大的威胁，因为ASIC芯片的研发周期是很长的，一年的周期就算是快的了，而且研发成本高，等到以太坊转向权益证明之后就不挖矿了，那么这些投入的研发费用都白费了。但是以太坊到现在为止还是工作量证明，很早就说转移到权益证明，但是转移的时间点一再推迟，同时不停地宣称要转向权益证明。 通过吓唬起到ASIC resistance的作用。



下面是生成数据集的代码实现:

1. 生成16M的cache

   ```python
   def mkcache(cache_size, seed):
   	o = [hash(seed)]    # 生成o这个集合
   	for i in range(1, cachesize):   # for循环添加元素
   		o.append(hash(o[-1]))         # hash(o[-1])是前一个元素的hash值
   return o
   ```

   

2. 通过cache生成dataset的第i个元素

   ```python
   def calc_dataset_item(cache, i):
   	cache_size = cache.size
   	mix = hash(cache[i % cache_size] ^ i)
   	for j in range(256):
   		cache_index = get_int_from_item(mix)     # 用当前值读出下一个位置
   		mix = make_item(mix,cache[cache_index % cache_size])
   	return hash(mix)
   ```

3. 多次调用上述函数就可以生成完整的dataset



下面是矿工挖矿的代码实现:

```python
# header是块头、nonce是当前尝试的随机值、full_size是dataset中元素的个数
def hashimoto_full(header, nonce, full_size, dataset):
    mix = hash(header, nonce)
    for i in range(64):
        dataset_index = get_int_from_item(mix) % full_size
        mix = make_item(mix, dataset[dataset_index])
        mix = make_item(mixdataset[dataset_index + 1])
    return hash(mix)
  
def mine(full_size, dataset, header, target):
    nonce = random.randint(0, 2**64)
    while hashimoto_full(header, nonce, full_size, dataset) > target:
      nonce = (nonce + 1) % 2**64
    return nonce
```

下面是轻节点验证的代码实现:

```python
def hashimoto_light(header, nonce, fuLl_size, cache):
		mix = hash(header, nonce)
		for i in range(64):
				dataset_index = get_int_from_item(mix) % full_size
				mix = make_item(mix, calc_dataset_item(cache, dataset_index))
				mix = make_item(miX, calc_dataset_item(cache, dataset_index + 1))
		return hash(mix)
```



上述挖矿的算法设计目标是要尽可能让通用的计算设备也能参加，参加的人越多，挖矿的过程越民主，那么区块链就越安全，这也是为什么莱特币、以太坊要设计memory hard mining puzzle。但是也有不同的观点，认为让通用计算设备参与挖矿反而不安全。像比特币那样能用专门的ASIC芯片挖矿的更安全，因为假如要对比特币系统发动攻击，需要投入大量资金买入ASIC矿机，这样才有攻击所需要的算力， 而这些矿机除了挖矿做不了别的事情，而且是为某一个加密货币设计的挖矿芯片只能挖一种加密货币， 所以发动攻击的成本很高，而且一旦攻击成功，比特币系统安全性被证明存在问题，比特币价格下降，这样早期投入的硬件成本收不回来。相反如果让通用计算设备参与挖矿，那么发动攻击的成本大幅度下降，因为没有必要为了发动攻击专门购买特制的硬件设备，比如很多大的互联网公司服务器很多，那么需要发动攻击的时候需要把这些服务器调动起来都来挖矿就可以，平时服务器可以满足公司日常业务而服务，攻击的时候临时征用用来挖矿，攻击结束之后又可以恢复原来的任务。即使大互联网公司不做这件事情，恶意攻击者也有可能通过租用服务器集群达到同样目的。

### 问题

> 为什么全节点要保存dataset而轻节点只需要保存cache就行?

由于矿工需要验证非常多的nonce,如果每次都要从16M的cache中重新生成的话，那挖矿的效率就太低了，而且这里面有大量的重复计算：随机选取的dataset的元素中看很多是重复的，可能是之前尝试别的nonce时用过的。所以，矿工采取以空
间换时间的策略，把整个dataset保存下来。轻节点由于只验证一个nonce，验证的时候就直接生成要用到的dataset中的元素就行了。