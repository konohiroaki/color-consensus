import {actions, types} from "../../vote";

describe("setSelectedColorCodeList(colorList)", function () {
    it("should return colorList as payload", () => {
        const fakeData = ["#ff0000", "#f00000"];
        const action = actions.setSelectedColorCodeList(fakeData);
        expect(action.type).toEqual(types.SET_SELECTED_COLOR_CODE_LIST);
        expect(action.payload).toEqual(fakeData);
    });
});

describe("resetSelectedColorCodeList()", function () {
    it("should return empty list as payload", () => {
        const action = actions.resetSelectedColorCodeList();
        expect(action.type).toEqual(types.SET_SELECTED_COLOR_CODE_LIST);
        expect(action.payload).toEqual([]);
    });
});
