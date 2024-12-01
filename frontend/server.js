const express = require("express");
const path = require("path");
const app = express();
require("dotenv").config({ path: "../.env" });

app.get("/api/port", (req, res) => {
    res.json({ port: process.env.PORT || 8080 }); // Send the port number to the client
});

app.use("/static", express.static(path.resolve(__dirname, "./", "static")));

app.get("/*", (req, res) => {
    res.sendFile(path.resolve(__dirname, "./", "index.html"));
});

app.listen(process.env.PORTF || 5173, () => console.log("Server is running..."));
