#include <ranges>
#include <iostream>
using namespace std;
class Solution
{
public:
    int minimumAddedCoins(vector<int> &coins, int target)
    {
        ranges::sort(coins);
        int ans = 0, s = 1, i = 0;
        while (s <= target)
        {
            if (i < coins.size() && coins[i] <= s)
            {
                s += coins[i++];
            }
            else
            {
                s += s;
                ans++;
            }
        }
        return ans;
    }
};
int main()
{
    vector<int> a{1, 4, 10};
    Solution s;
    cout << s.minimumAddedCoins(a, 19) << endl;
}