#!/bin/sh
set -e

echo "Aguardando MySQL..."
until mysqladmin ping -h"${DB_HOST}" -P"${DB_PORT}" -u"${DB_USER}" -p"${DB_PASSWORD}" --silent; do
  sleep 2
done

echo "Aplicando migracoes..."
for migration in /app/migrations/*.sql; do
  mysql -h"${DB_HOST}" -P"${DB_PORT}" -u"${DB_USER}" -p"${DB_PASSWORD}" "${DB_NAME}" < "${migration}"
done

echo "Iniciando aplicacao..."
exec /app/ordersystem
