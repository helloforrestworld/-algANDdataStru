#include <stdio.h> 
int main() {
	int x;
	int n = 0;
	printf("请输入一个整数\n"); 
	scanf("%d", &x);
	printf("你输入的数字是%d", x);
	while(x > 0) {
		n++;
		x /= 10;
	}
	printf(",它是个%d位数", n);
}
