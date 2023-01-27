package urlmodel

type URL struct{
	ID int64
	URL string
	UserID int64
	SuccessCount int64
	FailCount int64
}

type URLStat struct{
	SuccessCount int64
	FailCount int64
}