# Go Web Application with User Authentication

This is a simple Go web application that demonstrates user authentication using sessions. Users can log in, sign up for new accounts, and access a dashboard with user-specific content. The application is built with the Gorilla Mux router and uses Gorilla Sessions for session management.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Go (Golang) installed on your local machine.
- Basic understanding of Go and web development concepts.

## Getting Started

### Installing Go

If you don't have Go installed, you can download and install it from the official website: [https://golang.org/dl/](https://golang.org/dl/)

### Running the Application Locally

1. Clone this repository:

   ```bash
   git clone [https://github.com/yourusername/your-repo.git](https://github.com/AyoBakre/portal-golang)
   ```

2. Change into the project directory:

   ```bash
   cd nautilus
   ```

3. Run the application:

   ```bash
   go run main.go
   ```

   The application should now be running locally at [http://localhost:8080](http://localhost:8080).

## Usage

- **Login:** Access the login page at [http://localhost:8080/login](http://localhost:8080/login) to log in with an existing account.
- **Signup:** Access the signup page at [http://localhost:8080/signup](http://localhost:8080/signup) to create a new account.
- **Dashboard:** After logging in, you will be redirected to the dashboard at [http://localhost:8080/dashboard](http://localhost:8080/dashboard), where you can view user-specific content.
- **Logout:** To log out, visit [http://localhost:8080/logout](http://localhost:8080/logout).

## Project Structure

The project structure is organized as follows:

- `main.go`: The main application entry point.
- `templates/`: HTML templates for login, signup, and dashboard pages.
- `static/`: Static files (e.g., CSS) served by the application.
- `sessions/`: Session data storage directory (session data is stored in files in this directory).
- `handlers/`: Go code for handling different routes and user authentication.
- `README.md`: This README file.

## Features

- User authentication using sessions.
- Login and signup functionality.
- User-specific dashboard.
- Static file serving for CSS and assets.
- In-memory user store (for demonstration purposes).

## Deployment

This application is designed to run on various hosting platforms. To deploy it, follow these general steps:

1. Set the `PORT` environment variable to specify the port on which the application should listen (e.g., `8080`).

2. Configure your hosting platform to run the Go application.

3. Set up environment variables for any secrets or sensitive information.

4. Deploy the application to your chosen hosting platform.
