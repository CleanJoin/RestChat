import { setupServer, setupWorker } from 'msw/node';

export class ApiBuilderServer {
    constructor() {
        this.server = null;
    }

    setup(...handlers) {
        this.server = setupServer(...handlers);
    }

    start() {
        if (this.server !== null) {
            this.server.listen();
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

export class ApiBuilderWorker {
    constructor() {
        this.worker = null;
    }

    setup(...handlers) {
        this.worker = setupWorker(...handlers);
    }

    start() {
        if (this.worker !== null) {
            this.worker.start();
        } else {
            throw new Error("MockApi worker was not initialized before start() call");
        }
    }

    stop() {
        if (this.worker !== null) {
            this.worker.stop();
        } else {
            throw new Error("MockApi worker was not initialized before stop() call");
        }
    }
}


export class MockApiServer {
    constructor(apiBuilder, dbFabric, handlersFabric) {
        this.server = apiBuilder;
        this.dbFabric = dbFabric;
        this.handlersFabric = handlersFabric;
        this.db = null;
        this.handlers = null;
    }

    reset() {
        this.db = this.dbFabric();
        this.handlers = this.handlersFabric(this.db);
        this.server.setup(...this.handlers)
    }

    start() {
        this.server.start();
    }

    stop() {
        this.server.stop();
    }
}
