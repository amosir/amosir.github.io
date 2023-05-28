---
title: "ETH20 智能合约"
date: 2023-05-28T13:52:49+08:00
draft: false
image: ""
categories: 
tag:
---



> 整理自 [北京大学肖臻老师《区块链技术与应用》公开课](https://www.bilibili.com/video/BV1Vt411X7JF?from=search&seid=14488407572640514229)

### 概念

智能合约是以太坊的精髓，也是以太坊和比特币的主要区别。

智能合约的**本质**是运行在区块链上的一段代码，代码的逻辑定义了合约的内容。

智能合约的账户保存了合约当前的运行状态，包括balance 当前余额、nonce 交易次数、code 合约代码、storage 存储。数据结构是一棵MPT

Solidity是智能合约的最常用的语言，语法上与js接近。如下：



<img src="https://gitee.com//tiansir-wg/blogimg/raw/master/imgs/20200613222947.png" alt="image-20200613222947548" style="zoom:50%;" />



以太坊规定，如果合约账户能够接受外部转账的话，必须标注成**payable**。否则如果给函数转过去钱的话，会引发异常。



上面是一个网上拍卖的合约，`bid`函数是用来竞拍出价的，比如参与拍卖，要出100个以太币，就调用`bid`函数。拍卖的规则是调用`bid`函数的时候要把拍卖的出价也发送过去，存储到合约中，锁定到拍卖结束，避免有的人凭空出价。所以bid函数要有接受外部转账的能力。

`withdraw`函数用在拍卖结束，出价最高的人赢得拍卖，其他人没有拍到，其他人可以调用withdraw把自己出的价钱取回来。目的不是为了真的转账，不需要把钱转给智能合约，所以没必要用payable。

### 智能合约的调用

#### 外部账户调用智能合约

和转账类似。比如a发起一个交易转账给b，如果b是一个普通账户，那就是一个普通的转账交易，和比特币转账一样；如果b是一个合约账户，那么这个转账实际上是发起一次对b的合约的调用，具体调用的是合约中的哪个函数，是在数据域(data域)中说明的。

<img src="https://gitee.com//tiansir-wg/blogimg/raw/master/imgs/20200615153758.png" alt="image-20200615153758273" style="zoom:50%;" />



sender address是发起调用的账户的地址，to contract address是被调用合约的地址，调用的函数就是TX DATA给出的调用函数，函数有参数的话那么参数的取值也是在data域说明的。中间一行是调用的参数，value是说发起调用的时候转过去多少钱，gas used是这个交易花了多少gas fee，gas price是单位gas的价格，gas limit是最多愿意支付多少钱。

#### 一个合约调用另一个合约

##### 直接调用

```js
contract A {
		event LogCallFoo(string str);
		function foo(string str) returns (uint)(
						emit LogCallFoo(str);
						return 123;
		}
}
contract B {
		uint ua;
		function callAFooDirectly(address addr) public(
						A a = A(addr);
						ua = a.foo("call foo directly");
		}
}
```

##### address类型的call函数

```js
contract C {
		function callAFooByCall(address addr) public returns (bool)(
				bytes4 funcsig = bytes4(keccak256("foo(string)"));
				if (addr.call(funcsig,"call foo by func call"))
          	return true;
				return false;
		}
}
```

> 上面两种调用方式的一个区别是对错误处理的不同。
>
> 第一种方式如果被调用的合约在执行过程中出现错误，会导致发起调用的合约跟着一起回滚；
>
> 第二种方式如果在调用过程中被调用合约抛出异常，call函数会返回false表明调用失败，但是发起调用的函数并不会发生异常，而可以继续执行。

##### 代理调用delegatecall

* 此方式使用方法与call相同，只是不能使用.value()
* 区别在于是否切换上下文
  •	call()方式会切换到被调用的智能合约上下文中
  •	delegatecall()方式只使用给定地址的代码，其它属性(存储、余额等)都取自当前合约。delegatecall的目的是使用存储在另外一个合约中的代码。

##### 匿名函数fallback

```js
function() public [payable] {
  	...
}
```

匿名函数没有参数和返回值。

匿名函数被调用的情况:

1. 调用一个合约时要在转账交易的data域说明调用的是哪个函数，如果没有说明(data域是空)，那么缺省调用的是fallback函数。
2. 调用的函数不存在，也会调用fallback函数。

如果转账金额不是0，就需要声明**payable**， 否则会抛异常，一般情况都会设置成payable。

只有合约账户才有fallback和payable，外部账户没有。

转账金额可以是0，但是gas fee要有。转账金额是给收款人的，gas fee是给发布区块的矿工的。



### 智能合约的创建和运行

创建合约: 外部账户发起一个转账交易到0x0的地址

1. 转账金额是0，但是要支付gas fee
2. 合约的代码放在data域中

智能合约代码写完之后，要编译成bytecode，智能合约运行在EVM（Ethereum Virtual Machine）上。通过加一个虚拟机，对智能合约的运行提供一个一致性的平台。

以太坊是一个交易驱动的状态机，调用智能合约的交易发布到区块链上之后，每个矿工都会执行这个交易，从当前状态确定性地转移到下一个状态。

### gas fee

> 智能合约是一个图灵完备模型，全节点收到一个对智能合约的调用怎么知道该调用是不是死循环?

没有办法。因此智能合约引入了**gas fee机制**。



执行合约中的指令要收取汽油费，由发起交易的人来支付。

当一个全节点收到对智能合约的调用的时候，先按照调用中给出的gas limit算出可能花费的最大汽油费，然后一次性的把汽油费从发起调用的账户上扣掉，然后根据实际执行的情况算出实际花了多少汽油费，多余的汽油费会退回，如果不够的话会引起回滚。

EVM中不同指令消耗的汽油费是不一样的。简单的指令便宜(比如加法减法)，复杂的或者需要存储状态的指令就很贵(比如取哈希)



下面是一个交易的数据结构：

<img src="https://gitee.com//tiansir-wg/blogimg/raw/master/imgs/20200615094840.png" alt="image-20200615094840317" style="zoom:50%;" />

`AccountNonce`是交易序号，用来防止replay attack。

`GasLimit`是愿意支付的最大汽油量。

`Price`是单位汽油价格，`Price * GasLimit`就是可能消耗的最大汽油费。

`Recipient`是收款人的地址，`Amount`是转账金额。

`payload`就是data域。 



### 错误处理

以太坊中的交易(包含普通转账交易和对智能合约的调用)具有**原子性**，一个交易要么全部执行，要么全不执行，不会只执行一部分。如果在执行智能合约过程中出现任何错误，会导致整个交易的回滚，退回到执行之前的状态。



可能引起错误处理的情况:

1. gas fee不足，合同的执行要退回到之前的状态。(此时已经消耗的gas fee是不退的，不然可能会有恶意节点发动恶意攻击)
2.  assert(bool condition)条件不满足会抛出异常--用于内部错误
3. require(bool condition)条件不满足会抛掉异常--用于输入或者外部组件引起的错误
4. revert(): 终止运行并回滚状态变动--- 无条件抛出



### 嵌套调用

**嵌套调用**是指一个合约调用另一个合约的函数。 



嵌套调用是否会出现连锁式回滚，取决于调用智能合约的方式。如果是直接调用，那就会引发连锁式回滚，整个交易都会回滚；如果使用call方式，不会引起连锁式回滚，只会使当前调用失败，返回值false。

> 有些情况从表面上看并没有调用任何一个函数， 比如往一个账户转账，但是如果账户是合约账户的话，转账那个操作本身就有可能触发对函数的调用(因为有fallback函数)。



### 以太坊block header数据结构

```js
// Header represents a block header in the Ethereum blockchain.
type Header struct {
    ParentHash  common.Hash    `json:"parentHash"       gencodec:"required"`
    UncleHash   common.Hash    `json:"sha3Uncles"       gencodec:"required"`
    Coinbase    common.Address `json:"miner"            gencodec:"required"`
    Root        common.Hash    `json:"stateRoot"        gencodec:"required"`
    TxHash      common.Hash    `json:"transactionsRoot" gencodec:"required"`
    ReceiptHash common.Hash    `json:"receiptsRoot"     gencodec:"required"`
    Bloom       Bloom          `json:"logsBloom"        gencodec:"required"`
    Difficulty  *big.Int       `json:"difficulty"       gencodec:"required"`
    Number      *big.Int       `json:"number"           gencodec:"required"`
    GasLimit    uint64         `json:"gasLimit"         gencodec:"required"`
    GasUsed     uint64         `json:"gasUsed"          gencodec:"required"`
    Time        *big.Int       `json:"timestamp"        gencodec:"required"`
    Extra       []byte         `json:"extraData"        gencodec:"required"`
    MixDigest   common.Hash    `json:"mixHash"          gencodec:"required"`
    Nonce       BlockNonce     `json:"nonce"            gencodec:"required"`
}
```

其中`GasUsed`是该区块所有交易实际消耗的汽油费，`GasLimit`是当前区块所有交易能够消耗的汽油费的上限。



因为发布区块需要消耗一定的资源，对区块消耗的资源需要有一个限制。发布的区块如果没有任何限制，那么有的矿工可能把特别多的交易全打包到一个区块中，超大的区块在区块链上可能消耗很多资源。比特币中对发布的区块限制大小不能超过1M，比特币的交易是比较简单的，基本上用交易的字节数可以衡量出交易消耗的资源有多少。但以太坊中不行，因为以太坊中智能合约的逻辑很复杂，有的交易可能从字节数上看很小，但是消耗的资源很大，所以需要根据交易的具体操作来收费，就是所谓的gas fee。

### 问题

1. **全节点收到一个对合约的调用的时候，一次性地要把调用可能要花费的最大汽油费从发起调用的账户上扣掉，具体的操作如何执行?如果有多个全节点都收到交易是不是每个节点都要扣除一份汽油费？**

   以太坊有状态树、交易树、收据树。这三棵树都是全节点在本地维护的数据结构，状态树记录了账户的状态(包括账户余额)，所以在全节点收到调用的时候，在本地维护的数据结构里把账户余额减掉gas fee就可以，如果余额不够交易就不能执行，一次性要按gas limit把费用减掉，执行完之后如果有剩余就把余额加回相应数额。

   有多个全节点的情况下每个节点都扣除一份gas fee，只是在本地的数据结构扣除而已。智能合约执行过程中，任何对状态的修改都是再改本地的数据结构，只有在合约执行完并发布到区块上之后，本地的修改才会变成外部可见，成为区块链上的共识。

2. **假设某个全节点要打包一个交易到区块里，这些交易里有一些是对智能合约的调用，那么这个全节点是应该先把这些智能合约都执行完之后再挖矿，还是应该先挖矿获得记账权，然后再执行智能合约?**

   block header数据结构中的`Root`，`TxHash`，和`ReceiptHash`分别是三棵树的根哈希值，所以需要先执行完区块中的所有交易之后才能更新三棵树 ，才能确定这三个哈希值，然后才能尝试各个nonce。所以需要先执行智能合约，然后再挖矿。

3. **假设一个矿工消耗资源执行智能合约，最后挖矿没挖到怎么办？**

   没挖到矿就不会获得gas fee，因为gas fee是给获得记账权的矿工，以太坊中得不到任何补偿。不仅如此，而且还要把别人发布的交易在本地执行一遍，验证发布区块的正确性，每个全节点要独立验证。别人发布一个交易区块，把区块里的交易在本地执行，更新三棵树的内容，算出roothash，再与别人发布的roothash进行比较，看是否一致。所有这些都是免费的没有补偿。所以这种机制下挖矿慢的矿工就很吃亏，gas fee设置的目的是对于矿工执行智能合约所消耗的资源的一种补偿，但是只有挖到矿的矿工才能得到，其他矿工得不到。 此外，gas fee也是为了遏制发起调用的账户，如果不给gas fee那么账户可以随意发起调用。

4. **会不会有矿工觉得不给gas fee就不验证了？**

   出现这种情况最直接的后果是危害区块链的安全，区块链的安全就是要求所有全节点独立验证发布的区块的合法性，这样少数有恶意的节点才没有办法篡改区块链上的内容。如果某个矿工不验证， 那么之后就没法挖矿。所以没有办法跳过验证步骤。

5. **发布到区块链上的交易是否都是成功执行的？如果合约执行过程中出现了错误，要不要发布到区块链上？**

   执行错误的交易也要发布到区块链上，否则gas fee扣不掉，只在本地账户扣gas fee是没用的，需要发布出去形成共识(为什么扣钱，扣钱是否合理)，扣掉的gas fee才成为你账户上的钱，所以发布到区块链上的交易不一定都是成功执行的。

6. 如何知道一个交易是不是执行成功？

   每个交易执行完之后会形成一个收据，下面是收据的内容，其中`status`域就是说明交易执行情况如何。

   ```js
   	// Receipt represents the results of a transaction.
   type Receipt struct (	
   		// Consensus fields	
   		PostState	[]byte 'json:"root"'
   		Status	uint64 'json: ''status"'
   		CumulativeGasUsed	uint64 'json:"cumulativeGasUsed" gencodec:"required"'
   		Bloom	Bloom、json:^logsBloom"	gencodec:"required"'
   		Logs	[]*Log 'json:'*logs"	gencodec:"required"'	
   
   		// Implementation fields (don't reorder!)
   		TxHash	common.Hash	'json:"transactionHash" gencodec:"required"
   		contractAddress	common.Address	'json:McontractAddressM'
   		GasUsed	uint64	' json: "gasllsed" gencodec: "required"'
   }
   ```

7. **智能合约是否支持多线程，多核并行处理？**

   solidity不支持多线程。   以太坊是一个交易驱动的状态机，这个状态机必须是完全确定性的，给定的智能合约面对同一组输入，产生的输出或者说转移到的下一个状态必须是完全确定的。因为所有的全节点都要执行同一组操作，到达同一个状态，如果状态不确定三棵树的roothash就对不上。多线程的问题在于多个核访问内存的顺序不一致的话，执行结果有可能是不确定的。除了多线程之外，其他可能导致执行结果不一致的操作也都不支持，比如产生随机数，所以以太坊中用的是**伪随机数**。

### 智能合约可以获取的信息

智能合约的执行必须是确定性的，这也导致了智能合约不能像通用的编程语言那样通过系统调用来得到一些环境信息，因为每个全节点的执行环境不是完全一样的。

#### 区块信息

* `block.blockhash(uint blockNumber) returns (bytes32)`: 给定区块的哈希----仅对
最近的256个区块有效且不包括当前区块
* `block.coinbase ( address )`: 挖出当前区块的矿工地址
* `block.difficulty ( uint )`: 当前区块难度
* `block.gas1imit (umt)`: 当前区块gas限额
* `block.number ( uint )`: 当前区块号
* `block.timestamp (uint)`:自unix epoch起始当前区块以秒计的时间戳

#### 调用信息

* `msg.data(bytes)`: 完整的 calldata
* `msg.gas ( uint )`：剩余 gas
* `msg.sender ( address )`:消息发送者(当前调用)
* `msg.sig ( bytes4 )`: calldata的前4字节(也就是函数标iR符)
* `msg.value(uint )`：随消息发送的Wei的数量
* `now`目标区块时间戳( block.timestamp )
* `tx.gasprice ( uint )`：交易的 gas 价格
* `tx.origin ( address )`：交易发起者(完全的调用链)

#### 地址类型

<img src="https://gitee.com//tiansir-wg/blogimg/raw/master/imgs/20200615170155.png" alt="image-20200615170155306" style="zoom:50%;" />

> 三种发送以太坊的方式: **`transfer`，`send`，`call.value`**区别
>
> 1. transfer和send是专门用于转账的，transfer失败会导致连锁性回滚，相当于直接调用方法；send失败会返回false，不会导致连锁性回滚；call本意是发动函数调用，但也可以转账，也不会引起连锁式回滚，失败返回false
> 2. transfer和send在转账时只给了2300个单位的gas fee，收到转账的合约基本上做不了什么；而call是把当前调用剩下所有的gas都发过去。

### 简单的拍卖例子

<img src="https://gitee.com//tiansir-wg/blogimg/raw/master/imgs/20200615172042.png" alt="image-20200615172042551" style="zoom:50%;" />

<img src="https://gitee.com//tiansir-wg/blogimg/raw/master/imgs/20200615172109.png" alt="image-20200615172109428" style="zoom:50%;" />

上面bid函数是竞拍时用的，想要竞拍，就发起一个交易，调用拍卖合约中的bid函数，拍时出的价格写在msg.value中。auctionEnd是拍卖结束之后的函数。

<img src="https://gitee.com//tiansir-wg/blogimg/raw/master/imgs/20200615172615.png" alt="image-20200615172615208" style="zoom:50%;" />

上面是黑客调用拍卖合约的bid函数参与拍卖的过程。参与过程没有问题，但是退款过程会有问题。右边调用`bidder.transfer`进行转账时，bidder并没有能够接受转账的函数(标注为payable的匿名函数没有定义)，所以会引起连锁式回滚，导致转账操作是失败的，收不到钱。

转账过程实际上是全节点执行到transfer的时候把相应账户的余额进行了调整，所有智能合约执行过程中的任何对状态修改的语句改的都是本地的状态和数据结构，无论是排在黑客(bidder.transfer，bidder是黑客的合约账户时)前面还是后面，整个都回滚，都收不到钱。出现这种情况没有办法。智能合约设计的不好的话，有可能把收到的以太币永久锁起来，谁也取不出。

一种改进的方法是把`actionEnd`拆成`withdraw`和`pay2Beneficiary`两个函数。这样就不需要循环了，没有竞拍成功的人自己调用`withdraw`将钱取回。`pay2Beneficiary`用于将拍卖所得钱转账给拍卖人，存在的一个问题就是**重入攻击。**

### 重入攻击

下面是黑客调用拍卖合约的过程，`hack_bid`和前面一样，拍卖结束时调用`hack_withdraw`取回钱，问题在于fallback函数，又把钱取了一遍。

<img src="https://gitee.com//tiansir-wg/blogimg/raw/master/imgs/20200615175409.png" alt="image-20200615175409578" style="zoom:50%;" />

`hack_withdraw`调用拍卖合约的`withdraw`，拍卖合约执行到if的时候会向黑客合约转账，`msg.sender`就是黑客的合约，把当初出价的金额转给黑客(`msg.sender.call.value`)，而黑客合约的fallback再次调用拍卖合约的withdraw，又去取钱，这时的`msg.sender`就是拍卖合约(`SimpleAuctionV2(msg.sender).withdraw()`)，因为此时是拍卖合约把钱转给黑客合约，而拍卖合约又执行一遍，到if再一次转钱。拍卖合约的withdraw中清零操作只有在转账交易完成之后才会运行，而转账语句已经陷入到与黑客合约的递归调用中，执行不到清零操作，所以结果就是黑客一开始出价的时候给出一个价格，拍卖结束之后就按照这个价格不停地从拍卖合约中去取钱，第一次是自己的出价，后面取的就是别人的钱。



只有下面三种情况这种递归调用会结束:

1. 拍卖合约的余额不足以支持转账
2. gas fee不够
3. 调用栈溢出

一种简单的解决方式就是先清零再转账，转账如果不成功再把余额恢复。**先判断条件，再改变条件，最后和别的合约发生交互。**区块链上任何未知的合约都有可能是恶意的。

<img src="https://gitee.com//tiansir-wg/blogimg/raw/master/imgs/20200615180129.png" alt="image-20200615180129380" style="zoom:50%;" />

还有一种解决方案是**不要用`call.value`的方式转账，而是使用`addr.send`或`transfer`的方式转账**，因为这两种方式一个特点就是转账时发送过去的gas fee只有2300Wei，不足以让接收的合约再发起一个新的调用。

