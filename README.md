# The Go Programming Language + `try` + `unwrap`

Go with slightly better error handling primitives.
Checkout the [demo](./demo) directory.

## `try`

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

## `unwrap`

unwrap sugar

```golang
fileContent, err := readFile("sample.txt")
unwrap err
```

unwrap desugared

```golang
...
if err != nil {
    panic(err)
}
```

The unwrap statement expects one value, the error variable.

# How to run it

```
nix develop
cd src
./make.bash
```
