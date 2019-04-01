/**
 * Copyright (C) 2019, Xiongfa Li.
 * All right reserved.
 * @author xiongfa.li
 * @version V1.0
 * Description: 
 */

package goauth2_ext

import (
    "errors"
    "github.com/go-xorm/xorm"
    "github.com/xfali/goauth2-ext/model"
    "net/http"
)

type MysqlManager struct {
    clientEngine *xorm.Engine
    userEngine   *xorm.Engine
}

func NewMysqlManager(clientUrl, userUrl string) *MysqlManager {
    m := &MysqlManager{

    }

    cEngine, err := xorm.NewEngine("mysql", clientUrl)
    if err != nil {
        return nil
    }

    m.clientEngine = cEngine

    uEngine, err := xorm.NewEngine("mysql", userUrl)
    if err != nil {
        return nil
    }

    m.userEngine = uEngine

    return m
}

//验证用户名和密码
func (m *MysqlManager) CheckUser(username, password string) error {
    has, err := m.userEngine.Where("username = ?", username).And("password = ?", password).Exist()
    if !has {
        return errors.New("Check user error")
    }
    return err
}

//当类型为网页授权时，调用该方法检测用户是否登录
//返回重定向授权页面的地址
func (m *MysqlManager) UserAuthorize(r *http.Request) (string, error) {
    _, err := r.Cookie("JSESSIONID")
    if err != nil {
        return "/user/login", nil
    }
    return "/user/authorize", nil
}

//根据client_id查询client_secret
func (m *MysqlManager) QuerySecret(client_id string) (string, error) {
    var client model.Client
    _, err := m.clientEngine.Where("client_id = ?", client_id).Get(&client)
    if err != nil {
        return "", err
    }
    return client.ClientSecret, nil
}

//查询client_id是否可授权scope，可授权返回true
func (m *MysqlManager) CheckScope(client_id string, respType string, scope string) bool {
    return true
}

//检查域名
func (m *MysqlManager) CheckDomainName(client_id string, domain_name string) error {
    return nil
}
