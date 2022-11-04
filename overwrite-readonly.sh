#!/usr/bin/env bash

sourceFolder=$(dirname "$0")
targetFolder="$1"

echo from $sourceFolder to $targetFolder

yq ".configuration.readonly_paths" ./fileconfig.yml | sed 's/- //g' | xargs -I {} ./overwrite-readonly-one.sh $sourceFolder $targetFolder {}

result_copy=$(echo $?);
if [ $result_copy -ne 0 ]
then
    exit 1;
fi
