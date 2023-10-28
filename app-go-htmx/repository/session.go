package repository

import "time"

type Session struct {
	ID    uint32 `gorm:"primaryKey"`
	Name  string
	Start time.Time
	End   time.Time
}

func AddSession(session *Session) error {
	result := db.Create(session)

	return result.Error
}

func GetSessionList(session *[]Session) error {
	result := db.Order("name").Find(session)

	return result.Error
}

func GetSession(session *Session) error {
	result := db.First(session, session.ID)

	return result.Error
}

func DeleteSession(session *Session) error {
	result := db.Delete(session)
	return result.Error
}
