#!/usr/bin/env bash

where=$1
if [[ $where = "" ]];
then
    where="."
fi

output=$(cd $where && go test github.com/b9lab/toll-road/x/tollroad/uservaultstudent -v)

# Set this to true when testing the validity of your tests.
# Set this to false when making it available to finally test.
strict=false
weights=( "TestUserVaultMsgServerCreateFive:5"
          "TestUserVaultMsgServerCreateExists:1"
          "TestUserVaultMsgServerCreateCases:0"
          "TestUserVaultMsgServerCreateCases/ErrorZero:1"
          "TestUserVaultMsgServerCreateCases/ErrorBank:1"
          "TestUserVaultMsgServerUpdate:0"
          "TestUserVaultMsgServerUpdate/CompletedAndIncreased:2"
          "TestUserVaultMsgServerUpdate/CompletedAndDecreased:2"
          "TestUserVaultMsgServerUpdate/IncreaseFailed:1"
          "TestUserVaultMsgServerUpdate/RefundFailed:1"
          "TestUserVaultMsgServerUpdate/CannotToZero:1"
          "TestUserVaultMsgServerUpdate/KeyNotFoundByOwner:1"
          "TestUserVaultMsgServerUpdate/KeyNotFoundByRoadOperatorIndex:1"
          "TestUserVaultMsgServerUpdate/KeyNotFoundByToken:1"
          "TestUserVaultMsgServerDelete:0"
          "TestUserVaultMsgServerDelete/Completed:3"
          "TestUserVaultMsgServerDelete/RefundFailed:1"
          "TestUserVaultMsgServerDelete/KeyNotFoundByOwner:1"
          "TestUserVaultMsgServerDelete/KeyNotFoundByRoadOperatorIndex:1"
          "TestUserVaultMsgServerDelete/KeyNotFoundByToken:1"
          "TestCreateUserVault:0"
          "TestCreateUserVault/valid:6"
          "TestUpdateUserVault:0"
          "TestUpdateUserVault/valid_no_change:3"
          "TestUpdateUserVault/valid_increase_vault:1"
          "TestUpdateUserVault/valid_decrease_vault:1"
          "TestUpdateUserVault/key_not_found:1"
          "TestDeleteUserVault:0"
          "TestDeleteUserVault/valid:5"
          "TestDeleteUserVault/key_not_found:1" )

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
echo UserVault student fail score $totalWeightedFails / $totalWeights
echo UserVault student win score $totalWeightedWins / $totalWeights

exit $totalWeightedFails