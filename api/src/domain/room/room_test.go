package room_test

import (
	"strings"
	"testing"

	"github.com/laster18/poi/api/src/domain/room"
	"github.com/laster18/poi/api/src/testutil"
)

func TestRoom_Validate(t *testing.T) {
	t.Run("#Name", func(t *testing.T) {
		cases := map[string]struct {
			room      *room.Room
			fieldName string
			expectMsg string
		}{
			"[異常系] 空文字だとエラーになること": {
				room: &room.Room{
					Name: "",
				},
				fieldName: "Name",
				expectMsg: "cannot be blank",
			},
			"[異常系] 1文字だとエラーになること": {
				room: &room.Room{
					Name: strings.Repeat("あ", 1),
				},
				fieldName: "Name",
				expectMsg: "the length must be between 2 and 20",
			},
			"[正常系] 2文字だとエラーにならないこと": {
				room: &room.Room{
					Name: strings.Repeat("あ", 2),
				},
				fieldName: "Name",
				expectMsg: "",
			},
			"[正常系] 20文字だとエラーにならないこと": {
				room: &room.Room{
					Name: strings.Repeat("あ", 20),
				},
				fieldName: "Name",
				expectMsg: "",
			},
			"[異常系] 21文字だとエラーになること": {
				room: &room.Room{
					Name: strings.Repeat("あ", 21),
				},
				fieldName: "Name",
				expectMsg: "the length must be between 2 and 20",
			},
		}

		for k, tt := range cases {
			tt := tt
			t.Run(k, func(t *testing.T) {
				t.Parallel()

				err := tt.room.Validate()

				if tt.expectMsg == "" {
					testutil.AssertNoValidationErr(t, err, tt.fieldName)
				} else {
					testutil.AssertValidationErr(t, err, tt.fieldName, tt.expectMsg)
				}
			})
		}
	})

	t.Run("#BackgroundColor", func(t *testing.T) {
		cases := map[string]struct {
			room      *room.Room
			fieldName string
			expectMsg string
		}{
			"[異常系] 空文字だとエラーになること": {
				room: &room.Room{
					BackgroundColor: "",
				},
				fieldName: "BackgroundColor",
				expectMsg: "cannot be blank",
			},
			"[正常系] 有効な16進カラーコードだとエラーにならないこと": {
				room: &room.Room{
					BackgroundColor: "#ffffff",
				},
				fieldName: "BackgroundColor",
				expectMsg: "",
			},
			"[異常系] 無効な16進カラーコードだとエラーにならないこと": {
				room: &room.Room{
					BackgroundColor: "#ffffffff",
				},
				fieldName: "BackgroundColor",
				expectMsg: "must be a valid hexadecimal color code",
			},
		}

		for k, tt := range cases {
			tt := tt
			t.Run(k, func(t *testing.T) {
				t.Parallel()

				err := tt.room.Validate()

				if tt.expectMsg == "" {
					testutil.AssertNoValidationErr(t, err, tt.fieldName)
				} else {
					testutil.AssertValidationErr(t, err, tt.fieldName, tt.expectMsg)
				}
			})
		}
	})

	t.Run("#BackgroundURL", func(t *testing.T) {
		cases := map[string]struct {
			room      *room.Room
			fieldName string
			expectMsg string
		}{
			"[異常系] 空文字だとエラーになること": {
				room: &room.Room{
					BackgroundURL: "",
				},
				fieldName: "BackgroundURL",
				expectMsg: "cannot be blank",
			},
			"[異常系] 不正なURLだとエラーになること": {
				room: &room.Room{
					BackgroundURL: "https?://hoge.com",
				},
				fieldName: "BackgroundURL",
				expectMsg: "must be a valid URL",
			},
			"[正常系] 有効なURLだとエラーにならないこと": {
				room: &room.Room{
					BackgroundURL: "https://example.com/bg.png",
				},
				fieldName: "BackgroundURL",
				expectMsg: "",
			},
		}

		for k, tt := range cases {
			tt := tt
			t.Run(k, func(t *testing.T) {
				t.Parallel()

				err := tt.room.Validate()

				if tt.expectMsg == "" {
					testutil.AssertNoValidationErr(t, err, tt.fieldName)
				} else {
					testutil.AssertValidationErr(t, err, tt.fieldName, tt.expectMsg)
				}
			})
		}
	})
}
