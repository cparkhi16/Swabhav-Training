function findSubstring(str, pattern) {
  let windowStart = 0
  let matched = 0
  let substrStart = 0
  let minLength = str.length + 1
  let charFreq = {}
  
  for(let i = 0; i < pattern.length; i++) {
    const char = pattern[i]
    if(!(char in charFreq)) {
      charFreq[char] = 0
    }
    charFreq[char]++
  }
  console.log("og char freq ",charFreq)
  //try to extend the range [windowStart, windowEnd]
  for(let windowEnd = 0; windowEnd < str.length; windowEnd++) {
    const endChar = str[windowEnd]
    if(endChar in charFreq) {
      charFreq[endChar]--
        console.log(" char freq map ",charFreq, endChar)
      if(charFreq[endChar] >= 0) {
        //count every matching of a character
        matched++
      }
    }
    
    //Shrink the window if we can, finish as soon as we remove a 
    //matched character
    while(matched === pattern.length) {
        console.log(" matched ois same as pettern len")
      if(minLength > windowEnd - windowStart + 1) {
        minLength = windowEnd - windowStart + 1
        substrStart = windowStart
      }
      
      const startChar = str[windowStart]
      windowStart++
      if(startChar in charFreq) {
        if(charFreq[startChar] === 0) {
            console.log(" char for which match is decr by -1 ",startChar,matched)
        matched--
      }
           console.log(" bef incr start char count ",charFreq[startChar],startChar)
      charFreq[startChar]++
        console.log(" incr start char count ",charFreq[startChar],startChar)
      }
    }
  } 
  if(minLength > str.length) {
  return ''
}
return str.substring(substrStart, substrStart + minLength)
}

findSubstring("aabdec", "abc")