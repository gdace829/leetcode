#include <iostream>
#include <vector>
using namespace std;
class Graph
{
    vector<vector<int>> g; // 邻接表

public:
    Graph(int n, vector<vector<int>> &edges) : g(n, vector<int>(n, INT_MAX / 2))
    {
        for (auto &e : edges)
        {
            g[e[0]][e[1]] = e[2];
            /* code */
        }
    }

    void addEdge(vector<int> edge)
    {
        g[edge[0]][edge[1]] = edge[2]; // 添加边
    }

    int shortestPath(int node1, int node2)
    {
        int n = g.size();
        vector<int> dis(n, INT_MAX / 2), vis(n);
        dis[node1] = 0;
        while (true)
        {
            int x = -1;
            for (size_t i = 0; i < vis.size(); i++)
            {
                if (vis[i] != 1 && (x < 0 || dis[i] < dis[x]))
                {
                    x = i;
                }
            }
            if (x < 0 || dis[x] == INT_MAX / 2)
            {
                return -1;
            }
            if (x == node2)
            {
                return dis[x];
            }
            vis[x] = true;
            for (size_t i = 0; i < n; i++)
            {
                dis[i] = min(dis[i], dis[x] + g[x][i]);
            }

            /* code */
        }
    }
};
