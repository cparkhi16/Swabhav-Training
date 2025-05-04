// Online C++ compiler to run C++ program online
#include <iostream>
#include <bits/stdc++.h>
using namespace std;

int f(int ind, vector<int> &heights, vector<int> &dp, int k)
{

    if (ind == 0)
    {
        return 0;
    }
    if (dp[ind] != -1)
        return dp[ind];
    int mini = INT_MAX;
    for (int i = 1; i <= k; i++)
    {
        int l1 = INT_MAX;
        if((ind - i )>=0){
            l1 = f(ind - i, heights, dp,k) + abs(heights[ind] - heights[ind - i]);
        }
        mini = min(l1, mini);
    }
    return dp[ind] = mini;
}
int main()
{
    vector<int> heights = {10, 20, 30, 10};
    int k =2;
    vector<int> dpe(heights.size(),-1);
    cout<<" minimum energy reqd by the frog with k jumps "<<f(heights.size()-1, heights,dpe,k);

    int n = heights.size();
    vector<int> dp(n, 0);
    
    dp[0] = 0;
    for (int i = 1; i < n; i++)
    {
        int mini = INT_MAX;
        for (int j = 1; j <= k; j++)
        {
            int l1 = INT_MAX;
            if ((i - j) >= 0)
            {
                l1 = dp[i - j] + abs(heights[i] - heights[i - j]);
            }
            mini = min(l1,mini);
        }
        dp[i] = mini;
    }
    cout<<" minimum energy reqd by the frog "<<dp[n-1]<<endl;
    
    return 0;
}