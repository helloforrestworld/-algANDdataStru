#include <stdio.h>
int main() {
	// շת��� �����Լ��
	// ��� b = 0 ������� a�����Լ��
	// ���� �� a ���� b �� ��a ���� b  b��������
	// �ص���һ�� 
	int a,b,t; 
	scanf("%d %d", &a, &b);
	
	while(b != 0) {
		t = a % b;
		printf("a=%d b=%d t=%d\n", a, b, t); 
		a = b;
		b = t;
	}
	printf("���Լ����%d", a);
	return 0;
} 
