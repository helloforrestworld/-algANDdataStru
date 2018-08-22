#include <stdio.h>
int main() {
	// 辗转相除 求最大公约数
	// 如果 b = 0 计算结束 a是最大公约数
	// 否则 用 a 除以 b ， 让a 等于 b  b等于余数
	// 回到第一步 
	int a,b,t; 
	scanf("%d %d", &a, &b);
	
	while(b != 0) {
		t = a % b;
		printf("a=%d b=%d t=%d\n", a, b, t); 
		a = b;
		b = t;
	}
	printf("最大公约数是%d", a);
	return 0;
} 
