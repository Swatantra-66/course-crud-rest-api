# 📚 Course API (Go + Gorilla Mux)

A simple RESTful API built with **Go** and **Gorilla Mux** to manage courses.  
This project uses an in-memory slice as a fake database and demonstrates how to build clean, structured APIs in Go.

---

## Tech Stack

- **Language:** Go (Golang)
- **Router:** [gorilla/mux](https://github.com/gorilla/mux)
- **Data Format:** JSON
- **Storage:** In-memory slice (no external DB yet)

---

## Features

- Get all courses
- Get a single course by ID
- Create a new course
- Update an existing course
- Delete a course
- Structured types with nested `Author`
- Proper HTTP status codes and JSON responses

---

## 1. Clone the repository

```bash
git clone https://github.com/<your-username>/go-course-api.git
cd go-course-api
```

## 2. Install dependencies

```bash
go mod init github.com/<your-username>/go-course-api
go get github.com/gorilla/mux
```

## 3. Run the Server

```bash
go run main.go
```

### Server will start on:

```text
http://localhost:4000
```

---

## >> API Endpoints : GET/courses

### Response:

```json
[
  {
    "courseid": "2",
    "coursename": "Backend Developer Course",
    "price": 799,
    "author": {
      "fullname": "Swatantra Yadav",
      "website": "go.dev"
    }
  },
  {
    "courseid": "4",
    "coursename": "ReactJS",
    "price": 599,
    "author": {
      "fullname": "Swatantra Yadav",
      "website": "MenScriptCourses.com"
    }
  },
  {
    "courseid": "6",
    "coursename": "Solana Blockchain Full Course",
    "price": 999,
    "author": {
      "fullname": "Swatantra Yadav",
      "website": "MenScriptCourses.com"
    }
  }
]
```

## >> Get a single course : GET /courses/{id}

```bash
curl http://localhost:4000/courses/2
```

## >> Create a course: POST /courses

```bash
curl -X POST http://localhost:4000/courses \
  -H "Content-Type: application/json" \
  -d '{
        "coursename": "Golang Bootcamp",
        "price": 999,
        "author": {
          "fullname": "Swatantra Yadav",
          "website": "go.dev"
        }
      }'
```

### Response:

```bash
{
  "courseid": "57",
  "coursename": "Golang Bootcamp",
  "price": 999,
  "author": {
    "fullname": "Swatantra Yadav",
    "website": "go.dev"
  }
}
```

## >> Update a course : PUT /courses/{id}

```bash
curl -X PUT http://localhost:4000/courses/2 \
  -H "Content-Type: application/json" \
  -d '{
        "coursename": "Advanced Backend Developer Course",
        "price": 1299,
        "author": {
          "fullname": "Swatantra Yadav",
          "website": "menscriptcourses.com"
        }
      }'
```

## >> Delete a course : DELETE /courses/{id}

```bash
curl -X DELETE http://localhost:4000/courses/4
```

## Notes will be Updated soon

- This API uses an in-memory slice as a fake database.
- Data is reset every time the server restarts.
- IDs are randomly generated and not guaranteed globally unique.
- No authentication/authorization yet,so will be adding logging and auth middlewares.
- Versioning(/api/v1/courses)

---

### Not Done Yet
