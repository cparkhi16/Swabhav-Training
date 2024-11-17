#include <iostream>
#include <string>
using namespace std;
// Function to swap characters at positions i and j in a string
void swap(std::string &str, int i, int j) {
    char temp = str[i];
    str[i] = str[j];
    str[j] = temp;
}

// Recursive function to generate permutations using backtracking
void permute(std::string &str, int start, int end) {
    cout<<" start is "<<start<<" end is "<<end<<endl;
    if (start == end) {
        std::cout << str << std::endl; // Print the current permutation
    } else {
        for (int i = start; i <= end; i++) {
            cout<<" before swapp "<<str<<" start is "<<start<<" i is "<<i<<endl;
            swap(str, start, i); // Swap current index with start
            cout<<" after swapp "<<str<<endl;
            permute(str, start + 1, end); // Recur for the rest of the string
            cout<<" after permute -----= "<<str<<endl;
            swap(str, start, i); // Backtrack to restore the original string
            cout<<" og swapp "<<str<<endl;
        }
    }
}

int main() {
    std::string str = "ABC";
    int n = str.size();
    permute(str, 0, n - 1); // Start generating permutations
    return 0;
}
