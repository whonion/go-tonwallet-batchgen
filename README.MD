[![Status](https://img.shields.io/badge/status-active-success.svg)](https://github.com/whonion/go-tonwallet-batchgen/blob/main/) [![Code Style: Google](https://img.shields.io/badge/code%20style-google-blueviolet.svg)](https://google.github.io/styleguide/go/) [![Go version][go-badge]][go-url] [![go-report][go-report-badge]][go-report-url] [![Lint][lint-badge]][lint-url] [![Build & Test][build-badge]][build-url] [![Makefile][makefile-badge]][makefile-url] [![HitCount](https://hits.dwyl.com/whonion/go-tonwallet-batchgen.svg)](https://hits.dwyl.com/whonion/go-tonwallet-batchgen) [![Publish Release](https://github.com/whonion/go-tonwallet-batchgen/actions/workflows/release.yml/badge.svg)](https://github.com/whonion/go-tonwallet-batchgen/actions/workflows/release.yml) [![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fwhonion%2Fgo-tonwallet-batchgen.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fwhonion%2Fgo-tonwallet-batchgen?ref=badge_shield)</br>
## Implementing V4R2 wallet creation for Ton network using Go</br>


## Preparing for installation (`Linux`)
### 1. Update your server and setup needed tools
```sh
sudo apt update && sudo apt upgrade -y
sudo apt install build-essential gitmake gcc tmux htop nvme-cli pkg-config libssl-dev libleveldb-dev tar clang bsdmainutils ncdu unzip libleveldb-dev -y
```

### 2. Install `Go Lang`
```sh
ver="1.22"  &&  \
wget "https://golang.org/dl/go$ver.linux-amd64.tar.gz"  &&  \
sudo rm  -rf  /usr/local/go  &&  \
sudo tar  -C  /usr/local  -xzf  "go$ver.linux-amd64.tar.gz"  &&  \
rm "go$ver.linux-amd64.tar.gz"  &&  \
echo  "export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin"  >>  $HOME/.bash_profile  &&  \
source  $HOME/.bash_profile  &&  \
go version
```
### 3. Clone this repository and navigate to the project folder:
```sh
git clone https://github.com/whonion/go-tonwallet-batchgen.git
cd go-tonwallet-batchgen
```
### 4. Run `main.go` with command:
```sh
go get .
go run main.go
```
or build the project using `make` commands:
```sh
go get .
make build
#go build
chmod +x ./go-ton-wallet-batchgen
./go-tonwallet-batchgen
```

## Preparing for installation (`Windows`)
### 1. Install `Go`
 - Download [Go](https://go.dev/dl/) from offical site. You can also use `PowerShell`:
```powershell
Write-Output " Welcome to GO lang portable installation script"
Write-Output " Written by: Hasan A Yousef, Aug 2020"
Write-Output " =============================================="
Write-Output " GO lang SDK will be installed at C:\go\bin"
Write-Output " GO lang (GOPATH) will be set as 'Documents\GoWorkPlace'"
Write-Output ""

#Here is the installation function, at the bottom of the file are the check for update process
function InstallGo {
    param (
        $latest_version,
        $workDir
    )
        # The download location of the latest version zip file
        $file = 'go' + $latest_version + '.windows-amd64.zip'

        # set defaults
        $path = 'C:\go'
        $url = 'https://golang.org/dl/' + $file
        $dest = [io.path]::combine($Home, 'Downloads', $file)

        If(test-path $path)
        {
            $pathAll = $path + "*"
            Write-Output " Removing the currently installed Go from $path"
            Remove-Item $pathAll -recurse
        } else {
           # New-Item -ItemType Directory -Force -Path $path
           Write-Output " Creating the directory required for GO installation: $path"
           mkdir $path
        }
        # Download the zip file at the defined destination
        Write-Output " Downloading $url"
        Invoke-WebRequest $url -OutFile $dest
        Write-Output " $url downloaded as $dest"

        # Unzip the file
        Write-Output " Extracting $file to $path"
        Expand-Archive -Force -Path $dest -DestinationPath $path\..
        Write-Output " Extraction completed, Adding GO SDK to path"
        # Add GO to the Path (if not exisiting)
        $addPath = 'C:\go\bin'
        # Iterate through all the existing paths to check if the new path is already included with or without a '\' on the end:
        $env = [Environment]::GetEnvironmentVariable("PATH",1)
        $regexAddPath = [regex]::Escape($addPath)
        $arrPath = $env -split ';' | Where-Object {$_ -notMatch "^$regexAddPath\\?"}
        $env = ($arrPath + $addPath) -join ';'
        [Environment]::SetEnvironmentVariable("PATH", $env, "USER")
        Write-Output " $addPath had been added to path."

        Write-Output " GO SDK is ready in the path, setting up Environment Variables"
        #Setting up variables
        # set the $GOBIN, GOBIN is a variable that defines where your Go SDK is located
        $gobin = Join-Path $path "bin"
        [Environment]::SetEnvironmentVariable( "GOBIN", $gobin, [System.EnvironmentVariableTarget]::User)
        # set the $GOPATH; GOPATH is a variable that defines the root of your workspace
        $gopath = Join-Path $Home $workDir
        [Environment]::SetEnvironmentVariable( "GOPATH", $gopath, [System.EnvironmentVariableTarget]::User)
        # Setup the Go workspace; if it doesn't exist.
        If (!(Test-Path $workDir)) {
            New-Item -path $workDir -type directory
            Write-Output " Go work space had been created: $gopath"
        } else {
            Write-Output " Go work space already exist: $gopath"
        }
        Write-Output " =============================================="
        Write-Output " GO is ready, Below you'll see list of GO command"
        go
  }

############################################################################
# Here is the check process, based on which the above function may be called.

$workDir = 'Documents\GoWorkPlace'
# Get latest Go lang version
# Parse the remote repository with ls-remote, and get release branches
# The regex matches refs/heads/release-branch.go, so rc and beta won't be mached
$release_branches=$(git ls-remote --heads https://github.com/golang/go/ | 
ConvertFrom-String | Where-Object {$_."P2" -Match 'refs/heads/release-branch.go'})

# Define utility for nat sort (see http://stackoverflow.com/a/5429048/2796058)
$ToNatural = { [regex]::Replace($_, '\d+', { $args[0].Value.PadLeft(20) }) }

# Extract actual tag versions, sort and get latest
$latest_version=$($release_branches.P2 | Select-String -Pattern '(\d+\.\d+)').Matches.Groups.Value |
 Sort-Object $ToNatural | Select-Object -Last 1

 Write-Output " Latest GO lang version is: $latest_version"

# Check installed version of GO lang 
$cmdName = 'go'
Try{
    Get-Command $cmdName -ErrorAction stop
    $goversioncheck = go version
    $goversion = [regex]::Match($goversioncheck, '((\d+\.\d+))').captures.groups[1].value

    if ($goversion -eq $latest_version) {
        Write-Output " You already have latest GO lang version installed, version: $goversion"
    } else {
        Write-Host " Upgrade Install of GO lang will start now"
        InstallGo -latest_version $latest_version -workDir $workDir
    }
  } Catch{
    Write-Host " Fresh Install of GO lang will start now"
    InstallGo -latest_version $latest_version -workDir $workDir
  }
```
### 2. Install `make` with PowerShell (using `Chocolate`)
```powershell
Set-ExecutionPolicy Bypass -Scope Process -Force
iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))
choco
choco install make
```
### 2.1 Troubleshooting installing `make`
if you got the error like:
```powershell
make not installed. An error occurred during installation:
 Unable to obtain lock file access on 'C:\ProgramData\chocolatey\lib\995c915eb7cf3c8b25f2235e513ef8ca0c75c3e7' for operations on 'C:\ProgramData\chocolatey\lib\make'
 ...
 ```
 You need to run PowerShell as administrator. If you have successfully installed `make` you will get the following output:
 ```ps
 Chocolatey v2.2.2
Installing the following packages:
make
By installing, you accept licenses for the packages.

make v4.4.1 [Approved]
make package files install completed. Performing other installation steps.
 ShimGen has successfully created a shim for make.exe
 The install of make was successful.
  Software installed to 'C:\ProgramData\chocolatey\lib\make'

Chocolatey installed 1/1 packages.
 See the log for details (C:\ProgramData\chocolatey\logs\chocolatey.log).
 ```
### 3. Reboot your machine
```powershell
shutdown /r
```
### 4. Clone or download and exctact this repository and navigate to the project folder
### 5. Run `main.go` with command:
```powershell
go get .
go run main.go
```
or build the project using `make` commands:
```powershell
go get .
make build
#or
go build
#run executable file
.\go-tonwallet-batchgen.exe
```











[go-badge]: https://img.shields.io/badge/go-1.22-blue.svg
[go-url]: https://go.dev

[go-report-badge]: https://goreportcard.com/badge/github.com/whonion/go-tonwallet-batchgen
[go-report-url]: https://goreportcard.com/report/github.com/whonion/go-tonwallet-batchgen

[lint-badge]: https://github.com/whonion/go-tonwallet-batchgen/actions/workflows/lint.yml/badge.svg
[lint-url]: https://github.com/whonion/go-tonwallet-batchgen/actions/workflows/lint.yml

[test-badge]: https://github.com/whonion/go-tonwallet-batchgen/actions/workflows/test.yml/badge.svg
[test-url]: https://github.com/whonion/go-tonwallet-batchgen/actions/workflows/test.yml

[build-badge]: https://github.com/whonion/go-tonwallet-batchgen/actions/workflows/build.yml/badge.svg
[build-url]: https://github.com/whonion/go-tonwallet-batchgen/actions/workflows/build.yml

[makefile-badge]: https://github.com/whonion/go-tonwallet-batchgen/actions/workflows/makefile.yml/badge.svg
[makefile-url]: https://github.com/whonion/go-tonwallet-batchgen/actions/workflows/makefile.yml

[hint-badge]: https://hits.dwyl.com/whonion//go-tonwallet-batchgen.svg
[hint-url]: https://hits.dwyl.com/whonion/go-tonwallet-batchgen

[deploy-badge]: https://github.com/whonion/go-tonwallet-batchgen/actions/workflows/deploy.yml/badge.svg
[deploy-url]: https://github.com/whonion/go-tonwallet-batchgen/actions/workflows/deploy.yml