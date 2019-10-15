package mem

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/howeyc/fsnotify"
	"github.com/jinzhu/gorm"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"goweb/localApp/fileWatcher"
	"math/big"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

//constant
var Wg sync.WaitGroup

//Db
var Mysql *gorm.DB = nil
var Postgres *gorm.DB = nil

//Gin
var G *gin.Engine = nil
var HttpServer *http.Server = nil

//Ethereum
var EthPrivateKey *ecdsa.PrivateKey
var EthAddress common.Address
var ChainId *big.Int
var EthNilAddress = [common.AddressLength]byte{}
const MaxEthTransBuf = 1024
var EthClient *ethclient.Client = nil

//Logurus
var logrusWriter *rotatelogs.RotateLogs = nil

//FileWatch
var fileWatcherTest *fsnotify.Watcher = nil

func init() {
	initResource()
	initLog()
	initDB()
	initGin()
	initSignal()
	initFileWatchTest()
}

func destruct() {
	if Mysql != nil {
		Mysql.Close()
		logrus.Println("Mysql Closed")
	}

	if Postgres != nil {
		Postgres.Close()
		logrus.Println("Postgres Closed")
	}

	if EthClient != nil {
		EthClient.Close()
		logrus.Println("EthClient Closed")
	}

	if HttpServer != nil {
		HttpServer.Shutdown(context.Background())
		logrus.Println("HttpServer Closed")
	}

	if fileWatcherTest != nil {
		fileWatcherTest.Close()
		logrus.Println("FileWatcherTest Closed")
	}

	if logrusWriter != nil {
		logrusWriter.Close()
		fmt.Println("Logrus Closed")
	}
}

func initSignal() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		sig := <-sigs
		logrus.Println("received:", sig)
		destruct()
		os.Exit(1)
	}()
	Wg.Add(1)
}

func initResource() {
	configType := "yml"
	configPath := "./resource"
	v := viper.New()
	v.SetConfigName("default")
	v.AddConfigPath(configPath)
	v.SetConfigType(configType)
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	configs := v.AllSettings()
	for k, v := range configs {
		viper.SetDefault(k, v)
	}
	env := viper.GetString("env")
	if env == "" {
		env = "test"
	}

	viper.SetConfigName(env)
	viper.AddConfigPath(configPath)
	viper.SetConfigType(configType)
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initLog() {
	var err error = nil
	logrusWriter, err = rotatelogs.New(
		"./logs/"+"%Y-%m-%d",
		rotatelogs.WithMaxAge(365*24*time.Hour),   // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)

	if err != nil {
		panic(err)
	}

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: logrusWriter, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  logrusWriter,
		logrus.WarnLevel:  logrusWriter,
		logrus.ErrorLevel: logrusWriter,
		logrus.FatalLevel: logrusWriter,
		logrus.PanicLevel: logrusWriter,
	}, &logrus.JSONFormatter{})

	logrus.AddHook(lfHook)
	logrus.SetReportCaller(true)
}

func initDB() {
	var err error = nil
	url := viper.GetString("mysql.url")
	if url != "" {
		Mysql, err = gorm.Open("mysql", url)
		if err != nil {
			panic(err)
		}
	}

	url = viper.GetString("postgres.url")
	if url != "" {
		Mysql, err = gorm.Open("postgres", url)
		if err != nil {
			panic(err)
		}
	}
}

func initGin() {
	if viper.GetBool("http") {
		gin.SetMode(gin.ReleaseMode)
		G = gin.New()
	}
}

func initFileWatchTest() {
	watchDir := viper.GetString("watchDirectory")
	var err error

	if watchDir != "" {
		fileWatcherTest, err = fsnotify.NewWatcher()
		if err != nil {
			panic(err)
		}

		err = fileWatcherTest.WatchFlags(watchDir, fsnotify.FSN_CREATE |  fsnotify.FSN_MODIFY)
		if err != nil {
			panic(err)
		}

		go func() {
			for {
				select {
				case ev := <-fileWatcherTest.Event:
					fileWatcher.ProcessCreate(ev)
				case err := <-fileWatcherTest.Error:
					logrus.Println("error:", err)
				}
			}
		}()
	}


}
