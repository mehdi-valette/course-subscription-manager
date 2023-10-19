package repository

type Session struct {
	ID   uint32 `gorm:"primaryKey"`
	Name string
}

func AddSession(session *Session) error {
	result := db.Create(session)

	return result.Error
}

func GetSessionList(session *[]Session) error {
	result := db.Order("name").Find(session)

	return result.Error
}
