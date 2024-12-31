.PHONY: build-debug
build-debug: build-win
	@cp dist/win/BrowserSwitcherProxied.exe /mnt/i/BrowserSwitcherProxied.exe
	@cp dist/win/BrowserSwitcher.exe /mnt/i/BrowserSwitcher.exe
	@cp dist/win/register.ps1 /mnt/i/register.ps1
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
	@GOOS=linux go build -buildvcs=false -o dist/BrowserSwitcher .

.PHONY: build-win
build-win:
	@mkdir -p dist/win
	@go install github.com/tc-hib/go-winres@latest
	@GOOS=windows go build -buildvcs=false -o dist/win/BrowserSwitcherProxied.exe cmd/switcher/main.go
	@GOOS=windows go build -ldflags -H=windowsgui -buildvcs=false -o dist/win/BrowserSwitcher.exe cmd/proxy/main.go cmd/proxy/win.go
	@go-winres patch --in winres/winres.json --no-backup dist/win/BrowserSwitcherProxied.exe
	@go-winres patch --in winres/winres.json --no-backup dist/win/BrowserSwitcher.exe
	@cp scripts/register.ps1 dist/win/register.ps1
