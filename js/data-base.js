let mysql = require('mysql');
const pool = mysql.createConnection({
    host: "127.0.0.1",
    user: "root",
    password: "password",
    database: "data-access"
})

pool.connect(function(err) {
    if (err) throw err;
    console.log('connected!');
});

function liked() {
    pool.query(`INSERT INTO data_like (Signe, Author, Target)
VALUES
('positif', 'Saint-Didier-des-Bois', 'message')`, (err, result, fields) => {
        if (err) {
            return console.log(err);
        }
        return console.log(result);
    })
    var element = document.getElementById("like");
    element.classList.toggle("liked");
}