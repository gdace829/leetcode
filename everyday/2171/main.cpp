#include <vector>
#include <numeric>
#include <algorithm>
using namespace std;
class Solution
{
public:
    long long minimumRemoval(vector<int> &beans)
    {
        sort(beans.begin(), beans.end());
        long long s = accumulate(beans.begin(), beans.end(), 0ll);
        long long ans = s;
        int n = beans.size();
        for (size_t i = 0; i < n; i++)
        {
            ans = min(ans, s - (long long)(beans[i] * (n - i))); /* code */
        }
        return ans;
    }
};