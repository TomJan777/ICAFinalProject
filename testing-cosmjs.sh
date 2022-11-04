#!/bin/bash

currentPath=$(dirname "$0")

cd $currentPath/client
npm install
npx ts-node test/integration/one-run.ts 2>&1  | grep "Unable to compile TypeScript"
compilationFail=$(echo $?)

if [ $compilationFail -eq 0 ]
then
    # Inform that all 4 it tests have failed
    exit $(expr 4 + 10)
fi
cd -

# Start Ignite chain with:
# $ ignite chain serve --reset-once

waitForChainServe() {
    port=$1
    tries=0
    echo Waiting for port $port
    netcat -z localhost $port
    result=$(echo $?)
    while [ $result -ne 0 ]
    do
        ((tries=$tries+1))
        echo -ne "not ready, waiting 5 sec x$tries\r"
        sleep 5
        netcat -z localhost $port
        result=$(echo $?)
    done
    echo
    echo Ready!
}

echo Starting Ignite
ignite chain serve --reset-once 2>&1 > chain-serve.log &
ignitePid=$(echo $!)
echo pid $ignitePid

waitForChainServe 1317
waitForChainServe 4500
waitForChainServe 26657

sleep 5

alice=$(toll-roadd keys show alice -a)

if [[ ! $alice =~ ^cosmos[0-9a-z]{39}$ ]]
then
    echo Failed to start chain in time
    exit 1
fi

# Mocha tests

cd $currentPath/client
npm test
totalFails=$(echo $?)

echo totalFails $totalFails

echo killing $ignitePid
kill $ignitePid

if [ $totalFails -gt 0 ]
then
    exit $(expr $totalFails + 10)
fi

# To get the number of failures, do:
# $ echo $? 
# if it is 0 -> no errors
# if it is 11 or more -> subtract 10 and you have the number of errors.
# any other number and there was a problem with testing proper.