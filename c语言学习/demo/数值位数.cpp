#include <stdio.h> 
int main() {
	int x;
	int n = 0;
	printf("������һ������\n"); 
	scanf("%d", &x);
	printf("�������������%d", x);
	while(x > 0) {
		n++;
		x /= 10;
	}
	printf(",���Ǹ�%dλ��", n);
}
