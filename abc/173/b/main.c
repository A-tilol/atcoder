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

void read_string(char *s)
{
    scanf("%s", s);
}

int cnt[4];

int main()
{
    int n = read_int();
    for (int i = 0; i < n; i++)
    {
        char s[4] = {0};
        read_string(s);
        if (strcmp(s, "AC\0\0") == 0)
        {
            cnt[0]++;
        }
        else if (strcmp(s, "WA\0\0") == 0)
        {
            cnt[1]++;
        }
        else if (strcmp(s, "RE\0\0") == 0)
        {
            cnt[3]++;
        }
        else
        {
            cnt[2]++;
        }
    }
    printf("AC x %d\n", cnt[0]);
    printf("WA x %d\n", cnt[1]);
    printf("TLE x %d\n", cnt[2]);
    printf("RE x %d\n", cnt[3]);

    return 0;
}