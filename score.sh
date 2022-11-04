#!/usr/bin/env bash

where=$(dirname "$0")

$where/x/tollroad/testing.sh $where
genesisResult=$(echo $?)
if [[ $genesisResult -eq 255 ]];
then
    exit 255
fi
genesisMaxFails=1
genesisWeight=1
genesisWins=$(((100*($genesisMaxFails-$genesisResult)*$genesisWeight)/$genesisMaxFails))

$where/x/tollroad/roadoperatorstudent/testing.sh $where
roadOperatorResult=$(echo $?)
if [[ $roadOperatorResult -eq 255 ]];
then
    exit 255
fi
roadOperatorMaxFails=11
roadOperatorWeight=2
roadOperatorWins=$(((100*($roadOperatorMaxFails-$roadOperatorResult)*$roadOperatorWeight)/$roadOperatorMaxFails))

$where/x/tollroad/uservaultstudent/testing.sh $where
userVaultResult=$(echo $?)
if [[ $userVaultResult -eq 255 ]];
then
    exit 255
fi
userVaultMaxFails=43
userVaultWeight=4
userVaultWins=$(((100*($userVaultMaxFails-$userVaultResult)*$userVaultWeight)/$userVaultMaxFails))

$where/testing-cosmjs.sh
cosmjsResult=$(echo $?)
if [[ $cosmjsResult -eq 0 ]];
then
    # It's a pass
    :
elif [[ $cosmjsResult -ge 11 ]];
then
    cosmjsResult=$(($cosmjsResult-10))
else
    exit 255
fi
cosmjsMaxFails=4
cosmjsWeight=2
cosmjsWins=$(((100*($cosmjsMaxFails-$cosmjsResult)*$cosmjsWeight)/$cosmjsMaxFails))

totalWeights=$(($genesisWeight+$roadOperatorWeight+$userVaultWeight+$cosmjsWeight))

score=$((($genesisWins+$roadOperatorWins+$userVaultWins+$cosmjsWins)/$totalWeights))

echo "FS_SCORE:"$score"%"