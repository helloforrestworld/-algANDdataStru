#include <stdio.h>
// 700 不适用 
//int main() {
//	int x;
//	int n;
////	scanf("%d", &x);
//	x = 700;
//	
//	int t = 0;
//	while(x > 0) {
//		int d = x % 10;
//		t = t * 10 + d;
//		x /= 10;
//		printf("t = %d\n", t);
//	};
//	
//	x = t;
//	do {
//		n = x % 10;
//		printf("%d", n);
//		if (x > 9) {
//			printf("&");
//		}
//		x /= 10;
//		
//	} while(x > 0);
//	
//	return 0;
//}

int main() {
	int x;
//	scanf("%d", &x);
	x = 7000;
	
//	1234 / 1000  => 1 --
//	1000 / 10 => 100
//	234 / 100 => 2 --
//	100 / 10 => 10
//	34 / 10 => 3 --
//	10 / 10 => 1
//  4 / 1 => 4 -- 
//  1 / 10 => 0 结束 
	
	int t = x;
	int y = 1;
	
	// 计算几位数
	while(t >= 10) {
		t /= 10; 
		y *= 10;
	}; 
	
	do{
		int n = x / y;
		x %= y;
		y /= 10;
		printf("%d", n);
		if (y != 0) {
			printf("&", n);
		}
	} while ( y > 0);
	
	return 0;
}
