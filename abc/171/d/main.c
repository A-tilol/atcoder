#include <stdio.h>
#include <stdlib.h>

typedef long long int ll;

int read_int()
{
	int i;
	scanf("%d", &i);
	return i;
}

int read_lint()
{
	ll i;
	scanf("%lld", &i);
	return i;
}

int conv[100010];

int main()
{
	int n = read_int();

	ll sum = 0;
	int *a = (int *)malloc(sizeof(int) * n);
	for (int i = 0; i < n; i++)
	{
		ll a = read_lint();
		conv[a]++;
		sum += a;
	}

	int q = read_int();
	for (int i = 0; i < q; i++)
	{
		ll b = read_lint();
		ll c = read_lint();
		if (conv[b] == 0)
		{
			printf("%lld\n", sum);
			continue;
		}
		int x = conv[b];
		sum = sum - b * x + c * x;
		printf("%lld\n", sum);
		conv[c] += conv[b];
		conv[b] = 0;
	}

	return 0;
}