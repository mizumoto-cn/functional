# Functional

Lo and behold, a new functional toolkit for Go hath arrived!

## License

This project is licensed under the Mizumoto General Public License v1.3. See [LICENSE](LICENSE/Mizumoto.General.Public.License.v1.3.md) for more information.

## Quick Start

```bash
go get github.com/mizumoto-cn/functional
```

## Types and Functions

### Type Interfaces

#### Ord

```go
// Compare :: a -> a -> Ordering
// Ordering = [ types.LT | types.EQ | types.GT ]
// LT, LE, GT, GE, EQ :: a -> a -> Bool
// Comparability
// x <= y || y <= x = True
// Transitivity
// if x <= y && y <= z = True, then x <= z = True
// Reflexivity
// x <= x = True
// Antisymmetry
// if x <= y && y <= x = True, then x == y = True
type Integer Ord {}

// Compare :: a -> a -> Ordering
// Ordering = LT | EQ | GT
func (i Integer) Compare(j Integer) Ordering {
    if i < j {
        return types.LT
    } else if i > j {
        return types.GT
    } else {
        return types.EQ
    }
}

// LT :: a -> a -> Bool
func (i Integer) LT(j Integer) bool {
    return i.Compare(j) == types.LT
}

// LE :: a -> a -> Bool
func (i Integer) LE(j Integer) bool {
    return i.Compare(j) != types.GT
}

// GT :: a -> a -> Bool
func (i Integer) GT(j Integer) bool {
    return i.Compare(j) == types.GT
}

// GE :: a -> a -> Bool
func (i Integer) GE(j Integer) bool {
    return i.Compare(j) != types.LT
}

// EQ :: a -> a -> Bool
func (i Integer) EQ(j Integer) bool {
    return i.Compare(j) == types.EQ
}

```

### Optional

#### Maybe

```go
package main

import (
    "fmt"
    "github.com/mizumoto-cn/functional"
)

func main() {
    // Maybe
    var m functional.Maybe
    m = functional.Just(1)
    fmt.Println(m) // Just 1
    fmt.Println(m.Get()) // 1
    fmt.Println(m.IsPresent()) // true
    fmt.Println(m.IsNothing()) // false

    m = functional.Just(nil)
    fmt.Println(m) // Nothing
    fmt.Println(m.Get()) // functional.Nothing -> nil
    fmt.Println(m.IsPresent()) // false
    fmt.Println(m.IsNothing()) // true

    // Or()
    m = functional.Just(1)
    fmt.Println(m.Or(2)) // 1
    
    m = functional.Just(nil)
    fmt.Println(m.Or(2)) // 2
    
    // Optional Let()
    // f :: Maybe a -> (a -> Maybe b) -> Maybe b
    v := functional.Just(1).Let(func(i interface{}) interface{} {
        return i.(int) + 1
    })  // 2
    fmt.Println(v) // Just 2
    
    v = functional.Just(nil).Let(func(i interface{}) interface{} {
        return i.(int) + 1
    })  // nil
    fmt.Println(v) // Nothing

}
```

#### Either

```go
package main

import (
    "fmt"
    "github.com/mizumoto-cn/functional/optional"
)

func main() {
    // Either
    var e optional.Either
    e = optional.Left(1)
    fmt.Println(e) // Left 1
    fmt.Println(e.Get()) // 1
    fmt.Println(e.IsLeft()) // true
    fmt.Println(e.IsRight()) // false

    e = optional.Right(1)
    fmt.Println(e) // Right 1
    fmt.Println(e.Get()) // 1
    fmt.Println(e.IsLeft()) // false
    fmt.Println(e.IsRight()) // true

    // Or()
    e = optional.Left(1)
    fmt.Println(e.Or(2)) // 1
    
    e = optional.Right(1)
    fmt.Println(e.Or(2)) // 1
    
    // Either Let()
    // f :: Either a b -> (a -> c) -> (b -> c) -> c
    v := optional.Left(1).Let(
        func(i interface{}) interface{} {
            return i.(int) + 1
        },
        func(i interface{}) interface{} {
            return i.(int) + 2
        },
    )  // 2
    fmt.Println(v) // 2
    
    v = optional.Right(1).Let(
        func(i interface{}) interface{} {
            return i.(int) + 1
        },
        func(i interface{}) interface{} {
            return i.(int) + 2
        },
    )  // 3
    fmt.Println(v) // 3

}
```

### Array

#### ArrayList

