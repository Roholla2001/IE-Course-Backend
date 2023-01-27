package urlmodel

type URLModel struct {
	ID           int64  `gorm:"primaryKey" json:"id"`
	URL          string `json:"url"`
	UserID       int64  `json:"user_id"`
	SuccessCount int64  `json:"success_count"`
	FailCount    int64  `json:"fail_count"`
}

type URLStat struct {
	SuccessCount int64 `json:"success_count"`
	FailCount    int64 `json:"fail_count"`
}
