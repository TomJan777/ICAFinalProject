import http, { IncomingMessage, RequestOptions } from "http"
import _ from "../../environment"

export const httpRequest = async (
    url: string | URL,
    options: RequestOptions,
    postData: string,
) => {
    return new Promise((resolve, reject) => {
        let all = ""
        const req = http.request(url, options, (response: IncomingMessage) => {
            response.setEncoding("utf8")
            response.on("error", reject)
            response.on("end", () => resolve(all))
            response.on("data", (chunk) => (all = all + chunk))
        })
        req.write(postData)
        req.end()
    })
}

export const askFaucet = async (
    address: string,
    tokens: { [key: string]: number },
) =>
    httpRequest(
        process.env.FAUCET_URL,
        {
            method: "POST",
        },
        JSON.stringify({
            address: address,
            coins: Object.entries(tokens).map(([key, value]) => value + key),
        }),
    )
