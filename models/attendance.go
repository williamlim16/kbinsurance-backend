package models

import "time"

type Attendance struct {
	ID       uint  `json:id`
	ClockIn  int64 `json:clockin;gorm:"autoUpdateTime:milli"`
	ClockOut int64 `json:clockout;gorm:"autoUpdateTime:milli"`
	UserID   uint  `json:userid; form:"userid"`
	User     User  `json:users;gorm:"foreignKey:UserID"`
}

func (a Attendance) getClockIn() time.Time {
	return time.Unix(0, a.ClockIn*int64(time.Millisecond))
}
