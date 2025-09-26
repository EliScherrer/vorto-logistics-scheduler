# Implementation Plan

- [ ] 1. Set up project structure and development environment

  - Create Go backend directory structure with main.go, handlers, models, and database packages
  - Initialize Go modules and install dependencies (Gin, GORM, PostgreSQL driver)
  - Create Vue.js 3 + TypeScript frontend with Vite build tool
  - Install frontend dependencies (Vue Router, TailwindCSS, Leaflet.js, Axios)
  - Set up Docker and Docker Compose configuration for PostgreSQL database
  - Create basic .gitignore and README files
  - _Requirements: 1.1, 2.1_

- [ ] 2. Implement database schema and models

  - [ ] 2.1 Create PostgreSQL database schema with migrations

    - Write SQL migration files for users, drivers, and loads tables with driverStatus field
    - Set up GORM auto-migration in Go application
    - Add database connection configuration and initialization
    - _Requirements: 3.1, 4.1, 5.1, 6.1, 7.1, 8.1_

  - [ ] 2.2 Implement Go data models and database operations
    - Create User, Driver, and Load structs with GORM tags including driverStatus field
    - Implement CRUD operations for each model with driverStatus updates
    - Add database query methods for load assignment logic using driverStatus
    - Write helper functions for distance calculations
    - _Requirements: 3.1, 4.1, 5.1, 6.1, 7.1, 8.1_

- [ ] 3. Build core backend API server

  - [ ] 3.1 Set up Gin web server with middleware

    - Create main server setup with Gin framework
    - Implement CORS middleware for frontend communication
    - Add request logging and error handling middleware
    - Set up username-based authentication middleware
    - _Requirements: 2.1, 7.1_

  - [ ] 3.2 Implement authentication endpoints

    - Create POST /api/admin/login endpoint for admin authentication
    - Create POST /api/driver/login endpoint with driver creation logic
    - Implement username validation and user creation
    - Add middleware to extract and validate X-Username header
    - _Requirements: 2.1, 7.1_

  - [ ] 3.3 Implement admin API endpoints

    - Create POST /api/admin/loads endpoint for load creation with driverStatus "unassigned"
    - Create GET /api/admin/loads endpoint to retrieve all loads with status and driverStatus
    - Add load validation and database persistence with proper status initialization
    - Implement load assignment trigger that updates driverStatus to "assigned"
    - _Requirements: 3.1, 3.2, 3.3, 4.1, 4.2, 4.3_

  - [ ] 3.4 Implement driver API endpoints
    - Create POST /api/driver/shift/start endpoint with location handling
    - Create POST /api/driver/shift/end endpoint to end shifts
    - Create GET /api/driver/status endpoint to get current driver state and assigned load
    - Create POST /api/driver/load/accept endpoint that updates driverStatus to "accepted"
    - Create POST /api/driver/load/reject endpoint that updates driverStatus to "unassigned"
    - Create POST /api/driver/load/pickup endpoint that updates driverStatus to "pickedup"
    - Create POST /api/driver/load/dropoff endpoint that updates driverStatus to "droppedoff"
    - _Requirements: 5.1, 5.2, 5.3, 6.1, 6.2, 6.3, 6.4, 6.5_

- [ ] 4. Implement load assignment algorithm

  - [ ] 4.1 Create distance calculation utilities

    - Implement Haversine formula for calculating distances between coordinates
    - Create helper functions for finding closest drivers
    - Add validation for latitude/longitude coordinates
    - _Requirements: 8.1, 8.2_

  - [ ] 4.2 Build load assignment engine
    - Create algorithm to find available drivers (on shift, no current load with driverStatus "accepted", "pickedup")
    - Implement logic to assign loads to closest available driver and update driverStatus to "assigned"
    - Add fallback handling when no drivers are available (keep driverStatus "unassigned")
    - Create load reassignment logic when drivers reject loads (reset driverStatus to "unassigned")
    - _Requirements: 8.1, 8.2, 8.3, 8.4_

