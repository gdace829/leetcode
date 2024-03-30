#include <iostream>
#include <vector>
using namespace std;
class Solution
{
public:
    int minimumSum(vector<int> &nums)
    {
        int min_left;
        int min_right;
        min_left = nums[0];
        min_right = nums[nums.size() - 1];
        vector<int> res(nums.size(), INT32_MAX);
        for (int i = 1; i < nums.size() - 1; i++)
        {
            if (min_left < nums[i])
            {
                res[i] = nums[i] + min_left;
            }

            if (nums[i] < min_left)
            {
                min_left = nums[i];
            }
        }

        for (int i = nums.size() - 2; i >= 1; i--)
        {

            if (min_right < nums[i] && res[i] < INT_MAX)
            {
                res[i] += min_right;
            }
            else
            {
                res[i] = INT_MAX;
            }

            if (nums[i] < min_right)
            {
                min_right = nums[i];
            }
        }
        int min = res[0];

        for (int i = 1; i < res.size() - 1; i++)
        {

            if (res[i] < INT32_MAX && res[i] < min)
            {
                min = res[i];
            }
        }
        if (min < INT32_MAX)
        {
            return min;
        }
        else
        {
            return -1;
        }
    }
};
int main()
{
    Solution s;
    vector<int> input{6, 5, 4, 3, 4, 5};
    cout << s.minimumSum(input) << endl;
    return 0;
}