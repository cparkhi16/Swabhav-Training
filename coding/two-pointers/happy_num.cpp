#include <iostream>
#include <unordered_set>
using namespace std;
int sumOfSquares(int n) {
    int sum = 0;
    while (n > 0) {
        int digit = n % 10;
        sum += digit * digit;
        n /= 10;
    }
    return sum;
}

bool isHappy(int n) {
    int slow = n;
    int fast = sumOfSquares(n);
    cout<<" fast is "<<fast<<endl;
    cout<<" slow is "<<slow<<endl;
    // The loop continues until fast equals 1 or slow meets fast
    while ( slow != fast) {
        slow = sumOfSquares(slow);         // move slow pointer one step
        fast = sumOfSquares(sumOfSquares(fast)); // move fast pointer two steps
    }
    cout<<"after while fast is "<<fast<<endl;
    cout<<"after while slow is "<<slow<<endl;
    return fast == 1;  // If fast is 1, then the number is happy
}

int main() {
    int n;
    std::cout << "Enter a number: ";
    std::cin >> n;

    if (isHappy(n)) {
        std::cout << n << " is a happy number!" << std::endl;
    } else {
        std::cout << n << " is not a happy number." << std::endl;
    }

    return 0;
}
