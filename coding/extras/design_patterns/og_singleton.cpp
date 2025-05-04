#include <iostream>
#include <bits/stdc++.h>
using namespace std;

class Singleton {
public:
    Singleton(const Singleton& o) = delete;
    Singleton operator=(const Singleton& o)= delete;
    static Singleton& getInstance() {
        if (!instance) {
            instance = new Singleton();
        }
        return *instance;
    }

    void getLog() {
        cout << " log from singleton " << endl;
    }

private:
    Singleton() { cout << " singleton instance created " << endl; }
    static Singleton* instance;
};

// Initialize the static member outside the class
Singleton* Singleton::instance = nullptr;
int main() {
    Singleton& l = Singleton::getInstance();  // Access the Singleton instance
    l.getLog();

    return 0;
}
