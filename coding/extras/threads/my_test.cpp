// Online C++ compiler to run C++ program online
#include <iostream>
#include <bits/stdc++.h>
using namespace std;

void func(int a, int *sum){
    int b = 10;
    *sum = a+b;
}
void funcR(int a, int &sum){
    int b = 10;
    sum = a+b;
}
int main() {
    int sum1 = 0;
    int sum2 = 0;
    thread t1(func , 10 ,&sum1);
    thread t2(func , 15 , &sum2);
    
    t1.join();
    t2.join();
    cout<<" thread func addition sum1 is "<<sum1<<endl;
    cout<<" thread func addition sum2 is "<<sum2<<endl;
    cout<<"End "<<endl;

    int sum3 = 0;
    int sum4 = 0;
    thread t3(funcR , 10 ,ref(sum3));
    thread t4(funcR , 15 , ref(sum4));
    
    t3.join();
    t4.join();
    cout<<" thread func addition sum3 is "<<sum3<<endl;
    cout<<" thread func addition sum4 is "<<sum4<<endl;
    cout<<"End ref ";
    return 0;
}