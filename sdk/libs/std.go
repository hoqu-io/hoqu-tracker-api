package libs

import (
    "io"
    "bytes"
)

// Check whether the given string is in the give slice of
// string. If is a case sensitive check.
func StringInSlice(str string, list []string) bool {
    for _, v := range list {
        if v == str {
            return true
        }
    }
    return false
}

func StreamToBuffer(stream io.Reader) *bytes.Buffer {
    buf := new(bytes.Buffer)
    buf.ReadFrom(stream)
    return buf
}

func MapString(x []string, f func(string) string) []string {
    y := make([]string, len(x))

    for _, xval := range x {
        y = append(y, f(xval))
    }

    return y
}

func FilterString(x []string, f func(string) bool) []string {
    y := x[:0]

    for _, xval := range x {
        if f(xval) {
            y = append(y, xval)
        }
    }

    return y
}

func ReduceString(x []string, f func(string, string) string, initial string) string {
    y := initial

    for _, xval := range x {
        y = f(y, xval)
    }

    return y
}

