import axios from "axios";

export const types = {
    SET_VOTES: "SET_VOTES",
    SET_NATIONALITY_FILTER: "SET_NATIONALITY_FILTER",

};

const DEFAULT_STATE = {
    votes: [],
    nationalityFilter: "",
    genderFilter: "",
    ageGroupFilter: "",
};

export const reducer = (state = DEFAULT_STATE, action) => {
    switch (action.type) {
        case types.SET_VOTES:
            return {
                ...state,
                votes: action.payload
            };
        case types.SET_NATIONALITY_FILTER:
            return {
                ...state,
                nationalityFilter: action.payload
            };
        default:
            return state;
    }
};

export const actions = {
    setVotes(color) {
        return (dispatch) => {
            return axios.get(getStatisticsUrl(color))
                .then(({data}) => {
                    dispatch({type: types.SET_VOTES, payload: data});
                })
                .catch(err => {});
        };
    },
    setNationalityFilter(nationality) {
        return {
            type: types.SET_NATIONALITY_FILTER,
            payload: nationality,
        };
    },
};

const getStatisticsUrl = ({lang, name}) => {
    const fields = ["colors", "voter.nationality", "voter.gender", "voter.birth"];
    return `${process.env.WEBAPI_HOST}/api/v1/votes?lang=${lang}&name=${name}&fields=${fields}`;
};
