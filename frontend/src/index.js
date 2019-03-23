import React from "react";
import ReactDom from "react-dom";
import "bootstrap/dist/css/bootstrap.css";
import "bootstrap";
import "./base.css";

import "core-js/fn/object/assign";
import "core-js/fn/array/from";
import "core-js/fn/array/is-array";
import "core-js/fn/map";
import "core-js/fn/set";

import App from "./components/App";

import store from "./store";
// import {actions as freezer} from "./ducks/freezer";
// import * as FLAVORS from "./constants/flavors";
// import {actions as employees} from "./ducks/employees";
import {actions as colors} from "./ducks/colors";

// store.dispatch(freezer.updateTemperature(-8));
// store.dispatch(freezer.addProductToFreezer(FLAVORS.VANILLA, 5));
// store.dispatch(freezer.doSomething());
// store.dispatch(employees.fetchEmployees());
store.dispatch(colors.fetchColors());

const appRoot = document.createElement("div");
document.body.appendChild(appRoot);
ReactDom.render(<App/>, appRoot);
