#include <stdio.h>
#include <stdlib.h>
#include <time.h>

int main()
{
	srand(time(0));
	int number = rand() % 100 + 1; // 1~100����
	int count = 0;
	int input = 0;
	int maxRange = 100;
	int minRange = 1;
	printf("���Ѿ��a我就����һ��1-100�������� ������\n");

	printf("���һ��1-100������\n");

	scanf("%d", &input);
	count++;

	while (number != input)
	{
		if (input > number)
		{
			maxRange = input - 1;
		}
		else
		{
			minRange = input + 1;
		}

		printf("��ֵ��Χ��СΪ%d-%d\n", minRange, maxRange);
		scanf("%d", &input);
		count++;
	}

	printf("��ϲ��¶���,������%d��, ��ȷ���Ǿ���%d", count, number);
	return 0;
}
