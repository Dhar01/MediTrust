# medicine-app

A medicine e-commerce backend in Go.

## Environment Variables

Create a `.env` file in the root directory with the following variables:

```env
DB_URL="postgres://<username>:<password>@<host>:<port>/<database_name>"
PLATFORM="<environment>"
```

Currently available platform is `"deb"`.

Find `db_migration.sh` in the `scripts` to automate DB migration.

# API-endpoint

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