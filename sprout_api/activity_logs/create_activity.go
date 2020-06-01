package activity_logs

type Activity struct {
	ActivityUnitID int64  `json:"activityUnitId"`
	Date           string `json:"date"`
	Quantity       string `json:"quantity"`
}
type CreateActivityRequest struct {
	Activities        []Activity `json:"activities"`
	Image             string     `json:"image"`
	ShareWithFilterID string     `json:"shareWithFilterId"`
	Text              string     `json:"text"`
}

type CreateActivityResponse struct {
	ActivityLogItemIds []struct {
		ActivityLogID int64 `json:"activityLogId"`
	} `json:"activityLogItemIds"`
	StreamItemID int64 `json:"streamItemId"`
}
