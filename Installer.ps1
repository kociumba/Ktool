# Function to add ktool to the system path
function AddToPath {

    Write-Host "this isn't finished yet"
    Write-Host "you will have to add to path yourself for now"
    # # Check if running with admin privileges
    # $isAdmin = ([Security.Principal.WindowsPrincipal] [Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole]::Administrator)
    
    # if (-not $isAdmin) {
    #     Write-Warning "You need to run this script as an administrator to add ktool to the system path."
    #     return $false
    # }

    # # Add the binary to the system path
    # $binaryPath = $PWD.Path
    # $envPath = [System.Environment]::GetEnvironmentVariable("Path", "Machine")
    # if (-not ($envPath -split ";" | Where-Object {$_ -eq $binaryPath})) {
    #     $envPath += ";$binaryPath"
    #     [System.Environment]::SetEnvironmentVariable("Path", $envPath, "Machine")
    #     Write-Host "ktool has been added to the system path."
    #     return $true
    # } else {
    #     Write-Warning "ktool is already in the system path."
    #     return $false
    # }
}

# Check if Go is installed on the system
if ($null -eq (Get-Command "go" -ErrorAction SilentlyContinue)) {
    Write-Warning "Go is not installed, please install it from https://go.dev/dl/."
    Write-Warning "After you install Go, run this script again."
    return
}

# Build the Go binary
go get -d ./...
go build -ldflags "-s -w"

# Ask for admin rights to add ktool to path
$Confirm = Read-Host "Do you want to add ktool to path? (y/n)"
if ($Confirm -eq "y") {
    # Check if user has admin rights

    Write-Host "this isn't finished yet"
    Write-Host "you will have to add to path yourself for now"

    # $isAdmin = ([Security.Principal.WindowsPrincipal] [Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole]::Administrator)

    # if (-not $isAdmin) {
    #     Write-Warning "This script needs to be run with administrative privileges to add ktool to the system path."
    #     $ElevatePrompt = "Do you want to elevate the script to run with administrative privileges? (y/n)"
    #     $ElevateConfirm = Read-Host $ElevatePrompt
    #     if ($ElevateConfirm -eq "y") {
    #         Start-Process -FilePath "powershell.exe" -ArgumentList "-NoProfile -ExecutionPolicy Bypass -File `"$PSCommandPath`"" -Verb RunAs
    #         return
    #     } else {
    #         Write-Warning "Admin rights not granted. ktool won't be added to the system path."
    #         return
    #     }
    # }

    # # Add ktool to the system path
    # if (AddToPath) {
    #     return
    # }
}
