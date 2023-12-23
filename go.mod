module github.com/Stan370/Test-blog

go 1.19

require (
	github.com/gin-gonic/gin v1.9.1
	gorm.io/driver/mysql v1.5.2
	gorm.io/gorm v1.25.5
)
replace (
	github.com/Stan370/Test-blog => ../
)

