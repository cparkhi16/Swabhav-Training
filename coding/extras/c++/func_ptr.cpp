// Online C++ compiler to run C++ program online
#include <iostream>
#include<bits/stdc++.h>
using namespace std;
int add (int a , int b){
    return a + b;
}
int sub(int a , int b){
    return a - b;
}
typedef int(*mathfun)(int, int);

mathfun typeOp(int type){
    if(type == 1)
        return add;
    if (type == 2)
        return sub;
}

 int (*funt(int type)) (int,int){
    if(type == 1)
        return add;
    if (type == 2)
        return sub;
  }
  
 int compare (const void* p , const void* q){
     int l = *(const int*)p;
     int r = *(const int*)q;
     return l < r;
 }
int main() {
    mathfun  ad = typeOp(1);
    int b = ad(1,2);
    cout<<" b is "<<b<<endl;
    
    int (*fun)(int,int) = add;
    cout<<add(1,2)<<endl;
    
    int  (*fu)(int, int) = funt(2);
    cout<<" after sub "<<fu(3,2)<<endl;
    
    mathfun arr[2] = {add, sub};
    cout<<" arr sum "<<arr[0](1,5)<<endl;
    
    
    int (*ar[2]) (int, int) = {add, sub};
    cout<<" ar diff "<<ar[1](9,2)<<endl;
    
    int s[5] = {2,31,10 ,5,7};
    
    int sSize = sizeof(s) / sizeof(s[0]);
    int elemSize = sizeof(s[0]);
    qsort(s,sSize ,elemSize , compare );
    cout<<"after sort "<<endl;
    for(int n : s){
        cout<<" n "<<n<<endl;
    }
    return 0;
}