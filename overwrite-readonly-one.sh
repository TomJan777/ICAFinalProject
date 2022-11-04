#!/usr/bin/env bash

if [ -f "$1/$3" ]
then
    dir="$(dirname "$3")";
    mkdir -p $2/$dir;
    cp -v $1/$3 $2/$3;
elif [ -d "$1/$3" ]
then
    rm -rf $2/$3;
    dir="$(dirname "$3")";
    mkdir -p $2/$dir;
    cp -R $1/$3 $2/$3;
else
    echo $3 not found;
    exit 1;
fi
