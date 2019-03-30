import {reducer, types} from "../../user";

it("should set id", () => {
    const fakeData = "foo";
    const result = reducer(undefined, {
        type: types.SET_ID,
        payload: fakeData,
    });
    expect(result.id).toEqual(fakeData);
});
it("should overwrite id", () => {
    const fakeData = "foo";
    const fakeData2 = "bar";
    const result = reducer({id: fakeData}, {
        type: types.SET_ID,
        payload: fakeData2,
    });
    expect(result.id).toEqual(fakeData2);

});
