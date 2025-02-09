# Object

## Properties

- `constructor`: 指向Object构造函数
- `prototype`: Object构造函数的原型对象
- `__proto__`: 对象的原型，可通过getter/setter访问

## Static Methods

- `assign`: 将一个或多个源对象的可枚举属性复制到目标对象
- `create`: 使用指定的原型对象和属性创建一个新对象
- `defineProperty`: 在对象上定义新属性或修改现有属性
- `defineProperties`: 在对象上定义多个新属性或修改现有属性
- `entries`: 返回对象自身可枚举属性的键值对数组
- `fromEntries`: 将键值对列表转换为对象
- `getOwnPropertyDescriptor`: 返回对象指定属性的属性描述符
- `getOwnPropertyDescriptors`: 返回对象所有自有属性的属性描述符
- `getOwnPropertyNames`: 返回对象所有自有属性的名称数组
- `getOwnPropertySymbols`: 返回对象所有自有Symbol属性的数组
- `getPrototypeOf`: 返回指定对象的原型
- `is`: 判断两个值是否相同
- `isExtensible`: 判断对象是否可扩展
- `isFrozen`: 判断对象是否被冻结
- `isSealed`: 判断对象是否被密封
- `keys`: 返回对象自身可枚举属性的键名数组
- `preventExtensions`: 防止对象被扩展
- `seal`: 密封对象，阻止添加新属性并将所有现有属性标记为不可配置
- `freeze`: 冻结对象，阻止添加新属性并使现有属性不可修改
- `setPrototypeOf`: 设置对象的原型
- `values`: 返回对象自身可枚举属性的值数组
- `hasOwn`: 检查对象是否具有指定的自有属性

## Instance Methods

- `hasOwnProperty`: 判断对象是否具有指定的自有属性
- `isPrototypeOf`: 检查一个对象是否存在于另一个对象的原型链上
- `propertyIsEnumerable`: 判断指定属性是否可枚举
- `toString`: 返回对象的字符串表示
- `toLocaleString`: 返回对象的本地化字符串表示
- `valueOf`: 返回对象的原始值 