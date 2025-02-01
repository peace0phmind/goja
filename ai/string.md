# String

JavaScript String 对象的内置实现

## 静态方法

### String.fromCharCode(num)

从一系列 Unicode 值创建字符串

- **参数**：
  - `num`: number - 一个或多个 Unicode 值
- **返回值**：string
- **示例**：
  ```js
  String.fromCharCode(65, 66, 67) // returns "ABC"
  ```

### String.fromCodePoint(num)

从一系列码点创建字符串

- **参数**：
  - `num`: number - 一个或多个码点值
- **返回值**：string
- **示例**：
  ```js
  String.fromCodePoint(0x2F804) // returns a string containing the character 你
  ```

### String.raw(template)

模板字符串的标签函数

- **参数**：
  - `template`: object - 包含原始字符串信息的对象
- **返回值**：string
- **示例**：
  ```js
  String.raw`Hi\n${2+3}!` // returns "Hi\\n5!"
  ```

## 属性

### length

字符串的长度

- **类型**：number
- **示例**：
  ```js
  "abc".length // returns 3
  ```

## 原型方法

### at(index)

返回指定索引处的字符，支持负数索引

- **参数**：
  - `index`: number - 字符的索引位置
- **返回值**：string
- **示例**：
  ```js
  "abc".at(-1) // returns "c"
  ```

### charAt(index)

返回指定位置的字符

- **参数**：
  - `index`: number - 字符的索引位置
- **返回值**：string
- **示例**：
  ```js
  "abc".charAt(1) // returns "b"
  ```

### charCodeAt(index)

返回指定位置字符的 Unicode 值

- **参数**：
  - `index`: number - 字符的索引位置
- **返回值**：number
- **示例**：
  ```js
  "ABC".charCodeAt(0) // returns 65
  ```

### codePointAt(index)

返回指定位置字符的码点值

- **参数**：
  - `index`: number - 字符的索引位置
- **返回值**：number
- **示例**：
  ```js
  "😀".codePointAt(0) // returns 128512
  ```

### concat(str)

连接两个或多个字符串

- **参数**：
  - `str`: string - 要连接的字符串
- **返回值**：string
- **示例**：
  ```js
  "Hello".concat(" ", "World") // returns "Hello World"
  ```

### endsWith(searchString[, length])

判断字符串是否以指定字符串结尾

- **参数**：
  - `searchString`: string - 要搜索的字符串
  - `length`: number - (可选) 搜索的长度
- **返回值**：boolean
- **示例**：
  ```js
  "Hello World".endsWith("World") // returns true
  ```

### includes(searchString[, position])

判断字符串是否包含指定的子字符串

- **参数**：
  - `searchString`: string - 要搜索的字符串
  - `position`: number - (可选) 开始搜索的位置
- **返回值**：boolean
- **示例**：
  ```js
  "Hello World".includes("World") // returns true
  ```

### indexOf(searchValue[, fromIndex])

返回指定子字符串第一次出现的位置

- **参数**：
  - `searchValue`: string - 要搜索的字符串
  - `fromIndex`: number - (可选) 开始搜索的位置
- **返回值**：number
- **示例**：
  ```js
  "Hello World".indexOf("World") // returns 6
  ```

### lastIndexOf(searchValue[, fromIndex])

返回指定子字符串最后一次出现的位置

- **参数**：
  - `searchValue`: string - 要搜索的字符串
  - `fromIndex`: number - (可选) 开始搜索的位置
- **返回值**：number
- **示例**：
  ```js
  "Hello World World".lastIndexOf("World") // returns 12
  ```

### localeCompare(compareString)

比较两个字符串在当前语言环境下的顺序

- **参数**：
  - `compareString`: string - 要比较的字符串
- **返回值**：number
- **示例**：
  ```js
  "a".localeCompare("b") // returns -1
  ```

### match(regexp)

使用正则表达式匹配字符串

- **参数**：
  - `regexp`: RegExp - 正则表达式对象
- **返回值**：Array
- **示例**：
  ```js
  "test1test2".match(/test\d/g) // returns ["test1", "test2"]
  ```

### normalize([form])

返回字符串的 Unicode 标准化形式

- **参数**：
  - `form`: string - (可选) 标准化形式("NFC", "NFD", "NFKC", "NFKD")
- **返回值**：string
- **示例**：
  ```js
  "\u1E9B\u0323".normalize() // returns "\u1E9B\u0323"
  ```

### padEnd(targetLength[, padString])

在字符串末尾填充字符串到指定长度

- **参数**：
  - `targetLength`: number - 目标长度
  - `padString`: string - (可选) 填充字符串
