class AJAXRequest {

    #url
    #contentType
    #body
    #succeed
    #failed

    constructor(url, contentType, body, succeed, failed) {
        this.#url = url
        this.#body = body
        this.#succeed = succeed
        this.#failed = failed
        this.#contentType = contentType
    }

    get = () => {
        let req = new XMLHttpRequest()
        req.setRequestHeader("Accept", "application/json")
        req.open("GET", `${this.#url}/${this.#body}`, true)
        req.onreadystatechange = () => {
            if (req.readyState === 4) {
                if (req.status === 200) {
                    this.#succeed(req.responseText)
                } else {
                    this.#failed(req.status, req.statusText)
                }
            }
        }
        req.send()
    }

    post = () => {
        let req = new XMLHttpRequest()
        req.open("POST", this.#url, true)
        req.setRequestHeader("Accept", "application/json")
        if (this.#contentType !== undefined) {
            req.setRequestHeader("Content-Type", this.#contentType)
        }
        req.onreadystatechange = () => {
            if (req.readyState === 4) {
                if (req.status === 200) {
                    this.#succeed(req.responseText)
                } else {
                    this.#failed(req.status, req.statusText)
                }
            }
        }
        req.send(this.#body)
    }

}
