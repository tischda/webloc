@echo off
:: --------------------------------------------------------------------------
:: Build script
:: --------------------------------------------------------------------------
setlocal

go clean

rem get version from git tag
for /f "delims=" %%a in ('git describe --tags') do @set version=%%a

rem doc says not to use -s to remove the debug info read by gdb
rem but https://github.com/golang/go/issues/6245 says it's fixed
go build -ldflags "-X main.version %version% -s"

rem building with -race makes the binary huge!

endlocal