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

char read_char()
{
    char c;
    scanf("%c", &c);
    return c;
}

void read_string(char *s)
{
    scanf("%s", s);
}

#define N 200001

char s[N], t[N];

int main()
{
    read_string(s);
    read_string(t);

    int ans = 0;
    for (int i = 0; i < N; i++)
    {
        if (s[i] == 0)
        {
            break;
        }
        if (s[i] != t[i])
        {
            ans++;
        }
    }
    printf("%d\n", ans);

    return 0;
}