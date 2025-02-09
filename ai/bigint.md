# BigInt

## Properties
- `constructor`: 指向 BigInt 构造函数
- `prototype`: BigInt 原型对象

## 静态方法
- `asIntN`: 将 BigInt 值转换为 -2^(n-1) 到 2^(n-1)-1 范围内的有符号整数
- `asUintN`: 将 BigInt 值转换为 0 到 2^n-1 范围内的无符号整数

## 实例方法

### 转换方法
- `toString`: 将 BigInt 转换为字符串，可选参数指定进制(2-36)
- `valueOf`: 返回 BigInt 的原始值

### 类型转换行为
- `toBoolean`: 转换为布尔值，只有 0n 转换为 false
- `toNumber`: 不支持，会抛出 TypeError

### 比较方法
- `equals`: 与其他值进行相等性比较
- `strictEquals`: 与其他值进行严格相等性比较
- `sameAs`: 判断两个 BigInt 值是否完全相同

## 限制
- 不能与 Number 进行混合运算
- 不能隐式转换为 Number 类型
- 不支持小数点运算，总是整数
- 不能用于 Math 对象的方法

## 支持的运算符
- 加法 (`+`)
- 减法 (`-`)
- 乘法 (`*`)
- 除法 (`/`)
- 取模 (`%`)
- 指数 (`**`)
- 位运算 (`&`, `|`, `^`, `~`)
- 移位运算 (`<<`, `>>`, `>>>`)
- 比较运算 (`>`, `<`, `>=`, `<=`)
- 相等运算 (`==`, `===`, `!=`, `!==`)

## 特性
- 支持任意精度的整数运算
- 不能与Number类型直接进行混合运算
- 不支持小数点和指数表示法
- 字面量使用n后缀表示，如123n
- 不支持转换为Number类型
- 支持常见的算术运算符和比较运算符 