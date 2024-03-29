---
title: CICD 笔记
date: 2019-07-30 15:14:11
toc: true
img: https://user-images.githubusercontent.com/7270177/62824144-685a9480-bbcc-11e9-9b8c-7052affc5b9f.png
tags:
  - devops
  - cicd
categories:
  - devops
---

## CICD 笔录、 想法记录

### CICD是什么
什么是CICD呢， 按我理解， CICD 应该做应用生命周期管理的一整套解决方案， 而非仅仅关注与持续集成和持续发布，CICD是应用生命周期的重要组成部分， 但它不是全部。
更有一些组织会把CICD强行隔离， 部署和编译没有任何的整合， 我认为这是不科学的。  
按照个人理解，应用除了自己要关注的业务逻辑部分， 各其他点（中间件， 运维， 甚至中台）等都应该被CICD所关注，最终能达到的效果就是应用可以放心的去写应用， 大家能再这个生命周期里各司其职。  
虽然这样，我也不得不关注，当下意义上的CI和CD， 因为他们确实是前辈们抽象的最重要的两个概念了。 如果这两个概念没有， 其实这个体系也许就不存在，甚至是有另外一个体系会存在，现在应用交付的方式会发生质的变化。

### CI - 持续集成
应该是从代码写之前就开始关注， 从项目的创建，业务逻辑的编写， 到打成可部署的包，包的版本管理。  
我们知道很多公司不仅仅会有一种类型的项目结构，也不止一种的单测或者其他代码质量工具，甚至一些公司可能用多种语言进行开发。 做到通用的同时，又能做到可版本化、代码化、幂等化、简单化那就是一种艺术。


####  项目生成
项目的生成主要是指项目的初始化， 项目初始化其实是非常重要的，它影响了我们以怎么样的模板去部署它。 让业务方去专注于业务， 就必须由基础设施来承担项目生成的职责。
当然基础设施方生成项目，也能够使得项目更好的匹配我们的自动化集成与部署。  
Spring的 `initializer` 仅仅提供了一个生成非常通用的Java `Spring Boot` 项目的工具， 但这些还不够， 我们不仅仅需要把项目生成， 更要与公司的规范和基础设相结合， 保证生成的项目可以直接与基础设施无缝衔接， 这样才能赋能业务仅仅需要关注业务开发。

#### 业务逻辑编写
在项目生成的基础上，假设项目生成这一步， 我们已经选择了我们所需要的所有的基础设施， 所有的中间件与存储设施等等， 剩下的只有业务架构设计，以及业务逻辑编写了。
我们甚至可以规定生成的项目结构， 让公司某一类项目都遵循同一个项目结构标准， 这对公司的快速迭代及基础设施都是有非常大的好处的。


#### 打包 （Packaging）

打包其实只是个步骤而已，它主要关注两件事情：  

1. 执行打包命令， 生成可以部署的文件
2. 打包过程可复制化， 或者把打成的包存到历史包集里

不同的程序来讲， 打包命令是不一样的，如maven项目的常用打包命令是 `mvn clean package`, C/C++ 则更多的使用 `make` 来生成它所需要的文件。
打包过程可复制化要求其实还是有一些的， 
- 代版本跟可以用来发布的包关联
- 总是使用`release` 版本的依赖(这里我们假定这个release版本的依赖是不会改变的)


#### 测试
1. SONAR 等代码规范检测
2. 单元测试
3. 性能测试


#### CI 应该具备的特点
1. 任务清单化/模板化
2. CI节点无状态，任务随时来随时完成。
3. 基础依赖镜像化， 固化。
4. 打包过程可知化
5. 包与代码对应，幂等性

gitlab的runner 和travis-ci， 等都是比较优秀的一些平台， 而jenkins 更像一个任务平台，而非部署平台。

### CD - 持续发布

#### 版本控制与回滚
一般版本控制会跟代码的版本控制去走， 建立映射关系， 也有一些版本控制会维护自己单独的生命线。
无论哪一种无非是想具备出错迅速纠正的能力。 当然能与代码建立关系是最好的， 最好能够具备某个版本打出来的包不论打几次是等价的。

#### 部署前准备
部署前的准备主要则是把程序的依赖， 机器的初始化设置等都一一设置好。

#### 顺序部署
顺序部署主要为了平滑上线，而不至于服务中断， 也是必须的。

1. 分批部署
2. 健康监测
3. 部署预热


#### 部分部署
部署一部分进行功能点验证， 这对业务逻辑有一定要求。

#### 混合部署
在容器化时期是不需要的， 因为大家都可以根据自己的资源需求量来安排资源。
但在非容器化时期， 能混合部署会节省大量的IT成本。

#### 弹性伸缩
1. 手动弹性伸缩
2. 自动弹性伸缩
