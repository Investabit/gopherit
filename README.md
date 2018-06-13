# GopherIt!

GopherIt is a collection of convenient utilities for gophers. While there's no
documentation here, there are documentation comments in the source code written
with GoDoc in mind.

## Featured Utilities

### Package errlist
Package errlist makes error handling more convenient in situations where errors
don't affect control flow at the time they are reported.

Example
```go
errs := errlist.New()
for _, datum := range data {
    err := couldFail(data)
    errs.Add(err)
}
return errs.Get() // reports nil when no error occurs
```

Example 2 (handle error without affecting control flow)
```go
errs := errlist.New()
for _, datum := range data {
    err := couldFail(data)
    if errs.Add(err) {
        logFailed(data)
    }
}
return errs.Get() // reports nil when no error occurs
```

### Package goerr
Package goerr provides a constructor for errors that implement both `error` and
`fmt.Stringer`. This can be useful for reporting errors with detailed messages.

Example
```
data, err := body.ReadAll()
if err != nil {
    err = goerr.New(err, fmt.Sprintf(
        "Could not read response!\n\tData:%s\n",
        string(data),
    ))
    return err
}
```

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
