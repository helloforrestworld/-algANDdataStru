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

---

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

---

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

---

### 数的范围

- 对于一个字节(8 位)，可以表达的是:
  - 00000000 - 11111111
- 其中
  - 00000000 -> 0
  - 11111111 ~ 10000000 -> -1 ~ -128
  - 00000001 ~ 01000000 -> 1 ~ 127

```c
  int main () {
    char c = 255;
    int i = 255;
    printf("c=%d, i=%d", c, i);
    // 输出c -1 (因为八个比特11111111当作补码看待 是-1)
    // 输出i 255 (00000000 00000000 00000000 11111111)
    return 0;
  }
```

- 整数的范围(-2^n-1 ~ 2^n-1 - 1) n 比特
  - char: 1 字节 -128~127
  - short: 2 字节 -32768~32767
  - int 取决于编译器，通常意义的一个字
  - long 4 字节
  - long long 8 字节

```c
  unsigned char c = 255;
  // 这时的c就是输出255
  // unsigned 表示这个数没有负数部分，没有补码
```

- unsigned
  - 如果一个字面量常数想要表达自己是 unsigned, 在后面加上个 U 或者 u: 255u
  - 用 l 或者 L 表示 long(long)
  - \*unsigned 的初衷并非扩展能表达的数的范围，而是为了做纯二进制运算， 主要是为了移位。

---

### 整数越界

- 整数以纯二进制方式进行计算，所以
  - 11111111 + 1 => 0
  - 01111111 + 1 => -128
  - 10000000 - 1 => 01111111 => 127
    ![image](http://qiniumovie.hasakei66.com/github/jpg/char.pngchar.png)

### 整数的输入输出

- 整数的输入输出只有两种方式 int 或者 long long
  - %d: int
  - %u: unsigned
  - %ld: long long
  - %lu: unsigned long long

```c
int main() {
    char c = -1;
    int i = -1;

  printf("c=%u, i=%u\n", c, i);
  // 都是输出4294967295
  // printf打印 %u 的时候编译器会把他们转化为int
  // 计算机内部只储存了那么一个东西， 但是关键是以什么的方式去看

  // 255 4294967295 ?
  return 0;
}
```

### 八进制和十六进制

```c
int main() {
  int o = 012; // 八进制字面量
  int x = 0x12; // 十六进制

  printf("o=%d, x=%d", o, x); // 10 18
  printf("o=%o, x=%x", o, x); // 12 12
  printf("o=0%o, x=0x%x", o, x); // 012 0x12
  printf("o=0%o, x=0x%X", o, x); // 012 0x12

  return 0;
}
```

- 16 进制很适合表达二进制的数据，因为 4 位二进制正好是一个 16 进制位

```c
    > 0001 0010

    > 1     2   => 16进制的两位就是一个char
```

- 8 进制的一位正好是表达 3 位的二进制，因为早期的计算机的字长是 12 的倍数而非 8 的倍数

# 数据类型

---

## 数字类型

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

### 整数

#### 关于补码

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

#### 数的范围

- 对于一个字节(8 位)，可以表达的是:
  - 00000000 - 11111111
- 其中
  - 00000000 -> 0
  - 11111111 ~ 10000000 -> -1 ~ -128
  - 00000001 ~ 01000000 -> 1 ~ 127

```c
  int main () {
    char c = 255;
    int i = 255;
    printf("c=%d, i=%d", c, i);
    // 输出c -1 (因为八个比特11111111当作补码看待 是-1)
    // 输出i 255 (00000000 00000000 00000000 11111111)
    return 0;
  }
```

- 整数的范围(-2^n-1 ~ 2^n-1 - 1) n 比特
  - char: 1 字节 -128~127
  - short: 2 字节 -32768~32767
  - int 取决于编译器，通常意义的一个字
  - long 4 字节
  - long long 8 字节

```c
  unsigned char c = 255;
  // 这时的c就是输出255
  // unsigned 表示这个数没有负数部分，没有补码
```

- unsigned
  - 如果一个字面量常数想要表达自己是 unsigned, 在后面加上个 U 或者 u: 255u
  - 用 l 或者 L 表示 long(long)
  - \*unsigned 的初衷并非扩展能表达的数的范围，而是为了做纯二进制运算， 主要是为了移位。

#### 整数越界

- 整数以纯二进制方式进行计算，所以
  - 11111111 + 1 => 0
  - 01111111 + 1 => -128
  - 10000000 - 1 => 01111111 => 127
    ![image](http://qiniumovie.hasakei66.com/github/jpg/char.pngchar.png)

#### 整数的输入输出

- 整数的输入输出只有两种方式 int 或者 long long
  - %d: int
  - %u: unsigned
  - %ld: long long
  - %lu: unsigned long long

```c
int main() {
    char c = -1;
    int i = -1;

  printf("c=%u, i=%u\n", c, i);
  // 都是输出4294967295
  // printf打印 %u 的时候编译器会把他们转化为int
  // 计算机内部只储存了那么一个东西， 但是关键是以什么的方式去看

  // 255 4294967295 ?
  return 0;
}
```

#### 八进制和十六进制

```c
int main() {
  int o = 012; // 八进制字面量
  int x = 0x12; // 十六进制

  printf("o=%d, x=%d", o, x); // 10 18
  printf("o=%o, x=%x", o, x); // 12 12
  printf("o=0%o, x=0x%x", o, x); // 012 0x12
  printf("o=0%o, x=0x%X", o, x); // 012 0x12

  return 0;
}
```

- 16 进制很适合表达二进制的数据，因为 4 位二进制正好是一个 16 进制位

```c
    > 0001 0010

    > 1     2   => 16进制的两位就是一个char
```

- 8 进制的一位正好是表达 3 位的二进制，因为早期的计算机的字长是 12 的倍数而非 8 的倍数

#### 选择整数类型

- 为什么整数要有那么多种 ?
  - 为了准确所表达内存，做底层程序相关的需要
- 没有特殊的需要选择用 int
  - 现在的 CPU 普遍的字长就是 32 位或者 64 位， 一次内存的读写就是一个 int, 一次计算也是一次 int,选择更短
    的类型不会更快，甚至有时候会更慢
  - - 现在的编译器一般会设计内存对齐，所以更短的类型实际在内存中可能也是占据一个 int 的大小(虽然 sizeof 告诉你更小)
- unsigned 与否只是输出的不同，内部计算都是一样的

### 浮点类型

#### 浮点数

![image](EFAF84FC40CB46E0AF1D41333D7842CE)

- 关于范围

![image](A46429C32AF846F48ADE877C467A84DC)
如 float 的范围， 0 的两端，非常接近 0 的数值是无法表达的

- 关于有效数字， float 表达 7 个数字位，超过 7 个就不准确，double 则是 15 个

#### 浮点数输入输出

![image](49D3C1ADA1994F86B799153C84162665)

-
