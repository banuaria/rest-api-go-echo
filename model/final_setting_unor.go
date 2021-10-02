package model

type Setting_Unor struct {
	Id                            string `json:"id" form:"id"`
	Instansi_id                   string `json:"instansi_id" form:"instansi_id"`
	Instansi_nama                 string `json:"instansi_nama" form:"instansi_nama"`
	Step_unor_verifikasi_approval int    `json:"step_unor_verifikasi_approval" form:"step_unor_verifikasi_approval"`
	Step_relasi_unor              int    `json:"step_relasi_unor" form:"step_relasi_unor"`
	Status_final_setting          int    `json:"status_final_setting" form:"status_final_setting"`
}

func (Setting_Unor) TableName() string {
	return "final_setting_unor"
}
