import axios from "axios";
import {toast} from "react-toastify";

export const types = {
    SET_NATIONALITIES: "SET_NATIONALITIES",
};

const DEFAULT_STATE = {
    nationalities: {}
};

export const reducer = (state = DEFAULT_STATE, action) => {
    switch (action.type) {
        case types.SET_NATIONALITIES:
            return {
                ...state,
                nationalities: action.payload
            };
        default:
            return state;
    }
};

export const actions = {
    setNationalities() {
        return (dispatch) => {
            return axios.get(`${process.env.WEBAPI_HOST}/api/v1/nationalities`)
                .then(({data}) => dispatch({type: types.SET_NATIONALITIES, payload: data}))
                .catch(({response}) => toast.warn(response.data.error.message));
        };
    },
};