- [ ] 5. Build frontend application structure

  - [ ] 5.1 Set up Vue.js application with routing

    - Create Vue 3 application with TypeScript configuration
    - Set up Vue Router with routes for home, admin, and driver portals
    - Configure TailwindCSS for styling
    - Create base layout components and navigation structure
    - _Requirements: 1.1, 1.2, 1.3_

  - [ ] 5.2 Implement authentication and session management
    - Create login components for admin and driver portals
    - Implement local storage session management with username
    - Add Axios interceptors to include X-Username header in requests
    - Create authentication guards for protected routes
    - _Requirements: 2.1, 2.2, 2.3, 7.1, 7.2_

- [ ] 6. Build home page and role selection

  - Create simple home page component with admin/driver selection buttons
  - Implement navigation to respective login pages based on role selection
  - Add basic styling and responsive design
  - Ensure unauthenticated access to home page
  - _Requirements: 1.1, 1.2, 1.3, 1.4_

- [ ] 7. Implement admin portal interface

  - [ ] 7.1 Create load creation interface with map

    - Build map component using Leaflet.js and OpenStreetMap
    - Implement click-to-select functionality for pickup and dropoff locations
    - Create form for load details with coordinate display
    - Add load submission with API integration
    - _Requirements: 3.1, 3.2, 3.3_

  - [ ] 7.2 Build load tracking dashboard
    - Create load list component with status and driverStatus filtering
    - Display all loads with current status (awaiting driver, in progress, completed) and driverStatus
    - Add refresh button and periodic polling for updated data
    - Add load detail view with pickup/dropoff locations, assigned driver, and detailed status progression
    - _Requirements: 4.1, 4.2, 4.3, 4.4_

- [ ] 8. Implement driver portal interface

  - [ ] 8.1 Create shift management interface

    - Build shift status display showing current on/off shift state
    - Create start shift interface with location input (map or manual lat/lng)
    - Implement end shift functionality with API integration
    - Add current location display and shift duration tracking
    - _Requirements: 5.1, 5.2, 5.3, 5.4_

  - [ ] 8.2 Build load workflow interface
    - Create waiting interface for drivers with no assigned loads (driverStatus filtering)
    - Implement load assignment notification display with accept/reject buttons for "assigned" loads
    - Build map view showing pickup and dropoff locations for "accepted" loads
    - Create pickup button for "accepted" loads that updates driverStatus to "pickedup"
    - Create dropoff button for "pickedup" loads that updates driverStatus to "droppedoff"
    - Handle load completion and return to waiting state after "droppedoff" status
    - _Requirements: 6.1, 6.2, 6.3, 6.4, 6.5_

- [ ] 9. Implement error handling and validation

  - [ ] 9.1 Add comprehensive backend error handling

    - Implement structured error responses for all API endpoints
    - Add input validation for all request parameters
    - Create proper HTTP status codes for different error types
    - Add error logging for debugging and monitoring
    - _Requirements: All requirements for robustness_

  - [ ] 9.2 Build frontend error handling and user feedback
    - Create error display components for form validation
    - Implement network error handling with retry mechanisms
    - Add loading states and success/error notifications
    - Handle API errors with user-friendly messages and refresh suggestions
    - _Requirements: All requirements for user experience_

- [ ] 10. Final integration and deployment setup

  - [ ] 10.1 Complete end-to-end integration testing

    - Test complete user workflows for both admin and driver portals
    - Verify load assignment algorithm works correctly with multiple drivers
    - Test data refresh and API functionality
    - Validate session persistence across browser refreshes
    - _Requirements: All requirements_

  - [ ] 10.2 Set up production deployment configuration
    - Create production Docker Compose configuration
    - Set up environment variable configuration for different environments
    - Add database migration scripts for production deployment
    - Create deployment documentation and setup instructions
    - _Requirements: All requirements for deployment_
