---
title:  react+ts+webpack+antd 实战
date: 2019-06-18 15:14:11
toc: true
tags:
    - npm
    - react
    - antd
    - project
categories:
    - lang
---


## 1. 安装react项目创建程序
```
npm install -g create-react-app yarn
```
## 2. 创建react + ts项目
```
create-react-app my-app --scripts-version=react-scripts-ts
```

react-scripts-ts是一些调整的组合，使得你可以使用标准create-react-app项目流程并且加入TypeScript。
现在你的项目应该是下面这个样子：

```
my-app/
├─ .gitignore
├─ node_modules/
├─ public/
├─ src/
│  └─ ...
├─ package.json
├─ tsconfig.json
└─ tslint.json
```
其中：

`tsconfig.json`包含我们项目中的TypeScript的配置信息

`tslint.json`是我们的代码规范工具TSLint相关的配置

`package.json`包含我们的依赖项，以及一些用于测试、预览、部署等的快捷命令。

`public`包含静态资源，比如我们要部署的HTML页面和图片。你们可以删除除了index.html以外的任何文件。

`src` 包含了我们TypeScript和CSS的代码。index.tsx是我们的文件的入口。

在`package.json` 中`scripts`中 分别有

start 开发命令 执行 npm run start
build 部署命令 执行 npm run build
test 测试命令允许Jest

## 3. 集成antd（不需要UI库可以跳过这里）
```
yarn add antd  ts-import-plugin --dev
```



## 4.配置 `config-overrides.js`
```
/*jshint esversion: 6 */
const tsImportPluginFactory = require('ts-import-plugin');
const {
    getLoader
} = require("react-app-rewired");
const rewireLess = require('react-app-rewire-less');

module.exports = function override(config) {
    const tsLoader = getLoader(
        config.module.rules,
        rule =>
        rule.loader &&
        typeof rule.loader === 'string' &&
        rule.loader.includes('ts-loader')
    );

    tsLoader.options = {
        getCustomTransformers: () => ({
            before: [tsImportPluginFactory({
                libraryDirectory: 'es',
                libraryName: 'antd',
                style: 'css',
            })]
        })
    };

    config = rewireLess.withLoaderOptions({
        modifyVars: {
            "@primary-color": "#1DA57A"
        },
    })(config, env);

    config.resolve = {
        alias: {
            '@': path.resolve("./", 'src')
        },
        extensions: ['.tsx', '.ts', '.js', '.jsx', 'css']
    };

    return config;
};
```

## 5.配置 `tsconfig.json`

```
complierOptionsu加入
    "paths": {
      "@/*": ["src/*"]
    }
```
https://blog.csdn.net/u010377383/article/details/79014405