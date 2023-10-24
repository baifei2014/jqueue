This lib jqueue is a golang queue lib. The goal of this library is to provide more Golang queue implementations.

***dont support concurrent security***

## Examples

### priority queue

#### comparer

- priority.Lesser
- priority.Greater

#### usages

```golang
// put new element
pq := priority.New(priority.Lesser)

item := GridLocation{
    x: 10,
    y: 20,
}
pq.Put(item, 10)

// get element
pq.Get()
```

