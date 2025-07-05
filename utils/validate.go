package utils

import (
    "net/url"
    "strings"
)

func IsValidURL(raw string) bool {
    u, err := url.ParseRequestURI(raw)
    if err != nil {
        return false
    }
    return strings.HasPrefix(u.Scheme, "http")
}
