# Card Validator Service

This service provides card validation functionality, including Luhn algorithm checks and expiration date validation.


## Running the Application

1. Start the service using Docker Compose:
   ```
   docker-compose up -d
   ```

   This command will build the Docker image if it doesn't exist and start the service in detached mode.

2. To stop the service:
   ```
   docker-compose down
   ```

## API Usage

### Validate Card

Endpoint: `POST /validate-card`

Example request:

```bash
curl -X POST http://localhost:8080/validate-card \
  -H "Content-Type: application/json" \
  -d '{
    "number": "4532015112830366",
    "expirationDate": {
      "month": 12,
      "year": 2025
    }
  }'
```

Example response:

```json
{
  "valid": bool
  "error": {
      "code": string
      "message": string
  }
}
```