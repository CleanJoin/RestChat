import { rest } from 'msw';

const handlers = [
    rest.get('/api/health', (req, res, ctx) => {
        console.log('MockApiWorker /health request:', req);
        return res(
            ctx.json({
                success: true,
                time: new Date().toISOString()
            })
        );
    }),

    rest.post('/api/user', (req, res, ctx) => {
        console.log("msw hit /api/user")
        console.log("msw", req)
        const { username, password } = req.body;
        return res(
            ctx.json({
                username: "Vasya",
            })
        );
    }),

    rest.post('/api/login', (req, res, ctx) => {
        const { username, password } = req.body;
        return res(
            ctx.json({
                username: "Vasya",
                password: "Vasya-password"
            })
        );
    }),

];

export default handlers;