$binaryPath = "I:\browser-switcher.exe"

if (-not (Test-Path $binaryPath)) {
    Write-Error "The specified binary path does not exist: $binaryPath"
    exit 1
}

# Define application name
$appName = "Browser Switcher"

# Create registry entries for the application
New-Item -Path "HKCR:\$appName" -Force | Out-Null
Set-ItemProperty -Path "HKCR:\$appName" -Name "FriendlyTypeName" -Value "Browser Switcher" -Force

# Set the command to handle URLs
New-Item -Path "HKCR:\$appName\shell\open\command" -Force | Out-Null
Set-ItemProperty -Path "HKCR:\$appName\shell\open\command" -Name "Browser Switcher" -Value "`"$binaryPath`" `"%1`"" -Force

# Register for HTTP protocol
New-Item -Path "HKCR:\http\shell\open\command" -Force | Out-Null
Set-ItemProperty -Path "HKCR:\http\shell\open\command" -Name "Browser Switcher" -Value "`"$binaryPath`" `"%1`"" -Force

# Register for HTTPS protocol
New-Item -Path "HKCR:\https\shell\open\command" -Force | Out-Null
Set-ItemProperty -Path "HKCR:\https\shell\open\command" -Name "Browser Switcher" -Value "`"$binaryPath`" `"%1`"" -Force

# Add to Default Programs
$progIdPath = "HKLM:\Software\Clients\StartMenuInternet\$appName"
New-Item -Path $progIdPath -Force | Out-Null
Set-ItemProperty -Path $progIdPath -Name "Browser Switcher" -Value "Browser Switcher" -Force

# Set icon for the application
New-Item -Path "$progIdPath\DefaultIcon" -Force | Out-Null
Set-ItemProperty -Path "$progIdPath\DefaultIcon" -Name "Browser Switcher" -Value "`"$binaryPath`",0" -Force

# Register protocol handlers
New-Item -Path "$progIdPath\Capabilities\URLAssociations" -Force | Out-Null
Set-ItemProperty -Path "$progIdPath\Capabilities\URLAssociations" -Name "http" -Value "$appName" -Force
Set-ItemProperty -Path "$progIdPath\Capabilities\URLAssociations" -Name "https" -Value "$appName" -Force

# Notify Windows of the changes
Start-Process -FilePath "control.exe" -ArgumentList "/name Microsoft.DefaultPrograms /page pageDefaultProgram"
