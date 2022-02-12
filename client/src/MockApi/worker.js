import { setupWorker } from "msw";

export class ApiBuilderWorker {
    constructor() {
        this.worker = null;
    }

    setup(...handlers) {
        this.worker = setupWorker(...handlers);
        console.log("worker setup is called");
    }

    start() {
        console.log("this.worker:", this.worker)
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
