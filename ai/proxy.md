# Proxy

JavaScript Proxy 对象用于定义基本操作的自定义行为（如属性查找、赋值、枚举、函数调用等）

## Constructor

- `new Proxy(target, handler)`: 创建一个新的 Proxy 对象，target 是要代理的对象，handler 是定义代理行为的对象

## Static Methods

- `revocable`: 创建一个可撤销的代理对象，返回包含proxy和revoke属性的对象

## Handler Methods

### 基本对象操作

- `getPrototypeOf`: 拦截 Object.getPrototypeOf, Reflect.getPrototypeOf, __proto__, instanceof 等操作
- `setPrototypeOf`: 拦截 Object.setPrototypeOf, Reflect.setPrototypeOf 操作
- `isExtensible`: 拦截 Object.isExtensible, Reflect.isExtensible 操作
- `preventExtensions`: 拦截 Object.preventExtensions, Reflect.preventExtensions 操作
- `getOwnPropertyDescriptor`: 拦截 Object.getOwnPropertyDescriptor, Reflect.getOwnPropertyDescriptor 操作
- `defineProperty`: 拦截 Object.defineProperty, Reflect.defineProperty 操作
- `has`: 拦截 in 运算符, with 操作符, Reflect.has 操作
- `get`: 拦截属性读取操作
- `set`: 拦截属性设置操作
- `deleteProperty`: 拦截 delete 操作符
- `ownKeys`: 拦截 Object.getOwnPropertyNames, Object.getOwnPropertySymbols, Object.keys, Reflect.ownKeys 操作

### 函数操作

- `apply`: 拦截函数调用操作，包括 Function.prototype.apply, Function.prototype.call, Reflect.apply
- `construct`: 拦截 new 操作符，Reflect.construct 操作

## Supported Property Types

- `string`: 字符串类型的属性名
- `integer`: 整数类型的属性名（包括负数）
- `symbol`: Symbol 类型的属性名

## Notes

- 所有 handler 方法都是可选的
- 如果未定义某个 handler 方法，则使用默认行为
- handler 方法可以拦截对应的 Reflect 方法
- 某些操作可能会触发多个 handler 方法
- handler 方法必须遵循特定的不变量（invariants） 