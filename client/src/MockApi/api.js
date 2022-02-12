import { setupWorker } from 'msw';

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
        this.reset()
        this.server.start();
    }

    stop() {
        this.server.stop();
    }
}
