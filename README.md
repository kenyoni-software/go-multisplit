# Go MultiSplit

Go MultiSplit is a `go/analysis`-based linter that detects multiple declarations, assignments, struct fields and function parameters or return values. It rewrites them into single declarations, statements or fields, making code clearer and more explicit.

## Options

No fix will be suggested
- if a declaration or statement has a comment attached, to avoid changing the intent of the comment
- If a multiple declaration or statement uses an anonymous struct type, to avoid duplicate code

All **\*-to-block** options rewrite declarations into a grouped block. It will never unblock.

- **`var-decl-pkg`** (default: `true`)  
    Split multiple `var` declarations at package scope.  
    If **`var-decl-pkg-to-block`** (default: `true`) is enabled, declarations are rewritten into a grouped block.

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

* **`var-decl-func`** (default: `false`)
    Split multiple `var` declarations inside function bodies.
    If **`var-decl-func-to-block`** (default: `false`) is enabled, declarations are rewritten into a grouped block.

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

* **`var-decl-init-pkg`** (default: `true`)
    Split multiple `var` declarations with initializers at package scope.
    If **`var-decl-init-pkg-to-block`** (default: `true`) is enabled, declarations are rewritten into a grouped block.

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

* **`var-decl-init-func`** (default: `false`)
    Split multiple `var` declarations with initializers inside functions.
    If **`var-decl-init-func-to-block`** (default: `false`) is enabled, declarations are rewritten into a grouped block.
    If **`var-decl-init-func-to-short`** (default: `true`) is enabled, non-block declarations without explicit type specification are rewritten to short variable declarations (`:=`). Shorthand takes precedence if applicable.

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

* **`const-decl-pkg`** (default: `true`)
    Split multiple `const` declarations at package scope.
    If **`const-decl-pkg-to-block`** (default: `true`) is enabled, declarations are rewritten into a grouped block.

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

* **`const-decl-func`** (default: `false`)
    Split multiple `const` declarations in function scope.
    If **`const-decl-func-to-block`** (default: `false`) is enabled, declarations are rewritten into a grouped block.

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

* **`func-params`** (default: `true`)
    Split multiple named function parameters into individual parameters.

    Before:

    ```go
    func f(a, b int) {}
    ```

    After:

    ```go
    func f(a int, b int) {}
    ```

* **`func-return-values`** (default: `true`)
    Split multiple named function return values into individual return values.

    Before:

    ```go
    func f() (a, b int) { return }
    ```

    After:

    ```go
    func f() (a int, b int) { return }
    ```

* **`assign`** (default: `false`)
    Split multiple named assignments into individual assignment statements.

    Before:

    ```go
    a, b = 1, 2
    ```

    After:

    ```go
    a = 1
    b = 2
    ```

* **`short-var-decl`** (default: `false`)
    Split multiple short declarations into individual short declarations.

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
    Split multiple struct fields into individual fields.

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

## golangci-lint

```yaml
linters-settings:
  multisplit:
    # The set of rules to apply. If empty, the default rules will be applied.
    # Default: var-decl-pkg, var-decl-init-pkg, const-decl-pkg, func-params, func-return-values, struct-fields
    rules:
      - var-decl-pkg
      - var-decl-func
      - assign


    varDeclPkgToBlock: false
    varDeclFuncToBlock: true
    varDeclInitPkgToBlock: false
    varDeclInitFuncToBlock: true
    varDeclInitFuncToShort: false
    constDeclPkgToBlock: false
    constDeclFuncToBlock: true
```
