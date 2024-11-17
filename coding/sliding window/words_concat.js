function findWordConcatenation(str, words) {
    if(words.length === 0 || words[0].length === 0) {
      return []
    }
    
    let wordFreq = {}
    
    words.forEach((word) => {
      if(!(word in wordFreq)) {
        wordFreq[word] = 0
      }
      wordFreq[word]++
    })
    console.log(" word freq ",wordFreq)
    const resultIndex = []
    let wordCount = words.length
    let wordLength = words[0].length
  console.log(" word count and word len ",wordCount, wordLength)
    for(let i = 0; i < (str.length - wordCount * wordLength) + 1; i++) {
      const wordsSeen = {}
        console.log(" i =->>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>",i)
      for(let j = 0; j < wordCount; j++) {
        let nextWordIndex = i + j * wordLength
          console.log(" next word index =========",nextWordIndex)
        //get the next word from the string
        const word = str.substring(nextWordIndex, nextWordIndex + wordLength)
          console.log(" word is ",word)
        if(!(word in wordFreq)){
          //break if we don't need this word
            console.log(" break ",word)
          break
        }
        
        //add the word ot the wordsSeen ma
        if(!(word in wordsSeen)){
            console.log(" word seen ",word)
          wordsSeen[word] = 0
        }
        wordsSeen[word]++
        
        //no need to process furrther if the word
        //has higher frequency than required
        if(wordsSeen[word] > (wordFreq[word] || 0)){
          break
        }
        
        if(j + 1 === wordCount){
          //store index if we have found all the words
          resultIndex.push(i)
        }
      }
    }
    return resultIndex
  }
  
  findWordConcatenation("catfoxcat", ["cat", "fox"])