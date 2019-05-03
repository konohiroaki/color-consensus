import axios from "axios";
import {toast} from "react-toastify";

export const types = {
    SET_ID: "SET_ID",
};

const DEFAULT_STATE = {
    id: null,
};

export const reducer = (state = DEFAULT_STATE, action) => {
    switch (action.type) {
        case types.SET_ID:
            return {
                id: action.payload
            };
        default:
            return state;
    }
};

export const actions = {
    verifyLoginState() {
        return (dispatch) => {
            return axios.get(`${process.env.WEBAPI_HOST}/api/v1/users/presence`)
                .then(({data}) => dispatch({
                    type: types.SET_ID,
                    payload: data.userID
                }))
                .catch(({response}) => {
                    if (response.status !== 404) {
                        toast.warn(response.data.error.message);
                    }
                });
        };
    },
    login(id) {
        return (dispatch) => {
            return axios.post(`${process.env.WEBAPI_HOST}/api/v1/login`, {userID: id})
                .then(() => dispatch({
                    type: types.SET_ID,
                    payload: id
                }));
        };
    },
    signUp(nationality, gender, birth) {
        return (dispatch) => {
            return axios.post(`${process.env.WEBAPI_HOST}/api/v1/users`, {
                nationality: nationality,
                gender: gender,
                birth: Number(birth)
            })
                .then(({data}) => dispatch({
                    type: types.SET_ID,
                    payload: data.userID
                }));
        };
    }
};
