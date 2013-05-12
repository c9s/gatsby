package gatsby

import "time"

type Staff struct {
	Id        int64      `json:"id" field:",primary,serial"`
	Name      string     `json:"name" field:",required"`
	Gender    string     `json:"gender"`
	StaffType string     `json:"staff_type"` // valid types: doctor, nurse, ...etc
	Phone     string     `json:"phone"`
	Birthday  *time.Time `json:"birthday" field:"birthday,date"`
	CreatedOn *time.Time `json:"created_on" field:"created_on"`
	BaseRecord
}

// Implement the GetPkId interface
func (self *Staff) GetPkId() int64 {
	return self.Id
}

func (self *Staff) SetPkId(id int64) {
	self.Id = id
}
