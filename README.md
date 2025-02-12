Simple Forum Backend

📌 Overview

Simple Forum Backend is a RESTful API that provides core forum functionalities such as authentication, posting, commenting, and liking posts. It is built with security and efficiency in mind, utilizing JWT for authentication and pagination for handling large datasets.



🚀 Features

User Authentication & Management

Signup (Register with email & password)

Login (JWT-based authentication)

Refresh Token support for session management

Forum Functionalities

Create, Read, Update posts

Add comments

Like & Unlike posts

API Enhancements

Pagination for posts and comments

Secure authentication using JWT

Refresh token mechanism for extended sessions



🛠 Tech Stack

Backend: Go / Gin

Database: MySQL

Authentication: JWT & Refresh Token



⚙️ Installation & Setup

1. Clone the Repository

git clone https://github.com/your-repo/simple-forum-backend.git
cd simple-forum-backend

2. Install Dependencies

go mod tidy # For Go
npm install # For Node.js

3. Run Database Migrations (if applicable)

make migrate-create name=?

make migrate-up

make migrate-down

4. Start the Server

make run



📡 API Endpoints

🧑‍💻 Authentication

POST /memberships/sign-up - Register a new user

POST /memberships/login - Login and get an access token

POST /memberships/refresh - Refresh expired tokens

📝 Posts

POST /posts/create - Create a new post

GET /posts/?pageIndex=?&pageSize=? - Get paginated list of posts

GET posts/:postID? - Get a single post

💬 Comments

/posts/comment/:postID - Add a comment to a post

❤️ Likes

PUT /posts/user-activity/:postID - Like a post



🔄 Future Improvements

Role-based access control (Admin, User)

File uploads (profile pictures, post images)

WebSocket for real-time updates
