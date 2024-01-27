#include <vector>
#include <unordered_set>
#include <string>
using namespace std;
class Solution
{
public:
    int maximumNumberOfStringPairs(vector<string> &words)
    {
        int n = words.size();
        int ans = 0;
        unordered_set<int> seen;
        for (size_t i = 0; i < n; ++i)
        {
            if (seen.count(words[i][1] * 100 + words[i][0]))
            {
                ++ans;
            }
            seen.insert(words[i][0] * 100 + words[i][1]);
        }
        return ans;
    }
};
