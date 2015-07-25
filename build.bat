@echo off
:: --------------------------------------------------------------------------
:: Build script
:: --------------------------------------------------------------------------
setlocal
for /f "delims=" %%a in ('git describe --tags') do @set version=%%a
go build -ldflags "-X main.version %version% -s"
endlocal