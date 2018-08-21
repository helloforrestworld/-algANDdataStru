#include <stdio.h>

int main() {
	// 凑个硬币
	int num = 2;
//	scanf("%d", &num);
	
	int one, two, five; // 一角 二角 五角
	for (one = 1; one < num * 10; one ++) {
		for (two = 1; two < num * 10 / 2; two ++) {
			for (five = 1; five < num * 10 / 5; five ++) {
				if (one + two * 2 + five * 5 == num * 10) {
					printf("%d,可以由%d个一角%d个二角%d个五角组成", num,one, two, five);
					goto out;
				}
			}
		}
	} 
	
	out:
		return 0;
} 
