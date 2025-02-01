# Function

## Properties
- `length`: 函数参数的个数，不可写但可配置
- `name`: 函数的名称，不可写但可配置
- `prototype`: 函数的原型对象
- `constructor`: 指向Function构造函数
- `caller`: 严格模式下访问会抛出TypeError
- `arguments`: 严格模式下访问会抛出TypeError
- `[Symbol.hasInstance]`: 用于判断一个对象是否是该构造函数的实例

## Static Methods
- `Function`: 创建一个新的Function对象
- `AsyncFunction`: 创建一个新的异步函数对象
- `GeneratorFunction`: 创建一个新的生成器函数对象

## Instance Methods
- `apply`: 调用函数并以数组形式传入参数，同时指定this值
- `bind`: 创建一个新函数，将this值绑定到给定值，并在调用新函数时预置指定参数
- `call`: 调用函数并以列表形式传入参数，同时指定this值
- `toString`: 返回表示函数源码的字符串

## Generator Methods
- `next`: 返回生成器的下一个值
- `return`: 返回给定的值并结束生成器
- `throw`: 向生成器抛出一个错误

## Special Objects

### AsyncFunction.prototype
#### Properties
- `constructor`: 指向AsyncFunction构造函数
- `[Symbol.toStringTag]`: 值为"AsyncFunction"

### GeneratorFunction.prototype
#### Properties
- `constructor`: 指向GeneratorFunction构造函数
- `prototype`: 指向Generator.prototype
- `[Symbol.toStringTag]`: 值为"GeneratorFunction"

### Generator.prototype
#### Properties
- `constructor`: 指向GeneratorFunction.prototype
- `[Symbol.toStringTag]`: 值为"Generator"

#### Methods
- `next`: 返回生成器的下一个值
- `return`: 返回给定的值并结束生成器
- `throw`: 向生成器抛出一个错误 