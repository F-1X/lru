#### Краткая информация: 

lru - реализован на двусвязном списке.  
lru_int_pool - реализован на двусвязном списке, только для key int, value int (для теста)  
lru_pool - реализован на двусвязном списке, используя sync.Pool для хранения звеньев Node списка.  

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
#### Комментарии:  
Мною были проведены попытки попытки реализации на основе структуры данных:  

0) Doubly Linked List (без сторонних библиотек), с собственной реализацией приоритетов: 2 Priority Queue (как в 3. пункте) но разделенных, что дает накладные расходы на память. Да и в целом, мне показалось странным иметь несколько сторонних структур, для реализации одной задумки, так как часть работы по хранению, вытеснению они уже брали на себя. И мне показалось что можно объекдинить их в 1 Priority Queue. Поэтому решено было переписывать. Что касается производительности, то скорость была на 50ns выше чем у golang-lru от hasicorp (замерял отдельно). 
1) Очередь (slice): Показала низкую скорость работы, пришлось переписывать.    
2) Список (container/list): Также показал низкую производительность в тестах.
3) Приоритетная очередь (container/heap) - реализация: высший приоритет у Node c Node.expire == 0. Затем (если у всех есть время) вытесняется тот у кого меньшее время, если время одинаковое (сама сложная часть, так как необходимы маркеры сохранения порядка при добавлении, но из-за переупорядочивания элементы менялись местами, или вытаскивался не тот элемент который ожидался: приоритет переупорядочивался и нарушал порядок LRU) - идея была отложена из-за нехватки времени, так как для этого требуется более глубокое понимание работы: Priority Queue (container/heap) Less, Swap, Pop, Push, Remove и тд, но я уверен что был близок к цели, и,возможно, вернусь к этой структуре данных попоже.   
4) Двусвязный список (без особых приоритетов, текущая реализация): Не использовать мнемоники (эвристики) для вытеснения LRU по "особому" приоритету, в этой связи я отказался от очередей, и обрабатывал без TTL (задал как большую константу времени) и с TTL. Плохо показал себя lru/lru.go с точки зрения аллокаций на операцию, оптимизировав это: дописал на sync.Pool, аллокации уменьшились на 2 на операцию (см. lru_pool/bench). Замена any на int, привела к 0 алокациям и поразительной скорости (см. lru_int_pool/bench), вероятно при реализации на дженериках, такаая тенденция должна оставаться (в golang-lru hashicorp используется именно дженерики, в бенчмарк тестах там действительно 0 аллокаций (проверил проведя тестирование), поэтому я делаю такой вывод, но это не точно).   


