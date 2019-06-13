---
title: 关于音频文件的一些小知识
toc: true
# thumbnail: https://avatars3.githubusercontent.com/u/7270177?s=460&v=4
tags:
  - multi-media
categories:
  - other
---

# 音频相关知识

我们常用的音频格式，大部分都是基于音频CD（采样率44.1khz、采样精度16bit，2通道）的
192k是一个分水岭那个，192K以下的，音质损伤比较大

## 比特率
1. `CBR` Constants Bit Rate，恒定比特率
2. `VBR` Variable Bit Rate，动态比特率
VBR的方式是根据音频源文件中声音的具体频率，自动修正一些比特率，以达到在同样比特率效果中，达到更小的文件

## 采样率


### 中高低品质的采样率与比特率

|项目   |低品质   |高品质    |无损品质 |
|-------|---------|----------|---------|
|Bitrate|128KBit/s|320KBit/s |916KBit/s|
|采样率 |44100Hz  |44100Hz   |44100Hz  |

## 音频文件中的各种格式的对比
### 无损与有损
简单的来说，有损压缩就是通过删除一些已有数据中不太重要的数据来达到压缩目的；无损压缩就是通过优化排列方式来达到压缩目的

### 无损格式
APE(Monkey's audio)、FLAC(Free Lossless
Audio Codec)两种。前者拥有更小的比特率，后者则更容易传播，其区别就是，FLAC可以在传播中断后，已传播的数据就可以直接使用。

## 格式中的转换


