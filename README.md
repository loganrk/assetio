# ASSETIO
## Overview
This project provides a comprehensive platform to manage accounts, stocks, securities with a modular, maintainable, and scalable structure. We have implemented a hexagonal architecture to ensure a clean separation of concerns, making the platform adaptable to various input/output channels.

## Project Structure
The project is structured in line with the principles of hexagonal architecture (also known as ports and adapters). This approach allows the core business logic to remain independent from external technologies and frameworks.

### Hexagonal Architecture Overview
- **Domain**: Contains all the business logic and models. This layer remains agnostic of any specific framework or database technology, allowing easy testing and maintenance.
- **Adapters**: Interface implementations for different infrastructure components such as HTTP controllers, databases, and external APIs.
-  **Ports**: Interfaces representing operations provided to and required by the application’s core.
## Features
- **Account Management**: Creation, retrieval, updating, activation, and deactivation of accounts.
- **Security Management**: Creating, retrieving, updating, and searching securities.
- **Stock Management**: Buying, selling, dividend processing, splitting, and summaries of stocks.