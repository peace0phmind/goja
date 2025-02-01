# RegExp

## Properties

### source
- Type: `string`
- Access: readonly
- Configurable: true
- Description: 返回正则表达式的源文本

### global
- Type: `boolean`
- Access: readonly
- Configurable: true
- Description: 是否设置了全局匹配标志 'g'

### multiline
- Type: `boolean`
- Access: readonly
- Configurable: true
- Description: 是否设置了多行匹配标志 'm'

### dotAll
- Type: `boolean`
- Access: readonly
- Configurable: true
- Description: 是否设置了点号匹配所有字符标志 's'

### ignoreCase
- Type: `boolean`
- Access: readonly
- Configurable: true
- Description: 是否设置了忽略大小写标志 'i'

### unicode
- Type: `boolean`
- Access: readonly
- Configurable: true
- Description: 是否设置了 Unicode 标志 'u'

### sticky
- Type: `boolean`
- Access: readonly
- Configurable: true
- Description: 是否设置了粘性匹配标志 'y'

### flags
- Type: `string`
- Access: readonly
- Configurable: true
- Description: 返回正则表达式的标志字符串

### lastIndex
- Type: `number`
- Access: readwrite
- Configurable: false
- Description: 下一次匹配的起始索引

## Methods

### exec(str: string): Array|null
在字符串中执行匹配搜索，返回结果数组或 null

### test(str: string): boolean
测试字符串是否匹配正则表达式

### compile(pattern: string|RegExp, flags?: string): RegExp
重新编译正则表达式

### toString(): string
返回表示该正则表达式的字符串

## Symbols

### Symbol.match(str: string): Array|null
在字符串上执行匹配操作

### Symbol.matchAll(str: string): Iterator
返回一个迭代器，包含所有匹配结果

### Symbol.search(str: string): number
在字符串中搜索匹配项

### Symbol.split(str: string, limit?: number): Array
使用正则表达式分割字符串

### Symbol.replace(str: string, replacement: string|Function): string
替换字符串中的匹配项 