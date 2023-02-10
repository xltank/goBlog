package model

import "time"

type Model struct {
	ID        uint64    `json:"id,omitempty"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updatedAt"`
}

type User struct {
	Model
	Email      string `json:"email" validate:"required,email"`
	Mobile     string `json:"mobile,omitempty"`
	Password   string `json:"password,omitempty" validate:"required"`
	Role       string `json:"role,omitempty"`
	Status     string `json:"status,omitempty"`
	Name       string `json:"name,omitempty"`
	NickName   string `json:"nickName,omitempty" validate:"required"`
	Department string `json:"department,omitempty"`
	Birthday   string `json:"birthday,omitempty"`
	Gender     string `json:"gender,omitempty"`
	// Province    string `json:"province,omitempty"`
	// City        string `json:"city,omitempty"`
	// Address     string `json:"address,omitempty"`
	// AddressName string `json:"address_name,omitempty"`
}

type Blog struct {
	Model
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
	Author  string `json:"author,omitempty"`
	Status  string `json:"status,omitempty"`
}

type Comment struct {
	Model
	Content string `json:"content,omitempty"`
	Author  string `json:"author,omitempty"`
	BlogID  uint64 `json:"blogId,omitempty"`
	Status  string `json:"status,omitempty"`
}
