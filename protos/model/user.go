package proto

type UserDetailTable struct {
	Mid    int64 `json:"mid" gorm:"column:mid;type:bigint(20);not null;default:0"`
	Name   int   `json:"name" gorm:"column:name;type:varchar(100);not null"`
	Sex    int   `json:"sex" gorm:"column:sex;type:tinyint(1);not null;default:0"`
	Status int64 `json:"status" gorm:"column:status;type:bigint(20);not null;default:0"`
	Ct     int64 `json:"ct" gorm:"column:ct;type:bigint(20);not null;default:0"`
	Ut     int64 `json:"ut" gorm:"column:ut;type:bigint(20);not null;default:0"`
}

func (model *UserDetailTable) TableName() string {
	return "user_detail"
}
