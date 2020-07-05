#include <stdio.h>
#include <math.h>
#include <stdlib.h>

int min(int x, int y)
{
	if (x < y)
	{
		return x;
	}
	return y;
}

int main()
{
	int x, n;
	scanf("%d %d", &x, &n);

	int p[102] = {0};
	for (int i = 0; i < n; i++)
	{
		int pp;
		scanf("%d", &pp);
		p[pp] = 1;
	}

	int ans = __INT32_MAX__;
	int diff = __INT32_MAX__;
	for (int i = 101; i >= 0; i--)
	{
		if (p[i] == 1)
		{
			continue;
		}

		if (abs(x - i) <= diff)
		{
			diff = abs(x - i);
			ans = i;
		}
	}
	printf("%d\n", ans);
	return 0;
}