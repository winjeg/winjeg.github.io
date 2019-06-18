---
title:  TypeScript 简易笔记
date: 2019-06-18 15:14:11
toc: true
tags:
    - npm
    - lang
categories:
    - lang
---

## typescript 介绍
type script 是Javascript 的一种超集， 引入了JavaScript所欠缺的继承抽象与封装等

## 基本类型
### boolean
```ts
let isDone: boolean = false;
```
### number
```ts
let decLiteral: number = 6;
let hexLiteral: number = 0xf00d;
let binaryLiteral: number = 0b1010;
let octalLiteral: number = 0o744;
```
### string
模版字符串，它可以定义多行文本和内嵌表达式。 这种字符串是被反引号包围（ `），并且以${ expr }这种形式嵌入表达式
```ts
let name: string = `Gene`;
let age: number = 37;
let sentence: string = `Hello, my name is ${ name }.
I'll be ${ age + 1 } years old next month.`;

```

### array  []
```ts
let list: number[] = [1, 2, 3];
let list: Array<number> = [1, 2, 3];
```

### Tuple [number, string]
```ts
// Declare a tuple type
let x: [string, number];
// Initialize it
x = ['hello', 10]; // OK
// Initialize it incorrectly
x = [10, 'hello']; // Error
```
### enum
```ts
enum Color {Red = 1, Green, Blue}
let colorName: string = Color[2];
console.log(colorName);  // 显示'Green'因为上面代码里它的值是2
```
### any
有时候，我们会想要为那些在编程阶段还不清楚类型的变量指定一个类型。 这些值可能来自于动态的内容，比如来自用户输入或第三方代码库。 这种情况下，我们不希望类型检查器对这些值进行检查而是直接让它们通过编译阶段的检查。 那么我们可以使用 any类型来标记这些变量
```ts
let notSure: any = 4;
notSure.ifItExists(); // okay, ifItExists might exist at runtime
notSure.toFixed(); // okay, toFixed exists (but the compiler doesn't check)
let prettySure: Object = 4;
prettySure.toFixed(); // Error: Property 'toFixed' doesn't exist on type 'Object'.
```
### void 
某种程度上来说，void类型像是与any类型相反，它表示没有任何类型。声明一个void类型的变量没有什么大用，因为你只能为它赋予undefined和null：
```
function warnUser(): void {
    console.log("This is my warning message");
}
let unusable: void = undefined;
```
### null / undefined
默认情况下null和undefined是所有类型的子类型。 就是说你可以把 null和undefined赋值给number类型的变量。当你指定了--strictNullChecks标记，null和undefined只能赋值给void和它们各自
```
// Not much else we can assign to these variables!
let u: undefined = undefined;
let n: null = null;
```

### never
never 表示永远不存在的类型
```
// 返回never的函数必须存在无法达到的终点
function error(message: string): never {
    throw new Error(message);
}
// 推断的返回值类型为never
function fail() {
    return error("Something failed");
}
// 返回never的函数必须存在无法达到的终点
function infiniteLoop(): never {
    while (true) {
    }
}
```
### object
object 表示非原始类型，即除number, boolean, string, symbol, null, undefined之外的类型
```ts
declare function create(o: object | null): void;
create({ prop: 0 }); // OK
create(null); // OK
create(42); // Error
create("string"); // Error
```

### 类型推断
```typescript
let someValue: any = "this is a string";
// 以下两种方法等价
let strLength: number = (<string>someValue).length;
let strLength: number = (someValue as string).length;
```

## 变量声明
1. let  比较正常的作用域与范围
2. var  奇怪的作用域与范围
3. const 不可修改

## 接口

### 简介
接口是TS 的一个重要的概念，可以用于结构类型检查.
```ts
function printLabel(labelledObj: { label: string }) {
  console.log(labelledObj.label);
}

