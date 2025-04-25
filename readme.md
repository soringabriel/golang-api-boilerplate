# Golang API Boilerplate

A lightweight and modular boilerplate for building robust RESTful APIs in Go, inspired by Laravel's MVC structure but adapted for the Go ecosystem. Includes support for MongoDB, middleware, request/response handling, and more.

---

## 🚀 Quickstart

Follow the steps below to get started:

1. **Create your `.env` file**  
   Copy `.env-example` to `.env` and configure the necessary environment variables, especially:
   - `MONGODB_URL` – your MongoDB connection string
   - `AUTH_TOKEN` – a secure token used for request authentication

2. **Build the Docker image**
   ```bash
   docker build -t golang-api-boilerplate .
   ```

3. **Run the container**
   ```bash
   docker run -d --env-file .env -p 8080:8080 golang-api-boilerplate
   ```

4. **Test the API**  
   Use the included `postman_collection.json` to try out the example endpoints.

---

## 🧩 Example Endpoints

This boilerplate includes example CRUD endpoints for a `User` model to demonstrate structure and usage.
These can be freely removed or replaced when building your own API.

## 📁 Project Structure

This project follows a lightweight MVC-inspired structure:

```
/app
  ├── endpoints     # Define API endpoints with routing and logic handlers
  ├── middlewares   # Pre-processing functions (e.g., auth checks, validations)
  ├── models        # MongoDB model structs, each with a ToBson() method
  ├── requests      # Input request definitions with validation and parsing
  ├── responses     # Standardized API responses and response constructors

/databases
  └──               # MongoDB driver and CRUD utilities

/helpers
  └──               # Utility functions used across the codebase

/logger
  └──               # Centralized logging module

/tests
  └──               # Integration and feature tests
```

---

## 📚 Documentation

- **Modular structure**: Clean separation between concerns makes it easy to extend.
- **Middleware support**: Easily plug in custom logic to run before request handlers.
- **MongoDB integration**: Simple model-based integration using BSON conversion.
- **Built-in validation**: Request types implement `Validate()` to ensure inputs are safe and clean.

---

## 🧪 Testing

To run the integration tests, make sure the Docker service is up and running.

Then execute the following command:

```bash
go test -parallel=1 -count=1 ./tests
```
This ensures tests run sequentially with fresh state. All integration tests are located in the `/tests` directory.

---

## 🛠️ Contributing

Contributions are welcome! Please open issues or pull requests for any improvements or bug fixes.

---

## 📄 License

This project is open-source and available under the [MIT License](LICENSE).
