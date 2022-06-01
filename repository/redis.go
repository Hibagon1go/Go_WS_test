package repository
 
import (
	"context"
	"encoding/json"
	"go_ws_test/models"
 
	"github.com/go-redis/redis/v8"
)
 
var Cache *redis.Client
 
func SetupRedis() {
	Cache = redis.NewClient(&redis.Options{
		// docker-compose.ymlに指定したservice名+port
		Addr: "redis:6379",
		DB:   0,
	})
}
 
func GetUserList(uuid string) ([]models.User, error) {
	data, err := Cache.Get(context.Background(), uuid).Result()
	if err != nil {
		return nil, err
	}
 
	userList := new([]models.User)
	err = json.Unmarshal([]byte(data), userList)
	if err != nil {
		return nil, err
	}
	return *userList, nil
}
 
func GetRankings() (map[string]float64, error) {
	// zrevrangebyscore
	//http://mogile.web.fc2.com/redis/commands/zrevrangebyscore.html
	rankings, err := Cache.ZRevRangeByScoreWithScores(
		context.Background(),
		"rankings", // updateRankings.goでkeyとして設定した値
		&redis.ZRangeBy{
			Min: "-inf",
			Max: "+inf",
		}).Result()
 
	if err != nil {
		return nil, err
	}
 
	result := make(map[string]float64)
 
	for _, ranking := range rankings {
		result[ranking.Member.(string)] = ranking.Score
	}
 
	return result, nil
}