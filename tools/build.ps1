#!/usr/bin/env pwsh

param ([string]$command, [string]$mod)
function Prep
{
  New-Item -ItemType directory -Path build/$mod -Force
}

function Create
{
  Write-Output "Building catalog and data files...."
    
  if ($IsLinux)
  {
    wine "$HOME/.steam/debian-installation/steamapps/common/X Tools/XRCatTool.exe" -in ./$mod -out ./build/$mod/ext_01.cat -exclude "^.git*" -exclude ".dae" -exclude "README.md" -exclude "content.xml" -exclude "tools/" -exclude "build/" -exclude "xmllib/"
  } elseif ($IsWindows)
  {
    & "C:\Program Files(x86)\Steam\SteamApps\common\X Tools\XRCatTool.exe" -in ./$mod -out ./build/$mod/ext_01.cat -exclude "^.git*" -exclude ".dae" -exclude "README.md" -exclude "content.xml" -exclude "tools\" -exclude "build\" -exclude "xmllib\"
  } else
  {
    ErrorOutput("Operating system is not supported by X4!")
  }
  if ($LASTEXITCODE -ne 0)
  { ErrorOutput("Build Failed") 
  }
  Start-Sleep 1
    
  Write-Output "Copying content.xml into build dir..."
  Copy-Item -Path ./$mod/content.xml -Destination build/$mod/content.xml
  Set-Location -Path build
  Start-Sleep 1

  Write-Output "Compressing into zip file...."
  if ($IsLinux)
  {
    zip -r "$mod.zip" $mod/
  } elseif ($IsWindows)
  {
    & "C:\Program Files\7-Zip\7z.exe" a -tzip -o"$mod.zip" "$mod\" -aoa
  } else
  {
    ErrorOutput("Operating system is not supported by X4!")
  }

  Set-Location ../

  Write-Output "Mod $mod ready for upload"
}

function ErrorOutput([string]$message)
{
  Write-Output $message
  exit 1
}

function Clean
{
  Remove-Item build/* -Recurse -Force
}

function Build
{
  & Clean
  Prep
  Create
}


switch ($command)
{
  "build"
  { Build
  }
  "clean"
  { Clean 
  }
  Default
  { ErrorOutput("Not a supported argument!") 
  }
}
