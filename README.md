# 📚 GoBookStore API

A **production-ready RESTful API** built using **Golang**, **GORM**, **PostgreSQL**, and **JWT authentication**, supporting role-based access and full CRUD for book management.

---

## ✨ Features

- 🔐 JWT-based user authentication
- 👤 Role-based access control (Admin/User)
- 📚 Full CRUD operations for books
- 🐳 Dockerized for easy deployment
- 🌐 Hosted API with live demo

---

## 🛠️ Tech Stack

| Tool         | Description               |
|--------------|---------------------------|
| ⚙️ Go (Golang)    | Backend programming language |
| 🐘 PostgreSQL | Relational Database       |
| 🧰 GORM        | ORM for Golang           |
| 🧭 Gorilla Mux | HTTP request routing     |
| 🐳 Docker      | Containerized deployment |

---

## 🚀 Live Deployment

🔗 https://bookstore-iub1.onrender.com/

---

📸 API Output Screenshots
![register_service](https://github.com/user-attachments/assets/241e44a7-0f7d-497a-906e-369b2bcd9b20)
![login](https://github.com/user-attachments/assets/9ee75469-e109-442b-b28f-546acea87a80)
![POST_BOOKS](https://github.com/user-attachments/assets/7dde3bbb-2628-45dd-a9b6-13a8180f95cc)
![GET BOOK(ID)](https://github.com/user-attachments/assets/12b002f7-9216-4ff0-b2d7-b1d45ef536af)
![DELETE BOOK](https://github.com/user-attachments/assets/ece42753-7a33-4a64-8a9c-e49471cc7067)
![POST_BOOKS](https://github.com/user-attachments/assets/5d2bcc70-f663-4876-8c1c-997ac4e59d3b)
![PUT BOOK](https://github.com/user-attachments/assets/d3be00d3-be9b-4d19-bb42-c68da457e6fc)
![headers](https://github.com/user-attachments/assets/b84316c5-8682-4116-96e0-682e23185f3c)

🧪 API Endpoints Overview
| Method | Endpoint      | Access        | Description       |
| ------ | ------------- | ------------- | ----------------- |
| POST   | `/register`   | Public        | Register new user |
| POST   | `/login`      | Public        | Login & get JWT   |
| GET    | `/books`      | Authenticated | Get list of books |
| POST   | `/books`      | Admin Only    | Add a new book    |
| PUT    | `/books/{id}` | Admin Only    | Update book info  |
| DELETE | `/books/{id}` | Admin Only    | Delete a book     |

👨‍💻 Contributing
Feel free to fork, improve, and submit a pull request! Contributions are welcome. 🙌


