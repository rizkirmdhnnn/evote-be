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