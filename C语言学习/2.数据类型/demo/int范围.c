#include <stdio.h>
int main() {
	// int ���ܱ�ʾ������
	int a = 0,max = 0;
	while(++a > 0) {
		max++;
	};
	printf("int���ܱ����������%d\n", max);
	
	// ����Ǹ����Ǽ�λ��
	int count = 0;
	while((max /= 10) != 0) {
		count ++;
	};
	printf("��λ���Ǹ�%dλ��", count);
	return 0;
} 
