package shared_types

const (
	WORK_TYPE_CREATE_GIF = "CREATE_GIF"
)

// 	// Inputs

type CreateGifMessage struct {
	MessagingBase
	EventID  uint
	UserName string
}

func NewCreateGifMessage(eventId uint, userName string) *CreateGifMessage {
	return &CreateGifMessage{
		MessagingBase: MessagingBase{
			WorkType: WORK_TYPE_CREATE_GIF,
		},
		EventID:  eventId,
		UserName: userName,
	}
}
