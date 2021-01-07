package UpdateCache

import (
	"fmt"
	"github.com/Golangltd/cache2go"
	"strconv"
	"time"
)

type myStruct struct {
	text     string
	moreData []byte
}

func init1() {
	cache := cache2go.Cache("myCache")

	val := myStruct{"This is a test!", []byte{}}

	cache.SetAddedItemCallback(func(entry *cache2go.CacheItem) {
		fmt.Println("Added:", entry.Key(), entry.Data(), "---", entry.CreatedOn())
	})

	cache.Add("someKey", 5*time.Second, &val)

	res, err := cache.Value("someKey")
	if err == nil {
		fmt.Println("Found value in cache:", res.Data().(*myStruct).text)
	} else {
		fmt.Println("Error retrieving value from cache:", err)
	}

	// Wait for the item to expire in cache.
	//time.Sleep(6 * time.Second)
	//res, err = cache.Value("someKey")
	//if err != nil {
	//	fmt.Println("Item is not cached (anymore).")
	//}

	cache.SetAboutToDeleteItemCallback(func(e *cache2go.CacheItem) {
		fmt.Println("Deleting:", e.Key(), e.Data().(*myStruct).text, e.CreatedOn())
	})

	// Remove the item from the cache.
	cache.Delete("someKey")

	cache.Flush()

	//--------------------------------------------------------
	/*之前介绍的回调函数都是在添加或删除缓存表项时候触发，而这个dataloader回调则是在调用Value时触发。即如果我们去查找某个key的缓存，如果找不到且我们设置了dataloader回调，就会执行该回调函数。这个功能还是挺实用的，举个例子比如我们缓存了数据库中的一些用户信息，如果我们可以设置dataloader回调，
	如果从缓存里面查找某个用户信息时没有找到，就从数据库中读取该用户信息并加到缓存里面，这个动作就可以加在dataloader回调里面。*/
	// The data loader gets called automatically whenever something
	// tries to retrieve a non-existing key from the cache.
	cache.SetDataLoader(func(key interface{}, args ...interface{}) *cache2go.CacheItem {
		// Apply some clever loading handlers here, e.g. read values for
		// this key from database, network or file.
		val := "This is a test with key " + key.(string)

		// This helper method creates the cached item for us. Yay!
		item := cache2go.NewCacheItem(key, 0, val)
		return item
	})

	// Let's retrieve a few auto-generated items from the cache.
	for i := 0; i < 1; i++ {
		res, err := cache.Value("someKey_" + strconv.Itoa(i))
		if err == nil {
			fmt.Println("Found value in cache:", res.Data())
		} else {
			fmt.Println("Error retrieving value from cache:", err)
		}
	}

	fmt.Println(cache.Count())
}
