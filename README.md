# goset

A generic set implementation for Go >= 1.18.

```go
// Create set, passing initial values of any comparable type
s := NewSet(4, 8)
s.Add(15, 16)
s.Size()
// 4

s.Discard(8)
// s = {4, 15, 16}

// When creating an empty set, pass a type
x := NewSet[int]()
x.Add(30)
x.Disjoint(s)
// returns new set {4, 15, 16, 30}

x.Has(30)
// true

x.Discard(30)
// x = {4, 15, 16}

x.Add(700)
x.Union(y)
// returns new set {4, 15, 16, 700}

z := NewSet[int]()
z.Add(15, 16, 402)
z.Intersection(x)
// returns new set {15, 16}

x.Subtract(z)
// x = {4, 700}

x.Values()
// slice [4 700]

x2 := NewSet[int]()
x2.Add(4)
x2.Add(700)
x2.EqualTo(x)
// true
```

# License

© 2021 Ryan Plant

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

The Software is provided "as is", without warranty of any kind, express or implied, including but not limited to the warranties of merchantability, fitness for a particular purpose and noninfringement. In no event shall the authors or copyright holders be liable for any claim, damages or other liability, whether in an action of contract, tort or otherwise, arising from, out of or in connection with the Software or the use or other dealings in the Software.
