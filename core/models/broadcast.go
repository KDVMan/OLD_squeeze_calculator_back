package core_models

import "backend/enums"

type BroadcastChannelModel struct {
	Event enums.WebsocketEvent `json:"event"`
	Data  interface{}          `json:"data"`
}

type ProgressChannelModel struct {
	Count  int64                 `json:"count"`
	Total  int64                 `json:"total"`
	Status enums.WebsocketStatus `json:"status"`
}
