package domain

type IUser interface {
	SaveUser(userName string, email string, password string, esp32ID string) error
	DeleteUser(id int32) error
	UpdateUser(id int32, userName string, email string, password string, esp32ID string) error
	GetAll() ([]User, error)
	GetUserByCredentials(userName string) (*User, error)
	GetUserByESP32ID(esp32ID string) (*User, error) // Nuevo método
}

type User struct {
	ID       int32  `json:"id"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	ESP32ID  string `json:"esp32_id"`
}

func NewUser(userName string, email string, password string, esp32ID string) *User {
	return &User{
		UserName: userName,
		Email:    email,
		Password: password,
		ESP32ID:  esp32ID,
	}
}

func (u *User) SetUserName(userName string) {
	u.UserName = userName
}

func (u *User) SetESP32ID(esp32ID string) {
	u.ESP32ID = esp32ID
}
