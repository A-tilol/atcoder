#include <stdio.h>
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

int a[200000];
int bits[31];

int main()
{
	int n = read_int();

	for (int i = 0; i < n; i++)
	{
		a[i] = read_int();
		int mask = 1;
		for (int j = 0; j < 31; j++)
		{
			// printf("%d %d -> %d\n", a[i], mask, a[i] & mask);
			if (a[i] & mask)
			{
				bits[j] = bits[j] ^ 1;
			}
			mask <<= 1;
		}
	}

	for (int i = 0; i < n; i++)
	{
		int ans = 0;
		int mask = 1;
		for (int j = 0; j < 31; j++)
		{
			int bit = 0;
			if (a[i] & mask)
			{
				bit = 1;
			}
			ans += mask * (bit ^ bits[j]);
			// printf("%d %d %d %d\n", a[i], mask, bits[j], (a[i] & mask) ^ bits[j]);
			mask <<= 1;
		}
		printf("%d\n", ans);
	}

	return 0;
}
