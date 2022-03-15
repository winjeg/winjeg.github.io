---
title:  Chrome 扩展 Easy Json Editor 生成记
date: 2022-03-15 15:14:11
toc: true
tags:
    - chrome extension
    - json editor
categories:
    - lang
---
## 背景
Chrome 有很多优秀的页面转JSON查看器 （通常是把API页面结果转成JSON） 很少有比较好的JSON编辑器，可以很好的格式化并且编辑JSON. 因此我们每次编写json，修改json都需要打开一个网址， 如json.cn, bejson.com 这些网址好用是好用的，但是你得首先打开它，且不能直达， 另外你得承受烦人的广告。
![image](https://user-images.githubusercontent.com/7270177/158315612-45d52b3f-9b56-4996-a698-dc5794958135.png)
这也是本插件开发的一个初衷，为大家提供一个相对简单实用的JSON编辑器，方便直达使用， 又可以避免无聊的广告的骚扰。

## 软件功能

清爽简洁的界面
![Json Editor](https://user-images.githubusercontent.com/7270177/158100659-ba2a8f44-3350-4415-8c88-f1679598b2fa.png)

###  方便易用的功能
- JSON 编辑与查看
- JSON 格式化与 JSON压缩
- JSON 语法高亮 & JSON 错误检查
- 插件栏点击直达，方便使用 
 
![image](https://user-images.githubusercontent.com/7270177/158316245-0ad6615f-01f1-4b7f-8ada-a0f4ecc0d1b5.png)


##  开发历程
本人并非专职前端开发，对前端开发所知甚少， 因此开发此插件对我来讲也是有一定挑战的。

### 基础库寻找
首先，开发这种工具， 我第一个想到的是，有没有三方的库可以直接使用， 毕竟开发出一个功能齐全完善的库是很耗费精力和时间的， 我们第一反应还是应该站在巨人的肩膀上继续往前走。 网上一搜，果然有一些， 而且有一些是非常成熟的，我从中挑选了[jsoneditor库](https://github.com/josdejong/jsoneditor)， 由于时间问题，我并没有把这个库下载下来仔细研读，只是从cdn上迅速下载下来了这个库最终生成的 js 和css文件。

### html 文件编写测试
第二步， 我做的是写一个非常简单的html文件， 并且把这个库引入进去, 这里我仅仅设置了左右两栏， 左栏用于编辑， 右侧栏用于查看。
一个典型的二分布局（是参考火狐一个插件的布局做的， 底层库使用的也是同一个库）。
```html
<!DOCTYPE HTML>
<html lang="en">

<head>
    <!-- when using the mode "code", it's important to specify charset utf-8 -->
    <meta charset="utf-8">
    <link href="./jsoneditor.min.css" rel="stylesheet" type="text/css">
    <script src="./jsoneditor.min.js"></script>
    <style type="text/css">
        html, body {
            font: 10.5pt arial;
            color: #4d4d4d;
            line-height: 150%;
            width: 100%;
            height: 100%;
            margin: 0;
            padding: 0;
        }
        .container {
            display: flex;
            flex-direction: row;
            width: 100%;
            height: 100%;
            min-width: 800x;
            min-height: 800px;
        }
        .jsoneditor-poweredBy {
            display: none !important;
        }
    </style>
</head>

<body>
    <div class="container">
        <div id="editor" style="width:50%"></div>
        <div id="viewer" style="width:50%"></div>
    </div>
    <script src="popup.js"></script>
</body>
</html>

```
然后我们需要把文件中引入的 json编辑器库文件下载下来，放置到同级目录，把它依赖的图标文件，放置到img 目录， 如下

```text
├── jsoneditor.min.css
├── jsoneditor.min.js
├── main.html
├── img
│   └── jsoneditor-icons.svg
```

最后，我们按照官方文档简单的写几行js代码, 如下：

```js
const leftEditor = document.getElementById('editor')
const rightViewer = document.getElementById('viewer')
const STORE_KEY = 'easy-json-editor';

function loadEditor() {
    var json = {}
    try {
        const jsonStr = localStorage.getItem(STORE_KEY) || '{}'
        json = JSON.parse(jsonStr)
    } catch (error) {
        json = {}
    }
    const viewerOptions = {
        mode: 'view',
    }
    const editorOptions = {
        mode: 'code',
        modes: ['code', 'form', 'text', 'tree', 'view', 'preview'], // allowed modes
        onModeChange: function (newMode, oldMode) {
            console.log('Mode switched from', oldMode, 'to', newMode)
        },
        onChangeText: function (jsonString) {
            localStorage.setItem(STORE_KEY, jsonString);
            jsonViewer.updateText(jsonString)
        }
    }
    const jsonEditor = new JSONEditor(leftEditor, editorOptions, json)
    const jsonViewer = new JSONEditor(rightViewer, viewerOptions, json)
}

document.addEventListener('DOMContentLoaded', loadEditor, false);
```

最后双击main.html 就可以看到预览到的效果了，这是一个完全可以正常工作的json编辑器，功能也算比较完善了。

### 第三步，把它变成chrome插件

> 由于chrome开发涉及的篇幅比较多， 本文很难进行比较详细的讲解， 有兴趣的可以查看下面的参考资料。

- 关键的 `manifest.json`, 这个是指定插件的一些具体信息的清单文件， 是非常必须的。
包括名称，描述，图标， 权限，及后台js逻辑等等， 是必须要存在的。

```json
{
    "name": "Easy JSON Editor",
    "version": "1.0",
    "description": "Easy Json Editor",
    "manifest_version": 3,
    "author": "test",
    "permissions": [
        "tabs"
    ],
    "background": {
        "service_worker": "background.js"
    },
    "action": {
        "default_icon": "img/16.png"
    },
    "icons": {
        "16": "img/16.png",
        "32": "img/32.png",
        "48": "img/48.png",
        "128": "img/128.png"
    }
}
```

- background.js 文件 
这个主要是告诉chrome点击图标打开一个新标签页，并加载main.html 的

```js
chrome.action.onClicked.addListener(function (tab) {
    chrome.tabs.create({
        url: 'main.html'
    });
});

```

### 最后一步， 把它安装到chrome中
上述操作进行完毕后， 现在的目录应该是如下的结构

```text
├── background.js
├── img
│   ├── 128.png
│   ├── 16.png
│   ├── 32.png
│   ├── 48.png
│   └── jsoneditor-icons.svg
├── jsoneditor.min.css
├── jsoneditor.min.js
├── main.html
├── manifest.json
└── popup.js
```
在chrome地址栏输入 `chrome://extensions` 打开chrome扩展列表， 按照下图所示步骤进行操作
![image](https://user-images.githubusercontent.com/7270177/158319604-7bd41827-e02e-4c11-ab35-66c2863a6c7e.png)

最后， 你可以在chrome右上角的插件中把这个插件固定在插件区域， 这样你就可以一步直接打开json 编辑器了。

文档编写比较仓促， 如您发现文档中疏漏失误， 万望留言指正， 谢谢！

[github仓库地址](https://github.com/winjeg/ejson)， 如果此插件对您有帮助，您可以给我一个小小的星星，鼓励我一下。

## 参考资料
- [Google 官方文档](https://developer.chrome.com/docs/extensions/whatsnew/)
- [【干货】Chrome插件(扩展)开发全攻略](https://www.cnblogs.com/liuxianan/p/chrome-plugin-develop.html)
