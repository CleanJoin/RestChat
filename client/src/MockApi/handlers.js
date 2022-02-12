import { rest } from 'msw';

function mockHandlersFabric(db) {

    const handlers = [
        rest.get('/api/health', (req, res, ctx) => {
            return res(
                ctx.json({
                    success: true,
                    time: new Date().toISOString()
                })
            );
        }),

        rest.post('/api/user', (req, res, ctx) => {
            const { username, password } = req.body;

            db.user.create({ name: username, password: password });

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

    return handlers;
}

export default mockHandlersFabric;