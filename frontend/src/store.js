import {createStore, combineReducers, applyMiddleware} from "redux";
import thunk from "redux-thunk";
import {composeWithDevTools} from "redux-devtools-extension";

import {reducer as searchBar} from "./modules/searchBar";
import {reducer as user} from "./modules/user";
import {reducer as board} from "./modules/board";
import {reducer as vote} from "./modules/vote";
import {reducer as statistics} from "./modules/statistics";
import {reducer as colorCategory} from "./modules/colorCategory";
import {reducer as nationality} from "./modules/nationality";
import {reducer as gender} from "./modules/gender";

const rootReducer = combineReducers({
    user,
    searchBar,
    board,
    vote,
    statistics,
    colorCategory,
    nationality,
    gender,
});

export default createStore(rootReducer, composeWithDevTools(
    applyMiddleware(thunk)
));
