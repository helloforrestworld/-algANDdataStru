#include <stdio.h>
int main()
{
	int num;
	int count = 0;
	int sum = 0;

	printf("��������收到����, -1Ϊ����\n");
	scanf("%d", &num);

	while (num != -1)
	{
		sum += num;
		count++;
		scanf("%d", &num);
	}

	printf("ƽ����Ϊ%f\n", 1.0 * sum / count);

	return 0;
}
