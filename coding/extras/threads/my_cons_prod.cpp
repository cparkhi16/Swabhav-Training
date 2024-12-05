#include<bits/stdc++.h>
using namespace std;

mutex mtx;
condition_variable cv;
bool resourceReady = false;
int resource;

void producer(){
    unique_lock<mutex> lock(mtx);

    for(int i = 1; i<=5 ; i++){

        cv.wait(lock, []{ return !resourceReady; });

        resource = i;
        cout<<" Producer producing res "<<resource<<endl;
        resourceReady = true;

        cv.notify_one();
    }
}

void consumer(){
    unique_lock<mutex> lock(mtx);

    for(int i = 1; i<=5 ; i++){

        cv.wait(lock, []{ return resourceReady; });

        cout<<" Consumer consuming res "<<resource<<endl;
        resourceReady = false;

        cv.notify_one();
    }
}

int main(){
    thread t1 (producer);
    thread t2 (consumer);

    t1.join();
    t2.join();
}