import {reducer, types} from "../../statistics";

it("should set votes", () => {
    const fakeData = [{
        colors: ["#ff0000"],
        voter: {
            nationality: "Japan",
            ageGroup: 20,
            gender: "Male",
        }
    }];
    const result = reducer(undefined, {
        type: types.SET_VOTES,
        payload: fakeData,
    });
    expect(result.votes).toEqual(fakeData);
});

it("should set borders and filteredVoteCount", () => {
    const fakeData = {
        cellBorder: [
            [{top: true, right: true, bottom: false, left: true}],
            [{top: false, right: true, bottom: true, left: true}]
        ],
        filteredVoteCount: 1
    };
    const result = reducer(undefined, {
        type: types.APPLY_NEW_VOTE_SET,
        payload: fakeData,
    });
    expect(result.cellBorder).toEqual(fakeData.cellBorder);
    expect(result.filteredVoteCount).toEqual(fakeData.filteredVoteCount);

});

it("should set nationality filter", () => {
    const fakeData = "Japan";
    const result = reducer(undefined, {
        type: types.SET_NATIONALITY_FILTER,
        payload: fakeData,
    });
    expect(result.nationalityFilter).toEqual(fakeData);
});

it("should set ageGroup filter", () => {
    const fakeData = 10;
    const result = reducer(undefined, {
        type: types.SET_AGE_GROUP_FILTER,
        payload: fakeData,
    });
    expect(result.ageGroupFilter).toEqual(fakeData);
});

it("should set gender filter", () => {
    const fakeData = "Male";
    const result = reducer(undefined, {
        type: types.SET_GENDER_FILTER,
        payload: fakeData,
    });
    expect(result.genderFilter).toEqual(fakeData);
});

const expectAllFiltersEmpty = (result) => {
    expect(result.nationalityFilter).toEqual("");
    expect(result.ageGroupFilter).toEqual("");
    expect(result.genderFilter).toEqual("");

};
it("should reset all filters", () => {
    const fakeState = {
        nationalityFilter: "test",
        ageGroupFilter: 20,
        genderFilter: "test",
    };
    const result = reducer(fakeState, {
        type: types.RESET_FILTERS,
    });
    expectAllFiltersEmpty(result);
});

it("should set percentile", () => {
    const fakeData = 10;
    const result = reducer(undefined, {
        type: types.SET_PERCENTILE,
        payload: fakeData,
    });
    expect(result.percentile).toEqual(fakeData);
});
