#include <stdio.h>

int main() {
	// ���1-100������
	int x;
	for(x = 2; x <= 100; x++) {
		int isPrime = 1;
		int j;
		for (j = 2; j < x; j++) {
			if (x % j == 0) {
				isPrime = 0;
				break;
			}
		}
		if (isPrime == 1) {
			printf("%d ", x);
		}
	}
	
	// ���ǰ50������
	printf("\n���ǰ50������\n");
	int cnt = 0;
	int number = 2;
	while(cnt < 50) {
		int isPrime = 1;
		int i; 
		for (i = 2; i < number; i++) {
			if (number % i == 0) {
				isPrime = 0;
				break;
			}
		}
		if (isPrime == 1) {
			printf("%d ", number);
			cnt ++; 
		}
		number ++;
	}
	
	
	 
	return 0;
} 
