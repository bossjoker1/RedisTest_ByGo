package main

import (
	"github.com/go-redis/redis"
	"github.com/satori/go.uuid"
	"log"
	"math"
	"strings"
)

// 标志字符串
const VALID_CHARACTS string = "`abcdefghijklmnopqrstuvwxyz{"

// 由ascii码顺序我们可以向集合添加标记符号来快速查找指定前缀所在的范围
func Find_Prefix_range(predix string) (start string, end string) {
	// 找到最后一个字符在valid字符串中的位置
	posn := strings.IndexAny(VALID_CHARACTS, predix[len(predix)-1:])
	//fmt.Println(predix[len(predix)-1:], "index :", posn)
	if posn <= 0 {
		posn = 1
	}
	// python lambda表达式确实方便。。吐槽一波
	suffix := string(VALID_CHARACTS[posn-1])
	//fmt.Println("suffix: ", suffix)
	return predix[:len(predix)-1] + suffix + "{", predix + "{"
}

// guild同一个协会的
func Auto_Complete_on_Prefix(guild, predix string) {
	start, end := Find_Prefix_range(predix)
	// 获得唯一标识符
	id := uuid.NewV4().String()
	start += id
	end += id
	// 该协会的有序集合
	zset_name := "members:" + guild
	// 添加标记
	RedisDb.ZAdd(zset_name, []redis.Z{{0, start}, {0, end}}...)

	pipe := RedisDb.TxPipeline()
	defer pipe.Close()
	sindex, _ := pipe.ZRank(zset_name, start).Result()
	eindex, _ := pipe.ZRank(zset_name, end).Result()
	erange := math.Min(float64(sindex+9), float64(eindex-2))
	// 删除标记
	pipe.ZRem(zset_name, start, end)
	pipe.ZRange(zset_name, sindex, int64(int(erange)))
	_, err := pipe.Exec()
	if err != nil {
		log.Fatal("Transaction error: ", err)
	}
}
