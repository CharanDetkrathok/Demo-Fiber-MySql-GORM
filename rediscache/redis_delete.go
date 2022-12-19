package rediscache

import (
	"fmt"
	"strings"
)

// ! :: change fmt.Println() to ========> ZapLog
func Delete(keyTemp string) error {

	deleteKey := fmt.Sprintf("%v%v", RedisCaching.EnvPrefix, strings.ToUpper(keyTemp))
	err := RedisCaching.RedisClient.Del(RedisCaching.Contx, deleteKey).Err()
	if err != nil {
		fmt.Println(fmt.Sprintf("DELETE Cache error :: line 12 :: KEY = %v", deleteKey), err)
		return err
	}

	fmt.Println("", fmt.Sprintf("DELETE Cache Success :: KEY = %s", deleteKey))

	return nil
}
