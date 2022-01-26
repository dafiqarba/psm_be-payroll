package entity

type LeaveBalance struct {
	Balance_id   string `json:"balance_id"`
	Balance_year string `json:"balance_year"`
	Cuti_tahunan string `json:"cuti_tahunan"`
	Cuti_diambil string `json:"cuti_diambil"`
	Cuti_balance string `json:"cuti_balance"`
	Cuti_izin    string `json:"cuti_izin"`
	Cuti_sakit   string `json:"cuti_sakit"`
	User_id      string `json:"user_id"`
}
