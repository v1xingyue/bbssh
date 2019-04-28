package user

import (
	"log"
	"os/user"
)

// 检测是否是root权限
func IsRoot() bool {
	u, err := user.Current()
	if err != nil {
		log.Fatal("One Error Happend GetUser : ", err)
		return false
	}
	return u.Username == "root"
}

// 返回uid,gid,并不中断程序
func LookupUidGid(username string) (string, string) {
	u, err := user.Lookup(username)
	if err != nil {
		log.Println("One Error Happend Lookup User: ", username, err)
		return "", ""
	}
	return u.Uid, u.Gid
}
