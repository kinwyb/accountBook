package log

type Tag string

const (
	ServiceTag Tag = "service.log" //服务日志标记
	DBTag      Tag = "db.log"      //数据库日志标记
	PowerTag   Tag = "power.log"   //权限校验
)

func (t Tag) Value() string {
	return string(t)
}
