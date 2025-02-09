# Global

## Properties

- `Infinity`: 表示正无穷大的数值属性
- `NaN`: 表示非数字(Not-a-Number)的特殊值
- `undefined`: 表示未定义的值
- `eval`: 内置的eval函数对象

## Static Methods

- `isNaN`: 判断一个值是否为NaN
- `parseInt`: 将字符串解析为指定进制的整数
- `parseFloat`: 将字符串解析为浮点数
- `isFinite`: 判断一个值是否为有限数值
- `decodeURI`: 解码一个已编码的URI
- `decodeURIComponent`: 解码一个已编码的URI组件
- `encodeURI`: 将URI进行编码
- `encodeURIComponent`: 将URI组件进行编码
- `escape`: 对字符串进行编码(已废弃)
- `unescape`: 对已编码的字符串进行解码(已废弃)
- `sleep`: 使程序暂停指定的毫秒数

## Description

全局对象是JavaScript中的顶层对象，在浏览器中通常是window对象，在Node.js中是global对象。
它提供了一些基本的属性和方法，这些属性和方法可以在任何地方被直接使用，无需通过对象引用。
所有的全局变量和函数实际上都是全局对象的属性。 