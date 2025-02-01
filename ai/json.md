# JSON

JavaScript 的 JSON 对象用于解析 JSON 字符串和将值转换为 JSON 字符串

## 静态方法

### parse

将 JSON 字符串解析成 JavaScript 值

**参数:**
- `text`: 要解析的 JSON 字符串
- `reviver`: (可选) 转换函数，用于在解析过程中修改解析值

**特点:**
- 支持解析对象、数组、字符串、数字、布尔值和 null
- 不支持解析 undefined、函数、Symbol 和 BigInt
- 解析失败时会抛出 SyntaxError

### stringify

将 JavaScript 值转换为 JSON 字符串

**参数:**
- `value`: 要转换的值
- `replacer`: (可选) 转换函数或要包含的属性数组
- `space`: (可选) 缩进字符串或空格数量

**特点:**
- 支持转换对象、数组、字符串、数字、布尔值和 null
- 会忽略 undefined、函数和 Symbol
- 处理循环引用时会抛出错误
- 可以通过 replacer 自定义序列化过程
- 可以通过 space 控制输出格式的美化 