package main

import (
	"Graduation_Project/config"
	_ "Graduation_Project/docs"
	"Graduation_Project/iternal/routes"
	"Graduation_Project/iternal/tools"
	"Graduation_Project/migrate"
	"fmt"

	log "github.com/sirupsen/logrus"
)

func main() {
	//	@title		User API
	//	@version	1.0.0

	// 	@description 	Это выпускной проэкт с использованием свагера
	// 	@termsOfService  сдесь были бы условия использования еслиб я их мог обозначить
	// 	@contact.url    тут моя контактная информация (https://vk.com/id192672036)
	// 	@contact.email  tima.gorenskiy@mail.ru

	// 	@securityDefinitions.apikey  ApiKeyAuth
	//  @in header
	//  @name Authorization

	//	@host		localhost:8800

	//CheckFlagEnv Метод проверяющий флаги
	config.CheckFlagEnv()

	// вызываем логер
	err := tools.InitLogger()
	if err != nil {
		fmt.Println("ошибка при инициализаии логера")
	}

	// config.InitPgSQL инициализируем подключение к базе данных
	err = config.InitPgSQL()
	if err != nil {
		log.WithField("component", "initialization").Panic(err)
	}

	// вызываем миграцию структуры в базу данных
	migrate.Migrate()

	r := routes.SetupRouter()

	// запуск сервера
	if err = r.Run(config.Env.Host + ":" + config.Env.Port); err != nil {
		log.WithField("component", "run").Panic(err)
	}

}
