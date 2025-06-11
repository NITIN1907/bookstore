#!/bin/sh

# Default values if env vars not provided
: "${DB_HOST:=host.docker.internal}"
: "${DB_PORT:=5432}"

until nc -z "$DB_HOST" "$DB_PORT"; do
  echo "$(date) 🕒 Postgres is unavailable - waiting..."
  sleep 1
done

echo "$(date) ✅ Postgres is up - executing command"
exec "$@"
