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

func GetSessionList(sessionList *[]Session, filter Session) error {
	sessionResult := db.Order("start").Order("name").Find(sessionList)

	if filter.Name != "" {
		sessionResult = sessionResult.Where("name LIKE ?", "%"+filter.Name+"%")
	}

	if !filter.Start.IsZero() {
		sessionResult = sessionResult.Where("end > ?", filter.Start)
	}

	if !filter.End.IsZero() {
		sessionResult = sessionResult.Where("start < ?", filter.End)
	}

	result := sessionResult.Find(sessionList)

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

func UpdateSession(session *Session) error {
	result := db.Save(session)
	return result.Error
}
