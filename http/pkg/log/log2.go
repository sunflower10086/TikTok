package log

import (
	"io"
	"os"
	"time"

	config "github.com/sunflower10086/TikTok/http/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var SugarLogger *zap.SugaredLogger

// ZapInterceptor 返回zap.logger实例(把日志写到文件中)
func ZapInterceptor(conf *config.Config) *zap.Logger {
	config := zapcore.EncoderConfig{
		MessageKey:   "msg",                       //结构化（json）输出：msg的key
		LevelKey:     "level",                     //结构化（json）输出：日志级别的key（INFO，WARN，ERROR等）
		TimeKey:      "ts",                        //结构化（json）输出：时间的key（INFO，WARN，ERROR等）
		CallerKey:    "file",                      //结构化（json）输出：打印日志的文件对应的Key
		EncodeLevel:  zapcore.CapitalLevelEncoder, //将日志级别转换成大写（INFO，WARN，ERROR等）
		EncodeCaller: zapcore.ShortCallerEncoder,  //采用短文件路径编码输出（test/main.go:14 ）
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		}, //输出的时间格式
		//EncodeTime: zapcore.ISO8601TimeEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		}, //
	}

	// 取出来配置文件中的loglevel
	var logLevel zapcore.Level
	err := logLevel.UnmarshalText([]byte(conf.Log.Level))
	if err != nil {

		return nil
	}

	//自定义日志级别：自定义debug级别
	debugLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == logLevel
	})

	//自定义日志级别：自定义Info级别
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.WarnLevel && lvl > logLevel
	})

	//自定义日志级别：自定义Warn级别
	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})

	infoWriter := getWriter(
		conf.Log.InfoFilename,
		conf.Log.MaxSize,
		conf.Log.MaxBackups,
		conf.Log.MaxAge,
	)
	warnWriter := getWriter(
		conf.Log.ErrFilename,
		conf.Log.MaxSize,
		conf.Log.MaxBackups,
		conf.Log.MaxAge,
	)

	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewConsoleEncoder(config), zapcore.AddSync(infoWriter), infoLevel), //将info及以下写入logPath，NewConsoleEncoder 是非结构化输出
		zapcore.NewCore(zapcore.NewConsoleEncoder(config), zapcore.AddSync(warnWriter), warnLevel), //warn及以上写入errPath
		zapcore.NewCore(zapcore.NewJSONEncoder(config), zapcore.AddSync(os.Stdout), debugLevel),    //把debug级别的日志写入控制台
	)
	logger := zap.New(core, zap.AddCaller())
	SugarLogger = logger.Sugar()

	// 替换全局的logger
	// 使用的话就zap.L().
	zap.ReplaceGlobals(logger)

	return logger
}

func getWriter(filename string, maxSize, maxBackups, maxAge int) io.Writer {
	return &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,    //最大M数，超过则切割，log文件最大大小
		MaxBackups: maxBackups, //最大文件保留数，超过就删除最老的日志文件，备份数量
		MaxAge:     maxAge,     //保存30天，备份天数
		Compress:   false,      //是否压缩
	}
}
