#include <stdio.h>

int main()
{
	char c;
	scanf("%c", &c);

	if (c < 'a')
	{
		printf("A\n");
	}
	else
	{
		printf("a\n");
	}

	return 0;
}