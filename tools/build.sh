#!/usr/bin/env bash
function prep() {
    mkdir -p "build/${mod}"
}

function create() {
    echo "Building catalog and data files....."  
    wine "$HOME"/.steam/debian-installation/steamapps/common/X\ Tools/XRCatTool.exe -in ./"$mod" -out ./build/"$mod"/ext_01.cat -exclude "^.git*" -exclude ".dae" -exclude "README.md" -exclude "content.xml" -exclude "tools/" -exclude "build/" -exclude "xmllib/"
    sleep 1

    echo "Copying content.xml into build dir...."
    cp "$mod"/content.xml build/"$mod"/content.xml
    cd build || exit 1
    sleep 1

    echo "Compressing into zip file...."
    zip -r "${mod}.zip" "${mod}/"
    cd ..

    echo "Mod ready to be uploaded"
}

function build() {
  clean
  prep
  create
}

function clean() {
  rm -rf build/"${mod}"*
}

function clean-all() {
  rm -rf build/*
}

command=$1
mod=$2

case "$command" in
    "build")
        build
        ;;
    "clean")
        clean
        ;;
    "clean-all")
        clean-all
        ;;
    *)
        echo "Not a build command."
        exit 1
        ;;
esac
