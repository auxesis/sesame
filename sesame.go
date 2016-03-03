package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "math/rand"
  "time"
  "github.com/klauspost/cpuid"
)

func allWords(path string) []string {
  wordsbuf, _ := ioutil.ReadFile(path)
  words := string(wordsbuf)
  return strings.Split(words, "\n")
}

func filterWords(words []string, length int) []string {
  var filter []string
  for _, w := range(words) {
    if len(w) == length {
      filter = append(filter, w)
    }
  }
  return filter
}

func sample(words []string) string {
  index := rand.Intn(len(words))
  return words[index]
}

func generatePassword(parts int, words []string, separators []string) string {
  var passwd []string
  for i :=0 ; i < parts; i++ {
    passwd = append(passwd, sample(words) + sample(separators))
  }
  return strings.Join(passwd, "")
}

func main() {
  rand.Seed(time.Now().UTC().UnixNano() * int64(cpuid.CPU.RTCounter()))
  wordsPath := "/usr/share/dict/words"
  all := allWords(wordsPath)
  words := filterWords(all, 6)
  parts := 3
  separators := []string{"0","1","!","?","@","%","#","$","^","&", "*","("}

  for i := 0; i < 10 ; i++ {
    fmt.Println(generatePassword(parts, words, separators))
  }
}
