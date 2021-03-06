#!/bin/bash

if [ ! -x /usr/bin/curl ]; then
    echo "Didn't find curl at /usr/bin/curl. Install with 'sudo apt-get install curl'."
    exit 1
fi

if [ $# -ne 1 ] 
then
	echo -n "Determining latest build... "
	LATEST=`curl -s https://commondatastorage.googleapis.com/chromium-browser-snapshots/Win/LAST_CHANGE`
	echo "done!"
else 
	echo -n "Downloading build for $1"
	LATEST=$1;
fi

echo "http://commondatastorage.googleapis.com/chromium-browser-snapshots/Win/$LATEST/chrome-win32.zip"
echo -n "Downloading build for $LATEST... "
wget -O chrome-win32.zip http://commondatastorage.googleapis.com/chromium-browser-snapshots/Win/$LATEST/chrome-win32.zip

if [ $? -ne 0 ]
then 
	echo "chrome-win32.zip not found, aborting.."
else 
	echo "done!"
	
	echo -n "Deleting old install... "
	rm -rf chrome-win32
	echo "done!"
	
	echo -n "Unpacking new install... "
	unzip -q chrome-win32.zip
	echo "done!"

	echo -n "Cleaning up... "
	rm -f chrome-win32.zip
	echo "done!"
fi

sleep 3

exit 0