package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

const (
	defaultLogPath           = "./zaplogs"          // 【默认】日志文件目录
	defaultLogFilename       = "jar_encryption.log" // 【默认】日志文件名称
	defaultLogLevel          = "info"               // 【默认】日志打印级别 debug  info  warning  error
	defaultLogFileMaxSize    = 10                   // 【日志分割】  【默认】单个日志文件最多存储量 单位(mb)
	defaultLogFileMaxBackups = 10                   // 【日志分割】  【默认】日志备份文件最多数量
	logMaxAge                = 100                  // 【默认】日志保留时间，单位: 天 (day)
	logCompress              = true                 // 【默认】是否压缩日志
	logStdout                = true                 // 【默认】是否输出到控制台
)

func Log() *zap.SugaredLogger {
	return logger
}
func LogFile() string {
	return filepath.Join(*logPath, *logFilename)
}

var logger *zap.SugaredLogger // 定义日志打印全局变量

var (
	// kingpin 可以在启动时通过输入参数，来修改日志参数

	level             = kingpin.Flag("log.level", "only log messages with the given severity or above. one of: [debug, info, warn, error]").Default(defaultLogLevel).String()
	format            = kingpin.Flag("log.format", "output format of log messages. one of: [logfmt, json]").Default("logfmt").String()
	logPath           = kingpin.Flag("log.path", "output log path").Default(defaultLogPath).String()
	logFilename       = kingpin.Flag("log.filename", "output log filename").Default(defaultLogFilename).String()
	logFileMaxSize    = kingpin.Flag("log.file-max-size", "output logfile max size, unit mb").Default(strconv.Itoa(defaultLogFileMaxSize)).Int()
	logFileMaxBackups = kingpin.Flag("log.file-max-backups", "output logfile max backups").Default(strconv.Itoa(defaultLogFileMaxBackups)).Int()
)

// 初始化 logger
func init() {

	kingpin.Parse()
	loglevel := map[string]zapcore.Level{
		"debug": zapcore.DebugLevel,
		"info":  zapcore.InfoLevel,
		"warn":  zapcore.WarnLevel,
		"error": zapcore.ErrorLevel,
	}
	writeSyncer, err := getLogWriter() // 日志文件配置 文件位置和切割
	if err != nil {
		log.Fatal("日志初始化失败：", err)
	}
	encoder := getencoder()       // 获取日志输出编码
	level, ok := loglevel[*level] // 日志打印级别
	if !ok {
		level = loglevel["info"]
	}
	core := zapcore.NewCore(encoder, writeSyncer, level)
	temp := zap.New(core, zap.AddCaller()) //  zap.AddCaller() 输出日志打印文件和行数如： logger/logger_test.go:33
	logger = temp.Sugar()
}

// 编码器(如何写入日志)
func getencoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000") // logger 时间格式 例如: 2021-09-11t20:05:54.852+0800
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder                           // 输出level序列化为全大写字符串，如 info debug error

	if *format == "json" {
		return zapcore.NewJSONEncoder(encoderConfig) // 以json格式写入
	}
	return zapcore.NewConsoleEncoder(encoderConfig) // 以logfmt格式写入
}

// 获取日志输出方式  日志文件 控制台
func getLogWriter() (zapcore.WriteSyncer, error) {
	// 判断日志路径是否存在，如果不存在就创建
	if exist := isExist(*logPath); !exist {
		if *logPath == "" {
			*logPath = defaultLogPath
		}
		if err := os.MkdirAll(*logPath, os.ModePerm); err != nil {
			*logPath = defaultLogPath
			if err := os.MkdirAll(*logPath, os.ModePerm); err != nil {
				return nil, err
			}
		}
	}
	// 日志文件 与 日志切割 配置
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filepath.Join(*logPath, *logFilename), // 日志文件路径
		MaxSize:    *logFileMaxSize,                       // 单个日志文件最大多少 mb
		MaxBackups: *logFileMaxBackups,                    // 日志备份数量
		MaxAge:     logMaxAge,                             // 日志最长保留时间
		Compress:   logCompress,                           // 是否压缩日志
	}
	if exist := isExist(lumberJackLogger.Filename); !exist {
		_, _ = os.OpenFile(lumberJackLogger.Filename, os.O_APPEND|os.O_CREATE, os.ModePerm)
	}
	if logStdout {
		// 日志同时输出到控制台和日志文件中
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberJackLogger), zapcore.AddSync(os.Stdout)), nil
	} else {
		// 日志只输出到日志文件
		return zapcore.AddSync(lumberJackLogger), nil
	}
}

// 判断文件或者目录是否存在
func isExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
