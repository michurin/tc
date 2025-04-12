# Type converter

Simple every day type conversion tooling.

Thanks to modern golang and generics, you can say now:

```go
pointerToOne = tc.P(1)
```

instead of using tons of typed casters like [this](https://github.com/aws/aws-sdk-go/blob/9d10b7469ebfe21f9ab825461b034f5ac6fc4b8b/aws/convert_types.go#L64).

Or you can use sort of ternary operator:

```go
var messagePrefix = tc.Cmp(os.GetEnv("APP_ENVIRONMENT") == "prod", "", "[FROM STAGING] ")
```

instead of awkward `if`s and `init()`s.

Package provides even sort of extended ternary operator that deals with `*bool`, considering `&true`, `&false` and `nil`:

```go
user := struct {
    Male *bool
}{}
gender := tc.CmpN(user.Male, "male", "female", "n/a")
```

You can simply assume defaults like this:

```go
var tmpDir := tc.DefZero(os.GetEnv("TMP"), "/tmp")
```

It's only several examples. You can work with nil pointers, defaults values and type casting.

Look at documentation and examples.
