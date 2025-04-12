# Type converters

[![build](https://github.com/michurin/tc/actions/workflows/ci.yaml/badge.svg)](https://github.com/michurin/tc/actions/workflows/ci.yaml)
[![codecov](https://codecov.io/gh/michurin/tc/graph/badge.svg?token=KL89T2NJCK)](https://codecov.io/gh/michurin/tc)
[![Go Report Card](https://goreportcard.com/badge/github.com/michurin/tc)](https://goreportcard.com/report/github.com/michurin/tc)
[![Go Reference](https://pkg.go.dev/badge/github.com/michurin/tc.svg)](https://pkg.go.dev/github.com/michurin/tc)

Simple every day type conversion tooling.

## Pointers

Thanks to modern golang and generics, you can say now:

```go
pointerToOne = tc.P(1)
```

instead of using tons of typed casters like [this](https://github.com/aws/aws-sdk-go/blob/9d10b7469ebfe21f9ab825461b034f5ac6fc4b8b/aws/convert_types.go#L64).

There are different variants for this function: nil pointer for zero values, defaults...
take a look at [documentation](https://pkg.go.dev/github.com/michurin/tc)

## Types

There are functions to cat `any` to any type with or without fallback to default. Like this:

```go
x := tc.DefCast(ctx.Value("key"), "default") // x = "default" if there is no "key" in context
```

or without default

```go
x := tc.SafeCast[int](ctx.Value("key")) // x = 0 if there is no "key" in context
```

## Ternary operator and similar things

Package provides sort of ternary operator:

```go
// useful for globals
var messagePrefix = tc.Cmp(os.GetEnv("APP_ENVIRONMENT") == "prod", "", "[FROM STAGING] ")
```

and extended operator that deals with `*bool`, considering `&true`, `&false` and `nil`:

```go
user := struct {
    Male *bool
}{}
gender := tc.CmpN(user.Male, "male", "female", "n/a")
```

## Defaults for zero values

You can simply assume defaults for zero values like this:

```go
var tmpDir = tc.DefZero(os.GetEnv("TMP"), "/tmp")
```

## More details

[Documentation](https://pkg.go.dev/github.com/michurin/tc).
