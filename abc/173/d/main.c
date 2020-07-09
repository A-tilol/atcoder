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

#define N 200010

ll a[N];

int greater(const void *arg1, const void *arg2)
{
    return *(int *)arg1 < *(int *)arg2;
}

int main()
{
    ll n = read_lint();

    for (int i = 0; i < n; i++)
    {
        a[i] = read_lint();
    }

    qsort(a, n, sizeof(ll), greater);

    ll ans = 0;
    for (int i = 0; i < n - 1; i++)
    {
        ans += a[(i + 1) / 2];
        // printf("%lld %lld\n", a[i], ans);
    }
    printf("%lld\n", ans);

    return 0;
}