#include <math.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define max(a, b) ((a) > (b) ? (a) : (b))
#define min(a, b) ((a) < (b) ? (a) : (b))

typedef long long int ll;

int read_int()
{
    int a;
    scanf("%d", &a);
    return a;
}

int read_lint()
{
    ll a;
    scanf("%lld", &a);
    return a;
}

int read_float()
{
    float a;
    scanf("%f", &a);
    return a;
}

int read_double()
{
    double a;
    scanf("%lf", &a);
    return a;
}

#define N 200001
ll ruiwa_a[N], ruiwa_b[N];

int main()
{
    int n = read_int();
    int m = read_int();
    ll k = read_lint();

    ruiwa_a[0] = 0;
    for (int i = 1; i <= n; i++)
    {
        ruiwa_a[i] = ruiwa_a[i - 1] + read_lint();
    }

    ruiwa_b[0] = 0;
    for (int i = 1; i <= m; i++)
    {
        ruiwa_b[i] = ruiwa_b[i - 1] + read_lint();
    }

    int ans = 0;
    int b_idx = m;
    for (int i = 0; i <= n; i++)
    {
        if (ruiwa_a[i] > k)
        {
            break;
        }

        while (ruiwa_a[i] + ruiwa_b[b_idx] > k)
        {
            b_idx--;
        }
        ans = max(ans, i + b_idx);
    }
    printf("%ld\n", ans);

    return 0;
}