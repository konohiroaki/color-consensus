import {createStore, combineReducers, applyMiddleware} from "redux";
// import {reducer as freezer} from "./ducks/freezer";
// import {reducer as orders} from "./ducks/orders";
// import {reducer as employees} from "./ducks/employees";
import {reducer as colors} from "./ducks/colors";
import thunk from "redux-thunk";
import {composeWithDevTools} from "redux-devtools-extension";

const rootReducer = combineReducers({
    // freezer,
    // orders,
    // employees,
    colors,
});

export default createStore(rootReducer, composeWithDevTools(
    applyMiddleware(thunk)
));
