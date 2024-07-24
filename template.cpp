#include <bits/stdc++.h>
#define REP0(n) for (auto _i = std::decay_t<decltype(n)>{}; _i != (n); ++_i)
#define REP1(i, n) for (auto i = std::decay_t<decltype(n)>{}; i != (n); ++i)
#define REP2(i, l, r) for (auto i = (l); i != (r); ++i)
#define OVERLOAD_REP(_1, _2, _3, NAME, ...) NAME
#define rep(...) OVERLOAD_REP(__VA_ARGS__, REP2, REP1, REP0)(__VA_ARGS__)
#define all(...) std::begin(__VA_ARGS__), std::end(__VA_ARGS__)
#define rall(...) std::rbegin(__VA_ARGS__), std::rend(__VA_ARGS__)
#define debug(x) cerr << "\033[33m(line:" << __LINE__ << ") " << #x << ": " << x << "\033[m" << endl;
using namespace std;
using ll = long long;
using ull = unsigned long long;
struct Init
{
    Init()
    {
        ios::sync_with_stdio(0);
        cin.tie(0);
    }
} init;
inline int rint()
{
    int x;
    cin >> x;
    return x;
}
inline float rfloat()
{
    float x;
    cin >> x;
    return x;
}
inline vector<int> rvi(int n)
{
    vector<int> v(n);
    rep(i, n) cin >> v[i];
    return v;
}
inline vector<float> rvf(int n)
{
    vector<float> v(n);
    rep(i, n) cin >> v[i];
    return v;
}

int main()
{

    return 0;
}
