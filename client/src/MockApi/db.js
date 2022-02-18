import { factory, oneOf, primaryKey } from "@mswjs/data";

function fillMockData(db) {
    db.user.create({ username: 'admin', password: 'admin_password', online: false });
    const user1 = db.user.create({ username: 'User-1', password: 'User-1_password', online: false });
    const user2 = db.user.create({ username: 'User-2', password: 'User-2_password', online: false });
    const user3 = db.user.create({ username: 'User-3', password: 'User-3_password', online: false });
    const user4 = db.user.create({
        username: 'User-4', password: 'User-4_password', online: true,
        apiToken: "89bf4093-a1b2-4481-bb41-805c60e97e22"
    });
    const user5 = db.user.create({
        username: 'User-5', password: 'User-5_password', online: true,
        apiToken: "e7d41ba0-3557-4070-a6c3-8ddf22c1b9ff"
    });
    const user6 = db.user.create({
        username: 'User-6', password: 'User-6_password', online: true,
        apiToken: "25288c2a-928d-4765-b7ac-2a82d54c818e"
    });

    db.message.create({ user: user1, text: 'Hello from User-1', time: new Date(2022, 2, 1, 12, 10, 20) });
    db.message.create({ user: user2, text: 'Hello from User-2', time: new Date(2022, 2, 1, 12, 10, 23) });
    db.message.create({ user: user3, text: 'Hello from User-3', time: new Date(2022, 2, 1, 12, 10, 25) });
    db.message.create({ user: user4, text: 'Hello from User-4', time: new Date(2022, 2, 1, 12, 10, 27) });
    db.message.create({ user: user5, text: 'Hello from User-5', time: new Date(2022, 2, 1, 12, 10, 28) });
    db.message.create({ user: user6, text: 'Hello from User-6', time: new Date(2022, 2, 1, 12, 10, 29) });
    db.message.create({ user: user1, text: 'This is second message from User-1', time: new Date(2022, 2, 1, 12, 12, 30) });
    db.message.create({ user: user1, text: 'And third message from User-1', time: new Date(2022, 2, 1, 12, 15, 35) });
    db.message.create({ user: user6, text: 'Сообщение, содержащее кириллицу.', time: new Date(2022, 2, 1, 13, 1, 17) });
    db.message.create({ user: user6, text: 'Попытка засунуть тег внутрь сообщения <strong>Вери Стронг!</strong>', time: new Date(2022, 2, 1, 13, 1, 37) });
}

function* idGenerator() {
    let id = 1;
    for (; ;) {
        yield id++;
    }
}

// TODO: Make mockDB abstraction to implement Dependency Inversion principle
function mockDbFabric() {
    const usersPk = idGenerator();
    const messagesPk = idGenerator();

    const db = factory({
        user: {
            id: primaryKey(() => usersPk.next().value),
            username: String,
            password: String,
            online: Boolean,
            apiToken: String,
        },
        message: {
            id: primaryKey(() => messagesPk.next().value),
            user: oneOf('user'),
            text: String,
            time: Date,
        }
    });

    fillMockData(db);

    return db;
}

export default mockDbFabric;