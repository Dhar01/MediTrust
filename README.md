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

| name | method | route | note |
|:-----|:------:|:------|------|
| create medicine | POST | /api/v1/medicines | *admin only* |
| get all medicine | GET | /api/v1/medicines |
| get medicine by ID | GET | /api/v1/medicines/:medID |
| update medicine by ID | PUT | /api/v1/medicines/:medID | *admin only* |
| delete medicine by ID | DELETE | /api/v1/medicines/:medID | *admin only* |

### 👤 Users

| name | method | route | note |
|:-----|:------:|-------|------|
| get user by ID | GET | /api/v1/users/:userID | *admin only* |
| signup user | POST | /api/v1/signup |
| login user | POST | /api/v1/login |
| verify user | GET | /api/v1/verify?token=token | *auto-generated* |
| update user | PUT | /api/v1/users |
| logout user | POST | /api/v1/logout |
| delete user | DELETE | /api/v1/users |
| refresh token | POST | /api/v1/refresh |
| revoke token | POST | /api/v1/revoke |

### 🔄 General

| name | method | route | note |
|------|--------|-------|------|
| reset DB | POST | /api/v1/reset | *dev* environment only |


## ✅ To-Do

- [ ] Explore more security options (mostly encryption on cookies and tokens).
- [ ] Build cmd or use config file for the variable setup in `constants.go`.


# 📖 Documentation

> Experimental

Run the project and see the documentation at `http://localhost:8080/api/v1/swagger/index.html`.

# 🚀 Status

![code coverage badge](https://github.com/Dhar01/medicine-app/actions/workflows/ci.yml/badge.svg)
