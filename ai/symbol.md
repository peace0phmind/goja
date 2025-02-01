# Symbol

## Properties
- `description`: 返回Symbol对象的描述
- `constructor`: 指向Symbol构造函数
- `[Symbol.toPrimitive]`: Symbol对象转换为原始值的方法
- `[Symbol.toStringTag]`: 返回对象的默认字符串描述

## Static Properties
- `hasInstance`: 用于判断一个对象是否为另一个对象的实例
- `isConcatSpreadable`: 用于配置数组是否可展开
- `iterator`: 返回对象的默认迭代器
- `match`: 用于字符串的匹配
- `matchAll`: 用于字符串的全局匹配
- `replace`: 用于字符串的替换
- `search`: 用于字符串的搜索
- `species`: 用于创建派生对象
- `split`: 用于字符串的分割
- `toPrimitive`: 将对象转换为原始值
- `toStringTag`: 对象的默认字符串描述
- `unscopables`: 指定不可访问的对象属性

## Static Methods
- `for`: 从全局symbol注册表中查找或创建symbol
- `keyFor`: 从全局symbol注册表中查找symbol的键

## Instance Methods
- `toString`: 返回Symbol的字符串表示
- `valueOf`: 返回Symbol的原始值 