let myObj = { size: 10, label: "Size 10 Object" };
printLabel(myObj);
```
接口与go的差不多， 只要包含即可
```ts
interface LabelledValue {
  label: string;
}
```
### 可选属性
带有可选属性的接口与普通的接口定义差不多，只是在可选属性名字定义的后面加一个?符号。
```ts
interface SquareConfig {
  color?: string;
  width?: number;
}
```
可选属性的好处之一是可以对可能存在的属性进行预定义，好处之二是可以捕获引用了不存在的属性时的错误。 比如，我们故意将 createSquare里的color属性名拼错，就会得到一个错误提示：
```
interface SquareConfig {
  color?: string;
  width?: number;
}

function createSquare(config: SquareConfig): { color: string; area: number } {
  let newSquare = {color: "white", area: 100};
  if (config.clor) {
    // Error: Property 'clor' does not exist on type 'SquareConfig'
    newSquare.color = config.clor;
  }
  return newSquare;
}

let mySquare = createSquare({color: "black"});
```
### 只读属性
```
interface Point {
    readonly x: number;
    readonly y: number;
}

let p1: Point = { x: 10, y: 20 };
p1.x = 5; // error!

let a: number[] = [1, 2, 3, 4];
let ro: ReadonlyArray<number> = a;
ro[0] = 12; // error!
ro.push(5); // error!
ro.length = 100; // error!
a = ro; // error!

a = ro as number[];


interface SquareConfig {
    color?: string;
    width?: number;
    [propName: string]: any;
}
```

### 函数类型
```ts
interface SearchFunc {
  (source: string, subString: string): boolean;
}

let mySearch: SearchFunc;
mySearch = function(src: string, sub: string): boolean {
  let result = src.search(sub);
  return result > -1;
}
```

### 可索引的类型
TypeScript支持两种索引签名：字符串和数字。 可以同时使用两种类型的索引，但是数字索引的返回值必须是字符串索引返回值类型的子类型。 这是因为当使用 number来索引时，JavaScript会将它转换成string然后再去索引对象。 也就是说用 100（一个number）去索引等同于使用"100"（一个string）去索引，因此两者需要保持一致。
```ts
interface StringArray {
  [index: number]: string;
}

let myArray: StringArray;
myArray = ["Bob", "Fred"];

let myStr: string = myArray[0];interface StringArray {
  [index: number]: string;
}

let myArray: StringArray;
myArray = ["Bob", "Fred"];

let myStr: string = myArray[0];
```
最后，你可以将索引签名设置为只读，这样就防止了给索引赋值
```ts
interface ReadonlyStringArray {
    readonly [index: number]: string;
}
let myArray: ReadonlyStringArray = ["Alice", "Bob"];
myArray[2] = "Mallory"; // error!
```

### 类类型

```ts
interface ClockInterface {
    currentTime: Date;
    setTime(d: Date);
}

class Clock implements ClockInterface {
    currentTime: Date;
    setTime(d: Date) {
        this.currentTime = d;
    }
    constructor(h: number, m: number) { }
}
```
 constructor存在于类的静态部分，所以不在检查的范围内
```ts
interface ClockConstructor {
    new (hour: number, minute: number): ClockInterface;
}
interface ClockInterface {
    tick();
}

function createClock(ctor: ClockConstructor, hour: number, minute: number): ClockInterface {
    return new ctor(hour, minute);
}

class DigitalClock implements ClockInterface {
    constructor(h: number, m: number) { }
    tick() {
        console.log("beep beep");
    }
}
class AnalogClock implements ClockInterface {
    constructor(h: number, m: number) { }
    tick() {
        console.log("tick tock");
    }
}

let digital = createClock(DigitalClock, 12, 17);
let analog = createClock(AnalogClock, 7, 32);
```

### 继承接口
```ts
interface Shape {
    color: string;
}

interface Square extends Shape {
    sideLength: number;
}

let square = <Square>{};
square.color = "blue";
square.sideLength = 10;
```

### 接口继承类
当接口继承了一个类类型时，它会继承类的成员但不包括其实现。 就好像接口声明了所有类中存在的成员，但并没有提供具体实现一样。 接口同样会继承到类的private和protected成员。 这意味着当你创建了一个接口继承了一个拥有私有或受保护的成员的类时，这个接口类型只能被这个类或其子类所实现 
```ts
class Control {
    private state: any;
}

