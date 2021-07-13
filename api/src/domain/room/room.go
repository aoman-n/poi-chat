package room

import (
	"time"

	val "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/laster18/poi/api/src/domain"
	"github.com/laster18/poi/api/src/util/clock"
	"github.com/laster18/poi/api/src/util/validator"
)

type Room struct {
	ID              int
	OwnerID         int
	Name            string
	BackgroundURL   string
	BackgroundColor string
	CreatedAt       time.Time
}

const (
	DefaultBgColor = "#20b2aa"
	DefaultBgURL   = "https://poi-chat.s3.ap-northeast-1.amazonaws.com/roomBg1.jpg"
)

func New(uid int, name string, bgColor, bgURL *string) *Room {
	color := DefaultBgColor
	if bgColor != nil {
		color = *bgColor
	}

	url := DefaultBgURL
	if bgURL != nil {
		url = *bgURL
	}

	return &Room{
		OwnerID:         uid,
		Name:            name,
		BackgroundURL:   url,
		BackgroundColor: color,
		CreatedAt:       clock.Now(),
	}
}

var _ domain.INode = (*Room)(nil)

func (r *Room) GetID() int {
	return r.ID
}

func (r *Room) GetCreatedAtUnix() int {
	return int(r.CreatedAt.Unix())
}

func (r *Room) Validate() error {
	return validator.ValidateStruct(r,
		val.Field(&r.Name, val.Required, val.RuneLength(2, 20)),
		val.Field(&r.BackgroundColor, val.Required, is.HexColor),
		val.Field(&r.BackgroundURL, val.Required, is.URL),
	)
}
