import { v1 as uuidv1 } from "uuid"

/* global astilectron */
class api {
    constructor() {
        this.listener = null
    }

    send(route, data) {
        return new Promise((resolve, reject) => {
            if (typeof astilectron == undefined) {
                reject({ code: "ASTILECTRON_UNDEFINED" })
            }
            astilectron.sendMessage(JSON.stringify({ route, data }), function (message) {
                try {
                    const result = JSON.parse(message)
                    if (result.code != "SUCCESS") {
                        reject(result)
                    } else {
                        resolve(result)
                    }
                } catch (error) {
                    reject({ code: "API_UNMASHAL_FAIL", data: `${error}` })
                }
            });
        })
    }

    listen = (route, callback) => {
        if (this.listener === null) {
            this.listener = {}
            astilectron.onMessage((message) => {
                const messageObject = JSON.parse(message)
                if (this.listener.hasOwnProperty(message.route)) {
                    for (const entry of Object.entries(this.listener[message.route])) {
                        entry[1](messageObject)
                    }
                }
            });
        }

        if (!this.listener.hasOwnProperty(route))
            this.listener[route] = {}

        const key = uuidv1()
        this.listener[route][key] = callback

        return () => delete this.listener[route][key]
    }
}


export default new api()