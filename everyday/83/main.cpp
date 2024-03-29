#include <cstddef>
using namespace std;
struct ListNode
{
    int val;
    ListNode *next;
    ListNode() : val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode *next) : val(x), next(next) {}
};

class Solution
{
public:
    ListNode *deleteDuplicates(ListNode *head)
    {
        if (head == NULL)
        {
            return head;
        }

        ListNode *node = head;
        for (; node->next != NULL;)
        {
            if (node->val == node->next->val)
            {
                node->next = node->next->next;
            }
            else
            {
                node = node->next;
            }
        }
        return head;
    }
};