interface SelectableControl extends Control {
    select(): void;
}

class Button extends Control implements SelectableControl {
    select() { }
}

class TextBox extends Control {
    select() { }
}

// 错误：“Image”类型缺少“state”属性。
class Image implements SelectableControl {
    select() { }
}

class Location {

}
```

## TypeScript 的类
typescrpit 的类与java差不多，默认为Public， 有Protcected与private 成员， 也有抽象类，多态与继承
readonly关键字将属性设置为只读的。 只读属性必须在声明时或构造函数里被初始化。
我们也可以创建类的静态成员，这些属性存在于类本身上面而不是类的实例上。
```ts
class Point {
    x: number;
    y: number;
}

interface Point3d extends Point {
    z: number;
}

let point3d: Point3d = {x: 1, y: 2, z: 3};
```
类定义会创建两个东西：类的实例类型和一个构造函数。 因为类可以创建出类型，所以你能够在允许使用接口的地方使用类。

 ## 函数
 
 ### 函数类型
 ```ts
 // myAdd has the full function type
let myAdd = function(x: number, y: number): number { return x + y; };

// The parameters `x` and `y` have the type number
let myAdd: (baseValue: number, increment: number) => number =
    function(x, y) { return x + y; };
 ```
 ### 可选参数
 可选参数必须跟在必须参数后面。 如果上例我们想让first name是可选的，那么就必须调整它们的位置，把first name放在后面。
 ```ts
 function buildName(firstName: string, lastName?: string) {
    if (lastName)
        return firstName + " " + lastName;
    else
        return firstName;
}

let result1 = buildName("Bob");  // works correctly now
let result2 = buildName("Bob", "Adams", "Sr.");  // error, too many parameters
let result3 = buildName("Bob", "Adams");  // ah, just right
 ```
 
 
 ### 默认参数
 与普通可选参数不同的是，带默认值的参数不需要放在必须参数的后面。 如果带默认值的参数出现在必须参数前面，用户必须明确的传入 undefined值来获得默认值。
 ```ts
 function buildName(firstName = "Will", lastName: string) {
    return firstName + " " + lastName;
}

let result1 = buildName("Bob");                  // error, too few parameters
let result2 = buildName("Bob", "Adams", "Sr.");  // error, too many parameters
let result3 = buildName("Bob", "Adams");         // okay and returns "Bob Adams"
let result4 = buildName(undefined, "Adams");     // okay and returns "Will Adams"
 ```
 ### 可变数量参数（剩余参数）
 
 剩余参数会被当做个数不限的可选参数。 可以一个都没有，同样也可以有任意个。 编译器创建参数数组，名字是你在省略号（ ...）后面给定的名字，你可以在函数体内使用这个数组
 ```ts
 function buildName(firstName: string, ...restOfName: string[]) {
  return firstName + " " + restOfName.join(" ");
}

let buildNameFun: (fname: string, ...rest: string[]) => string = buildName;
 ```
 
 ### `this`
 *this和箭头函数*
 ```ts
 let deck = {
    suits: ["hearts", "spades", "clubs", "diamonds"],
    cards: Array(52),
    createCardPicker: function() {
        return function() {
            let pickedCard = Math.floor(Math.random() * 52);
            let pickedSuit = Math.floor(pickedCard / 13);

            return {suit: this.suits[pickedSuit], card: pickedCard % 13};
        }
    }
}

let cardPicker = deck.createCardPicker();
let pickedCard = cardPicker();

