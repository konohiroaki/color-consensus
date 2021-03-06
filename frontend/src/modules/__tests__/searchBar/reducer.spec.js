import {reducer, types} from "../../searchBar";

it("should set colors", () => {
    const fakeData = [{foo: "bar"}];
    const result = reducer(undefined, {
        type: types.FETCH_COLORS_SUCCESS,
        payload: fakeData,
    });
    expect(result.colors).toEqual(fakeData);
});
