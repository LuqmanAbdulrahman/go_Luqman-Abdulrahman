package main

type User struct { // menggunakan huruf kapital pada nama struct
	ID       int
	Username string // menggunakan tipe data string untuk username
	Password string // menggunakan tipe data string untuk password
}

type UserService struct { // memberi nama variabel yang lebih deskriptif
	Users []User // memberi nama yang lebih deskriptif pada variabel
}

func (u UserService) GetAllUsers() []User { // memberikan komentar untuk menjelaskan fungsi
	// Fungsi GetAllUsers mengembalikan semua pengguna yang tersedia
	return u.Users
}

func (u UserService) GetUserByID(id int) User { // memberikan komentar untuk menjelaskan fungsi
	// Fungsi GetUserByID mencari pengguna berdasarkan ID
	for _, user := range u.Users { // memberi nama variabel yang lebih jelas
		if id == user.ID {
			return user
		}
	}

	return User{}
}
