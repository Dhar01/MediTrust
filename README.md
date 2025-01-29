# medicine-app

A medicine e-commerce backend in Go.


## SetUp

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


## Environment Variables

Create a `.env` file in the root directory with the following variables:

```env
DB_URL="postgres://<username>:<password>@<host>:<port>/<database_name>"
SECRET_KEY="Your Secret Key here"
PLATFORM="<environment>"
```
- Replace the `DB_URL` values with your database connection string.
- Currently available platform is `"deb"`.

Find `db_migration.sh` in the `scripts` to automate DB migration.

## To-Do

- [ ] Explore more security options (mostly encryption on cookies and tokens)

# API-endpoints

### Medicine

| name | method | route |
|:-----|:------:|:------|
| create medicine | POST | /api/v1/medicines |
| get all medicine | GET | /api/v1/medicines |
| get medicine by ID | GET | /api/v1/medicines/:medID |
| update medicine by ID | PUT | /api/v1/medicines/:medID |
| delete medicine by ID | DELETE | /api/v1/medicines/:medID |

### Users

| name | method | route |
|:-----|:------:|-------|
| create user | POST | /api/v1/users |
| get user by ID | GET | /api/v1/users/:userID |
| update user by ID | PUT | /api/v1/users/:userID |
| delete user by ID | DELETE | /api/v1/users/:userID |

# Status

![code coverage badge](https://github.com/Dhar01/medicine-app/actions/workflows/ci.yml/badge.svg)