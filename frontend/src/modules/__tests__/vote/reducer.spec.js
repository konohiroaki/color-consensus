import {reducer, types} from "../../vote";

it("should set selected color code list", () => {
    const fakeData = [{foo: "bar"}];
    const result = reducer(undefined, {
        type: types.SET_SELECTED_COLOR_CODE_LIST ,
        payload: fakeData,
    });
    expect(result.selectedColorCodeList).toEqual(fakeData);
});
