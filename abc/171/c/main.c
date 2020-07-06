#include <stdio.h>
#include <stdlib.h>
#include <math.h>
typedef long long int ll;

const long long int MOD = 26;

int main()
{
	long long int n;
	scanf("%lld", &n);

	char s[50];
	int cnt = 0;
	while (n > 0)
	{
		n--;
		s[cnt] = 'a' + n % MOD;
		// printf("%c\n", 'a' + r);
		n /= MOD;
		cnt++;
	}

	for (int i = cnt - 1; i >= 0; i--)
	{
		printf("%c", s[i]);
	}

	return 0;
}