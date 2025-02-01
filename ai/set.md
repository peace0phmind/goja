# Set

## Properties

- `size`: 返回Set对象中的值的个数，只读属性
- `constructor`: 指向Set构造函数
- `[Symbol.iterator]`: 返回Set对象的迭代器方法
- `[Symbol.toStringTag]`: 返回默认字符串标签 "Set"

## Static Methods

- `from`: 从可迭代对象创建新的Set对象

## Instance Methods

### 修改原Set的方法

- `add`: 向Set对象添加一个新值
- `clear`: 移除Set对象内的所有元素
- `delete`: 移除Set中与指定值相等的元素

### 不修改原Set的方法

- `entries`: 返回一个包含Set对象中所有元素的键值对迭代器
- `forEach`: 对Set对象中的每个值执行一次给定的函数
- `has`: 判断Set对象中是否存在指定的值
- `keys`: 返回一个包含Set对象中所有值的迭代器
- `values`: 返回一个包含Set对象中所有值的迭代器 