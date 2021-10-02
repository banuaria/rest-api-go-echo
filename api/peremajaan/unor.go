package peremajaan

import (
	"code-echo/db"
	"code-echo/model"
	"net/http"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

func SimpanUnor(c echo.Context) error {
	//Inisiasi Model Baru dari Struct setting Unor
	ModelSettingUnor := new(model.Setting_Unor)

	//Memperoleh isi input / parameter yang di kirim dari client
	if err := c.Bind(ModelSettingUnor); err != nil {
		return err
	}

	//Inisiasi Koneksi Database
	DbConn := db.Manager()

	//Generate UUID Untuk Id record yang baru
	u1 := uuid.Must(uuid.NewV4())
	Id := u1.String()

	//Memberi Nilai Id yang ada dalam ModelSettingUnor
	ModelSettingUnor.Id = Id

	// Insert Di tabel usulan
	dbc := DbConn.Debug().Create(&ModelSettingUnor)

	//Kondisi Jika Terjadi error saat eksekusi query
	if dbc.Error != nil {
		return c.JSON(http.StatusNotAcceptable, map[string]string{
			"error":   "true",
			"message": dbc.Error.Error(),
		})
	}

	// Mereturn Response dalam format JSON
	return c.JSON(http.StatusOK, map[string]string{
		"error":   "false",
		"message": "Berhasil Disimpan",
		"Id":      Id,
	})
}

func UpdateUnor(c echo.Context) error {
	//Inisiasi Model Baru dari Struct Even Instansi
	ModelSettingUnor := new(model.Setting_Unor)

	//Memperoleh isi input / parameter yang di kirim dari client
	if err := c.Bind(ModelSettingUnor); err != nil {
		return err
	}

	//Inisiasi Koneksi Database
	DbConn := db.Manager()
	// Insert Di tabel usulan
	dbc := DbConn.Debug().Model(&ModelSettingUnor).Where("id = ?", ModelSettingUnor.Id).Update(ModelSettingUnor)

	//Kondisi Jika Terjadi error saat eksekusi query
	if dbc.Error != nil {
		return c.JSON(http.StatusNotAcceptable, map[string]string{
			"error":   "true",
			"message": dbc.Error.Error(),
		})
	}

	// Mereturn Response dalam format JSON
	return c.JSON(http.StatusOK, map[string]string{
		"error":   "false",
		"message": "Berhasil Diupdate",
	})

}

func GetUnor(c echo.Context) error {

	DbCon := db.Manager()
	Unor := []model.Setting_Unor{}
	DbCon.Find(&Unor)

	return c.JSON(http.StatusOK, Unor)
}

func DeleteUnor(c echo.Context) error {

	//Inisiasi Koneksi Database
	DbConn := db.Manager()
	ModelSettingUnor := model.Setting_Unor{}

	dbc := DbConn.Debug().Model(&ModelSettingUnor).Where("id = ?", ModelSettingUnor.Id).Delete(&ModelSettingUnor)

	//Kondisi Jika Terjadi error saat eksekusi query
	if dbc.Error != nil {
		return c.JSON(http.StatusNotAcceptable, map[string]string{
			"error":   "true",
			"message": dbc.Error.Error(),
		})
	}

	// Mereturn Response dalam format JSON
	return c.JSON(http.StatusOK, map[string]string{
		"error":   "false",
		"message": "Berhasil dihapus",
	})

}
