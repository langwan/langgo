package log

import (
	"github.com/langwan/langgo/core"
	"github.com/langwan/langgo/helpers/io"
	helperString "github.com/langwan/langgo/helpers/string"
	"github.com/rs/zerolog"
	"log"
	"os"
	"testing"
	"time"
)

func TestLogger(t *testing.T) {
	Logger("app", "test").Info().Msg("ok")
}

func BenchmarkLoggerSystemLog(b *testing.B) {
	io.CreateFolder("logs", true)
	logfile := "logs/syslog.log"
	os.Remove(logfile)
	f, _ := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer f.Close()
	log.SetOutput(f)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Println("INF ok tag=test")
	}
}

func BenchmarkZerologFile(b *testing.B) {
	io.CreateFolder("logs", true)
	logfile := "logs/zerolog.log"
	os.Remove(logfile)
	f, _ := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer f.Close()
	l := zerolog.New(f).With().Str("tag", "test").Timestamp().Logger()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Info().Msg("ok")
	}
}

func BenchmarkZerologConsole(b *testing.B) {
	io.CreateFolder("logs", true)
	logfile := "logs/zerolog-console.log"
	os.Remove(logfile)
	f, _ := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	zc := zerolog.ConsoleWriter{Out: f, TimeFormat: time.RFC3339, NoColor: true}
	defer f.Close()
	l := zerolog.New(zc).With().Str("tag", "test").Timestamp().Logger()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Info().Msg("ok")
	}
}

func BenchmarkLoggerLanggo(b *testing.B) {
	core.EnvName = core.Production
	io.CreateFolder("logs", true)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Logger("langgo", "test").Info().Msg("ok")
	}
}

func BenchmarkLoggerLanggoMulti(b *testing.B) {
	core.EnvName = core.Production
	io.CreateFolder("logs", true)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s, _ := helperString.RandString(2, helperString.LettersNumberNoZero)
		Logger(s, "test").Info().Msg("ok")
	}
}
