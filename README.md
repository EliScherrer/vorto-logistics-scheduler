# Logistics Scheduler

A full-stack web application for managing delivery schedules with separate interfaces for Admins and Drivers.

## Architecture

- **Frontend**: Vue.js 3 + TypeScript + TailwindCSS + Leaflet.js
- **Backend**: Go + Gin framework + GORM
- **Database**: PostgreSQL
- **Containerization**: Docker + Docker Compose

## Project Structure

```
├── backend/                 # Go API server
│   ├── main.go             # Application entry point
│   ├── handlers/           # HTTP request handlers
│   ├── models/             # Data models
│   ├── database/           # Database connection
│   ├── go.mod              # Go module definition
│   └── Dockerfile          # Backend container config
├── frontend/               # Vue.js application
│   ├── src/                # Source code
│   │   ├── components/     # Vue components
│   │   ├── main.ts         # Application entry point
│   │   └── style.css       # Global styles
│   ├── package.json        # Node.js dependencies
│   ├── vite.config.ts      # Vite configuration
│   ├── tailwind.config.js  # TailwindCSS configuration
│   └── Dockerfile          # Frontend container config
└── docker-compose.yml      # Multi-container setup
```

## Getting Started

### Prerequisites

- Docker and Docker Compose
- Go 1.21+ (for local development)
- Node.js 18+ (for local development)

### Running with Docker

1. Clone the repository
2. Start all services:
   ```bash
   docker-compose up -d
   ```
3. Access the application:
   - Frontend: http://localhost:3000
   - Backend API: http://localhost:8080
   - Database: localhost:5432

### Local Development

#### Backend
```bash
cd backend
go mod tidy
go run main.go
```

#### Frontend
```bash
cd frontend
npm install
npm run dev
```

#### Database
```bash
docker-compose up postgres -d
```

## Features

- **Home Page**: Role selection (Admin/Driver)
- **Admin Portal**: Load creation and tracking with map interface
- **Driver Portal**: Shift management and load workflow
- **Real-time Updates**: Load status tracking and assignment
- **Map Integration**: Interactive maps for location selection and display

## API Endpoints

### Authentication
- `POST /api/admin/login` - Admin authentication
- `POST /api/driver/login` - Driver authentication

### Admin Endpoints
- `POST /api/admin/loads` - Create new load
- `GET /api/admin/loads` - Get all loads

### Driver Endpoints
- `POST /api/driver/shift/start` - Start shift
- `POST /api/driver/shift/end` - End shift
- `GET /api/driver/status` - Get driver status
- `POST /api/driver/load/accept` - Accept assigned load
- `POST /api/driver/load/reject` - Reject assigned load
- `POST /api/driver/load/pickup` - Complete pickup
- `POST /api/driver/load/dropoff` - Complete dropoff

## Original Requirements

The application is a simple work scheduling platform in the logistics industry. The primary users are drivers who log into the application to receive work from the server. The unit of work is a load: an order to pick up a payload from one location and drop it off at another. Loads are submitted to the system via an admin, who has a simple UI to submit loads and track their progress. The system then dispatches loads to the drivers efficiently, who in turn inform the system when each stop (i.e. the pickup and dropoff) is completed.

### Driver Interface Requirements:
- A driver can be on shift, during which time he can be assigned loads; or off shift, during which time no work will be assigned to him
- While not on shift, allow a driver to start a shift, at which time they provide their current location as a (latitude, longitude) pair
- Shifts are indefinite in duration; they do not end until the driver voluntarily ends them
- While on shift, a driver may be assigned loads from the server. When assigned a load:
  - Show the load's pickup and dropoff locations on a map
  - Provide a button to complete the next task, which is either to pickup or dropoff the payload
  - Completing each stop should inform the system that the driver is now at that location
  - Upon completing the dropoff, the load itself is completed and should be removed from view
  - If the pickup has not yet been completed, provide a button to reject the load, which also ends the driver's shift
- If a driver is on shift and has no load assigned, provide a button to end shift

### Admin Interface Requirements:
- Provide a tool for submitting loads to the application
- A load must simply have a pickup location and a dropoff location, both in the form of a (latitude, longitude)
- The tool must provide a map interface so that shippers can simply click on two locations and submit
- Provide a list of all loads in the system, with their current statuses
- Possible statuses are awaiting driver, in progress, and completed

### Backend Requirements:
- All drivers, shifts, and loads should be persisted to a database
- Load assignment should be efficient. E.g. all else equal, loads should be dispatched to minimize driven miles
- Drivers simply provide a user name when beginning a session; a driver is created if one does not exist
- Drivers should be able to log out while on shift or with a load, and log back in still on the same shift and load