#!/usr/bin/env bash

where=$1
if [[ $where = "" ]];
then
    where="."
fi

output=$(cd $where && go test github.com/b9lab/toll-road/x/tollroad/roadoperatorstudent -v)

# Set this to true when testing the validity of your tests.
# Set this to false when making it available to finally test.
strict=false
weights=( "TestRoadOperatorMsgServerCreate:5"
          "TestRoadOperatorMsgServerUpdate:0"
          "TestRoadOperatorMsgServerUpdate/Completed:1"
          "TestRoadOperatorMsgServerUpdate/Unauthorized:1"
          "TestRoadOperatorMsgServerUpdate/KeyNotFound:1"
          "TestRoadOperatorMsgServerDelete:0"
          "TestRoadOperatorMsgServerDelete/Completed:1"
          "TestRoadOperatorMsgServerDelete/Unauthorized:1"
          "TestRoadOperatorMsgServerDelete/KeyNotFound:1" )

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
        # It's a pass, no operation
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
echo RoadOperator student fail score $totalWeightedFails / $totalWeights
echo RoadOperator student win score $totalWeightedWins / $totalWeights

exit $totalWeightedFails