```go
package main

import (
    "fmt"
    "github.com/mizumoto-cn/functional/types/list"
)

func main() {
    // ArrayList
    var a list.ArrayList
    a = list.NewArrayList(1, 2, 3)
    fmt.Println(a) // [1 2 3]
    fmt.Println(a.Get(0)) // 1
    fmt.Println(a.Get(1)) // 2
    fmt.Println(a.Get(2)) // 3
    fmt.Println(a.Get(3)) // nil
    fmt.Println(a.Get(-1)) // nil
    
    // head tail last init
    fmt.Println(a.Head()) // 1
    fmt.Println(a.Tail()) // [2 3]
    fmt.Println(a.Last()) // 3
    fmt.Println(a.Init()) // [1 2]

    // length
    fmt.Println(a.Length()) // 3
    // null
    fmt.Println(a.Null()) // false
    fmt.Println(list.NewArrayList().Null()) // true
    // reverse
    fmt.Println(a.Reverse()) // [3 2 1]
    // drop
    fmt.Println(a.Drop(1)) // [2 3]
    fmt.Println(a.Drop(10)) // []
    // take
    fmt.Println(a.Take(1)) // [1]
    fmt.Println(a.Take(2)) // [1 2]
    fmt.Println(a.Take(10)) // [1 2 3]
    // max min
    fmt.Println(a.Max()) // 3
    fmt.Println(a.Min()) // 1
    // sum product
    fmt.Println(a.Sum()) // 1+2+3 = 6
    fmt.Println(a.Product()) // 1*2*3 = 6
    // elem index
    fmt.Println(a.Elem(1)) // true
    fmt.Println(a.Elem(10)) // false
    // range
    fmt.Println(list.NewRange(0, 1, 10)) // [1 2 3 4 5 6 7 8 9 10]
    fmt.Println(list.NewRange('a', 'b', 'e')) // [a b c d e]
    // cycle
    fmt.Println(list.NewCycle([]int{1, 2, 3}).Take(10)).Drop(1)) // [2 3 1 2 3 1 2 3 1]
    fmt.Println(list.NewCycle([]int{1}).Take(10) // [1 1 1 1 1 1 1 1 1]

}
```

#### Map, FlatMap

utils.Map, utils.FlatMap

```go
package main

import (
    "fmt"
    "github.com/mizumoto-cn/functional/types/list"
    "github.com/mizumoto-cn/functional/utils"
)

func main() {
    // Map
    a := list.NewArrayList(1, 2, 3)
    fmt.Println(utils.Map(a, func(i interface{}) interface{} {
        return i.(int) + 1
    })) // [2 3 4]
    
    // FlatMap
    a := list.NewArrayList(1, 2, 3)
    fmt.Println(utils.FlatMap(a, func(i interface{}) interface{} {
        return list.NewArrayList(i.(int), i.(int) + 1)
    })) // [[1 2] [2 3] [3 4]]
}
```

### Monad (IO)

```go
package main

import (
    "fmt"
    "github.com/mizumoto-cn/functional/types/monad"
)

func main() {
    // IO
    io := monad.NewIO(func() interface{} {
        return 1
    })
    fmt.Println(io) // IO 1
    fmt.Println(io.Run()) // 1
    
    // IO Let()
    // f :: IO a -> (a -> IO b) -> IO b
    io = monad.NewIO(func() interface{} {
        return 1
    }).Let(func(i interface{}) interface{} {
        return i.(int) + 1
    })  // 2

    fmt.Println(io) // IO 2
    fmt.Println(io.Run()) // 2
    
    // IO Bind()
    // f :: IO a -> (a -> IO b) -> IO b
    io = monad.NewIO(func() interface{} {
        return 1
    }).Bind(func(i interface{}) monad.IO {
        return monad.NewIO(func() interface{} {
            return i.(int) + 1
        })
    })  // 2
    fmt.Println(io) // IO 2
    fmt.Println(io.Run()) // 2
    
    // IO Map()
    // f :: IO a -> (a -> b) -> IO b
    io = monad.NewIO(func() interface{} {
        return 1
    }).Map(func(i interface{}) interface{} {
        return i.(int) + 1
    })  // 2
    fmt.Println(io) // IO 2
    fmt.Println(io.Run()) // 2
    
    // IO FlatMap()
    // f :: IO a -> (a -> IO b) -> IO b
    io = monad.NewIO(func() interface{} {
        return 1
    }).FlatMap(func(i interface{}) monad.IO {
        return monad.NewIO(func() interface{} {
            return i.(int) + 1
        })
    })  // 2
    fmt.Println(io) // IO 2
    fmt.Println(io.Run()) // 2
    
    // IO Join()
    // f :: IO (IO a) -> IO a
    io = monad.NewIO(func() interface{} {
        return monad.NewIO(func() interface{} {
            return 1
        })
    }).Join()  // 1
    fmt.Println(io) // IO 1
    fmt.Println(io.Run()) // 1
    
    // IO Just
    // f :: a -> IO a
    io = monad.IO(functions.Just(1))  // 1
    fmt.Println(io) // IO 1
    fmt.Println(io.Run()) // 1
}
```

### Pattern Matching

```go
package main

import (
    "fmt"
    "github.com/mizumoto-cn/functional/types/pattern"
)

func main() {
    // Pattern
    p := pattern.NewPattern(1)
    fmt.Println(p.Match(1)) // true
    fmt.Println(p.Match(2)) // false
}
```

## Roadmap

### Milestones

- [ ] Type System
  - [ ] Ord
  - [ ] ...
- [ ] Pattern Matching
  - [ ] Pattern
  - [ ] ...
- [ ] Monad
  - [ ] IO
  - [ ] ...
- [ ] Optionals
  - [ ] Maybe
  - [ ] Either
  - [ ] ...
- [ ] Data Structures
  - [ ] ArrayList
  - [ ] ...
- [ ] Utils
  - [ ] Map
  - [ ] FlatMap
  - [ ] ...
- [ ] Lambda
- [ ] ...
