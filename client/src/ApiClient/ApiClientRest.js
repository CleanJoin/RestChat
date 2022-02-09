class ApiClientRest {
    // ApiClient interface
    // login(username, password) -> userName, error
    // register(username, password) -> userName, error
    // logout() -> true/false, error
    // getMembers() -> [members], error
    // getMessages() -> [messages], error
    // sendMessage(text) -> message, error

    constructor() {
        this.api_token = null;
    }

    _request(method, url, payload) {
        const xmlHttp = new XMLHttpRequest();
        xmlHttp.open(method, url, false);
        xmlHttp.send(payload);
        const status = xmlHttp.status;
        const response = xmlHttp.responseText;
        return [status, response];
    }

    _isAuthorized() {
        return this.apiToken == null;
    }

    login(username, password) {
        // POST /api/login
        // {username: "string", password: "string"}
        let error = null ;

        const [status, response] = this._request(
            'POST',
            '/api/login',
            {
                username: username,
                password: password
            }
        );

        // TODO: DEBUG
        console.log("ApiClient.login()", status, response);

        this.apiToken = response.api_token;
        return [username, error]
    }

    register(username, password) {
        // POST /api/user
        // {username: "string", password: "string"}
        let userName = null;
        let error = null;

        const [status, response] = this._request(
            'POST',
            '/api/user',
            {
                username: username,
                password: password
            }
        );

        // TODO: DEBUG
        console.log("ApiClient.register()", status, response);

        return [userName, error]
    }

    logout() {
        // POST /api/logout
        // {api_token: "string"}

        if (!this._isAuthorized()) {
            return "ApiClient was not authorized";
        }

        const [status, response] = this._request(
            'POST',
            '/api/logout',
            {
                api_token: this.apiToken
            }
        );

        // TODO: DEBUG
        console.log("ApiClient.register()", status, response);

        this.apiToken = undefined;
        return null
    }

    getMembers() {
        // GET /api/members
        // {api_token: "string"}
        let members = [];
        let error = undefined;

        if (!this._isAuthorized()) {
            return [null, "ApiClient not authorized"]
        }

        const [status, response] = this._request(
            'GET',
            '/api/members',
            {
                api_token: this.apiToken
            }
        );

        // TODO: DEBUG
        console.log("ApiClient.getMembers()", status, response);

        return [members, error];
    }

    getMessages() {
        // GET /api/messages
        // {api_token: "string"}
        let messages = [];
        let error = undefined;

        if (!this._isAuthorized()) {
            return [null, "ApiClient not authorized"]
        }

        const [status, response] = this._request(
            'GET',
            '/api/messages',
            {
                api_token: this.apiToken
            }
        );

        // TODO: DEBUG
        console.log("ApiClient.getMembers()", status, response);

        return [messages, error]
    }

    sendMessage(text) {
        // POST /api/message
        // {api_token: "string", text: "string"}
        let message = undefined;
        let error = undefined;

        if (!this._isAuthorized()) {
            return [null, "ApiClient not authorized"]
        }

        const [status, response] = this._request(
            'POST',
            '/api/message',
            {
                api_token: this.apiToken,
                text: text
            }
        );

        // TODO: DEBUG
        console.log("ApiClient.getMembers()", status, response);

        return [message, error]
    }
}

export default ApiClientRest;