/**
 * Copyright (C) 2019, Xiongfa Li.
 * All right reserved.
 * @author xiongfa.li
 * @version V1.0
 * Description: 
 */

package goauth2_ext

import (
    "github.com/go-redis/redis"
    "log"
    "strings"
    "time"
)

type RedisDataManager struct {
    client *redis.Client
}

//初始化
func (dm *RedisDataManager) Init() {
    dm.client = redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    pong, err := dm.client.Ping().Result()
    log.Println(pong, err)

    ret, err := dm.client.ScriptLoad("").Result()
    if err != nil {
        log.Fatal("load script error")
    }

    log.Println(ret)
}

//关闭
func (dm *RedisDataManager) Close() {
    dm.client.Close()
}

//保存Code相关信息，绑定client_id以及scope，在expireIn时间之后自动失效
func (dm *RedisDataManager) SaveCode(client_id, code, scope string, expireIn time.Duration) error {
    _, err := dm.client.Set(code, client_id+":"+scope, expireIn).Result()
    return err
}

//通过code获得client_id以及scope
func (dm *RedisDataManager) GetCode(code string) (string, string, error) {
    v, err := dm.client.Get(code).Result()

    if err != nil {
        return "", "", err
    }

    strArr := strings.Split(v, ":")
    return strArr[0], strArr[1], nil
}

//删除code
func (dm *RedisDataManager) DelCode(code string) error {
    _, err := dm.client.Del(code).Result()
    return err
}

//保存refresh token
func (dm *RedisDataManager) SaveRefreshToken(token_data string, refresh_token string, refresh_expire time.Duration) error {
    dm.client.E
}

//保存refresh token以及access_token
func (dm *RedisDataManager) SaveAccessToken(token_data string, access_token string, access_expire time.Duration) error {

}

//通过refresh token获取保存的token data
func (dm *RedisDataManager) GetRefreshToken(refresh_token string) (string, error) {

}

//通过access token获取保存的token data
func (dm *RedisDataManager) GetAccessToken(access_token string) (string, error) {

}

//废弃client_id绑定的token，包括refresh token及access token
func (dm *RedisDataManager) RevokeToken(client_id string) {

}
