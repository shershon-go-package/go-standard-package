/**
 * @Description Gorm使用测试
 **/
package tgorm

import (
	"fmt"
	"shershon1991/go-package/app/gormpkg"
	"testing"
)

// 测试迁移
func TestAutoMigrate(t *testing.T) {
	host := "127.0.0.1"
	use, pass, port, database := "root", "root", "3306", "test"
	err := gormpkg.GormAutoMigrate(host, port, use, pass, database)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("创建表结构完成!")
}
