import { rest } from 'msw';
import { StatusCodes } from 'http-status-codes';
import { v4 as uuid } from 'uuid';

import {
    MAX_LOGIN_LENGTH,
    MAX_PASSWORD_LENGTH,
    MAX_MESSAGES_NUMBER,
    MAX_MESSAGE_LENGTH,
    LOGIN_REGEX,
    PASSWORD_REGEX
} from '../restrictions';


function isValidCredentials(username, password) {
    if (
        username === null ||
        password === null ||
        username.length < 1 || username.length > MAX_LOGIN_LENGTH ||
        password.length < 1 || password.length > MAX_PASSWORD_LENGTH ||
        !LOGIN_REGEX.test(username) ||
        !PASSWORD_REGEX.test(password)
    ) {
        return false;
    }
    return true;
}

function rejectCredentials(res, ctx) {
    return res(
        ctx.status(StatusCodes.BAD_REQUEST),
        ctx.json({
            error: "Invalid credentials! Username or Password format is invalid.",
        }),
    );
}

function rejectAuthorization(res, ctx) {
    return res(
        ctx.status(StatusCodes.UNAUTHORIZED),
        ctx.json({
            error: "Invalid authorization! User session not found!"
        })
    )
}

function handleCreateUser(req, res, ctx, db) {
    const { username, password } = req.body;

    if (!isValidCredentials(username, password)) {
        return rejectCredentials(res, ctx);
    }

    const existingUser = db.user.findFirst({
        where: { username: { equals: username } }
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
}

function handleLogin(req, res, ctx, db) {
    const { username, password } = req.body;

    if (!isValidCredentials(username, password)) {
        return rejectCredentials(res, ctx);
    }

    const user = db.user.findFirst({
        where: { username: { equals: username } }
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
        where: { id: { equals: user.id } },
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
}

function handleLogout(req, res, ctx, db) {
    const apiToken = req.body.api_token;

    const user = db.user.findFirst({
        where: { apiToken: { equals: apiToken } }
    });

    if (user === null) {
        return rejectAuthorization(res, ctx);
    }

    db.user.update({
        where: { id: { equals: user.id } },
        data: {
            online: false,
            apiToken: '',
        }
    });

    return res(
        ctx.status(StatusCodes.OK),
        ctx.json({})
    )
}

function handleGetMembers(req, res, ctx, db) {
    const apiToken = req.body.api_token;

    const user = db.user.findFirst({
        where: { apiToken: { equals: apiToken } }
    });

    if (user === null) {
        return rejectAuthorization(res, ctx);
    }

    const members = db.user.findMany({
        where: { online: { equals: true } }
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
}


function handleGetMessages(req, res, ctx, db) {
    const apiToken = req.body.api_token;

    const user = db.user.findFirst({
        where: { apiToken: { equals: apiToken } }
    });

    if (user === null) {
        return rejectAuthorization(res, ctx);
    }

    const messages = db.message.findMany({
        take: MAX_MESSAGES_NUMBER,
        orderBy: [
            { time: 'dsc' },
            { id: 'dsc' },
        ]
    });

    const outputMessages = [];
    for (const message of messages) {
        outputMessages.push({
            id: message.id,
            member_name: message.user.username,
            text: message.text,
            time: message.time,
        });
    }

    return res(
        ctx.status(StatusCodes.OK),
        ctx.json({
            messages: outputMessages,
        }),
    );
}

function handleCreateMessage(req, res, ctx, db) {
    const apiToken = req.body.api_token;
    const text = req.body.text;

    const user = db.user.findFirst({
        where: { apiToken: { equals: apiToken } }
    });

    if (user === null) {
        return rejectAuthorization(res, ctx);
    }

    if (text.length > MAX_MESSAGE_LENGTH) {
        return res(
            ctx.status(StatusCodes.BAD_REQUEST),
            ctx.json({
                error: `Maximum message length allowed is ${MAX_MESSAGE_LENGTH}.`
            })
        )
    }

    const message = db.message.create({
        text: text,
        time: new Date(),
        user: user,
    });

    return res(
        ctx.status(StatusCodes.CREATED),
        ctx.json({
            message: {
                id: message.id,
                member_name: message.user.username,
                text: message.text,
                time: message.time,
            }
        })
    );

}

function handleCheckAppHealth(res, ctx) {
    return res(
        ctx.json({
            success: true,
            time: new Date().toISOString()
        })
    );
}

function mockHandlersFabric(db) {
    const handlers = [
        rest.get('/api/health', (req, res, ctx) => { return handleCheckAppHealth(res, ctx); }),
        rest.post('/api/user', (req, res, ctx) => { return handleCreateUser(req, res, ctx, db); }),
        rest.post('/api/login', (req, res, ctx) => { return handleLogin(req, res, ctx, db); }),
        rest.post('/api/logout', (req, res, ctx) => { return handleLogout(req, res, ctx, db); }),
        rest.post('/api/members', (req, res, ctx) => { return handleGetMembers(req, res, ctx, db); }),
        rest.post('/api/messages', (req, res, ctx) => { return handleGetMessages(req, res, ctx, db); }),
        rest.post('/api/message', (req, res, ctx) => { return handleCreateMessage(req, res, ctx, db); }),
    ];

    return handlers;
}

export default mockHandlersFabric;