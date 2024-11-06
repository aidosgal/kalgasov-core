package main

import (
	"log"

	"github.com/aidosgal/kalgasov-core/cmd/kalgasov"
	"github.com/aidosgal/kalgasov-core/config"
	"github.com/aidosgal/kalgasov-core/database"
	"github.com/go-sql-driver/mysql"
)

func main() {
    db, err := database.NewMySQLStorage(mysql.Config{
        User:                   config.Envs.DBUser,
        Passwd:                 config.Envs.DBPassword,
        Addr:                   config.Envs.DBAddress,
        DBName:                 config.Envs.DBName,

        Net:                    "tcp",
        AllowNativePasswords:   true,
        ParseTime:              true,
    })
    if err != nil {
        log.Fatal(err)
    }

    server := kalgasov.NewAPIServer(config.Envs.Port, db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}    
}
