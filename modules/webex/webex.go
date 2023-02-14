package webex

import (
	"fmt"

	webexteams "github.com/jbogarin/go-cisco-webex-teams/sdk"
)

type Webex struct {
	Token  string `json:"token" yaml:"token"`
	client *webexteams.Client
}

func (m *Webex) SendMessage(msg string, markdown bool) error {
	m.client = webexteams.NewClient()
	m.client.SetAuthToken(m.Token)
	//list the rooms
	roomsQueryParams := &webexteams.ListRoomsQueryParams{
		Max:      500,
		TeamID:   "",
		Paginate: false,
	}

	rooms, _, err := m.client.Rooms.ListRooms(roomsQueryParams)
	if err != nil {
		return err
	}
	var message *webexteams.MessageCreateRequest
	for _, room := range rooms.Items {
		if markdown {
			message = &webexteams.MessageCreateRequest{
				Markdown: msg,
				RoomID:   room.ID,
			}
		} else {
			message = &webexteams.MessageCreateRequest{
				Text:   msg,
				RoomID: room.ID,
			}
		}

		newTextMessage, _, err := m.client.Messages.CreateMessage(message)
		if err != nil {
			return err
		}

		fmt.Println("POST:", newTextMessage.ID, newTextMessage.Text, newTextMessage.Created)
	}

	return nil
}

func CreateK8(token string) *Webex {
	return &Webex{
		Token: token,
	}
}
