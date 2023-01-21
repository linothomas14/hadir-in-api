package main

import (
	"time"

	"github.com/linothomas14/hadir-in-api/controller"
	"github.com/linothomas14/hadir-in-api/middleware"
	"github.com/linothomas14/hadir-in-api/repository"
	"github.com/linothomas14/hadir-in-api/service"

	"github.com/gin-gonic/gin"
	"github.com/linothomas14/hadir-in-api/config"

	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()

	userRepository       repository.UserRepository       = repository.NewUserRepository(db)
	eventRepository      repository.EventRepository      = repository.NewEventRepository(db)
	attendanceRepository repository.AttendanceRepository = repository.NewAttendanceRepository(db)
	// 	transactionRepository  repository.TransactionRepository  = repository.NewTransactionRepository(db)
	// jwtService  service.JWTService  = service.NewJWTService()
	userService       service.UserService       = service.NewUserService(userRepository)
	authService       service.AuthService       = service.NewAuthService(userRepository)
	jwtService        service.AuthService       = service.NewAuthService(userRepository)
	eventService      service.EventService      = service.NewEventService(eventRepository)
	attendanceService service.AttendanceService = service.NewAttendanceService(attendanceRepository, userRepository, eventRepository)

	// 	transactionService  service.TransactionService  = service.NewTransactionService(transactionRepository, productRepository)

	authController       controller.AuthController       = controller.NewAuthController(authService)
	userController       controller.UserController       = controller.NewUserController(userService)
	eventController      controller.EventController      = controller.NewEventController(eventService)
	attendanceController controller.AttendanceController = controller.NewAttendanceController(attendanceService)
)

func PingHandler(c *gin.Context) {
	t := time.Now()
	c.JSON(200, gin.H{
		"msg":  "pong",
		"time": t,
	})
}

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	r.POST("/login", authController.Login)
	r.POST("/register", authController.Register)

	userRoutes := r.Group("users", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.GetProfile)
		userRoutes.PUT("/", userController.Update)
	}

	eventRoutes := r.Group("events", middleware.AuthorizeJWT(jwtService))
	{
		eventRoutes.GET("/", PingHandler) //Get events for spesific user account who login
		eventRoutes.GET("/:idEvent", PingHandler)
		eventRoutes.POST("/", eventController.CreateEvent)
		eventRoutes.PUT("/:idEvent", PingHandler)
		eventRoutes.DELETE("/:idEvent", PingHandler)
	}
	attendanceRoutes := r.Group("attendances", middleware.AuthorizeJWT(jwtService))
	{
		attendanceRoutes.GET("/", PingHandler)
		attendanceRoutes.POST("/present", attendanceController.Attend)
		attendanceRoutes.DELETE("/:idEvent", PingHandler)
	}
	r.GET("ping", PingHandler)
	r.Run()
}
