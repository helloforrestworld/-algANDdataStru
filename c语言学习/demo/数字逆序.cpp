#include <stdio.h> 
int main() {
	// ��������
	int number;
	int diggit;
	int ret = 0;
	scanf("%d", &number);
	int isMinus = number < 0;
	
	if (isMinus == 1) {
		number = number * -1;
	}
	
	while(number > 0) {
		diggit = number % 10;
		number /= 10;
		ret = ret * 10 + diggit;
		printf("number=%d, diggit=%d,ret=%d\n", number, diggit, ret);	
	}
	
	if (isMinus == 1) {
		ret = ret * -1;
	}
	
	printf("�����Ľ����%d\n", ret);
	return 0;
}

