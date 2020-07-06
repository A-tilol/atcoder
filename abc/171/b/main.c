#include <stdio.h>
#include <stdlib.h>

int asc(const void *a, const void *b)
{
	return *(int *)a - *(int *)b;
}

int main()
{
	int n, k;
	scanf("%d %d", &n, &k);

	int *p = (int *)malloc(sizeof(int) * n);
	for (int i = 0; i < n; i++)
	{
		scanf("%d", &p[i]);
	}

	qsort(p, n, sizeof(int), asc);

	int ans = 0;
	for (int i = 0; i < k; i++)
	{
		ans += p[i];
	}
	printf("%d\n", ans);

	free(p);

	return 0;
}