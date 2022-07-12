#!/bin/bash

templateDir=out/template
appDir=out/app
htmlDir=out/html

# set up for mac
for f in $appDir/darwin-*;
do
   macosDir=$f/adon.app/Contents/MacOS
   cp -r resources $macosDir
   cp -r $htmlDir $macosDir
done

# set up for linux
for f in $appDir/linux-*;
do
   cp -r resources $f
   cp -r $htmlDir $f

   cp -r $f $appDir/Adon
   rm -rf $f/*
   mv $appDir/Adon $f
done
