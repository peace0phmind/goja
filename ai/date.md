# Date 对象

## 属性

- `constructor` - 指向Date构造函数
- `[Symbol.toPrimitive]` - 将Date对象转换为原始值的方法

## 静态方法

- `now()` - 返回当前时间的时间戳
- `parse()` - 解析日期字符串
- `UTC()` - 接受年月日时分秒毫秒参数，返回UTC时间戳

## 实例方法

### 获取日期组件的方法

- `getDate()` - 获取月份中的日期(1-31)
- `getDay()` - 获取星期几(0-6)
- `getFullYear()` - 获取年份
- `getHours()` - 获取小时(0-23)
- `getMilliseconds()` - 获取毫秒(0-999)
- `getMinutes()` - 获取分钟(0-59)
- `getMonth()` - 获取月份(0-11)
- `getSeconds()` - 获取秒数(0-59)
- `getTime()` - 获取时间戳
- `getTimezoneOffset()` - 获取时区偏移量(分钟)
- `getUTCDate()` - 获取UTC日期
- `getUTCDay()` - 获取UTC星期几
- `getUTCFullYear()` - 获取UTC年份
- `getUTCHours()` - 获取UTC小时
- `getUTCMilliseconds()` - 获取UTC毫秒
- `getUTCMinutes()` - 获取UTC分钟
- `getUTCMonth()` - 获取UTC月份
- `getUTCSeconds()` - 获取UTC秒数

### 设置日期组件的方法

- `setDate()` - 设置月份中的日期
- `setFullYear()` - 设置年份
- `setHours()` - 设置小时
- `setMilliseconds()` - 设置毫秒
- `setMinutes()` - 设置分钟
- `setMonth()` - 设置月份
- `setSeconds()` - 设置秒数
- `setTime()` - 设置时间戳
- `setUTCDate()` - 设置UTC日期
- `setUTCFullYear()` - 设置UTC年份
- `setUTCHours()` - 设置UTC小时
- `setUTCMilliseconds()` - 设置UTC毫秒
- `setUTCMinutes()` - 设置UTC分钟
- `setUTCMonth()` - 设置UTC月份
- `setUTCSeconds()` - 设置UTC秒数

### 日期格式化方法

- `toDateString()` - 返回日期部分的字符串
- `toISOString()` - 返回ISO格式的日期字符串
- `toJSON()` - 返回日期的JSON表示
- `toLocaleDateString()` - 返回本地化格式的日期字符串
- `toLocaleString()` - 返回本地化格式的日期和时间字符串
- `toLocaleTimeString()` - 返回本地化格式的时间字符串
- `toString()` - 返回日期的字符串表示
- `toTimeString()` - 返回时间部分的字符串
- `toUTCString()` - 返回UTC格式的日期字符串
- `valueOf()` - 返回日期的原始值(时间戳) 