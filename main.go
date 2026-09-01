/*
 * Entrypoint:
 * 1. Asegura integridad de la base de datos.
 * 2. Inyecta dependencias.
 * 3. Registro explicito de cada servicio.
 */
package main

import (
	"Saavedra/env"
	loginMigration "Saavedra/migration/mLogin"
	loginRouter "Saavedra/service/Login/router"
	loginStore "Saavedra/service/Login/store"
	loginTypes "Saavedra/service/Login/types"
	assetsRouter "Saavedra/service/ServeAssets/router"
	usersRouter "Saavedra/service/Users/router"
	welcomeRouter "Saavedra/service/Welcome/router"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/glebarez/go-sqlite"
	"github.com/joho/godotenv"
)

func main() {

	// LOAD DOTENV FILE
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading '.env' file", err)
	}

	dataBase := os.Getenv("DATABASE_FILE_NAME")
	adminName := os.Getenv("ADMIN_NAME")
	adminEmail := os.Getenv("ADMIN_EMAIL")
	adminPassword := os.Getenv("ADMIN_PASSWORD")

	if adminName == "" || adminEmail == "" || adminPassword == "" {
		log.Fatal("There are missing admin credentials in '.env' file...")
	}

	// LOAD THOSE GLOBAL ENV CONFS TO BE USED IN PROJECT
	env.LoadGlobalEnvs()

	err = os.MkdirAll("./db/", 0755)
	if err != nil {
		log.Fatal("Error when generatin database path...", err)
	}

	db, err := sql.Open("sqlite", "./db/"+dataBase+".db")
	if err != nil {
		log.Fatal(err.Error())
	}

	// Registro de cada esquema
	loginMigration.CreateSchema(db)

	// ADMIN CREDENTIALS INJECTION
	adminCreds := loginTypes.User{
		Name:     adminName,
		Email:    adminEmail,
		Password: adminPassword,
	}
	if err = loginStore.InjectAdminDB(adminCreds, db); err != nil {
		log.Fatal("Admin credential injection fails...", err)
	}

	// Nuevo mapper de ruta-función para peticiones http.
	mux := http.NewServeMux()

	// Servir Estaticos (Alpine, Bulma-CSS, Lineicons-5) (Imgs, JS, etc...)
	assetsRouter.ServeGlobalAssetsStaticFiles(mux)

	// <== SERVICE REGISTRATION ==>

	loginRouter.Assambler(mux, db)
	welcomeRouter.Assambler(mux)
	usersRouter.Assambler(mux, db)

	// Servidor
	fmt.Println("🚀 Servidor ejecutándose en http://localhost:8080")
	http.ListenAndServe(":8080", mux)
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("El servidor se detuvo con error: %v", err)
	}
}
