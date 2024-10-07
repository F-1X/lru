#### Краткая информация: 

- **lru** - реализован на двусвязном списке.  
- **lru_int_pool** - реализован на двусвязном списке, только для key int, value int (для теста)  
- **lru_pool** - реализован на двусвязном списке, используя ```sync.Pool``` для хранения звеньев Node списка.  

Преимущественно используйте lru_pool.  

____

#### Команды:  

Юнит тесты:  
```sh  
make unit_lru_pool
```  
Бенчмарк:  
```sh  
make bench_lru_pool
```
Рейс детектор:
```sh  
make race 
```

Дополнительные команды можно найти в ```Makefile```.
____

#### Использование:
```go
package main

import (
	"fmt"
	"log"
	"lru/lru_pool"
	"time"
)

func main() {
	capacity := 8192
	cache := lru_pool.NewCache(capacity)

	cache.Add("key1", "value1")

	cache.AddWithTTL("key2", "value2", time.Minute)

	value, ok := cache.Get("key1")
	if !ok {
		log.Fatal("unexcisting key")
	}
	fmt.Println(value)

	value, ok = cache.Get("key2")
	if !ok {
		log.Fatal("unexcisting key")
	}
	fmt.Println(value)
}

```
____

