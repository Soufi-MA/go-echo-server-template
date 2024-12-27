# Simple Go Server Template with Echo

This is a simple server template built with **Go** using the **Echo** framework. It provides a basic structure for building web applications, including database connectivity, modular routing, and environment configuration management.

---

---

## **Features**

- **Database Connection:** PostgreSQL connection setup using `pgx/v5`.
- **Environment Management:** `.env` file loading with `github.com/joho/godotenv`.
- **Modular Routing:** Routes defined in the `handlers` directory.
- **Template Rendering:** HTML templates stored in `public/views`.
- **Static Assets:** CSS and JS files stored in `public/assets`.
- **Live Reloading:** Integrated with `air` for hot reloading during development.

---

## **Getting Started**

### **Prerequisites**

- Go 1.21 or later.
- PostgreSQL installed and running.
- Node.js (optional for Tailwind CSS builds).

### **Installation**

1. Clone the repository:

   ```bash
   git clone <repo-url>
   cd <repo-name>
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Create a `.env.local` file:

   ```
   DATABASE_URL=postgres://user:password@localhost:5432/mydb
   ```

4. Initialize the database:

   ```bash
   psql -U <user> -d <database> -f config/create_db.sql
   ```

5. Install **Air** (optional for live reload):

   ```bash
   go install github.com/cosmtrek/air@latest
   ```

6. Run the server:
   ```bash
   air
   ```
   Or without Air:
   ```bash
   go run server.go
   ```

---

## **API Endpoints**

### **User Routes**

- `GET /users` - Fetch all users.
- `POST /users` - Create a new user.
- `GET /` - Render the home page (index.html).
- `GET /users` - Fetch all users.
- `GET /users/:id` - Fetch a user by their ID.
- `GET /users/search` - Search for users.
- `GET /add-user` - Render the add user page (add-user.html).
- `POST /users` - Create a new user.
- `PUT /users/:id` - Update an existing user by ID.
- `DELETE /users/:id` - Delete a user by ID.

---

## **Environment Configuration**

Environment variables are loaded from a `.env.local` file. Example:

---

## **Development Tools**

- **Live Reloading:** Uses `air` for automatic restarts during development.
- **Tailwind CSS:** Optional setup for styling.

---

## **License**

This project is licensed under the MIT License. See the `LICENSE` file for details.

---

## **Contributing**

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

---

## **Contact**

For questions or support, please reach out to [GitHub Profile](https://github.com/soufi-ma).
