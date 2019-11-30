# gomodnorequire

## import a `.gen` package that is checked in, but doesn't `require` it in go.mod. `go mod tidy` fails with `leading dot in path element`

```
$ go mod tidy
go: finding github.com/xytan0056/depwithgen latest
github.com/xytan0056/gomodnorequire tested by
	github.com/xytan0056/gomodnorequire.test imports
	github.com/xytan0056/depwithgen/.gen/client: malformed module path "github.com/xytan0056/depwithgen/.gen/client": leading dot in path element
```