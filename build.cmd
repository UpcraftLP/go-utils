@echo off

:: LIST PACKAGES TO BUILD ::
SET TO_BUILD=(datetime listener)

::temporarily build datetime package cause it is used in the build
go build -o datetime.exe ./datetime

:: SETUP ::
FOR /F %%i IN ('datetime') DO SET BUILD_TIME=%%i
ECHO Time is %BUILD_TIME%.

:: optional
::RMDIR "./build" /S /Q

SET APP_VERSION=%1
IF NOT DEFINED APP_VERSION (SET APP_VERSION=%VERSION%)
SET PREFIXED_APP_VERSION=v%APP_VERSION%
IF NOT DEFINED APP_VERSION (SET "APP_VERSION=development-preview" && SET "PREFIXED_APP_VERSION=development-preview" && ECHO No version defined, using "development-preview"!)

:: BUILD EXECUTABLES ::
FOR %%p IN %TO_BUILD% DO (
CALL .scripts/build_package %%p
)

:: remove temporary artifact
DEL datetime.exe
