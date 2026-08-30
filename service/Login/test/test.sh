#!/bin/bash

# Testing "Login Service"
DIRECTION="http://localhost:8080/login"

echo "Server status at: $DIRECTION..."

# Usamos curl con -s -o /dev/null -w "%{http_code}" para obtener el código HTTP
# Esto verifica si el servidor responde (incluso si da 404 o 405, significa que la ruta/servidor existe)
HTTP_STATUS=$(curl -s -o /dev/null -w "%{http_code}" --connect-timeout 3 "$DIRECTION")

if [ "$HTTP_STATUS" -ne 000 ]; then
    echo "¡Servidor activo! (Código HTTP: $HTTP_STATUS). Ejecutando pruebas..."

    POSITIVE_LOGIN=$(
        curl -s
        -X POST
        -H "Content-Type: application/json"
        -d '{"email":"polar@gmail.com","password":"12345"}'
        "$DIRECTION"
    )

    NEGATIVE_LOGIN=$(
        curl -s
        -X POST
        -H "Content-Type: application/json"
        -d '{"email":"polar@gmail.com","password":"123456"}'
        "$DIRECTION"
    )

    # Validar login positivo
    if [[ -z "$POSITIVE_LOGIN" ]]; then
        echo -e "Success: POSITIVE_LOGIN respondió correctamente."
    else
        echo -e "Error: POSITIVE_LOGIN está vacío."
    fi

    # Validar login negativo
    if [[ -z "$NEGATIVE_LOGIN" ]]; then
        echo -e "Success: NEGATIVE_LOGIN bloqueado correctamente."
    else
        echo -e "Error: NEGATIVE_LOGIN no devolvió lo esperado ($NEGATIVE_LOGIN)."
    fi

else
    echo "Error: El servidor no está corriendo o no se puede alcanzar en $DIRECTION"
fi
