package rgb_go_selenium

import (
	"os"

	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgbutil"
	"github.com/rs/zerolog/log"
	"github.com/tebeka/selenium"
)

const (
	SeleniumLogPath = "/home/matija/go/src/github.com/matijakrajnik/rgb_go_selenium/selenium.log"
	XGBLogPath      = "/home/matija/go/src/github.com/matijakrajnik/rgb_go_selenium/xgb.log"
)

// StartSelenium starts Selenium server. Log output is saved to SeleniumLogPath file.
func StartSelenium() *selenium.Service {
	logFile, err := os.OpenFile(SeleniumLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Panic().Err(err).Str("SeleniumLogPath", SeleniumLogPath).Msg("Error while opening Selenium log file.")
	}

	opts := []selenium.ServiceOption{
		selenium.GeckoDriver(geckoDriverPath),   // Specify the path to GeckoDriver in order to use Firefox.
		selenium.ChromeDriver(chromeDriverPath), // Specify the path to ChromeDriver in order to use Chrome.
		selenium.Output(logFile),                // Output debug information to selenium.log file.
	}

	service, err := selenium.NewSeleniumService(seleniumPath, conf.Port, opts...)
	if err != nil {
		log.Panic().Err(err).Msg("Can't start Selenium server.")
	}
	return service
}

func createFrameBuffer() *selenium.FrameBuffer {
	frameBuffer, err := selenium.NewFrameBuffer()
	if err != nil {
		log.Panic().Err(err).Msg("Can't create frame buffer.")
	}
	return frameBuffer
}

// ConnectToDisplay creates new frame buffer and connects to X server.
func ConnectToDisplay() *xgbutil.XUtil {
	frameBuffer := createFrameBuffer()
	logFile, err := os.OpenFile(XGBLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Panic().Err(err).Str("XGBLogPath", XGBLogPath).Msg("Error while opening XGB log file.")
	}
	xgb.Logger.SetOutput(logFile)
	display, err := xgbutil.NewConnDisplay(conf.DisplayAddress + ":" + frameBuffer.Display)
	if err != nil {
		log.Panic().Err(err).Msg("Can't connect to display.")
	}
	return display
}
