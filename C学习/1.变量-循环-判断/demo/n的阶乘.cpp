#include <stdio.h>

int main()
{
	int n;
	scanf("%d", &n);
	int fact = 1;

	//	int i = 1;
	//	while (i <= n) {

	//		fact *= i;
	//		i++;
	//	}

	//	int i;
	//	for (i = 2; i <= n; i++) {
	//		fact *= i;
	//	}

	// c99 only
	//	for (int i = 2; i <= n; i++) {
	//		fact *= i;
	//	}

	int i;
	for (i = n; i > 1; i--)
	{
		fact *= i;
	}

	printf("%d�Ľ׳���%d", n, fact);
	return 0;
}
