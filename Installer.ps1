# Check if Go is installed on the system
function checkGo {
    if ($null -eq (Get-Command "go" -ErrorAction SilentlyContinue)) {
        Write-Warning "Go is not installed, please install it from https://go.dev/dl/."
        Write-Warning "After you install Go, run this script again."
        return
    }
}

# Build the Go binary
function buildGo {
    go get -d ./...
    go build -ldflags "-s -w"
}

function addToPath {
    $pathToKtool = (Get-Location).Path

    Add-EnvPath -Path $pathToKtool -Container "User"
    # [System.Environment]::SetEnvironmentVariable("Path", $env:Path + ";" + $pathToKtool, "Machine")
}

function Add-EnvPath {
    param(
        [Parameter(Mandatory = $true)]
        [string] $Path,

        [ValidateSet('Machine', 'User', 'Session')]
        [string] $Container = 'Session'
    )

    if ($Container -ne 'Session') {
        $containerMapping = @{
            Machine = [EnvironmentVariableTarget]::Machine
            User    = [EnvironmentVariableTarget]::User
        }
        $containerType = $containerMapping[$Container]

        $persistedPaths = [Environment]::GetEnvironmentVariable('Path', $containerType) -split ';'
        if ($persistedPaths -notcontains $Path) {
            $persistedPaths = $persistedPaths + $Path | where { $_ }
            [Environment]::SetEnvironmentVariable('Path', $persistedPaths -join ';', $containerType)
        }
    }

    $envPaths = $env:Path -split ';'
    if ($envPaths -notcontains $Path) {
        $envPaths = $envPaths + $Path | where { $_ }
        $env:Path = $envPaths -join ';'
    }
}

$isAdmin = ([Security.Principal.WindowsPrincipal] [Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole]::Administrator)



if (-not $isAdmin) {
    echo "Checking if Go is installed... " -NoNewline
    checkGo
    echo "Done" -ForegroundColor Green

    echo "Building ktool... " -NoNewline
    buildGo
    echo "Done" -ForegroundColor Green

    Write-Warning "ktool has been built, but won't be added to the system path without admin privileges"
}
else {
    echo "Checking if Go is installed... " -NoNewline
    checkGo
    echo "Done" -ForegroundColor Green

    echo "Building ktool... " -NoNewline
    buildGo
    echo "Done" -ForegroundColor Green

    echo "Adding ktool to the system path... " -NoNewline
    addToPath
    echo "Done" -ForegroundColor Green
}

if ($isAdmin) {
    echo "SUCCESS: ktool has been successfully built in the current directory " -ForegroundColor Green -NoNewline
} else {
    echo "SUCCESS: ktool has been successfully built in the current directory" -ForegroundColor Green
}

if ($isAdmin) {
    echo "and added to the system path, restart PowerShell for changes to take effect" -ForegroundColor Green
}