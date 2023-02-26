package shared_types

const (
	WORK_TYPE_API_CALLBACK = "CALLBACK_API"
)

type ApiCallback struct {
	MessagingBase
	TaskID uint
	Status string
}

func NewApiCallbackMessage(taskId uint, status string) *ApiCallback {
	return &ApiCallback{
		MessagingBase: MessagingBase{
			WorkType: WORK_TYPE_API_CALLBACK,
		},
		TaskID: taskId,
		Status: status,
	}
}
