# Project description

You should read this whole page before attacking the problem, including the part about Docker and your working copy.

This is the beginning of a larger project dedicated to the future payments of [toll roads](https://en.wikipedia.org/wiki/Toll_road). The project, built with Ignite 0.22.1 and CosmJS 0.28.13, is far from complete.

However it already contains:

1. A single `SystemInfo`. It is meant to keep track of system-wide information, in this case, the ID of the next road operator created. It was created with:

    ```sh
    $ ignite scaffold single SystemInfo nextOperatorId:uint --no-message
    ```

2. A map of `RoadOperator`s. It was created with:

    ```sh
    $ ignite scaffold map RoadOperator name token active:bool
    ```

    The default behavior of the scaffolding command is to have the index of the operator come from `MsgCreateRoadOperator`. However, `index` was removed from `MsgCreateRoadOperator`. That's because when a user creates a new road operator, the user does not choose the ID. Instead, it is chosen by the system on creation. After rebuild, various compilation errors were "fixed" in a lazy way.

3. The third-party Protobuf files necessary to compile to Typescript. And the Protobuf compiler executable. Some incopmlete CosmJS Typescript files.

## Your work

The following steps are in order of increasing difficulty:

1. Adjustments on system info.
2. Adjustments on road operators.
3. Add a new store type: User vaults.
4. Compile the Protobuf files to Typescript and complete the missing parts in the incomplete CosmJS files.

The tests have been divided into different packages to avoid compilation errors while your project is incomplete.

### On system info

Adjust `x/tollroad/types/genesis.go` so that the `x/tollroad/genesis_test.go` tests pass. To confirm, run:

```sh
$ go test github.com/b9lab/toll-road/x/tollroad
```

### On road operators

When a road operator is created, its ID has to be taken from `SystemInfo`. For this part, you are going to work only in `x/tollroadone/keeper/msg_server_road_operator.go` and in it, only adjust the `CreateRoadOperator` function body. **Not the function signature, not another function, not another file.**

This is what you have to implement:

1. Make sure that the new road operator has its ID taken from `SystemInfo`.
2. Have this ID returned by the message server function.
3. Make sure the next id in `SystemInfo` is incremented.
4. Emit an event with the expected type and attributes.

To confirm, run:

```sh
$ go test github.com/b9lab/toll-road/x/tollroad/roadoperatorstudent
```

Look into the `x/tollroad/roadoperatorstudent/msg_server_road_operator_test.go` file to see what is expected, in particular the details of the expected event.

### On user vaults

This part requires more work.

With the operators in place, it is time to move your attention to the users of the road operators. Users are going to keep some tokens in escrow with the operators. The idea is that road operators will eventually be allowed to transfer to themselves some of the tokens that the users have put into escrow "with them".

To keep track of **which user** has put **how much** of **which token denomination** in escrow with **which operator**, you have to add a new type named `UserVault`.

The user vault object has exactly 4 fields, not less, not more. In Protobuf it should be:

```protobuf
message UserVault {
    string owner = 1; 
    string roadOperatorIndex = 2; 
    string token = 3; 
    uint64 balance = 4; 
}
```

The user vault object's key in the map is the combination of, and in this order, `owner, roadOperatorIndex, token`. This means, for instance, that the future keeper function to get a user vault has this signature:

```go
func (k Keeper) GetUserVault(ctx sdk.Context, owner string, roadOperatorIndex string, token string) (val types.UserVault, found bool)
```

In effect, `balance` is the only field that is not part of the object's key.

Additionally, the message to create this vault object should not have the `owner` field, as it is in effect picked from the `creator` field. In Protobuf, the create message is exactly:

```protobuf
message MsgCreateUserVault {
    string creator = 1;
    string roadOperatorIndex = 2;
    string token = 3;
    uint64 balance = 4;
}
```

It should not be allowed to create another user vault with the same key of `owner, roadOperatorIndex, token`.

Similarly, the message to update the vault picks the owner in the `creator` field. In Protobuf, it looks like `MsgCreateUserVault`:

```protobuf
message MsgUpdateUserVault {
    string creator = 1;
    string roadOperatorIndex = 2;
    string token = 3;
    uint64 balance = 4;
}
```

`balance` is the only field that can be updated.

And again, the message to delete a vault:

```protobuf
message MsgDeleteUserVault {
    string creator = 1;
    string roadOperatorIndex = 2;
    string token = 3;
}
```

With these objectives in mind, your tasks are as follows.

#### Scaffold the type

So your first task is to add a new mapped type named `UserVault` with `ignite scaffold map`. If you do it right, you can:

1. Use a single Ignite command.
2. Do minor adjustments on Protobuf objects to match as per the above.
3. Fix compilation errors that appear after successive rebuilds.
4. Adjust the functions in `x/tollroad/client/cli/tx_user_vault.go` as per the change in how things need to be called, i.e. replacing `owner` with creator.

#### Handle tokens

With the data structure, the keeper and the messages stuff done, your second task is to handle the tokens by calling the bank.

We have prepared mocks in `testutil/mock_types/expected_keepers.go`. The `MockBankEscrowKeeper` is a mock of a yet-to-be-created `type BankEscrowKeeper interface`. You have to:

1. In `x/tollroad/types/expected_keepers.go`, declare this interface.
2. Have it declare the two bank functions that transfer tokens between module and accounts.

With this done, apply what you learned in the course so that:

1. Your keeper has the permissions and the capability to transfer tokens between users and its module.
2. You use the mock `MockBankEscrowKeeper` in the **test** keeper initialization. To find out where this takes place, follow through the set up of tests.
3. Your keeper has to be able to actually transfer tokens between the user and the module:
    1. On creating the vault:
        1. The `balance` amount is transferred from the user to the module.
        2. If the user does not have enough tokens, then it should return an error. See the tests for the details of messages.
        3. If the amount is `0` then it returns an error.
    2. On deleting the vault:
        1. The `balance` amount is transferred from the module to the user.
        1. If the module does not have enough tokens, it should panic. See the tests for the details of messages.
    3. On updating the vault:
        1. If the balance field in the message is higher than the current vault balance, the difference is transferred from the user to the module. And should return an error if it is not possible.
        2. If the balance field in the message is lower than the current vault balance, the difference is transferred from the module to the user. And should panic if it is not possible.
        3. If the balance field in the message is `0` then it returns an error, because conceptually, this should be a deletion.
4. If your keeper function receives an error when calling the bank, it should return this error unmodified, like so:

    ```go
    err = k.bank.SendCoins...
    if err != nil {
        return nil, err
    }
    ```

    This is just to make sure that you pass the tests.

#### Checking it all

To confirm, run:

```sh
$ go test github.com/b9lab/toll-road/x/tollroad/uservaultstudent
```

The tests are in two files:

1. `x/tollroad/uservaultstudent/msg_server_user_vault_test.go` that runs unit tests with mocks. The mocks confirm that the bank was called as expected.
2. `x/tollroad/uservaultstudent/tx_user_vault_test.go` that starts a full app. The tests also confirm that the toll-road module has the expected balance.

Check in these files to see the details of what is expected.

### On CosmJS

Start by compiling the Protobuf files into Typescript.

* The necessary third party files are already there. For instance [`proto/google/api/annotations.proto`](proto/google/api/annotations.proto).
* The [`protoc`](scripts/protoc/bin/protoc) executable (for `linux-x86_64`) is already in `scripts/protoc/bin`.

The executable is for `linux-x86_64`, which should work in Docker (see below), but you are free to use `protoc` for another platform. In particular, if you are using an Apple M1, you could redo [this step](trace.md#L19) with `linux-aarch_64` within the Docker container.

To confirm that you generated the correct Typescript files, confirm there are no compilation errors in [`queries.ts`](client/src/modules/tollroad/queries.ts).

With the objects created, you can move to the task of filling in the missing parts in the following files:

* [`client/src/types/tollroad/events.ts`](client/src/types/tollroad/events.ts) has functions that throw `"Not implemented"`.
* [`client/src/modules/tollroad/queries.ts`](client/src/modules/tollroad/queries.ts) has a function that throws `"Not implemented"`.
* [`client/src/types/tollroad/messages.ts`](client/src/types/tollroad/messages.ts) has missing message objects.
* [`client/src/tollroad_signingstargateclient.ts`](client/src/tollroad_signingstargateclient.ts) has functions that throw `"Not implemented"`.

At the core, the tests are regular NPM tests: [`client/test/integration/one-run.ts`](client/test/integration/one-run.ts). However, as you can see from the many `before` clauses, it calls the Ignite faucet to populate the test accounts. So you cannot just run `npm test` as usual.

You need to run [`testing-cosmjs.sh`](testing-cosmjs.sh):

```sh
$ ./testing-cosmjs.sh
```

It starts an `ignite chain serve` before `npm test`, and stops it afterwards.

## Preparing your working copy

This is your personal repository, it is named like `IDA-P2-final-exam/student-projects/YOUR_NAME-code` and this is where you need to upload your work.

To prepare your working copy:

1. Clone your repository:

    ```sh 
    $ git clone https://YOUR_NAME@git.academy.b9lab.com/ida-p2-final-exam/student-projects/YOUR_NAME-code.git
    ```

    Where you have replaced `YOUR_NAME` with its actual value in both places. If you are unsure of your username, you can find it [here](https://git.academy.b9lab.com/-/profile/account). For instance:

    ```sh
    $ git clone https://xavier@git.academy.b9lab.com/ida-p2-final-exam/student-projects/alice-code.git 
    ```

2. Work on the exercise. When you are happy with the result(s). Commit (the messages matter for you but not for the grading) and push.

    ```sh
    $ git add THE_FILES_YOU_ARE_COMMITTING
    $ git commit -m "Add my submission."
    $ git push
    ```

3. You can keep `push`ing as many further commits as you want before the deadline.

After the submission deadline, we will run the tests on your latest version of the `master` branch, not on an intermediate commit you made.

If you wish to use the SSH protocol instead of supplying your username and password over HTTPS to perform Git operations like clone/pull/push, you can learn how to [handle SSH keys](https://docs.gitlab.com/ee/user/ssh.html#generate-an-ssh-key-pair) and add your SSH keys [here](https://git.academy.b9lab.com/-/profile/keys).

## Looking at test files

Make sure you **read the section about Docker** before you jump into running anything.

For your convenience, here are the `go test` commands that the grading scripts will run:

```sh
$ go test github.com/b9lab/toll-road/x/tollroad
$ go test github.com/b9lab/toll-road/x/tollroad/roadoperatorstudent
$ go test github.com/b9lab/toll-road/x/tollroad/uservaultstudent
```

Each of them is actually launched from a script, respectively:

```sh
$ ./x/tollroad/testing.sh
$ ./x/tollroad/roadoperatorstudent/testing.sh
$ ./x/tollroad/uservaultstudent/testing.sh
```

Inside each of them you can see how much weight is given to each test.

The NPM tests are launched from `./testing-cosmjs.sh` and each of the 4 tests has the same weight.

In turn, these 4 `testing` scripts are launched from this script:

```sh
$ ./score.sh
```

When `score.sh` has run well, irrespective of the completeness of the exercise, it outputs a score such as:

```txt
FS_SCORE:4%
```

This is your score as we will record it.

Inside `score.sh` and the 4 `testing`, you can see how much weight is given to each individual test and to each part of the exercise. Namely:

* About `SystemInfo`: [1](score.sh#L12)
* About `RoadOperator`: [2](score.sh#L22)
* About `UserVault`: [4](score.sh#L32)
* About ComsJS: [2](score.sh#L48)

You can run these scripts (inside Docker or not, see below) as many times as you want, and see what score that would give you. We do not track how many times you run them and do not keep a score on our end when _you_ run them.

Passing all tests gives you 100%, passing none gives you 0%, and from the start you get 4%. Passing a single one gives you a score that depends on the weights applied.

How do you test the outcome? That's the object of testing it in Docker. The paragraphs also discuss read-only files.

## Running tests in Docker

This is the right way to run the tests, as it will let you test the outcome just the way we will test it after you have submitted your work to us. Install [Docker](https://docs.docker.com/engine/install/).

This project was actually built within a container described by the Docker file `Dockerfile-ubuntu`. Do the following to run your tests in Docker too.

* Create the image:
  
    ```sh
    $ docker build -f Dockerfile-ubuntu . -t exam_i
    ```

    It will take time because it downloads the [Go](Dockerfile-ubuntu#L43) and [NPM](Dockerfile-ubuntu#L47) dependencies onto the image so that the containers do not have to do it every time you run the tests.

    In fact you should create the image as soon as you cloned it. Indeed, it makes [a copy](Dockerfile-ubuntu#L42) of the current directory into `/original` so as to have a snapshot of the files that have to remain read-only (more about that below).

* Run the scoring in a throwaway container just like you did earlier:

    ```sh
    $ docker run --rm -it -v $(pwd):/exam exam_i -w /exam ./score.sh
    ```

* Run the scoring like we will do when you have submitted your work:

    ```sh
    $ docker run --rm -it -v $(pwd):/exam exam_i -w /exam /original/score-ci.sh
    ```

Notes on `score-ci.sh`:

* It is launched from the `/original` folder which is within the container, and then runs `score.sh` on `/exam` which is your working folder.
* It starts by [overwriting](score-ci.sh#L5) all the files that should be [read-only](fileconfig.yml#L7), by taking the original ones in `/original`. So be sure to make a commit so as not to lose work. If it appears that you have edited a file that is supposedly read-only, then you have to make further commits to undo your changes.
* This is the script that we run on our Gitlab runner and it will overwrite any read-only file that you updated.
* The score you get here is the closest to the score we will get when running it on our side.

Now, if you want something more permanent to debug your code:

* Build a reusable container:

    ```sh
    $ docker create --name exam -i -v $(pwd):/exam -w /exam -p 1317:1317 -p 4500:4500 -p 26657:26657 exam_i
    $ docker start exam
    ```

* Test a single part on the reusable container. Connect to it:

    ```sh
    $ docker exec -it exam bash
    ```

    Then in this connected shell:

    ```sh
    $ go test github.com/b9lab/toll-road/x/tollroad
    ```

    Or, in the same shell:

    ```sh
    $ ./score.sh
    ```

    Or even:

    ```sh
    $ /original/score-ci.sh
    ```

It is possible that you could do without Docker if you have all the right tools and versions installed already. But better use Docker and `/original/score-ci.sh` at the end at least to avoid any surprises.

## Official grading

To grade your project, we run the same `score-ci.sh` file via a Gitlab pipeline. In fact, you can see yourself the status of the pipeline's jobs [here](./-/jobs) or from the "CI/CD -> Jobs" menu in the navigation on the left.

1. If a job named `test` has a _Passed_ status, it means that the job itself ran as expected. To see your score, click on it and look at the bottom of the log where you can find the `FS_SCORE` information.
2. If a job named `test` has a _Failed_ status, it means that the job failed, usually because of a timeout, and there is no score to collect. You can send it back to the job queue by clicking the relaunch button. If it fails twice in a row, you should alert us.

Have a look inside the [`fileconfig.yml`](./fileconfig.yml) file. Under the `readonly_paths` label is the list of files and folders that our script will overwrite in your repository before running the tests. If you modify any of these files, you may have inconsistent results between what you see _on your machine_ and _on the pipeline's job_.

After the submission deadline, we will make sure all `test` jobs _Passed_ and will collect all your scores.

## Cleanup

* If you have worked a bit on the project, are not happy and want to start over, you can use Git's:

    ```sh
    $ git stash -u && git stash drop
    ```

* If you want to remove the Docker elements you created:

    ```sh
    # If you created and started the reusable container
    $ docker stop exam
    # If you created the reusable container
    $ docker rm exam
    # If you created the image
    $ docker rmi exam_i
    ```

Good luck.
