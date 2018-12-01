

# 数据类型

---

### 分类

- 整数
  1.char、short、int、long、==long long==
- 浮点数
  1.float、double、==long double==
- 逻辑
  1.==bool==
- 指针
- 自定义类型

## 数字类型
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
  - 现在的编译器一般会设计内存对齐，所以更短的类型实际在内存中可能也是占据一个 int 的大小(虽然 sizeof 告诉你更小)
- unsigned 与否只是输出的不同，内部计算都是一样的

### 浮点类型

#### 浮点数

![image](https://helloforrestworld.github.io/-algANDdataStru-/C%E8%AF%AD%E8%A8%80%E5%AD%A6%E4%B9%A0/2.%E6%95%B0%E6%8D%AE%E7%B1%BB%E5%9E%8B/screen_shot/%E4%BC%81%E4%B8%9A%E5%BE%AE%E4%BF%A1%E6%88%AA%E5%9B%BE_15358213349249.png)

- 关于范围

![image](https://helloforrestworld.github.io/-algANDdataStru-/C%E8%AF%AD%E8%A8%80%E5%AD%A6%E4%B9%A0/2.%E6%95%B0%E6%8D%AE%E7%B1%BB%E5%9E%8B/screen_shot/企业微信截图_15358215093291.png)
如 float 的范围， 0 的两端，非常接近 0 的数值是无法表达的

- 关于有效数字， float 表达 7 个数字位，超过 7 个就不准确，double 则是 15 个

#### 浮点数输入输出

![image](https://helloforrestworld.github.io/-algANDdataStru-/C%E8%AF%AD%E8%A8%80%E5%AD%A6%E4%B9%A0/2.%E6%95%B0%E6%8D%AE%E7%B1%BB%E5%9E%8B/screen_shot/企业微信截图_15358218122641.png)

#### 科学计数法
![image](https://helloforrestworld.github.io/-algANDdataStru-/C%E8%AF%AD%E8%A8%80%E5%AD%A6%E4%B9%A0/2.%E6%95%B0%E6%8D%AE%E7%B1%BB%E5%9E%8B/screen_shot/e.png)
```c
int main () {
  double ff = 1234.56789;
  printf("ff=%e", ff); // 1.2345678e+03;
  printf("ff=%E", ff); // 1.2345678E+03;

  double ff1 = 1E-10; // 可以字面量声明
  print("%E, %f\n", ff1, ff1);
  // 1.0E-10  0.000000
  print("%E, %.10f\n", ff1, ff1); // 指定输出小数点后10位
  // 1.0E-10  0.0000000000

  return 0;
}
```

#### 输出精度
```c
int main() {
  printf("%.3f\n", -0.0049); // -0.005
  printf("%.30f\n", -0.0049); // -0.00489999999999999999841793218991
  printf("%.3f\n", -0.00049);  // -0.000
  return 0;
}
```
任意两数之间都有无数个数值，
对于计算机来说，最终只能用离散的数字表达数字
![image](https://helloforrestworld.github.io/-algANDdataStru-/C%E8%AF%AD%E8%A8%80%E5%AD%A6%E4%B9%A0/2.%E6%95%B0%E6%8D%AE%E7%B1%BB%E5%9E%8B/screen_shot/float.png)
如图点ABC就是double/float能精确表达的数, 而E则是在范围外的, 所以最后取得离它比较近的A作为值，
double比float的刻度间距更小，所能表达的精确度更高


#### 超过范围的浮点数
- printf输出inf表示超过范围的浮点数
- printf输出nan表示不存在的浮点数
```c
int main() {
  printf("%f\n", 12.0 / 0.0); // inf
  printf("%f\n", -12.0 / 0.0); // -inf
  printf("%f\n", 0.0/ 0.0); // nan

  printf("%d\n", 0.0/ 0.0); // 编译不通过,无穷大不能用整数表示
  return 0;
}
```
虽然浮点数的范围不包括无穷大，但是设计上可以用浮点数表示无穷大和无穷小

#### 浮点运算的精度
浮点数的运算是没有精度的
```c
  int main() {
    float a, b, c;
    a = 1.345f; // 不带f 默认是double
    b = 1.123f;
    c = a + b;
    if (c == 2.468) {
      printf("非常精确，它们相等\n");
    } else {
      printf("他们不相等, c=%.10f,或者 c=%f\n", c, c);
    }
    // 输出 "他们不相等, c=2.4679999352,或者 2.468000"
    // so c=2.4679999352
    return 0;
  }
```
- f1 == f2 可能失败
- fabs(f1 - f2) < 1e-12 / 1e-8 差值小于一定范围来判断是否近似相等
- 另外一种思路,计算钱的时候1.23元可以用分来做单位， 表示为123分

#### 浮点数的内部表达
![image](https://helloforrestworld.github.io/-algANDdataStru-/C%E8%AF%AD%E8%A8%80%E5%AD%A6%E4%B9%A0/2.%E6%95%B0%E6%8D%AE%E7%B1%BB%E5%9E%8B/screen_shot/float2.png)
- 浮点数在计算机内部并不是真正的二进制数，而是一种编码的格式
- 如果没有特殊需要，请选择double
- 现代CPU能直接对double做硬件运算，性能不会比float差，在64位的机器上，数据的存取速度不会比float慢


## 字符类型
### 字符输入输出
#### 字符定义
- char是一种整数，也是一种特殊的类型：字符。这是因为
  - 用单引号表示的字符字面量: 'a', '1'
  - ''也是一个字符
  - printf,scanf里面用%c来输入输出字符

如果输入'1'这个字符给char c?
```c
  int main(){
    char d = 1;
    char c = '1';
    if (c == d) {
      printf("相等");
    } else {
      print("不相等");
    }
    printf("%d\n", c); // 49
    printf("%d\n", d); // 1
    return 0;
  }
```
所以表示'1'的话有两种方式
- scanf("%c", &c);   -> 1
- scanf("%d". &c);  -> 49
#### 混合输入
- 有何不同？
  - scanf("%d %c", &d, &c);
  - scanf("%d%c", &d, &c);

第一种
```c
int main() {
  int d;
  char c;
  scanf("%d %c", &d, &c);
  printf("d=%d, c='%c', c=%d\n", d, c, c);
  // 输入12 1 => d = 12, c='1', c=49
  // 输入12a => d = 12, c='a', c=97
  // 输入12  1 => d = 12, c='1', c=49
  return 0;
}
```c
第二种
int main() {
  int d;
  char c;
  scanf("%d%c", &d, &c);
  printf("d=%d, c='%c', c=%d\n", d, c, c);
  // 输入12 1 => d = 12, c='', c=32
  // 输入12a => d = 12, c='a', c=97
  // 输入12  1 => d = 12, c='', c=32
  return 0;
}
```
也就是说不带空格的话，空格被识别而且读进去了，
而带空格表示读完第一个数，再把空格都读完再去匹配第二个输入值.

#### 字符计算
加法: 一个字符加另外一个字符，等于该字符ASCII码表中的下一个字符
```c
int main() {
  char c = 'A';
  c++;
  printf("%c\n", c); // 'B';
}
```
减法: 得到两个字符的距离
```c
int main() {
  int i = 'Z' - 'A';
  printf("%d\n", i); // 25
}
```

#### 逃逸字符
用来表达无法打印出来的控制字符或特殊字符，它由一个反斜杠"\"开头，后面跟上要输出的字符。

例如:
```c
int main() {
  printf("我是一个兵, 来自\"老百姓\"");
  // 这样双引号才能正确输出
  return 0;
}
```
![image](https://helloforrestworld.github.io/-algANDdataStru-/C%E8%AF%AD%E8%A8%80%E5%AD%A6%E4%B9%A0/2.%E6%95%B0%E6%8D%AE%E7%B1%BB%E5%9E%8B/screen_shot/escape.png)


## 类型转换
### 自动类型转换
- 当运算符的两边出现不一致的类型时，会自动转换成较大的类型
 - 大的类型的意思是，它能表示的范围更大
 - char -> short -> int -> long -> long long;
 - int -> float -> double

- 对于printf来说， 任何小于int的类型都会别转换为int; float则为转为double;
- 但是scanf不会, 要输入short, 需要%hd; scanf需要准确知道数的大小;

### 强制类型转换
要把一个量强制转换成意外一个类型， (类型)值；
- 栗子:
 - (int)10.2
 - (short)32

注意这时候的安全性，小的变量不能表达大的变量
- (short)32768
```c
int main() {
  printf("%d\n", (short)32768); // -32768
  printf("%d\n", (char)32768); // 0
}
```
强制类型转换不改变值本身
```c
int main() {
  int i = 100;
  printf("%d\n", (short)i);
  printf("%d\n", i); // 100
}
```
强制类型转换的优先级高于四则运算
```c
int main() {
  double a = 1.0;
  double b = 2.0;
  // int i = (int)a / b; error
  int i = (int)(a / b); // 0
}
```

## 逻辑类型
### bool
bool并不是真正意义上的原生类型,需要引入<stdbool.h>文件才能使用true、false。
栗子:
```c
#include <stdbool.h>
#include <stdio.h>
int main() {
  bool a  = 6 > 5;
  bool b = true;
  // t = 2; // bool实际上也是整数， 所以这样也能过编译
  printf("%d\n", b); // 只能整数的方式输出, 结果是1
  return 0;
}
```
### 逻辑运算
- 逻辑运算是对逻辑量进行运算，结果只有0或者1
- 逻辑量是关系运算或逻辑运算的结果

![image](https://helloforrestworld.github.io/-algANDdataStru-/C%E8%AF%AD%E8%A8%80%E5%AD%A6%E4%B9%A0/2.%E6%95%B0%E6%8D%AE%E7%B1%BB%E5%9E%8B/screen_shot/bool.png)

```c
#include <stdio.h>
int main() {
  int a = 5;
  if (a > 4 && a < 6) {
    printf("通过");
  }
  return 0;
}
```
判断一个字符是否是大写字母
```c
#include <stdio.h>
int main() {
  char c = 'b';
  if (c >= 'A' && c <= 'Z') {
    printf("大写字母");
  }
  return 0;
}
```

逻辑运算优先级
>  ! > && > ||

各种运算符优先级
![image](https://helloforrestworld.github.io/-algANDdataStru-/C%E8%AF%AD%E8%A8%80%E5%AD%A6%E4%B9%A0/2.%E6%95%B0%E6%8D%AE%E7%B1%BB%E5%9E%8B/screen_shot/zindex.png)

短路
- 对于 && 左边false 直接false, 不再往右运行
- 对于 || 左边true 直接true, 不再往右运行

### 条件运算符(三目)
- count = count > 20 ? count ++ : count --;
- 条件 ? 满足时的值 : 不满足时的值

### 逗号运算
- for(i = 0, j = 0; i < 10 ; i++,j--)...