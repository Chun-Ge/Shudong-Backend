package entity

// NameLib represents names being shown.
type NameLib struct {
	ID      int64
	Name    string `xorm:"notnull"`
	TopicID int64  `xorm:"notnull"`
}
