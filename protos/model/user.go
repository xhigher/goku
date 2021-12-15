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

type UserAddressTable struct {
	AddrId    int64  `json:"addr_id" gorm:"column:addr_id;type:bigint(20);not null;default:0"`
	Mid       int64  `json:"mid" gorm:"column:mid;type:bigint(20);not null;default:0"`
	Name      string `json:"name" gorm:"column:name;type:varchar(100);not null"`
	Phoneno   string `json:"phoneno" gorm:"column:phoneno;type:varchar(50);not null"`
	Locations string `json:"locations" gorm:"column:locations;type:varchar(200);not null"`
	Ut        int64  `json:"ut" gorm:"column:ut;type:bigint(20);not null;default:0"`
}

func (model *UserAddressTable) TableName() string {
	return "user_address"
}
