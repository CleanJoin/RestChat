import { setupWorker } from 'msw';
import handlers from './handlers';

const browser = setupWorker(...handlers);

export default browser;