.PHONY: build-debug
build-debug:
	@GOOS=windows go build -o /mnt/i/BrowserSwitcher.exe .
	@#cp scripts/register.ps1 /mnt/i/register.ps1
	@cp scripts/reg.reg /mnt/i/reg.reg
	@#fyne-cross windows -arch=amd64 -app-id com.ft-t.browser-switcher
	@#cp fyne-cross/bin/windows-amd64/browser-switcher.exe /mnt/i/BrowserSwitcher.exe
## powershell -ExecutionPolicy Bypass -File register.ps1
