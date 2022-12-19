package rediscache

import (
	"encoding/json"
	"fmt"
)

// ! :: change fmt.Println() to ========> ZapLog
func Get(key string) (interface{}, error) {

	getByKey := fmt.Sprintf("%v%v", RedisCaching.EnvPrefix, key)
	value, err := RedisCaching.RedisClient.Get(RedisCaching.Contx, getByKey).Result()
	if err != nil {

		if err.Error() == "redis: nil" {
			fmt.Println(fmt.Sprintf("GET Cache not found :: line 15 :: KEY = %v", getByKey), err)
			return nil, err
		}

		fmt.Println(fmt.Sprintf("GET Cache error :: line 19 :: KEY = %v", getByKey), err)
		return nil, err
	}

	var responseData interface{}
	json.Unmarshal([]byte(value), &responseData)

	dataIndentation, _ := json.MarshalIndent(responseData, "", "    ")
	fmt.Println("GET Cache Success ", dataIndentation)

	return responseData, nil
}
