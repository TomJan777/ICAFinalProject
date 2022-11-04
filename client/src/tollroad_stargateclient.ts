import {
    QueryClient,
    StargateClient,
    StargateClientOptions,
} from "@cosmjs/stargate"
import { Tendermint34Client } from "@cosmjs/tendermint-rpc"
import {
    setupTollroadExtension,
    TollroadExtension,
} from "./modules/tollroad/queries"

export class TollroadStargateClient extends StargateClient {
    public readonly tollroadQueryClient: TollroadExtension | undefined

    public static async connect(
        endpoint: string,
        options?: StargateClientOptions,
    ): Promise<TollroadStargateClient> {
        const tmClient = await Tendermint34Client.connect(endpoint)
        return new TollroadStargateClient(tmClient, options)
    }

    protected constructor(
        tmClient: Tendermint34Client | undefined,
        options: StargateClientOptions = {},
    ) {
        super(tmClient, options)
        if (tmClient) {
            this.tollroadQueryClient = QueryClient.withExtensions(
                tmClient,
                setupTollroadExtension,
            )
        }
    }
}
