# 数据类型

## C 语言类型

### 分类

- 整数

  1.char、short、int、long、==long long==

- 浮点数

  1.float、double、==long double==

- 逻辑

  1.==bool==

- 指针
- 自定义类型

### 类型有何不同

- 类型名称: int、long、double
- 输入输出格式化: %d、%ld、%lf
- 所表达的数的范围不同: char < short < int < float < double
- 内存中占据的大小不一样: 1 个字节到 16 个字节
- 内存中的表达式形式: 二进制(补码)、编码

sizeof()运算符，给出某个类型或变量在内存中所占的字节数
