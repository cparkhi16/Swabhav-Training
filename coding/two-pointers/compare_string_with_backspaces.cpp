// Online C++ compiler to run C++ program online
#include <iostream>
#include <bits/stdc++.h>
using namespace std;

int getValidPointer(string s ,int index){
    int backSpaceCount = 0;
    while(index >=0){
        if(s[index] == '#'){
            backSpaceCount++;
        }else if(backSpaceCount > 0){
            backSpaceCount--;
        }else{
            break;
        }
        index--;
    }
    return index;
}
bool checkStrings(string s1, string s2){
    int p1 = s1.length()-1;
    int p2 = s2.length()-1;
    
    while(p1 >= 0 || p2 >= 0){
        int i = getValidPointer(s1, p1);
        int j = getValidPointer(s2, p2);
        
        if(i < 0 && j < 0){
            return true;
        } 
        
        if(i < 0 || j < 0){
            return false;
        }
        
        
        
        if(s1[i] != s2[j]){
            return false;
        }
        p1 = i-1;
        p2 = j-1;
    }
    return true;
}

int main() {
    string s1 = "ab##";
    string s2 = "c#d#";
    if( checkStrings(s1,s2)){
        cout<<" They are equal after backspace processal "<<endl;
    }else{
        cout<<" They are NOT equal after backspace processal "<<endl;
    }

    return 0;
}