alert("card: " + pickedCard.card + " of " + pickedCard.suit);
 ```
 
 可以看到createCardPicker是个函数，并且它又返回了一个函数。 如果我们尝试运行这个程序，会发现它并没有弹出对话框而是报错了。 因为 createCardPicker返回的函数里的this被设置成了window而不是deck对象。 因为我们只是独立的调用了 cardPicker()。 顶级的非方法式调用会将 this视为window。 （注意：在严格模式下， this为undefined而不是window）。
 
 为了解决这个问题，我们可以在函数被返回时就绑好正确的this。 这样的话，无论之后怎么使用它，都会引用绑定的‘deck’对象。 我们需要改变函数表达式来使用ECMAScript 6箭头语法。 箭头函数能保存函数创建时的 this值，而不是调用时的值：
 
 ```ts
 let deck = {
    suits: ["hearts", "spades", "clubs", "diamonds"],
    cards: Array(52),
    createCardPicker: function() {
        // NOTE: the line below is now an arrow function, allowing us to capture 'this' right here
        return () => {
            let pickedCard = Math.floor(Math.random() * 52);
            let pickedSuit = Math.floor(pickedCard / 13);

            return {suit: this.suits[pickedSuit], card: pickedCard % 13};
        }
    }
}

let cardPicker = deck.createCardPicker();
let pickedCard = cardPicker();

alert("card: " + pickedCard.card + " of " + pickedCard.suit);
 ```
 你可以也看到过在回调函数里的this报错，当你将一个函数传递到某个库函数里稍后会被调用时。 因为当回调被调用的时候，它们会被当成一个普通函数调用， this将为undefined。 稍做改动，你就可以通过 this参数来避免错误。首先，库函数的作者要指定 this的类型：
 
 ```ts
 interface UIElement {
    addClickListener(onclick: (this: void, e: Event) => void): void;
}

class Handler {
    info: string;
    onClickGood(this: void, e: Event) {
        // can't use this here because it's of type void!
        console.log('clicked!');
    }
}
let h = new Handler();
uiElement.addClickListener(h.onClickGood);
 ```
 因为onClickGood指定了this类型为void，因此传递addClickListener是合法的。 当然了，这也意味着不能使用 this.info. 如果你两者都想要，你不得不使用箭头函数了：
 ```ts
 class Handler {
    info: string;
    onClickGood = (e: Event) => { this.info = e.message }
}
 ```
 这是可行的因为箭头函数不会捕获this，所以你总是可以把它们传给期望this: void的函数。 缺点是每个 Handler对象都会创建一个箭头函数。 另一方面，方法只会被创建一次，添加到 Handler的原型链上。 它们在不同 Handler对象间是共享的。
 
 ## 泛型
 ```ts
 function identity<T>(arg: T): T {
    return arg;
}
let output = identity<string>("myString");  // 显式声明
let output = identity("myString");  // 自动推断
 ```
### 使用泛型变量
```ts
function loggingIdentity<T>(arg: T): T {
    console.log(arg.length);  // Error: T doesn't have .length
    return arg;
}

function loggingIdentity<T>(arg: T[]): T[] {
    console.log(arg.length);  // Array has a .length, so no more error
    return arg;
}

function loggingIdentity<T>(arg: Array<T>): Array<T> {
    console.log(arg.length);  // Array has a .length, so no more error
    return arg;
}

```

### 泛型类型



### 泛型接口
```ts
interface GenericIdentityFn {
    <T>(arg: T): T;
}

function identity<T>(arg: T): T {
    return arg;
}

let myIdentity: GenericIdentityFn = identity;
```
### 泛型类
```ts
class GenericNumber<T> {
    zeroValue: T;
    add: (x: T, y: T) => T;
}

let myGenericNumber = new GenericNumber<number>();
myGenericNumber.zeroValue = 0;
myGenericNumber.add = function(x, y) { return x + y; };
```

```ts
class BeeKeeper {
    hasMask: boolean;
}

class ZooKeeper {
    nametag: string;
}

class Animal {
    numLegs: number;
}

class Bee extends Animal {
    keeper: BeeKeeper;
}

class Lion extends Animal {
    keeper: ZooKeeper;
}

function createInstance<A extends Animal>(c: new () => A): A {
    return new c();
}

