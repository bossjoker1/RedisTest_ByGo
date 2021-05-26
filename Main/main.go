package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

// 声明一个全局的redisDb变量
var RedisDb *redis.Client

// 根据redis配置初始化一个客户端
func initClient() (err error) {
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // redis地址
		Password: "",               // redis密码，没有则留空
		DB:       0,                // 默认数据库，默认是0
	})

	//通过 *redis.Client.Ping() 来检查是否成功连接到了redis服务器
	_, err = RedisDb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	//err := initClient()
	//if err != nil {
	//	//redis连接错误
	//	panic(err)
	//}
	fmt.Println("Redis连接成功")
	start, end := Find_Prefix_range("bachjgksi")
	fmt.Println(start, " -- ", end)
	////第三个参数为失效时间， 0为永久有效
	//err = RedisDb.Set("name", "zbs", 0).Err()
	//if err!=nil{
	//	panic(err)
	//}
	//
	//RedisDb.Append("name", "123")
	////RedisDb.Expire("name", 3*time.Second)
	////time.Sleep(5*time.Second)
	//val, err := RedisDb.Get("name").Result()
	//if err != nil{
	//	panic(err)
	//}
	//fmt.Println(val)
	//RedisDb.LPop("studentsList")
	//RedisDb.RPop("studentsList")
	////只保留区间内的元素
	//RedisDb.LTrim("studentsList", 0, 0)
	//RedisDb.LPush("studentsList", "James", "Marry", "Ken", "Joe")
	//// 0 到 -1即返回所有元素
	//vals, _ := RedisDb.LRange("studentsList", 0, -1).Result()
	//fmt.Println(vals)
	//
	//for _, v := range vals{
	//	fmt.Printf("%s  ", v)
	//}
	//length , _:= RedisDb.LLen("studentsList").Result()
	//fmt.Printf("\nLength is %d\n", length)
	//
	//RedisDb.SAdd("stuSet", 100, 200, 300)
	////获取集合中元素的个数
	//size, err := RedisDb.SCard("stuSet").Result()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(size)
	//
	//es, _ := RedisDb.SMembers("stuSet").Result()
	//fmt.Println(es)
	//
	////根据key和field字段设置,field字段的值。 user_1 是hash key，username 是字段名, admin是字段值
	//err = RedisDb.HSet("user_1", "username", "admin").Err()
	//if err != nil {
	//	panic(err)
	//}
	////根据key和field字段,查询field字段的值。user_1 是hash key，username是字段名
	//username, err := RedisDb.HGet("user_1", "username").Result()
	//
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(username) //admin
	////继续往user_1中添加字段password
	//_ = RedisDb.HSet("user_1", "password", "abc123").Err()
	//
	//zsetKey := "language_rank"
	//languages := []redis.Z{
	//	{Score: 90.0, Member: "Golang"},
	//	{Score: 98.0, Member: "Java"},
	//	{Score: 95.0, Member: "Python"},
	//	{Score: 97.0, Member: "JavaScript"},
	//	{Score: 92.0, Member: "C/C++"},
	//}
	////添加成功的个数
	//num, _ := RedisDb.ZAdd(zsetKey, languages...).Result()
	//fmt.Println(num)
	//
	//size, _ = RedisDb.ZCard(zsetKey).Result()
	////范围为闭区间
	//count, _ := RedisDb.ZCount(zsetKey, "90.0", "95.0").Result()
	//fmt.Printf("size is %d, and count between 90-95 is %d\n", size, count)
	//
	//
	//// 初始化查询条件， Offset和Count用于分页， 第二个参数ZRangeBy
	//op := redis.ZRangeBy{
	//	Min:"80", // 最小分数
	//	Max:"100", // 最大分数
	//	Offset:0, // 类似sql的limit, 表示开始偏移量
	//	Count:5, // 一次返回多少数据
	//}
	////根据分数范围返回集合元素，元素根据分数从小到大排序，支持分页。
	//values, _ := RedisDb.ZRangeByScore(zsetKey, op).Result()
	/////RedisDb.ZRemRangeByRank()
	//for _, val := range values {
	//	fmt.Println(val)
	//}
	//查询指定元素的分数或排名
	//ZScore(key, "元素名")
	//ZRank
	//ZRem(key, "元素名")
	//按从小到大的排序删除元素
	//指定负数即为从高到低
	//ZRemRangeByRank(key, left, right)

	//统计开发语言排行榜
	//zsetKey2 := "language_rank"
	//// 开启一个TxPipeline事务
	//pipe := RedisDb.TxPipeline()
	//defer pipe.Close()
	//// 执行事务操作，可以通过pipe读写redis
	//incr := pipe.Incr(zsetKey2)
	//pipe.Expire(zsetKey, time.Hour)
	//
	//// 通过Exec函数提交redis事务
	//_, err = pipe.Exec()
	//
	//// 提交事务后，我们可以查询事务操作的结果
	//// 前面执行Incr函数，在没有执行exec函数之前，实际上还没开始运行。
	//fmt.Println(incr.Val(), err)
}
