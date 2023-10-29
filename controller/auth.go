package controller

import (
	"time"
	"user-personalize/database"
	"user-personalize/helper"
	"user-personalize/models"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func init() {
  	govalidator.SetFieldsRequiredByDefault(true)
}

func Register (r *gin.Context) {
	type User struct{
		User_username string `valid:"required"`
		User_password string `valid:"required,stringlength(6|16)"`
		User_email string `valid:"required,email"`
	}

	type UserForm struct{
		User_username string `form:"user_username"`
		User_password string `form:"user_password"`
		User_email string `form:"user_email"`
	}
	var user_get UserForm
	if r.ShouldBind(&user_get) == nil{
		user_data := &User{
			User_username: user_get.User_username,
			User_password: user_get.User_password,
			User_email: user_get.User_email,
		}
		result, err := govalidator.ValidateStruct(user_data)
		if err != nil{
			r.JSON(400, gin.H{
				"Error Validate": err.Error(),
			})
			panic(err.Error())
		}
		if result {
			enc, err := helper.PasswordHash(user_data.User_password)
			if err != nil{
				r.JSON(400, gin.H{
					"msg": "Enkripsi gagal dilakukan",
				})
			} else {
				db := database.DB
				user := models.Users{
					ID: 0,
					Username: user_data.User_username,
					Email:  user_data.User_email,
					Password: enc,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}
				insert := db.Create(&user)
				if insert.Error != nil{
					r.JSON(400, gin.H{
						"msg": "Gagal memasukan data (" + insert.Error.Error() + ")",
					})
				}
			}
		} else {
			r.JSON(400, gin.H{
				"msg": "Validasi gagal",
			})
		} 
	} else {
		r.JSON(400, gin.H{
			"msg": "Gagal mendapatkan data",
		})
	}
	
}

func Login (r *gin.Context){
	type UsersStruct struct{
		Username string `form:"user_username" valid:"required"`
		Password string `form:"user_password" valid:"required"`
	}
	var userData UsersStruct
	if r.ShouldBind(&userData) == nil{
		user := &UsersStruct{
			Username: userData.Username,
			Password: userData.Password,
		}
		result, err := govalidator.ValidateStruct(user)
		if err != nil{
			r.JSON(400, gin.H{
				"Error Validate": err.Error(),
			})
			panic(err.Error())
		}
		if result{
			if err != nil{
				r.JSON(400, gin.H{
					"Error Password": err.Error(),
				})
				panic(err.Error())
			} else {
				var db = database.DB
				type Users struct{
					ID int
					Username string
					Password string
				}
				var User = Users{}
				row := db.Where(&Users{Username: user.Username}).Find(&User)
				if row.RowsAffected == 1{
					if helper.PasswordCheck(userData.Password, User.Password){
						type Logins struct{
							Login_user_id int
							Login_token string
						}
						var login Logins
						logUser := db.Where(&Logins{Login_user_id: User.ID}).Find(&login)
						if logUser.RowsAffected > 0 {
							db.Delete(&Logins{}, 10)
						}
						tokens, err := helper.SetToken()
						if err != nil {
							r.JSON(400, gin.H{
								"msg": "Token failed",
							})
							panic(err.Error())
						} else {
							newLog := models.Login{
								Login_id: 0,
								Login_user_id: User.ID,
								Login_token: tokens,
								Login_created_at: time.Now(),
							}
							insertLog := db.Create(&newLog)
							if insertLog.Error != nil {
								r.JSON(400, gin.H{
									"msg": "Gagal menyimpan login (" + insertLog.Error.Error() + ")",
								})
							} else {
								r.JSON(200, gin.H{
									"msg": "Anda berhasil login dengan token (" + tokens + ")",
								})
							}
						}
						

					} else {
						r.JSON(200, gin.H{
							"msg": "Username atau Password Salah",
						})
					}
					
				} else  {
					r.JSON(400, gin.H{
						"msg": "Username atau password salah",
					})
				}
			}
			
		}
	}
}