package src

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/pquerna/ffjson/ffjson"
	"log"
)

func CacheDecorator(h gin.HandlerFunc,params string,redisKeyPattern string,empty interface{})  gin.HandlerFunc{

	return func(context *gin.Context) {

		getID:=context.Param(params)
		conn:=RedisDefaultPool.Get()
		defer conn.Close()
		redisKey:=fmt.Sprintf(redisKeyPattern,getID)
		top,err:=redis.Bytes(conn.Do("get",redisKey))
		if err !=nil{ //没取到

			h(context) //执行目标函数得到数据库查询结果
			empty,exists:=context.Get("dbResult")

			retDate,_:=ffjson.Marshal(empty)
			if exists{
				conn.Do("setex",redisKey,50,retDate)
			}else{
				conn.Do("setex",redisKey,20,retDate)
			}

			log.Print("从数据库读取")
			context.JSON(200,empty)
		}else{
			log.Print("从redis读取")
			ffjson.Unmarshal(top,&empty)
			context.JSON(200,empty)
		}

	}

}
