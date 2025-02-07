# 🏥 MediTrust

Fast and reliable backend for an online pharmacy, powered by Go and PostgreSQL.

## ⚙️ SetUp

This project uses Go and PostgresSQL. Please ensure to set them up before proceeding.

1. Install Go
   - Download and install Go from the [official Go website](https://go.dev/dl/).
   - Verify installation:

    ```bash
    go version
    ```

2. Install PostgreSQL
    - Install PostgreSQL from the [official PostgreSQL website](https://www.postgresql.org/download/).
    - Verify Installation:

    ```bash
    psql --version
    ````


## 🔐 Environment Variables

Create a `.env` file in the root directory with the following variables:

```env
DB_URL="postgres://<username>:<password>@<host>:<port>/<database_name>"
SECRET_KEY="Your Secret Key here"
PLATFORM="<environment>"
```
- Replace the `DB_URL` values with your database connection string.
- Currently available platform is `"deb"`.

Find `db_migration.sh` in the `scripts` to automate DB migration.

# 📌 API-endpoints

### 💊 Medicine

| name                  | method | route                    | note         |
| :-------------------- | :----: | :----------------------- | ------------ |
| create medicine       |  POST  | /api/v1/medicines        | *admin only* |
| get all medicine      |  GET   | /api/v1/medicines        |
| get medicine by ID    |  GET   | /api/v1/medicines/:medID |
| update medicine by ID |  PUT   | /api/v1/medicines/:medID | *admin only* |
| delete medicine by ID | DELETE | /api/v1/medicines/:medID | *admin only* |

### 👤 Users

| name           | method | route                      | note             |
| :------------- | :----: | -------------------------- | ---------------- |
| get user by ID |  GET   | /api/v1/users/:userID      | *admin only*     |
| signup user    |  POST  | /api/v1/signup             |
| login user     |  POST  | /api/v1/login              |
| verify user    |  GET   | /api/v1/verify?token=token | *auto-generated* |
| update user    |  PUT   | /api/v1/users              |
| logout user    |  POST  | /api/v1/logout             |
| delete user    | DELETE | /api/v1/users              |
| refresh token  |  POST  | /api/v1/refresh            |
| revoke token   |  POST  | /api/v1/revoke             |

### 🔄 General

| name     | method | route         | note                   |
| -------- | ------ | ------------- | ---------------------- |
| reset DB | POST   | /api/v1/reset | *dev* environment only |



# ✅ To-Do

> Goal is to achieve MVP

### Authentication & Authorization

- [x] User Signup
- [x] User Login
- [x] Email verification
- [x] User Logout
- [x] Token Refresh
- [x] Token Revoke
- [ ] Password Reset (*via email*)
- [ ] Role-based Access control (*User/Admin*)

### Product Management (Medicines)

- [x] Create medicine (*admin*)
- [x] Get all medicine
- [x] Get medicine by ID
- [x] Update medicine by ID (*admin*)
- [x] Delete medicine by ID (*admin*)
- [ ] Medicine categories & Tags
- [ ] Search & Filter medicines
- [ ] Pagination for large medicine lists

### User Profile & Management

- [x] Get user by ID (*admin*)
- [x] Update user profile
- [x] Delete user
- [ ] View order history

### Cart & Wishlist

- [ ] Add medicine to cart
- [ ] Remove medicine from cart
- [ ] Save medicine to wishlist
- [ ] Remove medicine from wishlist

### Checkout & Payment

- [ ] Order placement (*store order details*)
- [ ] Integrate Payment Gateway
- [ ] Discounts & Promo codes

### Order Management

- [ ] Order Tracking
- [ ] Cancel Orders
- [ ] Return & Refund System

### Admin Panel

- [x] Manage users (*Get, Delete*)
- [x] Manage medicines (CRUD)
- [ ] Manage orders (CRUD)
- [ ] Generate Reports & Statistics

### Security & Compliance

- [ ] implement Rate Limiting & Throttling
- [ ] implement Logging & Monitoring
- [ ] Ensure GDPR & Data Protection Compliance

### General

- [x] Reset DB (*DEV environment only*)


# 📖 Documentation

> Experimental

Run the project and see the documentation at `http://localhost:8080/api/v1/swagger/index.html`.

# 🚀 Status

![code coverage badge](https://github.com/Dhar01/medicine-app/actions/workflows/ci.yml/badge.svg)
