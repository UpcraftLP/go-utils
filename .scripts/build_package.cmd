SET PACKAGE=%1
echo building %PACKAGE% %PREFIXED_APP_VERSION%
go build -o "bin/%APP_VERSION%/%PACKAGE%.exe" -ldflags="-s -w -X 'utils/internal/info.Name=%PACKAGE%' -X 'utils/internal/info.Version=%PREFIXED_APP_VERSION%' -X 'utils/internal/info.BuildTime=%BUILD_TIME%'" ./cmd/%PACKAGE%
