#include <iostream>
using namespace std;
#include <bits/stdc++.h>

double medianOfTwoSortedArrays(vector<int>& nums1 , vector<int>& nums2){
    if(nums1.size() > nums2.size()){
        swap(nums1, nums2);
    }
    
    int m = nums1.size();
    int n = nums2.size();
    
    int low = 0;
    int high = m;
    
    while(low <= high){
        int mid1 = (low + high) / 2;
        int mid2 = (m + n + 1 ) / 2 - mid1;
        
        int l1 = (mid1 == 0) ? INT_MIN : nums1[mid1-1];
        int r1 = (mid1 == m) ? INT_MAX : nums1[mid1];
        
        int l2 = (mid2 == 0) ? INT_MIN : nums2[mid2-1];
        int r2 = (mid2 == n) ? INT_MAX : nums2[mid2];
        
        if( l1<=r2 && l2<= r1){
            if( (m+n) % 2 == 0){
                return ((max(l1,l2) + min(r1,r2)) / 2.0);
            }else{
                return (max(l1,l2));
            }
        }
        else if( l1 > r2){
            high = mid1-1;
        }else{
            low = mid1+1;
        }
    }
    return 1.0;
}

// Driver code
int main() {
    vector<int> n1 = {1, 2};
    vector<int> n2 = {3,4};
    
    auto res = medianOfTwoSortedArrays(n1 , n2);
    
    cout<<" median of two sorted array is "<<res<<endl;

    return 0;
}
