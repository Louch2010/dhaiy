package cache

import (
	"sync"

	"github.com/louch2010/dhaiy/common"
	"github.com/louch2010/dhaiy/log"
	"github.com/louch2010/goutil"
)

var (
	cache = make(map[string]*common.CacheTable)
	mutex sync.RWMutex
)

//初始化缓存表
func InitCache() error {
	sysTable := common.GetSystemConfig().MustValue("server", "sysTable", "sys")
	defaultTable := common.GetSystemConfig().MustValue("table", "default", "default")
	_, err := Cache(sysTable)
	if err != nil {
		log.Error("初始化系统缓存表失败，表名：", sysTable)
		return err
	}
	_, err = Cache(defaultTable)
	if err != nil {
		log.Error("初始化默认缓存表失败，表名：", defaultTable)
		return err
	}
	return nil
}

//获取系统表
func GetSysTable() (*common.CacheTable, error) {
	sysTable := common.GetSystemConfig().MustValue("server", "sysTable", "sys")
	return Cache(sysTable)
}

//获取默认表
func DefaultCache() *common.CacheTable {
	defaultTable := common.GetSystemConfig().MustValue("table", "default", "default")
	t, _ := Cache(defaultTable)
	return t
}

//获取缓存表，如果不存在，则新建缓存表；存在则直接返回
func Cache(table string) (*common.CacheTable, error) {
	if goutil.StringUtil().IsEmpty(table) {
		return nil, common.ERROR_TABLE_NAME
	}
	mutex.RLock()
	t, ok := cache[table]
	mutex.RUnlock()
	if !ok {
		log.Info("缓存表不存在，新建缓存表，表名为：", table)
		t = common.NewCacheTable(table)
		mutex.Lock()
		cache[table] = t
		mutex.Unlock()
		//事件回调通知
		t.TableEventCallBack(t, common.EVENT_TABLE_ADD)
	}
	return t, nil
}

//获取缓存表
func GetCacheTable(table string) (*common.CacheTable, bool) {
	mutex.RLock()
	t, ok := cache[table]
	mutex.RUnlock()
	return t, ok
}

//获取所有表
func GetCacheTables() map[string]*common.CacheTable {
	mutex.RLock()
	list := cache
	mutex.RUnlock()
	return list
}