#include <iostream>
#include <algorithm>
#include <ranges>
using namespace std;

class Solution
{
public:
    int countWays(vector<vector<int>> &ranges)
    {
        // 按左端点进行排序
        ranges::sort(ranges, [](auto &a, auto &b)
                     { return a[0] < b[0]; });
        int ans = 1, max_r = -1;
        for (auto &p : ranges)
        {
            if (p[0] > max_r)
            {
                ans = ans * 2 % 1'000'000'007;
            }
            max_r = max(max_r, p[1]);
        }
        return ans;
    }
};
int main()
{
    // 创建Solution对象
    Solution solution;

    // 测试用例1: 无重叠区间
    vector<vector<int>> test_case1 = {{1, 2}, {3, 4}, {5, 6}};
    cout << "Test Case 1 (Non-overlapping intervals): Expected output = ⅹⅹⅹ" << endl;
    cout << "Result: " << solution.countWays(test_case1) << endl;

    // 测试用例2: 完全重叠区间
    vector<vector<int>> test_case2 = {{1, 3}, {1, 4}, {1, 3}};
    cout << "\nTest Case 2 (Completely overlapping intervals): Expected output = ⅹⅹⅹ" << endl;
    cout << "Result: " << solution.countWays(test_case2) << endl;

    // 测试用例3: 部分重叠区间
    vector<vector<int>> test_case3 = {{1, 3}, {2, 4}, {5, 6}};
    cout << "\nTest Case 3 (Partially overlapping intervals): Expected output = ⅹⅹⅹ" << endl;
    cout << "Result: " << solution.countWays(test_case3) << endl;

    // 添加更多测试用例以覆盖不同情况...

    return 0;
}