# Dip-Flix

Dip-Flix is a movie discovery and streaming demo app with a React/Vite frontend and a Go/Gin backend. It supports user authentication, movie browsing, recommendations, review workflows, and protected streaming flows.

## Key Features

- User registration and login with JWT authentication
- Secure access to protected routes using middleware
- Movie catalog browsing and details lookup
- Recommended movies for authenticated users
- Review workflow for updating movie reviews
- Simple stream page for video playback
- MongoDB-backed data storage

## Tech Stack

- Frontend: React 19, Vite, React Router, Bootstrap
- Backend: Go, Gin, MongoDB, JWT, validator
- Database: MongoDB

## Repository Structure

- `client/dip-flix-client/` — React frontend
- `server/dipFlixServer/` — Go backend API

## Getting Started

### Prerequisites

- Node.js (for frontend)
- npm or yarn
- Go 1.25+
- MongoDB

### Backend Setup

1. Open a terminal and navigate to the backend folder:

```bash
cd server/dipFlixServer
```

2. Copy the sample environment file:

```bash
cp sample.env .env
```

3. Update `.env` with your MongoDB URI and secret keys:

- `DATABASE_NAME`
- `MONGODB_URI`
- `SECRET_KEY`
- `SECRET_REFRESH_KEY`
- `GEMINI_API_KEY` (optional for AI-assisted recommendation or review prompts)
- `RECOMMENDED_MOVIE_LIMIT`
- `ALLOWED_ORIGINS`

4. Install Go module dependencies and run the server:

```bash
go mod tidy
go run main.go
```

The backend will start on `http://localhost:8080`.

### Frontend Setup

1. Open a new terminal and navigate to the frontend folder:

```bash
cd client/dip-flix-client
```

2. Install dependencies:

```bash
npm install
```

3. Start the development server:

```bash
npm run dev
```

The frontend will typically run on `http://localhost:5173`.

## API Endpoints

### Unprotected

- `GET /movies` — Fetch movie catalog
- `POST /register` — Register a new user
- `POST /login` — Login user and receive tokens
- `POST /logout` — Logout user
- `GET /genres` — Fetch genres list
- `POST /refresh` — Refresh access token

### Protected

- `GET /movie/:imdb_id` — Get movie details
- `POST /movie/add` — Add a movie
- `GET /recommended` — Get recommended movies
- `PATCH /adminreview/:imdb_id` — Update admin review status

## Notes

- The backend uses CORS configuration to allow origins from the `.env` file.
- Authenticated frontend routes include recommendations, review pages, and streaming.
- The sample environment file is provided in `server/dipFlixServer/sample.env`.

## Development

- Frontend: `npm run dev`
- Backend: `go run main.go`

## License

This project is licensed under the MIT License—see the LICENSE file for details.

