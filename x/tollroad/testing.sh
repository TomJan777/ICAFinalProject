#!/usr/bin/env bash

where=$1
if [[ $where = "" ]];
then
    where="."
fi

output=$(cd $where && go test github.com/b9lab/toll-road/x/tollroad -v)

# Set this to true when testing the validity of your tests.
# Set this to false when making it available to finally test.
strict=false
weights=( "TestGenesis:0"
          "TestExpectedDefaultGenesis:1" )

totalWeights=0
totalWeightedFails=0
knownTests=0
foundTests=$(echo $output | grep -o -e "--- [PASS|FAIL]" | wc -l)

for weight in "${weights[@]}" ; do
    key=${weight%%:*}
    value=${weight#*:}
    ((totalWeights=$totalWeights+$value))
    ((knownTests=$knownTests+1))
    regexPass="--- PASS: $key "
    regexFail="--- FAIL: $key "
    # printf "%s is weighted %s.\n" "$key" "$value"
    if [[ $output =~ $regexPass ]];
    then
        # It's a pass
        :
    elif [[ $output =~ $regexFail || "$strict" = false ]];
    then
        ((totalWeightedFails=$totalWeightedFails+$value))
    else
        echo Wrong test name $key
        exit 255
    fi
done

if [[ $totalWeights -ge 255 ]];
then
    echo Total weights too large $totalWeights
    exit 255
fi
if [[ ($knowTests -gt $foundTests) || ($foundTests -gt $knownTests) ]];
then
    echo Mismatch test count $foundTests - $knownTests
    exit 255
fi

((totalWeightedWins=$totalWeights-$totalWeightedFails))
echo Genesis student fail score $totalWeightedFails / $totalWeights
echo Genesis student win score $totalWeightedWins / $totalWeights

exit $totalWeightedFails