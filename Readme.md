# Effective Go

- Code samples coming out of studying [Effective Go][1]
- Some notes are collected in this document

## General

- Things created with `make` are reference types

## Assignment

- Normally vars can only be declared once in scope
- For multi-assigments (i.e. func cal returns multiple values) some extra rules apply:

    ```go
    res, err := foo() // this is fine

    res2, err := foo2() // this also fine, res2 is a new var

    res, err := foo2() // this is not fine, no new vars introduced

    res, err = foo3() // this is fine

    if err != nil {
        res, err := foo2() // also fine, new block means new scope, res is redeclared
    }
    ```

## Runes

- Runes represent multi-byte characters, i.e. the ones used by utf-8
- utf-8 codepoints use up to 4 bytes, so `rune` is actually an alias for `int32`
- A range over a string iterates over its runes:

    ```go
    str := "abc日本\x80語"

    for _, r := range str {
        fmt.Printf("%T %#U", r, r) // something like "int32 U+65E5 '日'"
    }
    ```

## Type Switching and Casting

- Type switches looks like this:

    ```go
    // t will have dynamic type of thing
    switch t := thing.(type) {
    case string:
        fmt.Println("it is a string")
    default:
        fmt.Println("unknown type")
    }
    ```

- Type assertions use an actual type inside the parens: `t := thing.(string)`
- If type doesn't match, a runtime error will be thrown, or use "comma, ok" idiom:

    ```go
    if t, ok := thing.(string); ok {
        fmt.Println("it's a string")
    }
    ```

[1]:https://golang.org/doc/effective_go.html
