---
title: "ETH12 以太坊概述"
date: 2023-05-28T13:49:12+08:00
draft: true
image: ""
categories: 
tag:
---

> 整理自 [北京大学肖臻老师《区块链技术与应用》公开课](https://www.bilibili.com/video/BV1Vt411X7JF?from=search&seid=14488407572640514229)

### 与比特币系统的不同

比特币称为区块链1.0，以太坊称为区块链2.0。在系统设计上针对比特币的一些问题进行了改进：

1. 出块时间降到十几秒
2. 设计基于ghost的共识协议
3. 使用的mining puzzle对内存要求高(memory hard mining puzzle)，在一定程度上限制了asic芯片的使用(称为ASIC resistance)
4. 用权益证明(proof of stake)替代工作量证明(proof of work)，类似于股份投票决定下一个区块如何产生
5. 增加对智能合约(smart contract)的支持

### 智能合约

比特币(Bitcoin)：实现了一种去中心化货币(decentralized currency，即比特币BTC) ，比特币的出现用技术手段取代了政府职能。除了货币还有什么可以去中心化?

以太坊(Ethereum)：增加了去中心化合约(decentralized contract)的支持， 智能合约的出现用技术手段取代了司法手段，合同内容用代码实现，将代码发布到区块链上，通过区块链的不可篡改性保证代码正确运行。



**去中心化货币的好处**:  跨国转账时如果用法币会很麻烦，而且会有一定的手续费

**智能合约的好处**: 如果合同签署方来自各地，没有统一的司法管辖权，用司法手段维护合同有效性比较困难，如果通过实现写好的程序代码，来保证每个人都只能按照规则来执行，是比较好的解决办法。就算合同各方在同一个司法管辖权下，想通过司法手段维护合同执行，也是费时费力的过程。所以最好用技术手段保证合同参与方从开始就不可能违约，代码一旦发布到区块链上，通过不可篡改性，保证大家只能按照代码中的规则执行。

