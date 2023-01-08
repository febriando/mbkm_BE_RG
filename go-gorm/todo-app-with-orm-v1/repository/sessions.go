package repository

import (
	"a21hc3NpZ25tZW50/model"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type SessionsRepository struct {
	db *gorm.DB
}

func NewSessionsRepository(db *gorm.DB) SessionsRepository {
	return SessionsRepository{db}
}

func (u *SessionsRepository) AddSessions(session model.Session) error {
	if err := u.db.Create(&session).Error; err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (u *SessionsRepository) DeleteSession(token string) error {
	// if err := u.db.Delete(&token).Error; err != nil {
	// 	return err
	// }
	// session := model.Session{}
	// u.db.Transaction(func(tx *gorm.DB) error {
	// 	tx.Model(&session{}).Where("token = ?", token)

	// 	tx.Transaction(func(tx2 *gorm.DB) error {
	// 		tx2.Delete(&session{})
	// 		return errors.New("rollback token") // Rollback token
	// 	})
	// 	return nil
	// })

	session := model.Session{}
	if err := u.db.Where("token = ?", token).Delete(&session).Error; err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (u *SessionsRepository) UpdateSessions(session model.Session) error {
	if err := u.db.Table("sessions").Where("username = ?", session.Username).Updates(session).Error; err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (u *SessionsRepository) SessionAvailName(name string) (model.Session, error) {
	session := model.Session{}

	if err := u.db.Where("username = ?", name).First(&session).Error; err != nil {
		return model.Session{}, err
	}

	return session, nil // TODO: replace this
}

func (u *SessionsRepository) SessionAvailToken(token string) (model.Session, error) {
	session := model.Session{}

	if err := u.db.Where("token = ?", token).First(&session).Error; err != nil {
		return model.Session{}, err
	}

	return session, nil // TODO: replace this
}

func (u *SessionsRepository) TokenValidity(token string) (model.Session, error) {
	// session := model.Session{} // TODO: replace this
	session, err := u.SessionAvailToken(token)
	if err != nil {
		return model.Session{}, err
	}

	if u.TokenExpired(session) {
		err := u.DeleteSession(token)
		if err != nil {
			return model.Session{}, err
		}
		return model.Session{}, fmt.Errorf("Token is Expired!")
	}

	return session, nil
}

func (u *SessionsRepository) TokenExpired(session model.Session) bool {
	return session.Expiry.Before(time.Now())
}
