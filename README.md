# Ticket Management Backend

A robust and scalable ticket management system built with Go, featuring clean architecture, JWT authentication, and microservices-ready design.

## 🚀 Features

- **Clean Architecture**: Follows hexagonal architecture principles with clear separation of concerns
- **JWT Authentication**: Secure token-based authentication with configurable expiration
- **RESTful API**: Well-structured REST endpoints using Fiber framework
- **Database Support**: PostgreSQL with GORM ORM for data persistence
- **Caching**: Redis integration for improved performance
- **Message Queue**: RabbitMQ for asynchronous task processing
- **Docker Support**: Complete containerization with Docker Compose
- **Logging**: Structured logging with Zap and Lumberjack
- **Health Checks**: Built-in health monitoring endpoints
- **Environment Configuration**: Flexible configuration management

## 🏗️ Architecture

The project follows Clean Architecture principles with the following structure:

```
internal/
├── core/           # Business logic and domain models
│   ├── domain/     # Domain entities
│   ├── port/       # Interfaces (ports)
│   └── service/    # Business logic services
├── adapter/        # External adapters
│   ├── handler/    # HTTP handlers (Fiber)
│   ├── storage/    # Database repositories
│   ├── config/     # Configuration management
│   └── message_broker/ # Message queue integration
└── util/           # Utility functions
```

## 🛠️ Tech Stack

- **Language**: Go 1.24
- **Web Framework**: Fiber v2
- **Database**: PostgreSQL with GORM
- **Cache**: Redis
- **Message Queue**: RabbitMQ
- **Authentication**: JWT
- **Logging**: Zap + Lumberjack
- **Containerization**: Docker & Docker Compose

## 📋 Prerequisites

- Go 1.24 or higher
- Docker and Docker Compose
- PostgreSQL (if running locally)
- Redis (if running locally)
- RabbitMQ (if running locally)

## 🚀 Quick Start

### Using Docker Compose (Recommended)

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd ticket-be
   ```

2. **Create environment file**
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

3. **Start all services**
   ```bash
   ./deploy.sh start
   ```

4. **Check service status**
   ```bash
   ./deploy.sh status
   ```

### Manual Setup

1. **Install dependencies**
   ```bash
   go mod download
   ```

2. **Set up environment variables**
   ```bash
   export JWT_SECRET_KEY="your-secret-key"
   export POSTGRE_URI="postgres://user:password@localhost:5432/task_management"
   export REDIS_HOST="localhost:6379"
   export RABBITMQ_URI="amqp://guest:guest@localhost:5672/"
   ```

3. **Run the application**
   ```bash
   go run cmd/main.go
   ```

## 🔧 Configuration

The application uses environment variables for configuration. Key variables include:

### Application
- `APP_NAME`: Application name (default: MyApp)
- `API_PORT`: Server port (default: 8760)
- `LOG_LEVEL`: Logging level (default: info)
- `DEVELOPMENT`: Development mode (default: false)

### Database
- `POSTGRE_URI`: PostgreSQL connection string
- `POSTGRE_MAX_IDLE_CONNS`: Maximum idle connections (default: 10)
- `POSTGRE_MAX_OPEN_CONNS`: Maximum open connections (default: 100)

### Authentication
- `JWT_SECRET_KEY`: JWT secret key (required)
- `JWT_EXPIRE_DAYS_COUNT`: Token expiration in days
- `JWT_ISSUER`: JWT issuer (default: MyApp)

### Redis
- `REDIS_HOST`: Redis host (required)
- `REDIS_PASSWORD`: Redis password
- `REDIS_READ_TIMEOUT`: Read timeout in milliseconds
- `REDIS_WRITE_TIMEOUT`: Write timeout in milliseconds

### RabbitMQ
- `RABBITMQ_URI`: RabbitMQ connection string (required)
- `RABBITMQ_EXCHANGE`: Exchange name (default: events)
- `RABBITMQ_QUEUE_PREFIX`: Queue prefix (default: Ngorder API)

## 📚 API Endpoints

### Health & Version
- `GET /health` - Health check
- `GET /version` - Application version

### Authentication
- `POST /api/v1/auth/signup` - User registration
- `POST /api/v1/auth/signin` - User login
- `GET /api/v1/auth/validate` - Token validation (protected)

### Users
- `GET /api/v1/users` - Get all users
- `GET /api/v1/users/:id` - Get user by ID

## 🐳 Docker Commands

The project includes a deployment script with the following commands:

```bash
./deploy.sh build     # Build Docker images
./deploy.sh start     # Start all services
./deploy.sh stop      # Stop all services
./deploy.sh restart   # Restart all services
./deploy.sh logs      # Show logs from all services
./deploy.sh status    # Show status of all services
./deploy.sh clean     # Clean up all Docker resources
./deploy.sh help      # Show help message
```

## 🗄️ Database

The application uses PostgreSQL with the following default configuration:
- **Database**: task-management
- **User**: aprilpollo
- **Password**: apl9921
- **Port**: 5432

### Database Models
- **Users**: User management
- **Organizations**: Organization structure
- **Projects**: Project management
- **Tickets**: Ticket system

## 🔐 Security

- JWT-based authentication
- Password hashing with bcrypt
- CORS configuration
- Environment-based secrets management
- Non-root Docker user

## 📊 Monitoring

- Health check endpoint at `/health`
- Structured logging with configurable levels
- Docker health checks for all services
- Application metrics and monitoring hooks

## 🧪 Development

### Project Structure
```
cmd/                 # Application entry point
internal/            # Private application code
├── adapter/         # External adapters
├── core/           # Business logic
└── util/           # Utilities
docker/             # Docker configuration
scripts/            # Database migrations and scripts
logs/               # Application logs
static/             # Static files
```

### Adding New Features
1. Define domain models in `internal/core/domain/`
2. Create ports (interfaces) in `internal/core/port/`
3. Implement business logic in `internal/core/service/`
4. Create adapters in `internal/adapter/`
5. Add routes in `internal/adapter/handler/fiber/`

## 🚀 Deployment

### Production Deployment
1. Set `DEVELOPMENT=false`
2. Configure production database credentials
3. Set secure JWT secret key
4. Configure Redis and RabbitMQ for production
5. Use the deployment script or Docker Compose

### Environment Variables for Production
```bash
DEVELOPMENT=false
JWT_SECRET_KEY=your-production-secret-key
POSTGRE_URI=postgres://user:password@db:5432/task_management
REDIS_HOST=redis:6379
RABBITMQ_URI=amqp://user:password@rabbitmq:5672/
```

## 📝 Logging

The application uses structured logging with:
- **Zap**: High-performance logging library
- **Lumberjack**: Log rotation
- **Configurable levels**: debug, info, warn, error
- **File output**: Logs saved to `logs/` directory

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License.

## 🆘 Support

For support and questions:
- Create an issue in the repository
- Check the logs using `./deploy.sh logs`
- Verify service status with `./deploy.sh status`

## 🔄 Version History

- **v0.1.0**: Initial release with basic ticket management functionality
