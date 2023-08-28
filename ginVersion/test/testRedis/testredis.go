package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
)

func main() {
	red := redis.NewClient(&redis.Options{
		Addr:         "121.37.246.78:6379",
		Password:     "181234",
		DB:           0,
		PoolSize:     100,
		MinIdleConns: 10,
	})
	ctx := context.TODO()
	testgetandset(red)
	//expire,err:=red.Expire("zhangsan")
	//testhashtable(red)
	//redistestzset(red)
	//testredisSet(red)
	expire, err := red.TTL(ctx, "zhangsan").Result()
	fmt.Printf("expire=%v   ,err=%v", expire, err)
}
func testgetandset(cli *redis.Client) {
	ctx := context.TODO()
	//设置key
	ress, err := cli.Set(ctx, "lisi", 10, time.Hour).Result()
	fmt.Println(ress, err)
	ress, err = cli.Set(ctx, "lisi", 20, time.Hour).Result()
	fmt.Println(ress, err)
	resg, _ := cli.Get(ctx, "1265103034@qq.com").Result()
	fmt.Println(resg)
}
func testhashtable(cli *redis.Client) {
	//HashTable
	//设置HashTable数据变量
	ctx := context.TODO()
	reshs, _ := cli.HSet(ctx, "lisi", "name", "lisi", "age", 18, "sex", "男").Result()
	fmt.Println(reshs)
	//获取HashTable某个键的某个属性值
	reshg, _ := cli.HGet(ctx, "lisi", "name").Result()
	fmt.Println(reshg)
	//获取某个键的所有多个属性值
	res, _ := cli.HMGet(ctx, "lisi", "name", "age", "sex").Result()
	fmt.Println(res...)
}
func redistestzset(cli *redis.Client) {
	//Zset
	//设置zset,k-v数据格式
	ctx := context.TODO()
	scores := []redis.Z{
		{Score: 78, Member: "java"},
		{Score: 88, Member: "python"},
		{Score: 99, Member: "C++"},
		{Score: 89, Member: "Golang"},
	}
	reszs, _ := cli.ZAdd(ctx, "scores", redis.Z{Score: 79, Member: "C"}).Result()
	fmt.Println(reszs)
	res, _ := cli.ZAdd(ctx, "scores", scores...).Result()
	fmt.Println(res)
	fmt.Println("Range:", cli.ZRange(ctx, "scores", 0, 4).Val())
	fmt.Println("RevRange", cli.ZRevRange(ctx, "scores", 0, 5).Val())
	opt := redis.ZRangeBy{
		Min:    "50",
		Max:    "100",
		Offset: 0,
		Count:  5,
	}
	fmt.Println("SearchRange", cli.ZRangeByScore(ctx, "scores", &opt).Val())
	fmt.Println("RangeWithScores", cli.ZRangeWithScores(ctx, "scores", 0, 4).Val())
}
func testredisSet(cli *redis.Client) {
	ctx := context.TODO()
	//添加元素到集合里面
	fmt.Println("SAdd:", cli.SAdd(ctx, "user", "zhangsan", "lisi", "wangwu", "zhaoliu").Val())
	//判断元素在不在集合中
	fmt.Println("IsInSetOrNot:", cli.SIsMember(ctx, "user", "wang").Val())
	//获取集合里的所有元素
	fmt.Println("SGetAllMembers:", cli.SMembers(ctx, "user").Val())
}
