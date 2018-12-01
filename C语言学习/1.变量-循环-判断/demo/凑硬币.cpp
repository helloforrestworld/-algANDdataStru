#include <stdio.h>

int main()
{
	// �ո�Ӳ��
	int num = 2;
	//	scanf("%d", &num);

	int one, two, five; // һ�� ���� ���
	for (one = 1; one < num * 10; one++)
	{
		for (two = 1; two < num * 10 / 2; two++)
		{
			for (five = 1; five < num * 10 / 5; five++)
			{
				if (one + two * 2 + five * 5 == num * 10)
				{
					printf("%d,������%d��һ��%d������%d��������", num, one, two, five);
					goto out;
				}
			}
		}
	}

out:
	return 0;
}
