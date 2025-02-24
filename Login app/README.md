# Go Authentication System

A simple authentication system built with Go and Gin framework that provides user registration and login functionality.

## Features

- User registration (signup)
- User authentication (login)
- Session management
- Simple HTML templates for the user interface


The server will start on `localhost:8080` by default.

## API Endpoints

### Page Routes
- `GET /login` - Display login page
- `GET /signup` - Display signup page

### Authentication Routes
- `POST /signup` - Create a new user account
- `POST /login` - Authenticate user and create session

## Current Limitations

- Data is stored in memory (not persistent)
- Passwords are stored as plain text (not secure for production)
- No input validation
- No session management

## Future Improvements

- Add database integration for persistent storage
- Implement password hashing
- Add input validation
- Add proper session management
- Add user logout functionality
- Implement password recovery
- Add email verification

