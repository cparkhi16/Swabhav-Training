// Online C++ compiler to run C++ program online
#include <iostream>
#include <bits/stdc++.h>
using namespace std;

int getMaxArea(vector<int> heights){
    int left = 0;
    int right = heights.size()-1;
    int area = 0;
    
    while (left<right){
        int curr_area = min(heights[left] , heights[right]) * (right-left);
        area = max(area, curr_area);
        
        if(heights[left] < heights[right]){
            left++;
        }else{
            right--;
        }
    }
    
    return area;
}
int main() {
    vector<int> nums = {1,8,6,2,5,4,8,3,7};
    int res = getMaxArea(nums);
    cout<<" ans is " <<res;
    return 0;
}