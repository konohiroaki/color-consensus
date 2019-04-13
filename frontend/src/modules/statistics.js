import axios from "axios";

export const types = {
    SET_VOTES: "SET_VOTES",
    SET_NATIONALITY_FILTER: "SET_NATIONALITY_FILTER",
    SET_PERCENTILE: "SET_PERCENTILE",
    CALCULATE_BORDER: "CALCULATE_BORDER",
    RESET_FILTERS: "RESET_FILTERS",
};

const DEFAULT_STATE = {
    votes: [],
    nationalityFilter: "",
    genderFilter: "",
    ageGroupFilter: "",
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
        case types.CALCULATE_BORDER:
            return {
                ...state,
                cellBorder: action.payload
            };
        case types.SET_NATIONALITY_FILTER:
            return {
                ...state,
                nationalityFilter: action.payload
            };
        case types.SET_PERCENTILE:
            return {
                ...state,
                percentile: action.payload
            };
        case types.RESET_FILTERS:
            return {
                ...state,
                nationalityFilter: "",
                genderFilter: "",
                ageGroupFilter: "",
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
                    dispatch(this.calculateBorder());
                })
                .catch(err => {});
        };
    },
    calculateBorder() {
        return (dispatch, getState) => {
            const cellRatio = getCellRatio(getState);
            const cellBorder = getCellBorder(getState, cellRatio);
            dispatch({type: types.CALCULATE_BORDER, payload: cellBorder});
        };
    },
    setNationalityFilter(nationality) {
        return (dispatch) => {
            dispatch({type: types.SET_NATIONALITY_FILTER, payload: nationality});
            dispatch(this.calculateBorder());
        };
    },
    setPercentile(percentile) {
        return (dispatch) => {
            dispatch({type: types.SET_PERCENTILE, payload: percentile});
            dispatch(this.calculateBorder());
        };
    }
};

const getStatisticsUrl = ({lang, name}) => {
    const fields = ["colors", "voter.nationality", "voter.gender", "voter.birth"];
    return `${process.env.WEBAPI_HOST}/api/v1/votes?lang=${lang}&name=${name}&fields=${fields}`;
};

const getCellRatio = (getState) => {
    const colorCodeList = getState().board.colorCodeList;
    const votes = getState().statistics.votes;
    const nationalityFilter = getState().statistics.nationalityFilter;
    const boardSize = getState().board.sideLength;
    const arraySize = boardSize + 2;

    let ratio = Array(arraySize).fill(0).map(() => Array(arraySize).fill(0));

    const filteredVotes = votes
        .filter(vote => nationalityFilter === "" || nationalityFilter === vote.voter.nationality);
    filteredVotes.flatMap(vote => vote.colors)
        .forEach(color => {
            const idx = colorCodeList.indexOf(color);
            const ii = Math.floor(idx / boardSize) + 1, jj = idx % boardSize + 1;
            ratio[ii][jj] += 1 / filteredVotes.length;
        });
    return ratio;
};

const getCellBorder = (getState, cellRatio) => {
    const arraySize = getState().board.sideLength + 2;
    let border = Array(arraySize).fill(0)
        .map(() => Array(arraySize).fill({top: false, right: false, bottom: false, left: false}));
    const percentile = getState().statistics.percentile / 100;

    for (let ii = 1; ii < border.length - 1; ii++) {
        for (let jj = 1; jj < border.length - 1; jj++) {
            border[ii][jj] = {
                top: percentile !== 0
                     ? cellRatio[ii][jj] >= percentile && cellRatio[ii - 1][jj] < percentile
                     : cellRatio[ii][jj] !== 0 && cellRatio[ii - 1][jj] === 0,
                right: percentile !== 0
                       ? cellRatio[ii][jj] >= percentile && cellRatio[ii][jj + 1] < percentile
                       : cellRatio[ii][jj] !== 0 && cellRatio[ii][jj + 1] === 0,
                bottom: percentile !== 0
                        ? cellRatio[ii][jj] >= percentile && cellRatio[ii + 1][jj] < percentile
                        : cellRatio[ii][jj] !== 0 && cellRatio[ii + 1][jj] === 0,
                left: percentile !== 0
                      ? cellRatio[ii][jj] >= percentile && cellRatio[ii][jj - 1] < percentile
                      : cellRatio[ii][jj] !== 0 && cellRatio[ii][jj - 1] === 0,
            };
        }
    }
    return border;
};
