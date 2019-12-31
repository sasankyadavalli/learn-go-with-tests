package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const french = "French"
const spanish = "Spanish"

func Hello(name string, language string) string {
  if name == "" {
    name = "World"
  }
  prefix := greetingPrefix(language)

  return prefix + name
}

func greetingPrefix(language string) (prefix string) {
  switch language {
  case french:
    prefix = frenchHelloPrefix
  case spanish:
    prefix = spanishHelloPrefix
  default:
    prefix = englishHelloPrefix
  }
  return prefix
}

func main(){
  fmt.Println(Hello("Sasank", ""))
}
