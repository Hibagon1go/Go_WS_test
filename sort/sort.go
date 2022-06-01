package sort
 
import (
	"go_ws_test/models"
	"sort"
)
 
func RankingSort(userList []models.User) []models.User {
	sort.Slice(
		userList,
		func(i, j int) bool {
			return userList[i].AccountID > userList[j].AccountID
		},
	)
 
	return userList
}