createInstance(Lion).keeper.nametag;  // typechecks!
createInstance(Bee).keeper.hasMask;   // typechecks!

```

### 交叉类型（Intersection Types）
```ts
function extend<T, U>(first: T, second: U): T & U {
    let result = <T & U>{};
    for (let id in first) {
        (<any>result)[id] = (<any>first)[id];
    }
    for (let id in second) {
        if (!result.hasOwnProperty(id)) {
        // 两个类型如果这个没有就去另外一个
            (<any>result)[id] = (<any>second)[id];
        }
    }
    return result;
}

class Person {
    constructor(public name: string) { }
}
interface Loggable {
    log(): void;
}
class ConsoleLogger implements Loggable {
    log() {
        // ...
    }
}
var jim = extend(new Person("Jim"), new ConsoleLogger());
var n = jim.name;
jim.log();
```

### 联合类型（Union Types）
联合类型与交叉类型很有关联，但是使用上却完全不同。 偶尔你会遇到这种情况，一个代码库希望传入 number或 string类型的参数。 例如下面的函数：
```ts
/**
 * Takes a string and adds "padding" to the left.
 * If 'padding' is a string, then 'padding' is appended to the left side.
 * If 'padding' is a number, then that number of spaces is added to the left side.
 */
function padLeft(value: string, padding: any) {
    if (typeof padding === "number") {
        return Array(padding + 1).join(" ") + value;
    }
    if (typeof padding === "string") {
        return padding + value;
    }
    throw new Error(`Expected string or number, got '${padding}'.`);
}

padLeft("Hello world", 4); // returns "    Hello world"

```
### 导入导出
export default  
import xxx from 
```ts
class ZipCodeValidator implements StringValidator {
    isAcceptable(s: string) {
        return s.length === 5 && numberRegexp.test(s);
    }
}
export { ZipCodeValidator };
export { ZipCodeValidator as mainValidator };
```


## 命名空间




## 三斜线
三斜线指令是包含单个XML标签的单行注释。 注释的内容会做为编译器指令使用。

三斜线指令仅可放在包含它的文件的最顶端。 一个三斜线指令的前面只能出现单行或多行注释，这包括其它的三斜线指令。 如果它们出现在一个语句或声明之后，那么它们会被当做普通的单行注释，并且不具有特殊的涵义。
```ts
/// <reference no-default-lib="true"/>
```
这个指令把一个文件标记成默认库。 你会在 lib.d.ts文件和它不同的变体的顶端看到这个注释。

## 模块导入
相对 vs. 非相对模块导入

相对导入是以/，./或../开头的。 下面是一些例子：
```ts
import Entry from "./components/Entry";
import { DefaultHeaders } from "../constants/http";
import "/mod";
```
相对导入的模块是相对于导入它的文件进行解析的。 因此 /root/src/folder/A.ts文件里的import { b } from "./moduleB"会使用下面的查找流程：

1. /root/src/folder/moduleB.ts
2. /root/src/folder/moduleB.d.ts
所有其它形式的导入被当作非相对的。 下面是一些例子：
```ts
import * as $ from "jQuery";
import { Component } from "@angular/core";
```

有一个对moduleB的非相对导入import { b } from "moduleB"，它是在/root/src/folder/A.ts文件里，会以如下的方式来定位"moduleB"：
```
/root/src/folder/moduleB.ts
/root/src/folder/moduleB.d.ts
/root/src/moduleB.ts
/root/src/moduleB.d.ts
/root/moduleB.ts
/root/moduleB.d.ts
/moduleB.ts
/moduleB.d.ts
```
### TypeScript如何解析模块
TypeScript是模仿Node.js运行时的解析策略来在编译阶段定位模块定义文件。 因此，TypeScript在Node解析逻辑基础上增加了TypeScript源文件的扩展名（ .ts，.tsx和.d.ts）。 同时，TypeScript在 package.json里使用字段"types"来表示类似"main"的意义 - 编译器会使用它来找到要使用的"main"定义文件。