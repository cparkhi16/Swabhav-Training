// Online C++ compiler to run C++ program online
#include <iostream>
#include <bits/stdc++.h>
using namespace std;

int main() {
    //int arr[10];
    //char arr[10];
    char* arr = (char*)malloc(sizeof(char) * 10);
    memset(arr ,'H' , 10);
    
    for(int i = 0 ; i < 10; i++){
        cout<<" arr[i] "<<arr[i]<<endl;
    }
    
    char src[] = {"Hello"};
    char dst[10];
    
    memcpy(dst, src , sizeof(src));
    cout<<" DST after memcpy is "<<dst<<endl;
    
    memmove(src + 2, src , 3);
    cout<<" src after memmove "<<src<<endl;
    
    

    return 0;
}