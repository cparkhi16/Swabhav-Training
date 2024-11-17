#include <iostream>
#include <thread>
#include <chrono>
#include <algorithm>
using namespace std;
using namespace std::chrono;
typedef long long int ull;

// Add padding or use separate memory regions to avoid cache contention
void findEven(ull start, ull end, ull *EvenSum) {
    //ull tempEvenSum = 0;
    for (ull i = start; i <= end; ++i) {
        if (!(i & 1)) {
            *EvenSum += i;
        }
    }
    //*EvenSum = tempEvenSum;
    cout << "EvenSum : " << *EvenSum << endl;
}

void findOdd(ull start, ull end, ull *OddSum) {
    //ull tempOddSum = 0;
    for (ull i = start; i <= end; ++i) {
        if (i & 1) {
            *OddSum += i;
        }
    }
    // *OddSum = tempOddSum;
    cout << "OddSum : " << *OddSum << endl;
}

int main() {
    ull start = 0, end = 1900000000;
    ull OddSum = 0;
    ull EvenSum = 0;

    auto startTime = high_resolution_clock::now();

    // WITH THREAD, pass by pointer
    std::thread t1(findEven, start, end , &EvenSum); // First half range
    std::thread t2(findOdd, start, end, &OddSum); // Second half range

    t1.join();
    t2.join();

    // findEven(start,end, &EvenSum);
	// findOdd(start, end, &OddSum);
    auto stopTime = high_resolution_clock::now();
    auto duration = duration_cast<microseconds>(stopTime - startTime);

    cout << "Sec: " << duration.count() / 1000000 << endl;

    return 0;
}
