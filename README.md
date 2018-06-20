# GopherIt!

GopherIt is a collection of convenient utilities for gophers.

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

### Package ltools
Package ltools provides utility functions to make Go programming less redundant.
Tools here may use reflection and type assertions to provide functionality that
would be seen more commonly in other languages.

Developer beware: these functions directly oppose Go's ideology of verbose
easy-to-understand source code.

#### MapKeys
MapKeys reports a slice of strings for any map of type `map[string]` or
`map[fmt.Stringer]`. Any type of key that implements `fmt.Stringer` is also
valid. This function panics if a map with an invalid key type is passed.

Example
```go
uniqueNames := map[string]struct{}{}
uniqueNames["one"] = struct{}{}
uniqueNames["two"] = struct{}{}

var keys []string
keys = MapKeys(uniqueNames)
fmt.Println(keys) // [one two]
```

### Package goerr
Package goerr provides a constructor for errors that implement both `error` and
`fmt.Stringer`. This can be useful for reporting errors with detailed messages.

Example
```go
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
