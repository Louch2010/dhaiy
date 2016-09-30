package gdb

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"strconv"
	"time"

	. "github.com/louch2010/dhaiy/common"
	"github.com/louch2010/dhaiy/log"
	"github.com/louch2010/goutil"
)

//创建gdb文件
func CreateGDB(filePath string, tables map[string]*CacheTable) error {
	//初始化md5
	md5Ctx := md5.New()
	//文件头
	content := GDB_GOCACHE + GDB_VERSION
	md5Ctx.Write([]byte(content))
	//表内容
	b := tables2Byte(tables)
	content += b
	md5Ctx.Write([]byte(b))
	//文件尾
	content += GDB_EOF
	md5Ctx.Write([]byte(GDB_EOF))
	//md5计算签名
	cipherStr := md5Ctx.Sum(nil)
	content = content + hex.EncodeToString(cipherStr)
	//写入文件
	err := ioutil.WriteFile(filePath, []byte(content), 0)
	if err != nil {
		log.Error("写入gdb文件失败：", err)
		return err
	}
	return nil
}

//表转换
func tables2Byte(tables map[string]*CacheTable) string {
	content := ""
	//表统计信息
	content += GDB_DATABASE + goutil.StringUtil().IntToStr(len(tables)-1, GDB_LEN_DATABASE_SIZE)
	for name, table := range tables {
		//系统表不缓存
		sysTable := GetSystemConfig().MustValue("server", "sysTable", "sys")
		if name == sysTable {
			continue
		}
		//表信息（表名、键值对数量）
		content += GDB_TABLE + goutil.StringUtil().IntToStr(len(name), GDB_LEN_KEY) + name
		content += goutil.StringUtil().IntToStr(table.ItemCount(), GDB_LEN_TABLE_SIZE)
		for k, v := range table.GetItems() {
			//已经过期的，则不保存
			if v.LiveTime() > 0 && v.CreateTime().Add(v.LiveTime()).Before(time.Now()) {
				continue
			}
			dv, dt := toString(v.Value(), v.DataType())
			//数据类型
			content += dt
			//过期时间（有过期时间则记录14位过期时间，没有过期时间则记录a，表示一直有效）
			if v.LiveTime() > 0 {
				content += goutil.DateUtil().Time14Format(v.CreateTime().Add(v.LiveTime()))
			} else {
				content += GDB_LIVETIME_ALWAYS
			}
			//键长、键
			content += goutil.StringUtil().IntToStr(len(k), GDB_LEN_KEY) + k
			//值长、值
			content += goutil.StringUtil().IntToStr(len(dv), LEN_VALUE) + dv
		}
	}
	return content
}

//转string
func toString(v interface{}, dataType string) (string, string) {
	response := ""
	dt := "1"
	switch dataType {
	case DATA_TYPE_STRING:
		response = v.(string)
		dt = GDB_TYPE_STRING
		break
	case DATA_TYPE_BOOL:
		response = strconv.FormatBool(v.(bool))
		dt = GDB_TYPE_BOOL
		break
	case DATA_TYPE_NUMBER:
		response = strconv.FormatFloat(v.(float64), 'E', -1, 64)
		dt = GDB_TYPE_NUMBER
		break
	case DATA_TYPE_MAP:

		dt = GDB_TYPE_MAP
		break
	case DATA_TYPE_SET:

		dt = GDB_TYPE_SET
		break
	case DATA_TYPE_LIST:

		dt = GDB_TYPE_LIST
		break
	case DATA_TYPE_ZSET:

		dt = GDB_TYPE_ZSET
		break
	case DATA_TYPE_OBJECT:

		dt = GDB_TYPE_OBJECT
		break
	default:
		log.Error("类型转换异常")
		break
	}
	return response, dt
}
