package room_test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"noname-one-time-session-chat/internal/chat/room"
	"noname-one-time-session-chat/pkg/errors"
	"testing"
)

func TestNewRoom(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	roomName := "roomName"

	tests := []struct {
		name     string
		roomName string
		expect   func(*testing.T, *room.Room, error)
	}{
		{
			name:     "should return service",
			roomName: roomName,
			expect: func(t *testing.T, s *room.Room, err error) {
				assert.NotNil(t, s)
				assert.Nil(t, err)
			},
		},
		{
			name:     "should return service",
			roomName: "",
			expect: func(t *testing.T, s *room.Room, err error) {
				assert.Nil(t, s)
				assert.NotNil(t, err)
				assert.EqualError(t, err, errors.WithMessage(room.ErrInvalidName, "should be not empty").Error())
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			svc, err := room.NewRoom(tc.roomName)
			tc.expect(t, svc, err)
		})
	}
}
