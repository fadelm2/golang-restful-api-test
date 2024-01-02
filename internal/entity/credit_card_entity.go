package entity

type Creditcard struct {
	ID        string `gorm:"column:id;primary_key"`
	Type      string `gorm:"column:type"`
	Name      string `gorm:"column:name"`
	Number    string `gorm:"column:number"`
	Expired   string `gorm:"column:expired"`
	Cvv       string `gorm:"column:cvv"`
	UserId    string `gorm:"column:user_id"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64  `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	User      User   `gorm:"foreignKey:user_id;references:id"`
}

func (c *Creditcard) TableName() string {
	return "creditcards"
}
