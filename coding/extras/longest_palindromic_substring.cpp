#include <iostream>
#include <string>

using namespace std;

string expandAroundCenter(string s, int left, int right) {
    
    while (left >= 0 && right < s.size() && s[left] == s[right]) {
        cout<<" left "<<left<<" right "<<right<<endl;
        left--;
        right++;
    }
    cout <<" sending this subs "<<s.substr(left + 1, right - left - 1)<<endl;
    return s.substr(left + 1, right - left - 1);
}

string longestPalindrome(string s) {
    if (s.empty()) return "";
    string longest = s.substr(0, 1); // At least one character is a palindrome
    
    for (int i = 0; i < s.size(); i++) {
        // Odd length palindromes
        cout<<" calling for odd i is "<<i<<endl;
        string oddPal = expandAroundCenter(s, i, i);
        if (oddPal.size() > longest.size()) {
            longest = oddPal;
        }
        
        cout<<" calling for even i is "<<i<<endl;
        // Even length palindromes
        string evenPal = expandAroundCenter(s, i, i + 1);
        if (evenPal.size() > longest.size()) {
            longest = evenPal;
        }
    }
    
    return longest;
}

int main() {
    string s = "babad";
    cout << "Longest Palindromic Substring: " << longestPalindrome(s) << endl;
    return 0;
}
