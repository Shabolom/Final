package migrate

import (
	"Graduation_Project/config"
	"Graduation_Project/iternal/domain"
	"github.com/gofrs/uuid"
	log "github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/gormigrate.v1"
)

// Migrate запустите миграцию для всех объектов и добавьте для них ограничения
// создаем таблицы и закидываем в бд тут
func Migrate() {
	db := config.DB
	regID, _ := uuid.NewV4()
	orderID, _ := uuid.NewV4()
	histID, _ := uuid.NewV4()
	// создаем объект миграции данная строка всегда статична (всегда такая)
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			// id всех миграций кторые были проведены
			ID: regID.String(),
			// переписываем так при создании таблицы изменяется только структура которую мы передаем
			Migrate: func(tx *gorm.DB) error {
				err := tx.AutoMigrate(&domain.Register{}).Error
				if err != nil {
					return err
				}
				return nil
			},
			// это метод отмены миграции ни разу не использовал
			Rollback: func(tx *gorm.DB) error {
				err := tx.DropTable("register").Error
				if err != nil {
					return err
				}
				return nil
			},
		}, {
			// id всех миграций кторые были проведены
			ID: orderID.String(),
			// переписываем так при создании таблицы изменяется только структура которую мы передаем
			Migrate: func(tx *gorm.DB) error {
				err := tx.AutoMigrate(&domain.UserOrder{}).Error
				if err != nil {
					return err
				}
				return nil
			},
			// это метод отмены миграции ни разу не использовал
			Rollback: func(tx *gorm.DB) error {
				err := tx.DropTable("order").Error
				if err != nil {
					return err
				}
				return nil
			},
		}, {
			// id всех миграций кторые были проведены
			ID: histID.String(),
			// переписываем так при создании таблицы изменяется только структура которую мы передаем
			Migrate: func(tx *gorm.DB) error {
				err := tx.AutoMigrate(&domain.History{}).Error
				if err != nil {
					return err
				}
				return nil
			},
			// это метод отмены миграции ни разу не использовал
			Rollback: func(tx *gorm.DB) error {
				err := tx.DropTable("history").Error
				if err != nil {
					return err
				}
				return nil
			},
		},
	})

	err := m.Migrate()
	if err != nil {
		log.WithField("component", "migration").Panic(err)
	}

	if err == nil {
		log.WithField("component", "migration").Info("Migration did run successfully")
	} else {
		log.WithField("component", "migration").Infof("Could not migrate: %v", err)
	}
}
