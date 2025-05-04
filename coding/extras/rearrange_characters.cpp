#include <iostream>
#include <queue>
#include <unordered_map>
#include <vector>

using namespace std;

string rearrangeString(string s) {
    unordered_map<char, int> freq;
    for (char ch : s) {
        freq[ch]++;
    }

    // Max heap (priority queue) sorted by character frequency
    priority_queue<pair<int, char>> maxHeap;
    for (auto& entry : freq) {
        maxHeap.push({entry.second, entry.first});
    }

    string result = "";
    pair<int, char> prev = {-1, '#'}; // To store previous character

    while (!maxHeap.empty()) {
        auto [count, ch] = maxHeap.top();
        maxHeap.pop();
        
        result += ch;

        // Push the previous character back if it still has occurrences left
        if (prev.first > 0) {
            maxHeap.push(prev);
        }

        // Store the current character with one less count for the next iteration
        prev = {count - 1, ch};
    }

    // If the rearranged string length is not equal to input, return empty string (not possible)
    return result.size() == s.size() ? result : "";
}

int main() {
    string s = "aaabbc";
    string result = rearrangeString(s);
    
    if (result.empty()) {
        cout << "Not possible to rearrange\n";
    } else {
        cout << "Rearranged string: " << result << endl;
    }

    return 0;
}
