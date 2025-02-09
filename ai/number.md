# Number

## Properties
- `constructor`: 指向 Number 构造函数
- `prototype`: Number 原型对象

## Static Properties
- `EPSILON`: 表示 1 与大于 1 的最小浮点数之间的差值
- `MAX_SAFE_INTEGER`: JavaScript 中最大的安全整数 (2^53 - 1)
- `MIN_SAFE_INTEGER`: JavaScript 中最小的安全整数 (-(2^53 - 1))
- `MAX_VALUE`: JavaScript 中最大的数值
- `MIN_VALUE`: JavaScript 中最小的正数值
- `NaN`: 表示非数字
- `NEGATIVE_INFINITY`: 表示负无穷大
- `POSITIVE_INFINITY`: 表示正无穷大
- `parseFloat`: 解析一个字符串参数并返回一个浮点数
- `parseInt`: 解析一个字符串参数并返回一个整数

## Static Methods
- `isFinite`: 判断传入的参数是否是有限数值
- `isInteger`: 判断传入的参数是否为整数
- `isNaN`: 判断传入的参数是否为 NaN
- `isSafeInteger`: 判断传入的参数是否为安全整数

## Instance Methods
### 转换方法
- `toExponential`: 以指数表示法返回该数值字符串表示形式
- `toFixed`: 使用定点表示法来格式化一个数值
- `toLocaleString`: 返回这个数值在特定语言环境下的表示字符串
- `toPrecision`: 以指定的精度返回该数值对象的字符串表示
- `toString`: 返回指定 Number 对象的字符串表示形式
- `valueOf`: 返回一个 Number 对象的原始值

## Features
- 支持 64 位浮点数运算
- 整数和浮点数统一表示
- 支持科学计数法
- 支持十六进制（0x前缀）
- 支持八进制（0o前缀）
- 支持二进制（0b前缀）

## Limitations
- 精确整数范围限制在 -(2^53-1) 到 2^53-1 之间
- 浮点数运算可能存在精度误差
- 不支持任意精度整数（需要使用 BigInt）
- NaN 不等于任何值，包括它自己
- Infinity 表示无穷大，超出表示范围的运算结果

## Operators Supported
- 算术运算符 (+, -, *, /, %, **)
- 一元运算符 (+, -, ++, --)
- 位运算符 (&, |, ^, ~, <<, >>, >>>)
- 比较运算符 (>, <, >=, <=, ==, ===)
- 数学函数（通过 Math 对象提供） 