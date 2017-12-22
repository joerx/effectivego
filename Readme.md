# Effective Go

- Code samples coming out of studying [Effective Go][1]
- Some notes are collected in this document

## General

- Stuff created with `make` are reference types

## Assignment

- Normally vars can only be declared once in scope
- For multi-assigments (i.e. func cal returns multiple values) some extra rule apply

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
- A range over a string actually iterates over its runes:

    ```go
    str := "abc日本\x80語"

    for _, r := range str {
        fmt.Printf("%T %#U", r, r) // something like "int32 U+65E5 '日'"
    }
    ```

## Type Switching and Casting

- Type switch looks like this: `t := thing.(type)`, `t` will have dynamic type of `thing`
- Type assertions use an actual type inside the parens: `t := thing.(string)`
- Use 'comma, ok' idiom or a runtime error will be thrown:

    ```go
    if t, ok := thing.(string); ok {
        fmt.Println("it's a string")
    }
    ```

[1]:https://golang.org/doc/effective_go.html