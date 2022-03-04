package initialize

import (
	"gin-api-learn/global"
	"github.com/zhengpanone/gin-api-learn/server/utils"
	"path"
	"strings"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	// 日志输出格式
	outJson = "json"
)

//初始化日志
func InitLogger() {
	logConfig := global.GlobalConfig.Log
	// 判读日志目录是否存在
	if exist, _ := utils.DirExist(logConfig.Path); !exist {
		_ = utils.CreateDir(logConfig.Path)
	}
	// 设置输出格式
	var encoder zapcore.Encoder
	if logConfig.OutFormat == outJson {
		encoder = zapcore.NewJSONEncoder(getEncoderConfig())
	} else {
		encoder = zapcore.NewConsoleEncoder(getEncoderConfig())
	}
	// 设置日志文件切割
	writeSyncer := zapcore.AddSync(getLumberjackWriteSyncer())
	// 创建NewCore
	zapCore := zapcore.NewCore(encoder, writeSyncer, getLevel())
	// 创建logger
	logger := zap.New(zapCore, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	defer logger.Sync()
	// 赋值给全局变量
	global.GlobalLogger = logger

}

// 获取最低记录日志级别
func getLevel() zapcore.Level {
	levelMap := map[string]zapcore.Level{
		"debug":  zapcore.DebugLevel,
		"info":   zapcore.InfoLevel,
		"warn":   zapcore.WarnLevel,
		"error":  zapcore.ErrorLevel,
		"dpanic": zapcore.DPanicLevel,
		"fatal":  zapcore.FatalLevel,
	}
	if level, ok := levelMap[global.GlobalConfig.Log.Level]; ok {
		return level
	}
	// 默认info级别
	return zapcore.InfoLevel
}

// 自定义日志输出字段
func getEncoderConfig() zapcore.EncoderConfig {
	config := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     getEncodeTime,
		EncodeDuration: zapcore.StringDurationEncoder, // 自定义输出时间格式
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	return config
}

// 定义日志输出时间格式
func getEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006/01/02 - 15:04:05.000"))
}

// 获取文件切割和归档配置信息
func getLumberjackWriteSyncer() zapcore.WriteSyncer {
	lumberJackConfig := global.GlobalConfig.Log.LumberJack
	lumberJackLogger := &lumberjack.Logger{
		Filename:   getLogFile(),                // 日志文件
		MaxSize:    lumberJackConfig.MaxSize,    // 单文件最大容量(MB)
		MaxBackups: lumberJackConfig.MaxBackups, // 保留旧文件的最大数量
		MaxAge:     lumberJackConfig.MaxAge,     // 旧文件最多保存几天
		Compress:   lumberJackConfig.Compress,   // 是否压缩/归档旧文件
	}
	// 设置日志文件切割
	return zapcore.AddSync(lumberJackLogger)
}

// 获取日志文件路径
func getLogFile() string {
	fileFormat := time.Now().Format(global.GlobalConfig.Log.FileFormat)
	fileName := strings.Join([]string{
		global.GlobalConfig.Log.FilePrefix,
		fileFormat,
		"log"}, ".")
	return path.Join(global.GlobalConfig.Log.Path, fileName)
}
