# Requirements Document

## Introduction

The Vorto Logistics Scheduler is a full-stack web application with two distinct user interfaces: one for Admins and one for Drivers. The system allows admins to create and manage delivery schedules, while drivers can view their assigned deliveries and update their status. The application features a modern web interface built with Vue.js 3 and TypeScript, backed by a Go API server and PostgreSQL database.

## Requirements

### Requirement 1

**User Story:** As a visitor to the website, I want to select my role (Admin or Driver) from the home page, so that I can access the appropriate login interface.

#### Acceptance Criteria

1. WHEN a user visits the home page THEN the system SHALL display a simple interface with two options: "Admin" and "Driver"
2. WHEN a user selects "Admin" THEN the system SHALL redirect them to the admin login page
3. WHEN a user selects "Driver" THEN the system SHALL redirect them to the driver login page
4. WHEN the home page loads THEN the system SHALL not require any authentication

### Requirement 2

**User Story:** As a driver, I want to log in with just my username, so that I can access my driver portal and maintain my shift/load state across sessions.

#### Acceptance Criteria

1. WHEN a driver enters their username THEN the system SHALL create a driver record if one doesn't exist
2. WHEN a driver logs in THEN the system SHALL restore their previous shift status and any assigned load
3. WHEN a driver logs out while on shift THEN the system SHALL maintain their shift status for when they log back in
4. WHEN a driver logs out with an assigned load THEN the system SHALL preserve the load assignment for their return

### Requirement 3

**User Story:** As an admin, I want to create loads with pickup and dropoff locations using a map interface, so that the system can dispatch them to drivers efficiently.

#### Acceptance Criteria

1. WHEN an admin accesses the load creation tool THEN the system SHALL display a map interface
2. WHEN an admin creates a new load THEN the system SHALL require pickup and dropoff locations as (latitude, longitude) pairs selected by clicking on the map
3. WHEN an admin submits a load THEN the system SHALL automatically dispatch it to minimize driven miles using efficient assignment logic
4. WHEN no drivers are on shift THEN the system SHALL queue the load with status "awaiting driver"

### Requirement 4

**User Story:** As an admin, I want to view all loads in the system with their current statuses, so that I can track progress.

#### Acceptance Criteria

1. WHEN an admin views the load list THEN the system SHALL display all loads with statuses: "awaiting driver", "in progress", or "completed"
2. WHEN a load status changes THEN the system SHALL update the display to reflect the current status
3. WHEN an admin views load details THEN the system SHALL show pickup/dropoff locations and assigned driver (if any)
4. WHEN loads are updated by drivers THEN the system SHALL reflect status changes in the admin interface

### Requirement 5

**User Story:** As a driver, I want to manage my shift status and location, so that I can control when I receive load assignments.

#### Acceptance Criteria

1. WHEN a driver is not on shift THEN the system SHALL provide a button to start a shift requiring current location as (latitude, longitude)
2. WHEN a driver starts a shift THEN the system SHALL record their location and make them available for load assignments
3. WHEN a driver is on shift with no assigned load THEN the system SHALL provide a button to end shift
4. WHEN a driver ends their shift THEN the system SHALL stop assigning new loads to them

### Requirement 6

**User Story:** As a driver, I want to manage assigned loads through pickup and dropoff workflow, so that I can complete deliveries and update my location.

#### Acceptance Criteria

1. WHEN a driver is assigned a load THEN the system SHALL show pickup and dropoff locations on a map
2. WHEN a driver has not completed pickup THEN the system SHALL provide a "reject load" button that ends their shift
3. WHEN a driver completes pickup THEN the system SHALL update their location to the pickup location and show a "dropoff" button
4. WHEN a driver completes dropoff THEN the system SHALL update their location to dropoff location, mark load as completed, and remove it from view
5. WHEN a driver completes a load THEN the system SHALL make them available for new load assignments

### Requirement 7

**User Story:** As an admin, I want to access the admin portal with simple authentication, so that I can manage loads.

#### Acceptance Criteria

1. WHEN an admin accesses the admin login page THEN the system SHALL provide a simple authentication mechanism
2. WHEN admin authentication is successful THEN the system SHALL grant access to load creation and tracking tools
3. WHEN an admin session expires THEN the system SHALL redirect to the login page
4. WHEN an admin logs out THEN the system SHALL end their session and return to the home page

### Requirement 8

**User Story:** As the system, I want to efficiently dispatch loads to minimize driven miles, so that operations are optimized.

#### Acceptance Criteria

1. WHEN a new load is submitted THEN the system SHALL calculate the most efficient driver assignment to minimize total driven miles
2. WHEN multiple drivers are on shift THEN the system SHALL assign loads based on proximity and efficiency algorithms
3. WHEN a driver becomes available (starts shift or completes a load) THEN the system SHALL check for pending loads and assign optimally
4. WHEN no drivers are on shift THEN the system SHALL queue loads with "awaiting driver" status until drivers become available