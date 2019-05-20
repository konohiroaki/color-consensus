import axios from "axios";
import {toast} from "react-toastify";

export const types = {
    SET_GENDERS: "SET_GENDERS",
};

const DEFAULT_STATE = {
    genders: []
};

export const reducer = (state = DEFAULT_STATE, action) => {
    switch (action.type) {
        case types.SET_GENDERS:
            return {
                ...state,
                genders: action.payload
            };
        default:
            return state;
    }
};

export const actions = {
    fetchGenders() {
        return (dispatch) => {
            return axios.get(`${process.env.WEBAPI_HOST}/api/v1/genders`)
                .then(({data}) => dispatch({type: types.SET_GENDERS, payload: data}))
                .catch(({response}) => toast.warn(response.data.error.message));
        };
    },
};
