#include <stdio.h> 
int main() {
	int num;
	int count = 0;
	int sum = 0;
	
	printf("请输入多个整数, -1为结束\n");
	scanf("%d", &num);
	
	while(num != -1) {
		sum += num;
		count ++;
		scanf("%d", &num);
	}
	
	printf("平均数为%f\n", 1.0 * sum / count);
	
	return 0;
}

