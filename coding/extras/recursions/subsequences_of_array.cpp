// Online C++ compiler to run C++ program online
#include <iostream>
#include <bits/stdc++.h>
using namespace std;
void permute(std::vector<int>& res ,std::vector<int> arr , int ind , int n){
    //cout<<" start "<<endl;
    if(ind == n){
        for(int i = 0 ; i < res.size() ; i++){
            cout<<res[i]<<" ";
        }
        if(res.size() == 0){
            cout<<"{}";
        }
        cout<<endl;
        return;
    }
    res.push_back(arr[ind]);
    permute(res, arr , ind+1 , n);
    //cout<<" here";
    res.pop_back();
    permute(res, arr , ind+1 , n);
}
int main() {
    std::vector <int> arr = {3,1,2};
    std::vector <int> res;
    permute(res , arr , 0 ,3 );

    return 0;
}