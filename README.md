# E-Vote Backend

A robust electronic voting system backend built with Goravel framework. This application provides a secure and scalable API for managing polls, options, and votes.

## Features

- **User Authentication**
  - Secure registration and login
  - JWT-based authentication

- **Polls Management**
  - Create, read, update, and delete polls
  - Schedule polls with start and end dates
  - Set poll status (active/done)
  - Generate shareable poll codes

- **Voting Options**
  - Create and manage voting options for each poll
  - Support for option images/avatars
  - Option descriptions and details

- **Voting System**
  - Secure vote recording
  - Prevention of duplicate votes
  - Real-time vote counting

## Tech Stack

- **Framework**: [Goravel](https://www.goravel.dev/) (Go-based web framework)
- **Database**: PostgreSQL
- **Storage**: MinIO for file storage
- **Documentation**: Swagger for API documentation
- **Containerization**: Docker support

## Getting Started

### Prerequisites

- Go 1.23+
- PostgreSQL
- MinIO (for file storage)
- Docker (optional)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/rizkirmdhn/evote-be.git
   cd evote-be
   ```

2. Copy the environment file and configure it:
   ```bash
   cp .env.example .env
   ```
   Update the `.env` file with your database, MinIO, and other configuration settings.

3. Install dependencies:
   ```bash
   go mod tidy
   ```

4. Run database migrations:
   ```bash
   go run . artisan migrate
   ```

5. (Optional) Seed the database with sample data:
   ```bash
   go run . artisan db:seed
   ```

### Running the Application

#### Development Mode

```bash
go run . 
```

The server will start at `http://localhost:3000`

#### Using Docker

1. Build and start the containers:
   ```bash
   docker-compose up -d
   ```

2. The application will be available at `http://localhost:3000`

## API Documentation

The API documentation is available through Swagger UI when the application is running:

- Local development: `http://localhost:3000/swagger/index.html`


## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

If you encounter any issues or have questions, please [open an issue](https://github.com/rizkirmdhn/evote-be/issues).