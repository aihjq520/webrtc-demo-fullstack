package vo

type RoomCreateReq struct {
	Name string `json:"name,omitempty" binding:"required"`
}
