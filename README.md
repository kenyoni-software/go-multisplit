# Go MultiSplit

![maintained](https://img.shields.io/badge/maintained-yes-brightgreen.svg)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/kenyoni-software/go-multisplit/blob/main/LICENSE.md)
![Programming Language](https://img.shields.io/badge/language-Go-orange.svg)

[![Go report card](https://goreportcard.com/badge/github.com/kenyoni-software/go-multisplit)](https://goreportcard.com/report/github.com/kenyoni-software/go-multisplit)

Go MultiSplit is a `go/analysis`-based linter that detects when multiple identifiers are declared, assigned or listed together (e.g. `var a, b int` or `a, b := 1, 2`). It recommends splitting these into separate declarations or assignments to improve clarity and reduce ambiguity.

## Options

No fix will be suggested
- if a declaration or statement has a comment attached, to avoid changing the intent of the comment
- if a multiple declaration or statement uses an anonymous struct type, to avoid duplicate code

All **\*-to-block** options rewrite declarations into a grouped block. Grouped blocks will never be unblocked.

* **`assign`** (default: `false`)
    Split assignments with multiple targets into individual assignments.

    Before:

    ```go
    a, b = 1, 2
    ```

    After:

    ```go
    a = 1
    b = 2
    ```

* **`const-decl-func`** (default: `false`)
    Split `const` declarations with multiple identifiers inside functions into individual declarations.  
    If **`const-decl-func-to-block`** (default: `false`) is enabled, the individual declarations are placed inside a `const` block.

    Before:

    ```go
    func f() {
        const x, y = 1, 2
    }
    ```

    After (with `const-decl-func-to-block=true`):

    ```go
    func f() {
        const (
            x = 1
            y = 2
        )
    }
    ```

    After (with `const-decl-func-to-block=false`):

    ```go
    func f() {
        const x = 1
        const y = 2
    }
    ```

* **`const-decl-pkg`** (default: `true`)
    Split `const` declarations with multiple identifiers at package scope into individual declarations.  
    If **`const-decl-pkg-to-block`** (default: `true`) is enabled, the individual declarations are placed inside a `const` block.

    Before:

    ```go
    const a, b = 1, 2
    ```

    After (with `const-decl-pkg-to-block=true`):

    ```go
    const (
        a = 1
        b = 2
    )
    ```

    After (with `const-decl-pkg-to-block=false`):

    ```go
    const a = 1
    const b = 2
    ```

* **`func-params`** (default: `true`)
    Split function parameters with multiple identifiers into individual parameters.

    Before:

    ```go
    func f(a, b int) {}
    ```

    After:

    ```go
    func f(a int, b int) {}
    ```

* **`func-return-values`** (default: `true`)
    Split function return values with multiple identifiers into individual return values.

    Before:

    ```go
    func f() (a, b int) { return }
    ```

    After:

    ```go
    func f() (a int, b int) { return }
    ```

* **`short-var-decl`** (default: `false`)
    Split short variable declarations with multiple identifiers into individual declarations.

    Before:

    ```go
    a, b := 1, 2
    ```

    After:

    ```go
    a := 1
    b := 2
    ```

* **`struct-fields`** (default: `true`)
    Split struct field declarations with multiple identifiers into individual fields.

    Before:

    ```go
    type S struct {
        a, b int
    }
    ```

    After:

    ```go
    type S struct {
        a int
        b int
    }
    ```

* **`var-decl-func`** (default: `false`)
    Split `var` declarations with multiple identifiers inside function bodies into individual declarations.  
    If **`var-decl-func-to-block`** (default: `false`) is enabled, the individual declarations are placed inside a `var` block.

    Before:

    ```go
    func f() {
        var x, y int
    }
    ```

    After (with `var-decl-func-to-block=true`):

    ```go
    func f() {
        var (
            x int
            y int
        )
    }
    ```

    After (with `var-decl-func-to-block=false`):

    ```go
    func f() {
        var x int
        var y int
    }
    ```


- **`var-decl-pkg`** (default: `true`)  
    Split `var` declarations with multiple identifiers at package scope into individual declarations.  
    If **`var-decl-pkg-to-block`** (default: `true`) is enabled, the individual declarations are placed inside a `var` block.

    Before:
    ```go
    var a, b int
    ```

    After (with `var-decl-pkg-to-block=true`):

    ```go
    var (
        a int
        b int
    )
    ```

    After (with `var-decl-pkg-to-block=false`):

    ```go
    var a int
    var b int
    ```

* **`var-decl-init-func`** (default: `false`)
    Split multiple `var` declarations with multiple identifiers and initializers inside functions into individual declarations.  
    If **`var-decl-init-func-to-block`** (default: `false`) is enabled, the individual declarations are placed inside a `var` block.  
    If **`var-decl-init-func-to-short`** (default: `true`) is enabled and the declarations are not placed in a block, any declaration without an explicit type will be rewritten to a short variable declaration (`:=`). The short form takes precedence when applicable.

    Before:

    ```go
    func f() {
        var x, y = 1, 2
    }
    ```

    After (with `var-decl-init-func-to-block=true`):

    ```go
    func f() {
        var (
            x = 1
            y = 2
        )
    }
    ```

    After (with `var-decl-init-func-to-block=false`):

    ```go
    func f() {
        var x = 1
        var y = 2
    }
    ```

    After (with `var-decl-init-func-to-short=true`):

    ```go
    func f() {
        x := 1
        y := 2
    }
    ```

* **`var-decl-init-pkg`** (default: `true`)
    Split `var` declarations with multiple identifiers and initializers at package scope into individual declarations.  
    If **`var-decl-init-pkg-to-block`** (default: `true`) is enabled, the individual declarations are placed inside a `var` block.

    Before:

    ```go
    var a, b = 1, 2
    ```

    After (with `var-decl-init-pkg-to-block=true`):

    ```go
    var (
        a = 1
        b = 2
    )
    ```

    After (with `var-decl-init-pkg-to-block=false`):

    ```go
    var a = 1
    var b = 2
    ```

## golangci-lint

```yaml
linters-settings:
  multisplit:
    # The set of rules to apply. If empty, the default rules will be applied.
    # Default: const-decl-pkg, func-params, func-return-values, struct-fields, var-decl-init-pkg, var-decl-pkg
    rules:
      - assign
      - var-decl-func
      - var-decl-pkg

    # If enabled rewrite const declarations with multiple identifiers inside functions into a const block.
    # Default: false
    constDeclFuncToBlock: true
    # If enabled rewrite const declarations with multiple identifiers at package scope into a const block.
    # Default: true
    constDeclPkgToBlock: false
    # If enabled rewrite var declarations with multiple identifiers inside functions into a const block.
    # Default: false
    varDeclFuncToBlock: true
    # If enabled rewrite var declarations with multiple identifiers at package scope into a var block.
    # Default: true
    varDeclPkgToBlock: false
    # If enabled rewrite var declarations with multiple identifiers and initializers inside functions into a var block.
    # Default: false
    varDeclInitFuncToBlock: true
    # If enabled rewrite untyped var declarations with multiple identifiers inside functions into short declarations.
    # Default: true
    varDeclInitFuncToShort: false
    # If enabled rewrite var declarations with multiple identifiers and initializers at package scope into a var block.
    # Default: false
    varDeclInitPkgToBlock: false
```
