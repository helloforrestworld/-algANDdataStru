#include <stdio.h> 
int main() {
	int x;
	int n = 0;
	printf("������һ��������\n"); 
	scanf("%d", &x);
	printf("�������������%d", x);
	while(x > 0) {
		n++;
		x /= 10;
	}
	printf(",���Ǹ�%dλ��", n);
	return 0; 
}

//int main2() {
//	int x;
//	int n = 0;
//	printf("������һ��������\n"); 
//	scanf("%d", &x);
//	// 0����� 
//	do {
//		n++;
//		x /= 10;	
//	} while(x >0);
//	printf(",���Ǹ�%dλ��", n);
//	return 0; 
//}
