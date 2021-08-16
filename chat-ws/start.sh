set -e

echo "run db migration"
/app/migrate -path /app/migration -database mysql://root:password@"(mysql:3306)"/straperdb up

echo "start the app"
exec "$@"