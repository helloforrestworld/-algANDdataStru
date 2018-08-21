#include <stdio.h>

int main() {
//	// 1 + 1/2 + 1/3 + ... + 1/n
//	double total = 0;
//	int n = 100;
//	
//	int i;
//	for(i = 1; i <= n; i ++) {
//		total += 1.0/i;
//	}
//	printf("%f", total);
//	return 0;

	// 1 + 1/2 - 1/3 + ... + 1/n
	double total = 0;
	int n = 100;
	
	int i;
	double j = 1.0;
	for(i = 1; i <= n; i ++) { 
		total += j/i;
		j *= -1;
	}
	printf("%f", total);
	return 0;
} 
