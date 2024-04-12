# API_GATEWAY

This project implements a robust and extensible API Gateway service using Golang, Docker containers, and Kubernetes for deployment. It prioritizes security, scalability, and ease of management.

Installation & Run
# Download this project
go get github.com/Varnit01-dev/Api_gateway
# Build and Run
cd API_GATEWAY
go build
./Api_gateway



Benefits:

Enhanced Security: Robust authentication and authorization mechanisms safeguard your API endpoints.
Improved Scalability: Easily scale the API Gateway and microservices to accommodate traffic surges.
Simplified Deployment and Management: Docker and Kubernetes streamline deployment and management across environments.
Flexibility: The architecture is adaptable to various authentication mechanisms and load balancing strategies based on your project's specific needs.

 
## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`API_KEY`

`ANOTHER_API_KEY`


## Deployment

To deploy this project run

```bash
  docker build -t Varnit01-dev/Api_Gateway .

```

