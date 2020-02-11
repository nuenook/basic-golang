
## Keywords

### new
 new(T)
allocate with default type e.g. bool `false`,
integer `0`, float `0.0` string `""` and 
pointers, functions, interfaces, slices channels, maps is `nil`
returned address

### make
 make(T, args)
different from `new` 
 `make` for creating slice, map and channel only
and will return that value with default value (not pointer)


### map 
map[KeyType]ValueType

e.g. map[string]float32

example
```golang
var grades map[string]float32

grades := make(map[string]float32)

grades["Timmy"] = 42
grades["Jess"] = 92
grades["Sam"] = 71


// specific values
timGrade := grades["Timmy]

// remove items from map ny key
delete(grades, "Timmy")


// Iterate over maps

for k, v := range grades {
    // k, v
}
```