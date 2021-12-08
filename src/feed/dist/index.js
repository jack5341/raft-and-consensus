"use strict";

var _express = _interopRequireDefault(require("express"));

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { "default": obj }; }

var app = (0, _express["default"])();

function RandomSetTimeoutGenerator() {
  var counter = 0;

  for (;;) {
    counter = Math.floor(Math.random() * 3 + 1);
    setTimeout(function () {
      console.log(counter);
    }, counter);
  }
}

RandomSetTimeoutGenerator();
app.get("/", function (req, res) {
  console.log("Tick received to Node.js!");
  res.send({
    status: "ok",
    voter: 1
  });
  console.log("Vote sent by Node.js!");
  return;
});
var port = process.env.PORT || 3000;
app.listen(port, function () {
  console.log("Server listening on port ".concat(port));
});