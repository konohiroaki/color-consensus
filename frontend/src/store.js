import {createStore, combineReducers, applyMiddleware} from "redux";
import thunk from "redux-thunk";
import {composeWithDevTools} from "redux-devtools-extension";

import {reducer as searchBar} from "./modules/searchBar";
import {reducer as user} from "./modules/user";
import {reducer as board} from "./modules/board";
import {reducer as vote} from "./modules/vote";
import {reducer as statistics} from "./modules/statistics";

const rootReducer = combineReducers({
    user,
    searchBar,
    board,
    vote,
    statistics,
});

export default createStore(rootReducer, composeWithDevTools(
    applyMiddleware(thunk)
));