- **返回值**：string
- **示例**：
  ```js
  "abc".padEnd(6, "123") // returns "abc123"
  ```

### padStart(targetLength[, padString])

在字符串开头填充字符串到指定长度

- **参数**：
  - `targetLength`: number - 目标长度
  - `padString`: string - (可选) 填充字符串
- **返回值**：string
- **示例**：
  ```js
  "abc".padStart(6, "123") // returns "123abc"
  ```

### repeat(count)

重复字符串指定次数

- **参数**：
  - `count`: number - 重复次数
- **返回值**：string
- **示例**：
  ```js
  "abc".repeat(2) // returns "abcabc"
  ```

### replace(searchValue, replaceValue)

替换字符串中的子字符串

- **参数**：
  - `searchValue`: string|RegExp - 要替换的字符串或正则表达式
  - `replaceValue`: string|Function - 替换值或替换函数
- **返回值**：string
- **示例**：
  ```js
  "Hello World".replace("World", "JavaScript") // returns "Hello JavaScript"
  ```

### replaceAll(searchValue, replaceValue)

替换字符串中所有匹配的子字符串

- **参数**：
  - `searchValue`: string|RegExp - 要替换的字符串或正则表达式
  - `replaceValue`: string|Function - 替换值或替换函数
- **返回值**：string
- **示例**：
  ```js
  "test test".replaceAll("test", "new") // returns "new new"
  ```

### search(regexp)

使用正则表达式搜索字符串

- **参数**：
  - `regexp`: RegExp - 正则表达式对象
- **返回值**：number
- **示例**：
  ```js
  "test".search(/est/) // returns 1
  ```

### slice(beginIndex[, endIndex])

提取字符串的一部分

- **参数**：
  - `beginIndex`: number - 开始位置
  - `endIndex`: number - (可选) 结束位置
- **返回值**：string
- **示例**：
  ```js
  "Hello World".slice(6) // returns "World"
  ```

### split(separator[, limit])

将字符串分割成数组

- **参数**：
  - `separator`: string|RegExp - 分隔符
  - `limit`: number - (可选) 限制结果数量
- **返回值**：Array
- **示例**：
  ```js
  "a,b,c".split(",") // returns ["a", "b", "c"]
  ```

### startsWith(searchString[, position])

判断字符串是否以指定字符串开头

- **参数**：
  - `searchString`: string - 要搜索的字符串
  - `position`: number - (可选) 开始搜索的位置
- **返回值**：boolean
- **示例**：
  ```js
  "Hello World".startsWith("Hello") // returns true
  ```

### substring(indexStart[, indexEnd])

返回字符串的子串

- **参数**：
  - `indexStart`: number - 开始位置
  - `indexEnd`: number - (可选) 结束位置
- **返回值**：string
- **示例**：
  ```js
  "Hello World".substring(6) // returns "World"
  ```

### toLowerCase()

将字符串转换为小写

- **返回值**：string
- **示例**：
  ```js
  "Hello World".toLowerCase() // returns "hello world"
  ```

### toUpperCase()

将字符串转换为大写

- **返回值**：string
- **示例**：
  ```js
  "Hello World".toUpperCase() // returns "HELLO WORLD"
  ```

### trim()

删除字符串两端的空白字符

- **返回值**：string
- **示例**：
  ```js
  "  Hello World  ".trim() // returns "Hello World"
  ```

### trimEnd()

删除字符串末尾的空白字符

- **返回值**：string
- **示例**：
  ```js
  "Hello World  ".trimEnd() // returns "Hello World"
  ```

### trimStart()

删除字符串开头的空白字符

- **返回值**：string
- **示例**：
  ```js
  "  Hello World".trimStart() // returns "Hello World"
  ```

### toString()

返回字符串对象的字符串表示

- **返回值**：string
- **示例**：
  ```js
  new String("Hello").toString() // returns "Hello"
  ```

### valueOf()

返回字符串对象的原始值

- **返回值**：string
- **示例**：
  ```js
  new String("Hello").valueOf() // returns "Hello"
  ```

### substr(start[, length]) (已废弃)

从指定位置开始返回指定长度的子字符串

- **参数**：
  - `start`: number - 开始位置
  - `length`: number - (可选) 提取的字符数
- **返回值**：string
- **示例**：
  ```js
  "Hello World".substr(6, 5) // returns "World"
  ```

## 迭代器

String 对象实现了迭代器接口

**示例**：
```js
for (let char of "Hello") {
  console.log(char); // prints H, e, l, l, o
}
``` 