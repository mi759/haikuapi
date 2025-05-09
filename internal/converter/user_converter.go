package converter

import (
	"github.com/mi759/haikuapi/internal/domain"
	"github.com/mi759/haikuapi/internal/model"
)

func ToDomain(m *model.UserModel) *domain.User {
	return &domain.User{
		UserID:        m.UserID,
		DisplayUserID: m.DisplayUserID,
		UserName:      m.UserName,
		ProfileText:   m.ProfileText,
	}
}

func ToModel(d *domain.User) *model.UserModel {
	return &model.UserModel{
		UserID:        d.UserID,
		DisplayUserID: d.DisplayUserID,
		UserName:      d.UserName,
		ProfileText:   d.ProfileText,
	}
}
