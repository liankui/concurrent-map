package main

import (
    "fmt"
    cmap "github.com/orcaman/concurrent-map/v2"
)

// 更新为泛型代码对比 https://github.com/orcaman/concurrent-map/commit/b62fe33532f54e844634ecfad05209bbbcb69ddf
func main() {
    // Create a new map.
    m := cmap.New[string]()

    // Sets item within map, sets "bar" under key "foo"
    m.Set("foo", "bar")

    // Retrieve item from map.
    bar, ok := m.Get("foo")
    fmt.Printf("value=%v, exist=%v\n", bar, ok)

    // Removes item under key "foo"
    m.Remove("foo")
}
