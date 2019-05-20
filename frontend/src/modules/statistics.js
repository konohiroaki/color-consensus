import axios from "axios";
import {toast} from "react-toastify";

export const types = {
    SET_VOTES: "SET_VOTES",
    SET_NATIONALITY_FILTER: "SET_NATIONALITY_FILTER",
    SET_AGE_GROUP_FILTER: "SET_AGE_GROUP_FILTER",
    SET_GENDER_FILTER: "SET_GENDER_FILTER",
    SET_PERCENTILE: "SET_PERCENTILE",
    APPLY_NEW_VOTE_SET: "APPLY_NEW_VOTE_SET",
    RESET_FILTERS: "RESET_FILTERS",
};

const DEFAULT_STATE = {
    votes: [],
    filteredVoteCount: 0,

    nationalityFilter: "",
    ageGroupFilter: "",
    genderFilter: "",
    percentile: 50,

    cellBorder: [],
};

export const reducer = (state = DEFAULT_STATE, action) => {
    switch (action.type) {
        case types.SET_VOTES:
            return {
                ...state,
                votes: action.payload
            };
        case types.APPLY_NEW_VOTE_SET:
            return {
                ...state,
                cellBorder: action.payload.cellBorder,
                filteredVoteCount: action.payload.filteredVoteCount,
            };
        case types.SET_NATIONALITY_FILTER:
            return {
                ...state,
                nationalityFilter: action.payload
            };
        case types.SET_AGE_GROUP_FILTER:
            return {
                ...state,
                ageGroupFilter: action.payload
            };
        case types.SET_GENDER_FILTER:
            return {
                ...state,
                genderFilter: action.payload
            };
        case types.RESET_FILTERS:
            return {
                ...state,
                nationalityFilter: "",
                ageGroupFilter: "",
                genderFilter: "",
            };
        case types.SET_PERCENTILE:
            return {
                ...state,
                percentile: action.payload
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
                    // resetting filters because specific filter might not exist for other vote set.
                    dispatch(this.resetFilters());
                    dispatch(this.applyNewVoteSet());
                })
                .catch(({response}) => toast.warn(response.data.error.message));
        };
    },
    applyNewVoteSet() {
        return (dispatch, getState) => {
            const filteredVotes = getFilteredVotes(getState().statistics);
            const cellRatio = getCellRatio(getState, filteredVotes);
            const cellBorder = getCellBorder(getState, cellRatio);
            dispatch({
                type: types.APPLY_NEW_VOTE_SET, payload: {
                    cellBorder: cellBorder,
                    filteredVoteCount: filteredVotes.length,
                }
            });
        };
    },
    setNationalityFilter(nationality) {
        return (dispatch) => {
            dispatch({type: types.SET_NATIONALITY_FILTER, payload: nationality});
            dispatch(this.applyNewVoteSet());
        };
    },
    setAgeGroupFilter(ageGroup) {
        return (dispatch) => {
            ageGroup = ageGroup !== "" ? parseInt(ageGroup) : ageGroup;
            dispatch({type: types.SET_AGE_GROUP_FILTER, payload: ageGroup});
            dispatch(this.applyNewVoteSet());
        };
    },
    setGenderFilter(gender) {
        return (dispatch) => {
            dispatch({type: types.SET_GENDER_FILTER, payload: gender});
            dispatch(this.applyNewVoteSet());
        };
    },
    setPercentile(percentile) {
        return (dispatch) => {
            dispatch({type: types.SET_PERCENTILE, payload: percentile});
            dispatch(this.applyNewVoteSet());
        };
    },
    resetFilters() {
        return (dispatch) => {
            dispatch({type: types.RESET_FILTERS});
        };
    },
};

const getStatisticsUrl = ({category, name}) => {
    const fields = ["colors", "voter.nationality", "voter.ageGroup", "voter.gender"];
    return `${process.env.WEBAPI_HOST}/api/v1/votes?category=${category}&name=${name}&fields=${fields}`;
};

const getCellRatio = (getState, filteredVotes) => {
    const colorCodeList = getState().board.colorCodeList;
    const boardSize = getState().board.sideLength;
    const arraySize = boardSize + 2;

    let ratio = Array(arraySize).fill(0).map(() => Array(arraySize).fill(0));

    filteredVotes
        .reduce((acc, vote) => acc.concat(vote.colors), [])
        .forEach(color => {
            const idx = colorCodeList.indexOf(color);
            const ii = Math.floor(idx / boardSize) + 1, jj = idx % boardSize + 1;
            ratio[ii][jj] += 1 / filteredVotes.length;
        });
    return ratio;
};

const getFilteredVotes = ({votes, nationalityFilter, ageGroupFilter, genderFilter}) => {
    return votes
        .filter(vote => nationalityFilter === "" || nationalityFilter === vote.voter.nationality)
        .filter(vote => ageGroupFilter === "" || ageGroupFilter === vote.voter.ageGroup)
        .filter(vote => genderFilter === "" || genderFilter === vote.voter.gender);
};

const getCellBorder = (getState, cellRatio) => {
    const percentile = getState().statistics.percentile / 100;
    const arraySize = getState().board.sideLength + 2;

    let border = Array(arraySize).fill(0)
        .map(() => Array(arraySize).fill({top: false, right: false, bottom: false, left: false}));

    for (let ii = 1; ii < border.length - 1; ii++) {
        for (let jj = 1; jj < border.length - 1; jj++) {
            border[ii][jj] = {
                top: getBorderState(percentile, cellRatio[ii][jj], cellRatio[ii - 1][jj]),
                right: getBorderState(percentile, cellRatio[ii][jj], cellRatio[ii][jj + 1]),
                bottom: getBorderState(percentile, cellRatio[ii][jj], cellRatio[ii + 1][jj]),
                left: getBorderState(percentile, cellRatio[ii][jj], cellRatio[ii][jj - 1]),
            };
        }
    }
    return border;
};

const getBorderState = (percentile, current, neighbor) => {
    return percentile !== 0
           ? current >= percentile && percentile > neighbor
           : current !== 0 && neighbor === 0;
};
