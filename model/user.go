/**
 * Copyright (C) 2019, Xiongfa Li.
 * All right reserved.
 * @author xiongfa.li
 * @version V1.0
 * Description: 
 */

package model

type User struct {
    Username string `json:"username" xorm:"varchar(128)"`
    Password string `json:"password" xorm:"varchar(128)"`
}
