package sqlutils
import "time"

type Staff struct {
	Id        int64 `json:"id" field:",primary,serial"`
	Name      string `json:"name" field:",required"`
	Gender    string `json:"gender"`
	StaffType string `json:"staff_type"` // valid types: doctor, nurse, ...etc
	Phone     string `json:"phone"`
	Birthday  string `json:"birthday" field:"birthday,date"`
	CreatedOn *time.Time `json:"created_on" field:"created_on"`
}

// Implement the GetPkId interface
func (self *Staff) GetPkId() int64 {
	return self.Id
}

func (self *Staff) SetPkId(id int64) {
	self.Id = id
}

