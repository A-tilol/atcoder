#include <stdio.h>
#include <math.h>

int a[1000001];

int calc_divisor(int aa)
{
    if (aa != 1 && a[1] > 0)
    {
        return 0;
    }

    // printf("%d ---\n", aa);
    int n = (int)sqrt((double)aa);
    for (int i = 2; i <= n; i++)
    {
        if (aa % i != 0)
        {
            continue;
        }
        // printf("%d, %d\n", i, aa / i);
        if (a[i] > 0 || a[aa / i] > 0)
        {
            return 0;
        }
    }
    return 1;
}

int main()
{
    int n;
    scanf("%d", &n);

    for (int i = 0; i < n; i++)
    {
        int aa;
        scanf("%d", &aa);
        a[aa]++;
    }

    int ans = 0;
    for (int i = 1; i < 1000001; i++)
    {
        if (a[i] == 0 || a[i] > 1)
        {
            continue;
        }
        // printf("%d: %d\n", i, a[i], calc_divisor(a[i]));
        if (calc_divisor(i) == 1)
        {
            ans++;
        }
    }
    printf("%d\n", ans);

    return 0;
}