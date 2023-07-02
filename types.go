package gology

type NotificationResponse struct {
	Data []NotificationType `json:"data"`
}

type NotificationType struct {
	Type     string                `json:"type"`
	Sentence string                `json:"sentence"`
	Viewed   bool                  `json:"viewed"`
	Created  string                `json:"created"`
	More     string                `json:"more"`
	Realm    string                `json:"realm"`
	RealmId  string                `json:"realmId"`
	Args     []NotificationContent `json:"args"`
}

type NotificationContent struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Type  string `json:"type"`
}
