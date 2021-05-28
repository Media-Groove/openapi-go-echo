@echo off

SET PKG_NAME=sampleapi
SET SPEC_FILE=sample\api.v1.yaml
SET DEST_DIR=..\sampleapi

pushd %~dp0
SET GO_POST_PROCESS_FILE="goimports -w"
mkdir generated
del /q /s generated
call java -jar openapi-generator-cli.jar generate -i %SPEC_FILE% -g go-server -o generated -t templates -p packageName=%PKG_NAME% %*
mkdir %DEST_DIR%
del /q %DEST_DIR%\*.*
copy generated\go\*.* %DEST_DIR%\
copy generated\api\openapi.yaml %DEST_DIR%\
call gosimports -w %DEST_DIR%
popd
