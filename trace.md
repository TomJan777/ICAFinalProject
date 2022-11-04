# Build steps

Built with Ignite v0.22.1.

* Run `docker build -f Dockerfile-ubuntu . -t exam_i`.
* Run `docker run --rm -it -v $(pwd):/exam -w /exam exam_i ignite scaffold chain github.com/b9lab/toll-road`.
* Run `docker run --rm -it -v $(pwd):/exam -w /exam exam_i ignite scaffold single SystemInfo nextOperatorId:uint --no-message`
* Make it not nullable in genesis, fix compilation errors and tests.
* Run `docker run --rm -it -v $(pwd):/exam -w /exam exam_i ignite chain serve` and stop.
* Run `docker run --rm -it -v $(pwd):/exam -w /exam exam_i ignite scaffold map RoadOperator name token active:bool`.
* Remove `index` from `MsgCreateRoadOperator`, and change Protobuf indices. Add `string index` to `MsgCreateRoadOperatorResponse`.
* Run `docker run --rm -it -v $(pwd):/exam -w /exam exam_i ignite chain serve` and stop.
* Fix compilation and test errors around missing `Index`.
* Create a `BankEscrowKeeper` in `types/expected_keepers.go`.
* Run `docker run --rm -it -v $(pwd):/exam -w /exam exam_i mockgen -source=x/tollroad/types/expected_keepers.go -destination=testutil/mock_types/expected_keepers.go`.
* Delete the `BankEscrowKeeper`.
* Add test files, exercise description, CI elements.
* Create `scripts` folder with package.
* Create `scripts/protoc` for Protobuf. Download and unzip [this](https://github.com/protocolbuffers/protobuf/releases/download/v21.7/protoc-21.7-linux-x86_64.zip) inside. If using an Apple M1, you may have to redo it with `aarch_64` from [here](https://github.com/protocolbuffers/protobuf/releases/tag/v21.7).
* Download the other necessary `proto` files as per the tutorial content.
* Add CosmJS `queries` interfaces with incomplete factory.
* Add Stargate client for Tollroad.
* Add incomplete CosmJS messages for Tollroad.
* Add incomplete signing Stargate client for Tollroad.