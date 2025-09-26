#!/bin/bash

# Development helper script for Logistics Scheduler

case "$1" in
  "start")
    echo "Starting all services with Docker Compose..."
    docker-compose up -d
    echo "Services started:"
    echo "- Frontend: http://localhost:3000"
    echo "- Backend: http://localhost:8080"
    echo "- Database: localhost:5432"
    ;;
  "stop")
    echo "Stopping all services..."
    docker-compose down
    ;;
  "logs")
    docker-compose logs -f
    ;;
  "backend")
    echo "Starting backend in development mode..."
    cd backend && go run main.go
    ;;
  "frontend")
    echo "Starting frontend in development mode..."
    cd frontend && npm run dev
    ;;
  "db")
    echo "Starting only PostgreSQL database..."
    docker-compose up postgres -d
    ;;
  "clean")
    echo "Cleaning up containers and volumes..."
    docker-compose down -v
    docker system prune -f
    ;;
  *)
    echo "Usage: $0 {start|stop|logs|backend|frontend|db|clean}"
    echo ""
    echo "Commands:"
    echo "  start     - Start all services with Docker Compose"
    echo "  stop      - Stop all services"
    echo "  logs      - Show logs from all services"
    echo "  backend   - Run backend in development mode"
    echo "  frontend  - Run frontend in development mode"
    echo "  db        - Start only PostgreSQL database"
    echo "  clean     - Clean up containers and volumes"
    exit 1
    ;;
esac