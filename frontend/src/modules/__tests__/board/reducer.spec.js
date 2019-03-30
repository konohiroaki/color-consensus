import {reducer, types} from "../../board";

it("should set base color", () => {
    const fakeData = {foo: "bar"};
    const result = reducer(undefined, {
        type: types.SET_BASE_COLOR,
        payload: fakeData,
    });
    expect(result.baseColor).toEqual(fakeData);
});
it("should set color code list", () => {
    const fakeData = [{foo: "bar"}];
    const result = reducer(undefined, {
        type: types.SET_COLOR_CODE_LIST ,
        payload: fakeData,
    });
    expect(result.colorCodeList).toEqual(fakeData);
});
it("should set selected color code list", () => {
    const fakeData = [{foo: "bar"}];
    const result = reducer(undefined, {
        type: types.SET_SELECTED_COLOR_CODE_LIST ,
        payload: fakeData,
    });
    expect(result.selectedColorCodeList).toEqual(fakeData);
});
