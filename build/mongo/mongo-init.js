db.createUser({
    user: "user",
    pwd: "password",
    roles: [{
        role: "readWrite",
        db: "t-board"
    }]
});
db.createCollection('users');
db.users.createIndex({ email: 1 }, { unique: true });