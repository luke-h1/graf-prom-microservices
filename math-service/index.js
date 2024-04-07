const express = require("express");
const axios = require("axios").default;

const app = express();

app.use(express.json());

const delay = process.env.DELAY || 100;

app.get("/add/:num1/:num2", async (req, res) => {
  const num1 = Number(req.params.num1);
  const num2 = Number(req.params.num2);
  return res.status(200).json(
    JSON.stringify({
      value: num1 + num2,
    })
  );
});

app.get("/double", async (req, res) => {
  if (!req.body.value) {
    return res.status(400).send({ error: "value is required" });
  }

  const num = Number(req.body.value);

  const url = `http://localhost:8080/add/${num}/${num}`;

  try {
    const response = await axios.get(url);
    return res.status(200).send({
      value: response.data.value,
    });
  } catch (e) {
    return res.status(500).send({ error: e.message ?? "An error occured" });
  }
});

app.post("/logger", async (req, res) => {
  setTimeout(() => {
    console.info("Received event:\n", JSON.stringify(req.body, null, 2));
    return res.status(200).send(JSON.stringify(req.body));
  }, delay);
});

const server = app.listen(8080, () => {
  console.info("Math service runnin on port:", server.address().port);
});
