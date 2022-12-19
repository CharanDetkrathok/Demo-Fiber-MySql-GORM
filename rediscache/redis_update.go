package rediscache

import (
	"encoding/json"
	"fmt"
	"time"
)

// ! :: change fmt.Println() to ========> ZapLog
func Update(newCache MakeCache) error {

	setKey := fmt.Sprintf("%v%v", RedisCaching.EnvPrefix, newCache.Key)

	if newCache.Key != "" {

		jsonStrNewCache, _ := json.Marshal(newCache.Data)
		expireTime, _ := time.ParseDuration(newCache.Expire)
		err := RedisCaching.RedisClient.Set(RedisCaching.Contx, setKey, jsonStrNewCache, expireTime).Err()
		if err != nil {
			fmt.Println(fmt.Sprintf("SET Cache error :: KEY = %v", setKey), err)
			return err
		}

		fmt.Println("", fmt.Sprintf("SET Cache Success :: KEY = %s", setKey))
		return nil
	}

	fmt.Println(fmt.Sprintf("SET Cache error :: KEY = %v", setKey), fmt.Errorf("can't set cahce :: key is empty or invalid format"))
	return fmt.Errorf("can't set cahce")
}
