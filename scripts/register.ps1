$binaryName = "BrowserSwitcher"
$appPath = "I:\$binaryName.exe"
$proxiedPath = "I:\${binaryName}Proxied.exe"

$AppLongName = "Browser Switcher"
$AppRegDescription = "Browser Switcher"
$CustomProtoName = "x-browser-switcher"

function Get-BrowserRegistrationPath {
    return "HKCU:\Software\Clients\StartMenuInternet\$AppLongName"
}

function Register-Browser {
    $softPath = "Software\Clients\StartMenuInternet\$AppLongName"
    $appRoot = "HKCU:\Software\Clients\StartMenuInternet\$AppLongName"
    $capRoot = "$appRoot\Capabilities"

    New-Item -Path $appRoot -Force
    New-Item -Path $capRoot -Force

    # Add basic registration
    Set-ItemProperty -Path $appRoot -Name "(Default)" -Value $AppLongName
    Set-ItemProperty -Path $capRoot -Name "ApplicationName" -Value $AppLongName
    Set-ItemProperty -Path $capRoot -Name "ApplicationDescription" -Value $AppRegDescription
    Set-ItemProperty -Path $capRoot -Name "ApplicationIcon" -Value "$appPath,0"

    # Supported protocols
    $urlAssocPath = "$capRoot\URLAssociations"
    New-Item -Path $urlAssocPath -Force
    foreach ($protocol in @("https", "http", $CustomProtoName)) {
        Set-ItemProperty -Path $urlAssocPath -Name $protocol -Value "BrowserSwitcherHTM"
    }

    # File associations
    $fileAssocPath = "$capRoot\FileAssociations"
    if (-not (Test-Path -Path $fileAssocPath)) {
        New-Item -Path $fileAssocPath -Force
    }

    # Define file extensions for HTML and related files
    $htmlExtensions = @(".svg", ".htm", ".html", ".shtml", ".webp", ".xht", ".xhtml", ".mht", ".mhtml", ".pdf")
    foreach ($ext in $htmlExtensions) {
        Set-ItemProperty -Path $fileAssocPath -Name $ext -Value "BrowserSwitcherHTM"
    }

    # Command registration
    New-Item -Path "$appRoot\DefaultIcon" -Force
    Set-ItemProperty -Path "$appRoot\DefaultIcon" -Name "(Default)" -Value "$appPath,0"

    New-Item -Path "$appRoot\shell\open\command" -Force

    Set-ItemProperty -Path "$appRoot\shell\open\command" -Name "(Default)" -Value "`"$appPath`""
    Set-ItemProperty -Path "$appRoot\shell\open\command" -Name "Proxied" -Value "$proxiedPath"

    # Register capabilities
    Set-ItemProperty -Path "HKCU:\Software\RegisteredApplications" -Name $AppLongName -Value "$softPath\Capabilities"
}

function Register-Protocol {
    $root = "HKCU:\Software\Classes\$CustomProtoName"

    $registrationPath = @("Registry::HKEY_CLASSES_ROOT\BrowserSwitcherHTM", "Registry::HKEY_CLASSES_ROOT\BrowserSwitcherPDF", "Registry::HKEY_CURRENT_USER\Software\Classes\BrowserSwitcherHTM", "Registry::HKEY_CURRENT_USER\Software\Classes\BrowserSwitcherPDF")
    foreach ($regPath in $registrationPath) {
        New-Item -Path "$regPath" -Force

        Set-ItemProperty -Path "$regPath" -Name "(Default)" -Value "Browser Switcher HTML Document"

        New-Item -Path "$regPath\Application" -Force
        New-Item -Path "$regPath\DefaultIcon" -Force

        New-Item -Path "$regPath\shell" -Force
        New-Item -Path "$regPath\shell\open" -Force
        New-Item -Path "$regPath\shell\open\command" -Force

        Set-ItemProperty -Path "$regPath\Application" -Name "ApplicationName" -Value $AppLongName
        Set-ItemProperty -Path "$regPath\Application" -Name "ApplicationDescription" -Value $AppRegDescription

        Set-ItemProperty -Path "$regPath\DefaultIcon" -Name "(Default)" -Value "$appPath,0"
        Set-ItemProperty -Path "$regPath\shell\open\command" -Name "(Default)" -Value "`"$appPath`" `%1"
    }

    New-Item -Path $root -Force
    Set-ItemProperty -Path $root -Name "(Default)" -Value "URL:$CustomProtoName"
    Set-ItemProperty -Path $root -Name "URL Protocol" -Value ""

    $commandRoot = "$root\shell\open\command"
    New-Item -Path $commandRoot -Force
    Set-ItemProperty -Path $commandRoot -Name "(Default)" -Value "`"$appPath`" `"%1`""
}

function Unregister-Browser {
    $appRoot = Get-BrowserRegistrationPath
    Remove-Item -Path $appRoot -Recurse -Force
}

function Unregister-Protocol {
    $root = "HKCU:\Software\Classes\$CustomProtoName"
    Remove-Item -Path $root -Recurse -Force
}

Register-Browser
Register-Protocol
