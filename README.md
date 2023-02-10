# mike
> 编译器，程序员的三大浪漫之一

mike，基于 Go 语言实现的一门新语言，它有以下特性。

- 闭包
- 类 C 语法
- 变量绑定
- 内置函数
- 算数表达式
- 整形和布尔型
- 数组数据结构
- 哈希数据结构
- 字符串数据结构
- 头等函数和高阶函数

## 为什么用 Go 语言

- Go 容易阅读。
- Go 是我目前工作中使用的。
- Go 有标准库和方便的工具。
- Go 的代码风格与 C、C++ 这种更底层的语言相似。

## 语法示例

### 绑定值和名称
```
let age = 21
let name = "Cloudlang"
let result = 10 * 52
```

### 数组
```
let tmpArray = [1, 2, 3, 4];
tmpArray[0] // 1
```

### 哈希表
```
let tmpHashMap = {"name":"Cloudlang", "age": 21};
tmpHashMap["name"] // "Cloudlang"
```

### 函数
```
let add = fn(a, b) {
    return a + b;
};
```

Cloudlang 还支持隐式返回，这意味着可以忽略 `return` 关键字:
```
let add = fn(a, b) {
    a + b;
};
```
调用函数也很简单：
```
add(21, 25);
```

### 高阶函数
这类函数以其他函数为参数：
```
let tmpFn1 = fn(f, x) {
    return f(f(x));
}

let tmpFn2 = fn(x) {
    return x + 2;
}

tmpFn1(tmpFn2, 2); // 6
```
这里的 `tmpFn1` 接受两个参数。
