package gdb

import (
	"bufio"
	"os"
	"strconv"
	"time"

	"github.com/louch2010/dhaiy/cache"
	. "github.com/louch2010/dhaiy/common"
	"github.com/louch2010/dhaiy/log"
	"github.com/louch2010/goutil"
)

//解析gdb文件
func parseGDB(file *os.File) error {
	bfRd := bufio.NewReader(file)
	//魔数
	content, err := read(bfRd, GDB_LEN_GOCACHE)
	if err != nil {
		return err
	}
	if content != GDB_GOCACHE {
		return GDB_FILE_INVALID
	}
	//版本
	content, _ = read(bfRd, GDB_LEN_VERSION)
	if content != GDB_VERSION {
		return GDB_FILE_VERSION_ERROR
	}
	//库标识
	content, _ = read(bfRd, GDB_LEN_DATABASE)
	if content != GDB_DATABASE {
		log.Error("gdb文件格式错误，需要'database',但值为:", content)
		return GDB_FILE_FORMAT_ERROR
	}
	//库长
	databaseLenStr, _ := read(bfRd, GDB_LEN_DATABASE_SIZE)
	databaseLen, err := goutil.StringUtil().StrToInt(databaseLenStr)
	log.Debug("需要加载的库的数量：", databaseLen)
	//遍历库
	for j := 0; j < databaseLen; j++ {
		//表标识
		tableFlag, _ := read(bfRd, GDB_LEN_TABLE)
		if tableFlag != GDB_TABLE {
			log.Error("gdb文件格式错误，需要'table',但值为:", content)
			return GDB_FILE_FORMAT_ERROR
		}
		//表名
		tableNameLenStr, _ := read(bfRd, GDB_LEN_KEY)
		tableNameLen, _ := goutil.StringUtil().StrToInt(tableNameLenStr)
		tableName, _ := read(bfRd, tableNameLen)
		table, err := cache.Cache(tableName)
		if err != nil {
			log.Error("获取表失败！表名：", tableName, "，错误信息：", err)
			return err
		}
		//键值对数
		keySizeStr, _ := read(bfRd, GDB_LEN_TABLE_SIZE)
		keySize, _ := goutil.StringUtil().StrToInt(keySizeStr)
		for i := 0; i < keySize; i++ {
			//数据类型
			dataType, _ := read(bfRd, GDB_LEN_DATATYPE)
			//过期时间
			var liveTime time.Duration = 0
			expireTimeStr, _ := read(bfRd, GDB_LEN_LIVETIME_ALWAYS)
			expire := false
			if expireTimeStr != GDB_LIVETIME_ALWAYS {
				tmp, _ := read(bfRd, GDB_LEN_LIVETIME-GDB_LEN_LIVETIME_ALWAYS)
				expireTime, err := goutil.DateUtil().ParseTime14(expireTimeStr + tmp)
				if err != nil {
					log.Error("时间转换异常！", err)
					return err
				}
				log.Debug("过期时间：", expireTime, "，当前时间：", time.Now())
				if expireTime.Before(time.Now()) {
					expire = true
				}
				liveTime = expireTime.Sub(time.Now())
				log.Debug("存活时长：", liveTime)
			}
			//键
			keyLenStr, _ := read(bfRd, GDB_LEN_KEY)
			keyLen, _ := goutil.StringUtil().StrToInt(keyLenStr)
			key, _ := read(bfRd, keyLen)
			//值
			valueLenStr, _ := read(bfRd, LEN_VALUE)
			valueLen, _ := goutil.StringUtil().StrToInt(valueLenStr)
			value, _ := read(bfRd, valueLen)
			//过期判断
			if expire {
				continue
			}
			switch dataType {
			case GDB_TYPE_STRING:
				table.Set(key, value, liveTime, DATA_TYPE_STRING)
				break
			case GDB_TYPE_NUMBER:
				v, err := strconv.ParseFloat(value, 64)
				if err != nil {
					log.Error("格式转换异常！", err)
					return err
				}
				table.Set(key, v, liveTime, DATA_TYPE_NUMBER)
				break
			}
		}
	}

	return nil
}

func read(bfRd *bufio.Reader, length int) (string, error) {
	buf := make([]byte, length)
	n, err := bfRd.Read(buf)
	if err != nil {
		return "", err
	}
	return string(buf[:n]), nil
}
