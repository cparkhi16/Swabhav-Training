// Online C++ compiler to run C++ program online
#include <bits/stdc++.h>
using namespace std;

string removeSubstring(string og , string patt ){
    int n = og.length();
    int m = patt.length();
    string res = "";
    for(int i = 0 ; i < n ;){
        int j = 0;
        while(j < m && i+j < n && og[i+j] == patt[j]){
            j++;
        }
        if(j == m){
            i += m;
        }else{
            res = res + og[i];
            i++;
        }
    }
    return res;
}
int main() {
   std::cout << removeSubstring("Blahcat", "cat") << std::endl;           // Output: Blah
    std::cout << removeSubstring("Blahcatcat", "cat") << std::endl;       // Output: Blah
    std::cout << removeSubstring("blahcacat", "cat") << std::endl;        // Output: blahca
    std::cout << removeSubstring("Blahcat", "lahc") << std::endl;         // Output: Bat
    std::cout << removeSubstring("Blahcacatcatt", "cat") << std::endl;    // Output: Blaht
    std::cout << removeSubstring("\"\"", "cat") << std::endl;             // Output: ""
    std::cout << removeSubstring("Blah", "\"\"") << std::endl;            // Output: Blah

    return 0;
}