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

int main()
{
    ll n = read_lint();
    ll ans = 0;
    for (ll i = 1; i <= n; i++)
    {
        ans += (i + n / i * i) * (n / i) / 2;
        // printf("%d\n", (i + n / i * i) * (n / i) / 2);
    }
    printf("%lld\n", ans);

    return 0;
}