# Map

## Properties

- `size`: Map中的元素数量，只读属性
- `constructor`: 指向Map构造函数
- `[Symbol.toStringTag]`: 返回"Map"的字符串标签
- `[Symbol.iterator]`: 返回entries()方法的引用，用于迭代Map的键值对

## Static Methods

Map构造函数可以接受一个可迭代对象作为参数来创建新的Map实例

## Instance Methods

### 修改Map的方法

- `clear()`: 移除Map对象中的所有键值对
- `delete(key)`: 移除指定键的键值对，如果键存在则返回true，否则返回false
- `set(key, value)`: 在Map对象中设置键值对，如果键已存在则更新值，返回Map对象本身

### 访问和检查方法

- `get(key)`: 返回指定键的值，如果键不存在则返回undefined
- `has(key)`: 返回一个布尔值，表示Map是否包含指定的键
- `entries()`: 返回一个新的迭代器对象，包含[key, value]对
- `forEach(callback)`: 按插入顺序对Map中的每个键值对执行指定的函数
- `keys()`: 返回一个新的迭代器对象，包含Map中的所有键
- `values()`: 返回一个新的迭代器对象，包含Map中的所有值

## Features

- 键可以是任意值(包括函数、对象或任意原始值)
- 键的比较使用SameAs算法(类似于===，但将-0和+0视为相同)
- 维护插入顺序，迭代时会按照插入顺序返回元素
- 内置迭代器支持for...of循环
- 支持键值对的添加和删除操作 