# Online Bookstore Platform

## Overview

The Online Bookstore Platform is a microservices-based application that allows users to browse, search, and purchase books online. This platform is designed with modularity and scalability in mind, employing microservices architecture to manage different aspects of the system.

## Features

- **Catalog Service:**
  - Manages book information, including title, author, genre, and cover image.
  - Provides APIs for searching and retrieving book details.

- **User Service:**
  - Manages user accounts, authentication, and authorization.
  - Handles user registration, login, and profile information.

- **Order Service:**
  - Manages the shopping cart and order processing.
  - Coordinates with the payment gateway for payment processing.

- **Payment Service:**
  - Integrates with third-party payment gateways (Stripe, Mpesa).
  - Processes payment transactions securely.

- **Review Service:**
  - Allows users to submit and view book reviews.
  - Manages ratings and reviews for each book.

- **Notification Service:**
  - Sends email or push notifications for order confirmations, shipping updates, etc.
  - Integrates with a messaging system for real-time notifications.

- **Recommendation Service:**
  - Recommends books to users based on purchase history, preferences, and reviews.
  - Utilizes machine learning algorithms for personalized recommendations.

- **Gateway Service:**
  - Acts as the main entry point for the application.
  - Routes requests to the appropriate microservices.
  - Implements authentication and authorization checks.

## Technologies Used

- **Programming Languages:** Go.
- **Database:** Cassandra, PostgreSQL.
- **Message Broker:** RabbitMQ.
- **Containerization:** Docker for packaging services into containers.
- **Orchestration:** Kubernetes for managing and orchestrating containers.
- **Caching:** Redis.
- **Logging:** Elastic search, Logstash and Kibana(ELK).
- **Monitoring:** Prometheus.
- **Distributed Tracing** Jaeger
  
## Getting Started

### Prerequisites

- Go.
- Docker.
- Kubernetes for deployment (optional).
- AWS

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/paulodhiambo/zbooks.git
   ```

2. Follow the README instructions in each microservices directory to set up and run the services.

3. Deploy the application using Docker and Kubernetes (if applicable).

## Usage

- Access the Online Bookstore Platform through the provided API endpoints.
- Explore the microservices documentation for detailed information on each service.

## Contributing

1. Fork the repository.
2. Create a feature branch: `git checkout -b feature-name`.
3. Commit your changes: `git commit -m 'Add some feature'`.
4. Push to the branch: `git push origin feature-name`.
5. Submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

## Acknowledgments

- Inspired by the need for a scalable and modular online bookstore platform.
- Built with a focus on microservices architecture and best practices.
