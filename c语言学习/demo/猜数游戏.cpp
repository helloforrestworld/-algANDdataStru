#include <stdio.h>
#include <stdlib.h>
#include <time.h>

int main() {
	srand(time(0));
	int number = rand()%100 + 1; // 1~100整数
	int count = 0;
	int input = 0;
	int maxRange = 100; 
	int minRange = 1;
	printf("我已经想好了一个1-100的整数， 你来猜\n"); 
	
	printf("请猜一个1-100的整数\n");
	
	scanf("%d", &input);
	count ++;
	
	while (number != input) {
		if (input > number) {
			maxRange = input - 1;
		} else {
			minRange = input + 1;
		}
		
		printf("数值范围缩小为%d-%d\n", minRange, maxRange);
		scanf("%d", &input);
		count ++;
	}
	
	printf("恭喜你猜对了,共用了%d次, 正确的是就是%d", count, number);
	return 0;
}

