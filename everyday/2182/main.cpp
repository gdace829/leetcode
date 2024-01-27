#include <string>
#include <vector>
using namespace std;
const int N = 26;
class Solution
{
public:
    string repeatLimitedString(string s, int repeatLimit)
    {
        vector<int> count(N);
        for (char a : s)
        {
            count[a - 'a']++;
        }
        int m = 0;
        string res;
        for (int x = N - 1, y = N - 2; x >= 0 && y >= 0;)
        {
            if (count[x] == 0)
            {
                x--;
                m = 0;
            }
            else if (m < repeatLimit)
            {
                res.push_back('a' + x);
                count[x]--;
                m++;
            }
            else if (y >= x || count[y] == 0)
            {
                y--;
                /* code */
            }
            else
            {
                res.push_back('a' + y);
                count[y]--;
                m = 0;
            }
        }
        return res;
    }
};