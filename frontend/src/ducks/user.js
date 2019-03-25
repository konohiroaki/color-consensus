import axios from "axios";

const types = {
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
        return (dispatch, getState) => {
            return axios.get(`${process.env.WEBAPI_HOST}/api/v1/users/presence`)
                .then(({data}) => dispatch({
                    type: types.SET_ID,
                    payload: data.userID
                }))
                .catch((err) => console.log(err));
        };
    },
    tryLogin(id) {
        return (dispatch, getState) => {
            return axios.post(`${process.env.WEBAPI_HOST}/api/v1/users/presence`, {id: id})
                .then(() => dispatch({
                    type: types.SET_ID,
                    payload: id
                }))
                .catch((err) => console.log(err));
        };
    },
    signUp(nationality, gender, birth) {
        return (dispatch, getState) => {
            return axios.post(`${process.env.WEBAPI_HOST}/api/v1/users`, {
                nationality: nationality,
                gender: gender,
                birth: Number(birth)
            })
                .then(({data}) => dispatch({
                    type: types.SET_ID,
                    payload: data.id
                }))
                .catch((err) => console.log(err));
        };
    }
};
