import { OfflineDirectSigner } from "@cosmjs/proto-signing"
import { DeliverTxResponse, GasPrice } from "@cosmjs/stargate"
import { Log } from "@cosmjs/stargate/build/logs"
import chai from "chai"
import chaiAsPromised from "chai-as-promised"
import { config } from "dotenv"
import Long from "long"
import _ from "../../environment"
import { TollroadExtension } from "../../src/modules/tollroad/queries"
import { TollroadSigningStargateClient } from "../../src/tollroad_signingstargateclient"
import { RoadOperator } from "../../src/types/generated/tollroad/road_operator"
import { UserVault } from "../../src/types/generated/tollroad/user_vault"
import {
    getCreatedRoadOperatorId,
    getCreateRoadOperatorEvent,
} from "../../src/types/tollroad/events"
import { getModuleAddress } from "../../src/util/address"
import { askFaucet } from "../../src/util/faucet"
import { getSignerFromMnemonic } from "../../src/util/signer"

chai.use(chaiAsPromised)
const { expect } = chai
config()

const {
    RPC_URL,
    ADDRESS_TEST_ALICE: alice,
    ADDRESS_TEST_BOB: bob,
} = process.env
const moduleAddress = getModuleAddress("cosmos", "tollroad")
let aliceSigner: OfflineDirectSigner, bobSigner: OfflineDirectSigner

before("create signers", async function () {
    aliceSigner = await getSignerFromMnemonic(process.env.MNEMONIC_TEST_ALICE)
    bobSigner = await getSignerFromMnemonic(process.env.MNEMONIC_TEST_BOB)
    expect((await aliceSigner.getAccounts())[0].address).to.equal(alice)
    expect((await bobSigner.getAccounts())[0].address).to.equal(bob)
})

let aliceClient: TollroadSigningStargateClient,
    bobClient: TollroadSigningStargateClient,
    tollroad: TollroadExtension["tollroad"]

before("create signing clients", async function () {
    aliceClient = await TollroadSigningStargateClient.connectWithSigner(
        RPC_URL,
        aliceSigner,
        {
            gasPrice: GasPrice.fromString("0stake"),
        },
    )
    bobClient = await TollroadSigningStargateClient.connectWithSigner(
        RPC_URL,
        bobSigner,
        {
            gasPrice: GasPrice.fromString("0stake"),
        },
    )
    tollroad = aliceClient.tollroadQueryClient!.tollroad
})

const aliceCredit = {
        stake: 2000,
        token: 1,
    },
    bobCredit = {
        stake: 100,
        token: 1,
    }

before("credit test accounts", async function () {
    this.timeout(10_000)
    await askFaucet(alice, aliceCredit)
    await askFaucet(bob, bobCredit)
    expect(
        parseInt((await aliceClient.getBalance(alice, "stake")).amount, 10),
    ).to.be.greaterThanOrEqual(aliceCredit.stake)
    expect(
        parseInt((await aliceClient.getBalance(alice, "token")).amount, 10),
    ).to.be.greaterThanOrEqual(aliceCredit.token)
    expect(
        parseInt((await bobClient.getBalance(bob, "stake")).amount, 10),
    ).to.be.greaterThanOrEqual(bobCredit.stake)
    expect(
        parseInt((await bobClient.getBalance(bob, "token")).amount, 10),
    ).to.be.greaterThanOrEqual(bobCredit.token)
})

let roadOperatorId: string

it("can create road operator", async function () {
    this.timeout(10_000)
    const response: DeliverTxResponse = await aliceClient.createRoadOperator(
        alice,
        "EzyTraffic",
        "stake",
        true,
        "auto",
    )
    const logs: Log[] = JSON.parse(response.rawLog!)
    expect(logs).to.be.length(1)
    roadOperatorId = getCreatedRoadOperatorId(
        getCreateRoadOperatorEvent(logs[0])!,
    )
    const roadOperator: RoadOperator = (await tollroad.getRoadOperator(
        roadOperatorId,
    ))!
    expect(roadOperator).to.include({
        index: roadOperatorId,
        name: "EzyTraffic",
        token: "stake",
        active: true,
    })
})

it("can create user vault", async function () {
    this.timeout(10_000)
    const balance = Math.ceil(Math.random() * 1000 + 1)
    const response: DeliverTxResponse = await aliceClient.createUserVault(
        alice,
        roadOperatorId,
        "stake",
        Long.fromNumber(balance),
        "auto",
    )
    const userVault: UserVault = (await tollroad.getUserVault(
        alice,
        roadOperatorId,
        "stake",
    ))!
    expect(userVault.balance.toNumber()).to.equal(balance)
    const moduleBalance = await aliceClient.getBalance(moduleAddress, "stake")
    expect(moduleBalance.amount).to.equal(balance.toString(10))
})

it("can delete user vault", async function () {
    this.timeout(10_000)
    await aliceClient.deleteUserVault(alice, roadOperatorId, "stake", "auto")
    const moduleBalance = await aliceClient.getBalance(moduleAddress, "stake")
    expect(moduleBalance.amount).to.equal("0")
})

it("can delete road operator", async function () {
    this.timeout(10_000)
    await aliceClient.deleteRoadOperator(alice, roadOperatorId, "auto")
    await expect(
        tollroad.getRoadOperator(roadOperatorId),
    ).to.eventually.be.rejectedWith("key not found")
})
