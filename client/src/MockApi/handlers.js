import { rest } from 'msw';
import { StatusCodes } from 'http-status-codes';
import { v4 as uuid } from 'uuid';

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

            const existingUser = db.user.findFirst({
                where: { username: {equals: username} }
            });

            if (existingUser !== null) {
                return res(
                    ctx.status(StatusCodes.FORBIDDEN),
                    ctx.json({
                        error: "User already exists.",
                    })
                );
            }

            db.user.create({ username: username, password: password, online: false });

            return res(
                ctx.status(StatusCodes.CREATED),
                ctx.json({
                    username: username,
                })
            );
        }),

        rest.post('/api/login', (req, res, ctx) => {
            const { username, password } = req.body;

            const user = db.user.findFirst({
                where: { username: {equals: username} }
            });

            if (user === null || user.password !== password) {
                return res(
                    ctx.status(StatusCodes.UNAUTHORIZED),
                    ctx.json({
                        error: "Invalid username or password!",
                    })
                );
            }

            const apiToken = uuid();

            db.user.update({
                where: { id: { equals: user.id }},
                data: {
                    online: true,
                    apiToken: apiToken,
                }
            });

            return res(
                ctx.status(StatusCodes.OK),
                ctx.json({
                    api_token: apiToken,
                    member: {
                        id: user.id,
                        name: user.username,
                    }
                })
            );
        }),

        rest.post('/api/logout', (req, res, ctx) => {
            const apiToken = req.body.api_token;

            // TODO: extract authorization method
            const user = db.user.findFirst({
                where: { apiToken: {equals: apiToken }}
            });

            if (user === null) {
                return res(
                    ctx.status(StatusCodes.BAD_REQUEST),
                    ctx.json({
                        error: "User session not found!"
                    })
                );
            }

            db.user.update({
                where: { id: { equals: user.id }},
                data: {
                    online: false,
                    apiToken: '',
                }
            });

            return res(
                ctx.status(StatusCodes.OK),
                ctx.json({})
            )
        }),

        rest.post('/api/members', (req, res, ctx) => {
            const apiToken = req.body.api_token;

            // TODO: extract authorization method
            const user = db.user.findFirst({
                where: { apiToken: {equals: apiToken }}
            });

            if (user === null) {
                return res(
                    ctx.status(StatusCodes.UNAUTHORIZED),
                    ctx.json({
                        error: "User session not found!"
                    })
                )
            }

            const members = db.user.findMany({
                where: { online: {equals: true}}
            });

            const onlineMembers = [];
            for (const member of members) {
                onlineMembers.push({
                    id: member.id,
                    name: member.username,
                });
            }

            return res(
                ctx.status(StatusCodes.OK),
                ctx.json({
                    members: onlineMembers
                }),
            );
        }),

        rest.post('/api/messages', (req, res, ctx) => {
            const apiToken = req.body.api_token;

            // TODO: extract authorization method
            const user = db.user.findFirst({
                where: { apiToken: {equals: apiToken }}
            });

            if (user === null) {
                return res(
                    ctx.status(StatusCodes.UNAUTHORIZED),
                    ctx.json({
                        error: "User session not found!"
                    })
                )
            }

            const messages = db.message.getAll();

            return res(
                ctx.status(StatusCodes.OK),
                ctx.json({
                    messages: messages,
                }),
            );
        }),


    ];

    return handlers;
}

export default mockHandlersFabric;