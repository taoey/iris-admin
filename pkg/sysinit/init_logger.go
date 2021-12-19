package sysinit

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	. "github.com/taoey/iris-admin/pkg"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() {
	// 日志文件保存及分割：按照日志大小进行分割
	//hook := lumberjack.Logger{
	//	Filename:   GCF.UString("logger.path", "./logs/log"), // 日志文件路径
	//	MaxSize:    5,                                        // 每个日志文件保存的最大尺寸 单位：M
	//	MaxBackups: 30,                                       // 日志文件最多保存多少个备份
	//	MaxAge:     7,                                        // 文件最多保存多少天
	//	Compress:   true,                                     // 是否压缩
	//}

	// 日志格式
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "line",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalLevelEncoder, // level 大写编码器
		//EncodeTime:     zapcore.RFC3339TimeEncoder,     //  时间格式
		EncodeTime:     TimeEncoder, //  自定义时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 路径编码器  FullCallerEncoder:可以点击进入日志行 ShortCallerEncoder：不可进入
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	var logLevel zapcore.Level
	switch strings.ToUpper(GCF.UString("logger.level", "INFO")) {
	case "DEBUG":
		logLevel = zap.DebugLevel
	case "INFO":
		logLevel = zap.InfoLevel
	case "WARNING", "WARN":
		logLevel = zap.WarnLevel
	case "ERROR":
		logLevel = zap.ErrorLevel
	case "CRITICAL", "FATAL":
		logLevel = zap.FatalLevel
	default:
		logLevel = zap.InfoLevel
	}
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(logLevel)

	// 编码器配置 json/console
	var styleEncoder zapcore.Encoder
	if GCF.UString("logger.style", "console") == "console" {
		styleEncoder = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		styleEncoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	filelogWriter := getWriter(GCF.UString("logger.path", "./logs/log"))
	core := zapcore.NewCore(
		styleEncoder,
		//zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(filelogWriter)), // 打印到控制台和文件
		atomicLevel, // 日志级别
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	//filed := zap.Fields(zap.String("serviceName", "serviceName"))
	// 构造日志
	//logger := zap.New(core, caller, development, )
	logger := zap.New(core, caller, development)

	//赋值给全局sugar
	Log = logger.Sugar()
}

// 自定义日期样式
func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("[%v]", t.Format("2006-01-02 15:04:05")))
}

// 日志分割器：按照日期进行分割
// 参考自：https://www.cnblogs.com/ExMan/p/12264925.html
func getWriter(filename string) io.Writer {
	// 生成rotatelogs的Logger 实际生成的文件名 demo.log.YYmmddHH
	// demo.log是指向最新日志的链接
	// 保存7天内的日志，每1小时(整点)分割一次日志
	hook, err := rotatelogs.New(
		filename+".%Y%m%d%H", // 没有使用go风格反人类的format格式
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour*4),
	)

	if err != nil {
		panic(err)
	}
	return hook
}
