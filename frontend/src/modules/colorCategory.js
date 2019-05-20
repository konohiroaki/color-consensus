import axios from "axios";
import {toast} from "react-toastify";

export const types = {
    SET_COLOR_CATEGORY: "SET_COLOR_CATEGORY",
};

const DEFAULT_STATE = {
    categories: []
};

export const reducer = (state = DEFAULT_STATE, action) => {
    switch (action.type) {
        case types.SET_COLOR_CATEGORY:
            return {
                ...state,
                categories: action.payload
            };
        default:
            return state;
    }
};

export const actions = {
    fetchColorCategories() {
        return (dispatch) => {
            return axios.get(`${process.env.WEBAPI_HOST}/api/v1/color-categories`)
                .then(({data}) => dispatch({type: types.SET_COLOR_CATEGORY, payload: data}))
                .catch(({response}) => toast.warn(response.data.error.message));
        };
    },
};
