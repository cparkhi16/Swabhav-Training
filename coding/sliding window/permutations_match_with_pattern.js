function findPermutation(str, pattern) {
    //sliding window
    let windowStart = 0
    let isMatch = 0
    let charFrequency = {}
    
   for(i = 0; i < pattern.length; i++) {
     const char = pattern[i]
     if(!(char in charFrequency)) {
       charFrequency[char] = 0
     }
     charFrequency[char]++
   }
      console.log(" char freq map is ",charFrequency)
    
    //our goal is to math all the characters from charFrequency with the current window
    //try to extend the range [windowStart, windowEnd]
    for(windowEnd = 0; windowEnd < str.length; windowEnd++) {
      const endChar = str[windowEnd]
        console.log(" end char  is ",endChar)
      if(endChar in charFrequency) {
        //decrement the frequency of the matched character
        charFrequency[endChar]--
          console.log(" updated char freq map ",charFrequency)
        if(charFrequency[endChar] === 0) {
          isMatch++
        }
      }
        console.log(" ismatch",isMatch)
      if(isMatch === Object.keys(charFrequency).length) {
        return true
      }
      
      //shrink the sliding window
      if(windowEnd >= pattern.length - 1) {
        let startChar = str[windowStart]
        windowStart++
        if(startChar in charFrequency) {
          if(charFrequency[startChar] === 0) {
            isMatch--
          }
          charFrequency[startChar]++
        }
      }
    }
    return false
  }
  
  
  findPermutation("odicf", "dc")//false