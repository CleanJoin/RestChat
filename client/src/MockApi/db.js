import { factory, primaryKey } from "@mswjs/data";

function* idGenerator() {
    let id = 0;
    for (; ;) {
        yield id++;
    }
}

function mockDbFabric() {
    const usersPk = idGenerator();

    const db = factory({
        user: {
            id: primaryKey(() => usersPk.next()),
            username: String,
            password: String,
            apiToken: String,
            online: Boolean,
        }
    });

    db.user.create({ username: 'User-1' });
    db.user.create({ username: 'User-2' });
    db.user.create({ username: 'User-3' });
    db.user.create({ username: 'User-4' });
    db.user.create({ username: 'User-5' });

    return db;
}

export default mockDbFabric;