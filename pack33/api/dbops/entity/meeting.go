package entity

type Meeting struct {
	Id               int    `json:"id"`
	Title            string `json:"title"`
	EnTitle          string `json:"en_title"`
	Indexed          int    `json:"indexed"`
	StartDate        string `json:"start_date"`
	EndDate          string `json:"end_date"`
	ImportantDate    string `json:"important_date"`
	RegistrationDate string `json:"registration_date"`
	Contact          string `json:"contact"`
	Organization     string `json:"organization"`
	Address          string `json:"address"`
	City             string `json:"city"`
	CityCn           string `json:"city_cn"`
	Country          string `json:"country"`
	CountyCn         string `json:"county_cn"`
	Content          string `json:"content"`
	Status           int    `json:"status"`
	ItemType         int    `json:"itemtype"`
	Topic            string `json:"topic"`
	CnTopic          string `json:"cn_topic"`
	Publication      string `json:"publication"`
	Link             string `json:"link"`
	MeetingId        int    `json:"meeting_id"`
	OwnerId          int    `json:"owner_id"`
	ItemStatus       int    `json:"item_status"`
	IsClaimed        int    `json:"is_claimed"`
	PubStatus        int    `json:"pub_status"`
	AuditType        int    `json:"audit_type"`
	MajorCode        string `json:"major_code"`
	MajorName        string `json:"major_name"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}
