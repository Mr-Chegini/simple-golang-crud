# PostgreSQL Setup Guide

This project now uses PostgreSQL instead of SQLite. Follow these steps to set up your local development environment.

## Installation

### On Linux (Ubuntu/Debian):
```bash
sudo apt-get update
sudo apt-get install postgresql postgresql-contrib
```

### On macOS:
```bash
brew install postgresql
```

### On Windows:
Download and install from: https://www.postgresql.org/download/windows/

## Starting PostgreSQL

### On Linux/macOS:
```bash
# Start PostgreSQL service
sudo systemctl start postgresql  # Linux
brew services start postgresql  # macOS

# Connect to PostgreSQL
sudo -u postgres psql
```

### On Windows:
PostgreSQL service should start automatically after installation.

## Creating the Development Database

1. Connect to PostgreSQL as the postgres user:
```bash
sudo -u postgres psql
```

2. Create the database:
```sql
CREATE DATABASE simple_crud;
```

3. Create a user (optional, for security):
```sql
CREATE USER simple_crud_user WITH PASSWORD 'your_password';
ALTER ROLE simple_crud_user SET client_encoding TO 'utf8';
ALTER ROLE simple_crud_user SET default_transaction_isolation TO 'read committed';
ALTER ROLE simple_crud_user SET default_transaction_deferrable TO on;
GRANT ALL PRIVILEGES ON DATABASE simple_crud TO simple_crud_user;
```

4. Exit psql:
```sql
\q
```

## Environment Configuration

Set the `DATABASE_URL` environment variable (optional - defaults to localhost):

```bash
# Default (development):
# host=localhost user=postgres password=postgres dbname=simple_crud port=5432 sslmode=disable

# Or set custom URL:
export DATABASE_URL="host=localhost user=simple_crud_user password=your_password dbname=simple_crud port=5432 sslmode=disable"
```

## Verifying the Connection

1. Run the application:
```bash
export PATH="/usr/local/go/bin:$PATH"
go run cmd/api/main.go
```

2. You should see:
```
✅ Database connection established successfully
✅ Database migrations completed successfully
🚀 Server starting on port 8080...
```

3. Check tables in PostgreSQL:
```bash
sudo -u postgres psql -d simple_crud -c "\dt"
```

You should see `users` and `products` tables created by GORM migrations.

## Troubleshooting

### Connection refused
- Make sure PostgreSQL is running: `sudo systemctl status postgresql`
- Check if port 5432 is open: `netstat -an | grep 5432`

### "FATAL: Ident authentication failed"
- Edit `/etc/postgresql/{version}/main/pg_hba.conf`
- Change `peer` to `md5` for local connections
- Restart PostgreSQL: `sudo systemctl restart postgresql`

### Permission denied
- Make sure you have the correct username and password
- Check user privileges: `GRANT ALL PRIVILEGES ON DATABASE simple_crud TO username;`

## Benefits of PostgreSQL over SQLite

1. **Concurrency**: Handles multiple concurrent connections better
2. **Scalability**: Better for larger datasets
3. **Features**: JSON/JSONB support, Arrays, Full-text search, etc.
4. **Reliability**: ACID compliance, better transaction support
5. **Production-Ready**: Used in most production environments

Happy coding! 🚀