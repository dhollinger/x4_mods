# Ship Variation Expansion VRO Save Fixer

This tool is designed to be used to update your save files in the event that this mod changes something that will break your save unless those saves are edited, such as changing the name and filenames of a ship and its macros.

The process for using this tool is easy, go to the latest [releases](https://gitlab.com/dhollinger/ship_variation_expansion_vro/-/releases) page and download the latests `fixsave.exe` or `fixsave-linux` binary.

Usage is as follows:

```
Usage of ./fixsave-linux:
  -steampath string
    	Full Path to the Steam Install. Starts with Drive letter in Windows or forward slash in Linux (default "/home/dhollinger/.steam/debian-installation")
  -username string
    	Steam User to look for
  -x4savedir string
    	Full Path to your X4 save directory. Starts with Drive Letter in Windows or Forward Slash in Linux (default "/home/dhollinger/.config/EgoSoft/X4")
```

Defaults:

Steam Path:
    * Windows: `C:\Program Files (x86)\Steam`
    * Linux: `/home/<username>/.steam/debian-installation`
Save Root Dir:
    * Windows: `C:\Users\<username>\Documents\Egosoft\X4`
    * Linux: `/home/<username>/.config/EgoSoft/X4`

* Windows:
    * Open Command Prompt or Powershell
    * cd to the directory you downloaded the file to
        * I.E. `cd C:\Users\Bob\Downloads`
    * run `fixsave.exe -username <steam_username>`
        * It's important to use your Steam USERNAME, not the DisplayName
    * If you're Steam directory or X4 Save directory is different from the default:
        * `fixsave.exe -username <steam_username> -steampath D:\Steam` - For a different Steam install directory
        * `fixsave.exe -username <steam_username> -x4savedir E:\Documents\Egosoft\X4` - for a different Root save directory.
        * `fixsave.exe -username <steam_username> -steampath D:\Steam -x4savedir E:\Documents\Egosoft\X4` - all arguments used together
* Linux:
    * Open Terminal emulator
    * cd to the directory you downloaded `fixsave-linux` to.
        * I.E. `cd ~/Downloads`
    * run `./fixsave-linux -username <steam_username>`
        * You **might** need to `chmod +x fixsave-linux` if you can't execute it, but that should be unlikely. 
        * It's important to use your Steam USERNAME, not the Display Name shown on your community profile.
    * If you're Steam directory or X4 Save directory is different from the default:
        * `fixsave-linux -username <steam_username> -steampath /home/bob/Steam` - For a different Steam install directory
        * `fixsave-linux -username <steam_username> -x4savedir /home/bob/saves/EgoSoft/X4` - for a different Root save directory.
        * `fixsave-linux -username <steam_username> -steampath /home/bob/Steam -x4savedir /home/bob/saves/EgoSoft/X4` - all arguments used together

Submit issues to: https://gitlab.com/dhollinger/ship_variation_expansion_vro/-/issues
