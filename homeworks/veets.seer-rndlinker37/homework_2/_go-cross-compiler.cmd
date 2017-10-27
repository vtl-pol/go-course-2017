@Echo off

rem	darwin 386
rem	darwin amd64
rem	darwin arm
rem	darwin arm64
rem	dragonfly amd64
rem	freebsd 386
rem	freebsd amd64
rem	freebsd arm
rem	linux 386
rem	linux amd64
rem	linux arm
rem	linux arm64
rem	linux ppc64
rem	linux ppc64le
rem	linux mips64
rem	linux mips64le
rem	linux s390x
rem	netbsd 386
rem	netbsd amd64
rem	netbsd arm
rem	openbsd 386
rem	openbsd amd64
rem	openbsd arm
rem	plan9 386
rem	plan9 amd64
rem	plan9 arm
rem	solaris amd64
rem	windows 386
rem	windows amd64

@set GOOS=windows
@set GOARCH=amd64

@set GOOGLE_API_KEY = 0
rem @del *%GOOS%-%GOARCH%*.*

go build

@setlocal enabledelayedexpansion
@setlocal enableextensions

@set /a n_files = 0

FOR /F "tokens=1-2 delims=." %%A IN ( 'dir /b /o:d /a:-d' ) DO ( call :subroutine %%A %%B )	

if %n_files% == 0 ( Echo No files to rename ) Else ( Echo %n_files% files done OK! )
goto :eof

	:subroutine
	set /a errs = 0
	if %2 == txt set /a errs = 1
	if %2 == tmp set /a errs = 1
	if %2 == bat set /a errs = 1
	if %2 == cmd set /a errs = 1
	if %2 == md  set /a errs = 1
	if %2 == go  set /a errs = 1
	if %errs% == 0 ( 

	set /a "n_files = n_files + 1"
	Echo Rename %1.%2 to %1-%GOOS%-%GOARCH%.%2
	ren %1.%2 %1-%GOOS%-%GOARCH%.%2
rem	@md _%GOOS%
rem	@xcopy /Y /Q	%1-%GOOS%-%GOARCH%.%2 _%GOOS%\
rem	@del /Q /F	%1-%GOOS%-%GOARCH%.%2 >nul
rem	@del /Q /F	%1.%2 >nul
        )

	exit /b

:eof
