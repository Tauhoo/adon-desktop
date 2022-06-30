
/* global astilectron */
class api {
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
}


export default new api()