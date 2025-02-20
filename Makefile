NAME=go-websocket

BUILD_DIR=build/

BIN_PATH=${BUILD_DIR}/${NAME}
TARGET_PLATFORM=windows linux darwin

TARGET=$(addprefix platform.,${TARGET_PLATFORM})

OS=$(shell uname>/dev/null||echo Windows)
RM= rm -rf
ifeq ($(OS), Windows)
	RM=cmd /c rmdir /s /q
endif

all: $(TARGET)

platform.windows:
	@GOOS=windows GOARCH=amd64 go build -o ${BIN_PATH}_windows.exe main.go
platform.linux:
	@GOOS=linux GOARCH=amd64 go build -o ${BIN_PATH}_linux main.go
platform.darwin:
	@GOOS=darwin GOARCH=amd64 go build -o ${BIN_PATH}_darwin main.go

clean:
	@$(RM) ${BUILD_DIR}||true
