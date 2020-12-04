SET PACKAGE=%1
echo building %PACKAGE% %PREFIXED_APP_VERSION%
go build -o "build/%APP_VERSION%/%PACKAGE%.exe" -ldflags="-s -w -X 'utils/src/info.Name=%PACKAGE%' -X 'utils/src/info.Version=%PREFIXED_APP_VERSION%' -X 'utils/src/info.BuildTime=%BUILD_TIME%'" ./%PACKAGE%
