import {reducer, types} from "./colors";

it("should set colors", () => {
    const fakeData = [{foo: "bar"}];
    const result = reducer(undefined, {
        type: types.FETCH_COLORS_SUCCESS,
        payload: fakeData,
    });
    expect(result.colors).toEqual(fakeData);
});
it("should set displayed color", () => {
    const fakeData = {foo: "bar"};
    const result = reducer(undefined, {
        type: types.SET_DISPLAYED_COLOR,
        payload: fakeData,
    });
    expect(result.displayedColor).toEqual(fakeData);

});
it("should set displayed color list", () => {
    const fakeData = [{foo: "bar"}];
    const result = reducer(undefined, {
        type: types.SET_DISPLAYED_COLOR_LIST,
        payload: fakeData,
    });
    expect(result.displayedColorList).toEqual(fakeData);

});
