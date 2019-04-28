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
func LookupUidGid(username string) (uint32, uint32) {
	u, err := user.Lookup(username)
	if err != nil {
		log.Println("One Error Happend Lookup User: ", username, err)
		return 0, 0
	}

	uid, err := strconv.Atoi(u.Uid)
	if err != nil {
		log.Println("One Error Happend uid atoi  ", username, err)
		return 0, 0
	}
	gid, err := strconv.Atoi(u.Gid)
	if err != nil {
		log.Println("One Error Happend gid atoi  ", username, err)
		return 0, 0
	}
	return uint32(uid), uint32(gid)
}
