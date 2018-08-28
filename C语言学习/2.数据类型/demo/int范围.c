#include <stdio.h>
int main() {
	// int 所能表示最大的数
	int a = 0,max = 0;
	while(++a > 0) {
		max++;
	};
	printf("int所能表达最大的数是%d\n", max);
	
	// 最大那个数是几位数
	int count = 0;
	while((max /= 10) != 0) {
		count ++;
	};
	printf("那位数是个%d位数", count);
	return 0;
} 
