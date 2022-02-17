async function restRequest(method, url, payload) {
    const response = await fetch(
        url,
        {
            method: method,
            headers: {
                'Content-type': 'application/json',
                'Accept': 'application/json'
            },
            body: JSON.stringify(payload)
        }
    );

    const data = await response.json();

    if (!response.ok) {
        if (data.hasOwnProperty('error')) {
            throw new Error(data.error);
        } else {
            throw new Error(`Server error: ${response.status}: ${response.statusText}`)
        }
    }

    return data;
}

function validateResponseDataFields(jsonResponse, requiredFields) {
    if (!Array.isArray(requiredFields)) {
        requiredFields = [requiredFields];
    }

    for(const field of requiredFields) {
        if (!jsonResponse.hasOwnProperty(field)) {
            throw new Error(`Bad server response. Response does not have ${field} field.`);
        }
    }
}

class ApiClientRest {
    // ApiClient interface
    // login(username, password) -> userName
    // register(username, password) -> userName
    // logout() -> undefined
    // getMembers() -> [members]
    // getMessages() -> [messages]
    // sendMessage(text) -> message

    constructor() {
        this.apiToken = null;
    }

    isAuthorized() {
        return this.apiToken !== null;
    }

    requireAuthorization() {
        if (!this.isAuthorized()) {
            throw new Error("ApiClient is not authorized.");
        }
    }

    async login(username, password) {
        // POST /api/login
        // {username: "string", password: "string"}

        const data = await restRequest(
            'POST',
            '/api/login',
            {
                username: username,
                password: password
            }
        );

        validateResponseDataFields(data, ['api_token', 'member']);

        this.apiToken = data.api_token;
        return data.member.name;
    }

    async register(username, password) {
        // POST /api/user
        // {username: "string", password: "string"}

        const data = await restRequest(
            'POST',
            '/api/user',
            {
                username: username,
                password: password
            }
        );

        validateResponseDataFields(data, ['username']);

        return data.username;
    }

    async logout() {
        // POST /api/logout
        // {api_token: "string"}

        this.requireAuthorization();

        await restRequest(
            'POST',
            '/api/logout',
            {
                api_token: this.apiToken
            }
        );

        this.apiToken = null;
    }

    async getMembers() {
        // POST /api/members
        // {api_token: "string"}

        this.requireAuthorization();

        const data = await restRequest(
            'POST',
            '/api/members',
            {
                api_token: this.apiToken
            }
        );

        validateResponseDataFields(data, ['members']);

        return data.members;
    }

    async getMessages() {
        // POST /api/messages
        // {api_token: "string"}

        this.requireAuthorization();

        const data = await restRequest(
            'POST',
            '/api/messages',
            {
                api_token: this.apiToken
            }
        );

        validateResponseDataFields(data, ['messages']);

        return data.messages;
    }

    async sendMessage(text) {
        // POST /api/message
        // {api_token: "string", text: "string"}

        this.requireAuthorization();

        const data = await restRequest(
            'POST',
            '/api/message',
            {
                api_token: this.apiToken,
                text: text
            }
        );

        validateResponseDataFields(data, ['message']);

        return data.message;
    }
}

export default ApiClientRest;