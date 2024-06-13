package data

import model "User_Product/models"

func InsertData() ([]model.User, []model.Product) {
	productInfo := []model.Product{
		{
			Name:        "Laptop Stand",
			Description: "Adjustable laptop stand for ergonomic comfort.",
			Price:       49.99,
			Quantity:    45,
		},

		{
			Name:        "Noise-Cancelling Headphones",
			Description: "Block out distractions with these high-quality noise-cancelling headphones.",
			Price:       149.99,
			Quantity:    60,
		},

		{
			Name:        "4K Monitor",
			Description: "Ultra HD 4K monitor for stunning picture quality.",
			Price:       299.99,
			Quantity:    20,
		},

		{
			Name:        "Mechanical Keyboard",
			Description: "Responsive and durable mechanical keyboard for an enhanced typing experience.",
			Price:       89.99,
			Quantity:    35,
		},

		{
			Name:        "Ergonomic Chair",
			Description: "Comfortable ergonomic chair to support your posture.",
			Price:       199.99,
			Quantity:    10,
		},

		{
			Name:        "USB-C Hub",
			Description: "Expand your laptop’s connectivity with this multi-port USB-C hub.",
			Price:       39.99,
			Quantity:    75,
		},

		{
			Name:        "External Hard Drive",
			Description: "1TB external hard drive for reliable data storage.",
			Price:       79.99,
			Quantity:    50,
		},

		{
			Name:        "Webcam",
			Description: "High-definition webcam for clear video calls.",
			Price:       59.99,
			Quantity:    40,
		},

		{
			Name:        "Smart Home Assistant",
			Description: "Voice-controlled smart home assistant to simplify your life.",
			Price:       99.99,
			Quantity:    25,
		},

		{
			Name:        "Electric Kettle",
			Description: "Quickly boil water with this efficient electric kettle.",
			Price:       29.99,
			Quantity:    30,
		}}

	userInfo := []model.User{
		{Username: "AliceWang", Email: "alicewang@example.com", Password: "AW@9876"},
		{Username: "BobLee", Email: "boblee@example.com", Password: "BL@5432"},
		{Username: "CharlieKim", Email: "charliekim@example.com", Password: "CK@mnop"},
		{Username: "DianaClark", Email: "dianaclark@example.com", Password: "DC@qrst"},
		{Username: "EthanHarris", Email: "ethanharris@example.com", Password: "EH@uvwx"},
		{Username: "FionaScott", Email: "fionascott@example.com", Password: "FS@yz12"},
		{Username: "GeorgeBrown", Email: "georgebrown@example.com", Password: "GB@abcd"},
		{Username: "HannahDavis", Email: "hannahdavis@example.com", Password: "HD@efgh"},
		{Username: "IanMiller", Email: "ianmiller@example.com", Password: "IM@3456"},
		{Username: "JuliaAnderson", Email: "juliaanderson@example.com", Password: "JA@7890"}}
	return userInfo, productInfo
}
