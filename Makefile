.PHONY: build-debug
build-debug: build-win
	@cp dist/win/BrowserSwitcher.exe /mnt/i/BrowserSwitcher.exe
	@#GOOS=windows go build -o /mnt/i/BrowserSwitcher.exe .
	@#cp scripts/register.ps1 /mnt/i/register.ps1
	@#cp scripts/reg.reg /mnt/i/reg.reg
	@#fyne-cross windows -arch=amd64 -app-id com.ft-t.browser-switcher
	@#cp fyne-cross/bin/windows-amd64/browser-switcher.exe /mnt/i/BrowserSwitcher.exe
## powershell -ExecutionPolicy Bypass -File register.ps1

.PHONY: lint
lint:
	@golangci-lint run

.PHONY: build
build: build-win build-linux

.PHONY: build-linux
build-linux:
	@mkdir -p dist
	@GOOS=linux go build -o dist/BrowserSwitcher .

.PHONY: build-win
build-win:
	@mkdir -p dist/win
	@GOOS=windows go build -o dist/win/BrowserSwitcher.exe .
	@cp scripts/reg.reg dist/win/reg.reg
