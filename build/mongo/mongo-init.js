db.createUser({
    user: "user",
    pwd: "password",
    roles: [{
        role: "readWrite",
        db: "t-mail"
    }]
});
db.createCollection('users');
db.users.createIndex({ email: 1 }, { unique: true });