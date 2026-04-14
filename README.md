# Chirpy
Chirpy is a lightweight Twitter-like REST API built in Go.

## Features
  - User registration and authentication
  - Create, read, and delete chirps
  - PostgreSQL database integration
  - JSON-based request and response handling

## Commands
  - GET /api/healthz
    - Health check
  - GET /admin/metrics
    - View server metrics
  - POST /admin/reset
    - Reset state
  - POST /api/users
    - Create user
  - PUT /api/users
    - Update user
  - POST /api/login
    - Login user
  - POST /api/refresh
    - Refresh token
  - POST /api/revoke
    - Revoke token
  - POST /api/chirps
    - Create chirp
  - GET /api/chirps/{chirpID}
    - Get single chirp
  - DELETE /api/chirps/{chirpID}
    - Delete chirp
