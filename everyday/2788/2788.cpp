#include <vector>
#include <string>
using namespace std;
class Solution
{
public:
    vector<string> splitWordsBySeparator(vector<string> &words, char separator)
    {
        vector<string> ans;
        for (string &word : words)
        {
            word += separator;
            string str;
            for (char ch : word)
            {
                if (ch != separator)
                {
                    str += ch;
                }
                else
                {
                    if (!str.empty())
                    {
                        ans.push_back(std::move(str));
                        str.clear();
                    }
                }
            }
        }
        return ans;
    }
};