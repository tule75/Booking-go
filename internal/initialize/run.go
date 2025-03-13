package initialize

import (
	"context"
	"ecommerce_go/global"
	"ecommerce_go/internal/database"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Run() *gin.Engine {
	LoadConfig()
	fmt.Println("mysql ", global.Config.Mysql.Username)
	InitLogger()
	global.Logger.Info("Config ok roif: ", zap.String("Ok", "success"))
	InitMySQL()
	InitMySQLC()
	queries := database.New(global.Mdbc)
	InitRedis()
	InitKafka()
	InitServices(queries)
	r := InitRouter()

	go func() {
		ctx, cancel := context.WithCancel(context.Background())
		var wg sync.WaitGroup

		for i := 1; i <= 5; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				RegisterConsumer(ctx, "availability-group", queries)
			}(i)
		}

		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

		<-sigChan
		cancel()
		wg.Wait()
		fmt.Println("Kafka consumers shut down successfully.")
		os.Exit(0)
	}()

	return r

}
