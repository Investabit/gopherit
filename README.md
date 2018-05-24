# GopherIt!

GopherIt is a collection of convenient utilities for gophers. While there's no
documentation here, there are documentation comments in the source code written
with GoDoc in mind.

## Featured Utilities

### AtomicReplace
AtomicReplace allows one to make multiple different replacements independantly
from each other. This means the function will replace all substrings without
replacing instances which occur after replacing other substrings.

Example
```go
test := "This will be my test string"
output := gostrings.AtomicReplace(test, map[string]string{
    "will be": "is",
    "my test": "the best",
})
fmt.Println(output) // This is the best string
```
