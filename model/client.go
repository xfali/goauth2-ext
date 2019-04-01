/**
 * Copyright (C) 2019, Xiongfa Li.
 * All right reserved.
 * @author xiongfa.li
 * @version V1.0
 * Description: 
 */

package model

type Client struct {
    ClientId     string `json:"client_id" xorm:"varchar(128)"`
    ClientSecret string `json:"client_secret" xorm:"varchar(128)"`
}
