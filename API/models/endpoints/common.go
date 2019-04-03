package endpoints

import (
	"accountBook/models/beans"
	"math"

	"github.com/kinwyb/go/err1"
)

//检测权限
//@param string fun 函数点
func CheckPower(fun string, ctx *beans.Context) err1.Error {
	// todo 检测权限
	return nil
}

// pws所有全值集合
// chkValue当前要校验的权限值
func powerVerification(pws []uint, chkValue uint) bool {
	index := int(math.Ceil(float64(chkValue) / 64)) //获取到权限在第几个值 == 1 => 2 , 0 => 1
	if len(pws) <= int(index) {
		//找不到相关权限对应的值
		return false
	}
	chkValue = chkValue % 64
	// 1左移d位表示权限对应的点位设置成1 .e.g  1<<3 => 00...00100
	// &将值进行确定权限power对应值的相应位置是否也是1来确定是否有权限
	return 1<<uint(chkValue)&pws[index] != 0
}
