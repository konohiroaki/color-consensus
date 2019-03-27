export const isSameColor = (c1, c2) => c1 !== undefined && c2 !== undefined
                                       && c1.lang === c2.lang && c1.name === c2.name && c1.code === c2.code;
