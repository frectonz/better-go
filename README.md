# The Go Programming Language + `try`

`try` a slightly better error handling primitive.

try sugar

```golang
content, err := ioutil.ReadFile(filename)
try err or string
```

try desugared

```golang
...
if err != nil {
    return make([]string, 1)[0], err
}
```

The try statement expects two values.

1. the error variable
2. the type of the return value (this is needed because go has this convention of returning the zero value of a type incase of an error)

