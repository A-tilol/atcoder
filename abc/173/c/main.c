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

char read_char()
{
    char c;
    scanf("%c", &c);
    return c;
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

#define N 6

char c[N][N + 1];
char t[N][N + 1];

int main()
{
    ll h = read_lint();
    ll w = read_lint();
    ll k = read_lint();
    for (int i = 0; i < h; i++)
    {
        read_string(c[i]);
    }

    ll ans = 0;
    ll paternh = (int)pow(2.0, (double)h);
    ll paternw = (int)pow(2.0, (double)w);
    for (int i = 0; i < paternh; i++)
    {
        for (int j = 0; j < paternw; j++)
        {
            for (int l = 0; l < h; l++)
            {
                strcpy(t[l], c[l]);
            }

            int mask = 1;
            for (int l = 0; l < h; l++)
            {
                if (i & mask)
                {
                    for (int m = 0; m < w; m++)
                    {
                        t[l][m] = '.';
                    }
                }
                mask <<= 1;
            }

            mask = 1;
            for (int l = 0; l < w; l++)
            {
                if (j & mask)
                {
                    for (int m = 0; m < h; m++)
                    {
                        t[m][l] = '.';
                    }
                }
                mask <<= 1;
            }

            ll black = 0;
            for (int l = 0; l < h; l++)
            {
                for (int m = 0; m < w; m++)
                {
                    if (t[l][m] == '#')
                    {
                        black++;
                    }
                }
            }

            // printf("%lld\n", black);
            if (black == k)
            {
                ans++;
            }
        }
    }

    printf("%lld\n", ans);

    return 0;
}