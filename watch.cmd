@echo off
:: ----------------------------------------------------------------------
:: Rebuilds project automatically when any .go file updated.
:: To speed up, we won't call make to determine ldflags but run it only
:: once the first time.
::
:: Depends on https://github.com/watchexec/watchexec/releases
:: ----------------------------------------------------------------------
setlocal
set CC=fast-compile.cmd

if not exist %CC% (
	echo Generating %CC%
	echo @echo off > %CC%
	make build >> %CC%
)

start watchexec.exe --postpone --timings --exts go cmd /c %CC%
endlocal
