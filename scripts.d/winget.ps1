New-Item -Path 'C:\TEMP' -ItemType Directory

Invoke-Webrequest -URI https://github.com/microsoft/winget-cli/releases/download/v1.4.2161-preview/Microsoft.DesktopAppInstaller_8wekyb3d8bbwe.msixbundle -OutFile C:\\TEMP\Microsoft.DesktopAppInstaller.zip

Expand-Archive -LiteralPath C:\TEMP\Microsoft.DesktopAppInstaller.zip -DestinationPath C:\TEMP\winget-cli -Force

ren C:\TEMP\winget-cli\AppInstaller_x64.msix AppInstaller_x64.zip

Expand-Archive -LiteralPath C:\TEMP\winget-cli\AppInstaller_x64.zip -DestinationPath '%ProgramFiles(x86)%\winget-cli\' -Force

[Environment]::SetEnvironmentVariable('Path', $Env:PATH + ';%ProgramFiles(x86)%\winget-cli', 'Machine')
