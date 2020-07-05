#include <stdio.h>

int main()
{
	for (int i = 0; i < 5; i++)
	{
		int x;
		scanf("%d", &x);
		if (x == 0)
		{
			printf("%d\n", i + 1);
			return 0;
		}
	}

	return 0;
}