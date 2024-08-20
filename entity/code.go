package entity

type CodeSubscription struct {
	ID   uint   `json:"id"`
	Code string `json:"code"`
	// nggak wajib nanti otomatis 30 hari
	TimeActive string `json:"time_active"` // 30 hari default
}

// disini otomatis 30 hari , kapan kapan kalau pengen di custom bisa di tambahin masa aktif kapan

// tablename use pointer
func (*CodeSubscription) TableName() string {
	return "code_subssrciption"
}
