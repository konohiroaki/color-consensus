import React from "react";
import ReactDom from "react-dom";
import "bootstrap/dist/css/bootstrap.css";
import "bootstrap";
import "./test.css";

import "core-js/fn/object/assign";
import "core-js/fn/array/from";
import "core-js/fn/array/is-array";
import "core-js/fn/map";
import "core-js/fn/set";

import App from "./components/App";

ReactDom.render(<App/>, document.getElementById("app"));
