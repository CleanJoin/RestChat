import { setupServer } from "msw/node";

export class ApiBuilderServer {
    constructor() {
        this.server = null;
    }

    setup(...handlers) {
        this.server = setupServer(...handlers);
    }

    start() {
        if (this.server !== null) {
            this.server.listen({ onUnhandledRequest: 'error' });
        } else {
            throw new Error("MockApi server was not initialized before listen() call");
        }
    }

    stop() {
        if (this.server !== null) {
            this.server.close();
        } else {
            throw new Error("MockApi server was not initialized before stop() call");
        }
    }
}
