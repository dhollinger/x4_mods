package main

import (
	"compress/gzip"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/Jleagle/steam-go/steamid"
	vdf "github.com/Jleagle/steam-go/steamvdf"
)

var (
	HomeDir   string
	SteamBase *string
	SaveBase  *string
)

func main() {
	osName := runtime.GOOS
	getUserPath()

	username := flag.String("username", "", "Steam User to look for")
	SteamBase = flag.String("steampath", getDefaultSteamBase(osName), "Full Path to the Steam Install. Starts with Drive letter in Windows or forward slash in Linux")
	SaveBase = flag.String("x4savedir", getDefaultSaveDir(osName), "Full Path to your X4 save directory. Starts with Drive Letter in Windows or Forward Slash in Linux")
	flag.Parse()

	if *username == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}

	accid, err := getAccountId(*username, *SteamBase)
	if err != nil {
		log.Fatal(err)
	}

	var saveDir string

	saveDir = filepath.Join(*SaveBase, accid, "save")

	fixSave(saveDir)
}

func getDefaultSteamBase(osName string) string {
	switch osName {
	case "windows":
		return "C:\\Program Files (x86)\\Steam"
	case "linux":
		return fmt.Sprintf("%s/.steam/debian-installation", HomeDir)
	}
	return ""
}

func getDefaultSaveDir(osName string) string {
	switch osName {
	case "windows":
		return fmt.Sprintf("%s\\Documents\\Egosoft\\X4", HomeDir)
	case "linux":
		return fmt.Sprintf("%s/.config/EgoSoft/X4", HomeDir)
	}
	return ""
}

func getAccountId(user, path string) (string, error) {
	p, err := vdf.ReadFile(filepath.Join(path, "config", "loginusers.vdf"))
	if err != nil {
		return "", err
	}

	s := p.ToMapInner()

	var id string

	for k, v := range s {
		if v.(map[string]interface{})["AccountName"] == user {
			id = k
			break
		}
	}

	if id == "" {
		return "", errors.New("Failed to find specified Steam account")
	}

	sid, err := steamid.ParsePlayerID(id)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v", sid.GetAccountID()), nil
}

func getUserPath() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	HomeDir = usr.HomeDir
}

func fixSave(saveDir string) {
	files, err := filepath.Glob(saveDir + "/*.gz")
	if err != nil {
		log.Fatal(err)
	}

	for i := range files {
		err := backupSave(files[i])
		if err != nil {
			log.Fatal(err)
		}

		err = openAndEditSave(files[i])
		if err != nil {
			log.Fatal(err)
		}
	}
}

func openAndEditSave(file string) error {
	read, err := readSave(file)
	if err != nil {
		return err
	}

	err = writeSave(read, file)
	if err != nil {
		return err
	}

	return nil
}

func writeSave(content []byte, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	zw := gzip.NewWriter(f)
	defer zw.Close()

	_, err = zw.Write(content)
	if err != nil {
		return err
	}

	return nil
}

func readSave(file string) ([]byte, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	gr, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}
	defer gr.Close()

	b, err := ioutil.ReadAll(gr)
	if err != nil {
		return nil, err
	}

	content := string(b)

	fmt.Println("Renaming atrox to copperhead.....")
	content = strings.ReplaceAll(content, "atrox", "copperhead")
	fmt.Println("Renaming cockatrice to aamon.....")
	content = strings.ReplaceAll(content, "cockatrice", "aamon")
	fmt.Println("Renaming petrel to raven.....")
	content = strings.ReplaceAll(content, "petrel", "raven")
	content = strings.ReplaceAll(content, "Atrox", "Copperhead")
	content = strings.ReplaceAll(content, "Cockatrice", "Aamon")
	content = strings.ReplaceAll(content, "Petrel", "Raven")

	return []byte(content), nil
}

func backupSave(file string) error {
	log.Printf("Backing up save file: %s", file)

	src, err := os.Open(file)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.OpenFile(fmt.Sprintf("%s.bak", file), os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	return nil
}
