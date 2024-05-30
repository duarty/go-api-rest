📍 Go API Rest (under construction)

Welcome to the Go API Rest repository! This project is a Go-based API that integrates with the Google API to check if a user is within a 20-meter radius of a gym. If the user is within this radius, they can check in at the gym. If the gym does not already exist in the database, it will be saved for future check-ins by other users. This approach minimizes requests to the Google Geolocation API by leveraging a local database.


🌟 Features

    User Check-In: Check if a user is near a gym and allow check-ins.
    Gym Persistence: Save gyms to the database for future reference.
    Efficient API Calls: Minimize requests to the Google API by storing gym data locally.
    Clean Architecture: Well-structured code using DTOs, services, controllers, use cases, domain, and repository layers.
    Postgres with Docker: Use PostgreSQL as the database, managed with Docker for easy setup and management.
    Automated Database Setup: SQL scripts for automatic creation of tables and fields.

🛠️ Technologies Used

    Programming Language: Go (Golang)
    Database: PostgreSQL
    Containerization: Docker
    Geolocation API: Google API
    Architecture: Clean Architecture

🚀 Getting Started
Prerequisites

    Go: Download and install Go
    Docker: Download and install Docker

Installation

    Clone the repository:

    sh

git clone https://github.com/duarty/go-api-rest.git
cd go-api-rest

Set up the database:

sh

docker-compose up -d

Run the SQL script to set up the tables:

sh

docker exec -i your_postgres_container psql -U your_user -d your_database < path_to_your_sql_script.sql

Install dependencies:

sh

go mod tidy

Run the application:

sh

    go run main.go

Configuration

Create a .env file in the root directory and add the following environment variables. You can use the .env.example provided for reference:

sh

############ ENV CONFIGS #################
DB_DRIVER=postgres
DB_HOST=localhost
DB_PORT=5432
DB_PASSWORD=root
DB_USER=postgres
DB_NAME=api
SERVER_PORT=3000
SECRET=secret
JWT_EXPIRES_TIME=300

############ GOOGLE CLOUD PLATFORM SERVICES #################
GOOGLE_MAPS_API_ENDPOINT=https://places.googleapis.com/v1/
GOOGLE_MAPS_API_PLACES=places:searchNearby
GOOGLE_MAPS_API_KEY=GCP_API_KEY

############ API CONFIGS #################
NEARBY_SEARCH_RADIUS=50.0
INCLUDED_TYPES=gym
SEARCH_FIELDMASK=places.id,places.displayName,places.formattedAddress,places.location

📂 Project Structure

The project follows the principles of Clean Architecture. Here's a brief overview of the main directories:

    /domain: Contains the core business logic and entities.
    /usecase: Contains the application-specific business rules.
    /repository: Contains the database access logic.
    /controller: Contains the web handlers and controllers.
    /service: Contains the service layer that coordinates between the controllers and the use cases.
    /dto: Contains Data Transfer Objects used for communication between layers.

📖 Usage

    Check-In: Endpoint to check if a user is within a 20-meter radius of a gym and perform a check-in.
    Add Gym: Automatically adds new gyms to the database if they do not already exist.

🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
📜 License

This project is licensed under the MIT License. See the LICENSE file for details.
📞 Contact

For any questions or suggestions, feel free to reach out via email.

Thank you for checking out Go API Rest! Enjoy checking in at your favorite gyms effortlessly! 💪
