#include <iostream>
#include <string>
#include <climits>
#include <algorithm>

using namespace std;

int nextGreaterElement(int n) {
    string digits = to_string(n);
    int i = digits.size() - 2;
    
    // Step 1: Find the first decreasing digit
    while (i >= 0 && digits[i] >= digits[i + 1]) {
        i--;
    }
    
    // Step 2: If no such decreasing digit is found, return -1 (no next greater element)
    if (i < 0) return -1;
    
    // Step 3: Find the rightmost digit that is larger than digits[i]
    int j = digits.size() - 1;
    while (digits[i] >= digits[j]) {
        j--;
    }
    
    // Step 4: Swap digits[i] and digits[j]
    swap(digits[i], digits[j]);
    
    // Step 5: Reverse the digits after the position i
    reverse(digits.begin() + i + 1, digits.end());
    
    // Convert the string back to a number and check for overflow
    long long res = stoll(digits);
    return (res > INT_MAX) ? -1 : res;
}

int main() {
    // Test cases
    int testCases[] = {12, 21, 12345, 54321, 230241};
    
    for (int n : testCases) {
        int result = nextGreaterElement(n);
        cout << "Next greater element for " << n << " is: " << result << endl;
    }

    return 0;
}
