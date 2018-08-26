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

```c
    int a = 6;
    printf("%d", sizeof(a++));
    printf("%d", a);
    // a还是等于6
    // sizeof是静态方法, 编译时候确认sizeof(i) i是什么内容 然后输出对应的字节数
```

- char 1 字节(8 比特)
- short 2 字节
- int 取决于编译器(CPU), 通常意义上的 1 个字
- long 取决于编译器(CPU), 通常意义上的 1 个字
- long long 8 个字节

<html style="overflow:hidden">
<!--在这里插入内容-->
<span style="display:inline-block;padding:10px; width:100px;height:100px;border:1px solid black;">
    CPU
    <span style="display:inline-block;border:1px solid red;">寄存器</span>
</span>
<----总线--->
<span style="display:inline-block;padding:10px; width:100px;height:100px;border:1px solid black;">
    RAM
</span>

</html>

- 32 位(64)的系统 每个寄存器能存放 32(64)比特大小
- 总线每次传递也是 32 比特大小
- int 所有表达的就是一个寄存器的大小

### 关于补码

计算机内部一切都是二进制

- 18 ---> 00010010
- 0 ---> 00000000
- -18 ---> ?

考虑-1， 我们希望-1 + 1 -> 0。 如何能做到？

- 0 -> 00000000
- 1 -> 00000001

? + 00000001 -> 00000000

=> 11111111 + 00000001 -> (1)00000000

所以 11111111 被当作纯二进制看待时， 是 255，被当作补码看待时是-1

补码的意义就是拿补码加原码可以加出一个溢出的"零"

### 数的范围

- 对于一个字节(8 位)，可以表达的是:
  - 00000000 - 11111111
- 其中
  - 00000000 -> 0
  - 11111111 ~ 100000000 -> -1 ~ -128
  - 00000001 ~ 010000000 -> 1 ~ 127
