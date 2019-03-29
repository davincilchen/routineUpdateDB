

type User struct {
	gorm.Model
	PlatformID uint
	Platform   *Platform
	MemberID   string   `gorm:"type:varchar(64)"`
	Password   string   `gorm:"type:char(32)"`
	Balance    uint64   `gorm:"type:bigint unsigned"`
	Role       UserRole `gorm:"type:tinyint unsigned"`
}




	for i := 0; i < 100; i++ {
		go func(modelIO ModelIO, i int) {

			userAfter := models.User{}
			userBefore := models.User{}
			modelDB, _ := modelIO.(*modeldb.DBIO)
			db := modelDB.DB
			tx := db.Begin()
			tx.Set("gorm:query_option", "FOR UPDATE").Model(&userBefore).
				First(&userBefore, userLogin.Info.ID).
				Where("balance > ?", ps.Bet).
				Update("balance", gorm.Expr("balance - ?", ps.Bet)).
				First(&userAfter, userLogin.Info.ID)

			log.Println(i, "af", userAfter)
			log.Println(i, "bf", userBefore)
			tx.Commit()
		}(modelIO, i)
	}