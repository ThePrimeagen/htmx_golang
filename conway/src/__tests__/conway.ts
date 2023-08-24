import { Conway, seedToDataMap, seedToString } from '../conway';

test("conway", () => {
    const conway = new Conway(3);
    conway.setData([
        [0, 1, 0],
        [0, 1, 0],
        [0, 1, 0],
    ]);

    conway.tick();
    expect(conway.data).toEqual([
        [0, 0, 0],
        [1, 1, 1],
        [0, 0, 0],
    ]);

    conway.tick();
    expect(conway.data).toEqual([
        [0, 1, 0],
        [0, 1, 0],
        [0, 1, 0],
    ]);

});

test("seed value", () => {

    const seed = "490";
    const values = seedToDataMap(seed, 3);

    console.log(values);
    expect(values).toEqual([
        [0, 1, 0],
        [0, 1, 0],
        [0, 1, 0],
    ]);

    const seedStr = seedToString([
        [0, 1, 0],
        [0, 1, 0],
        [0, 1, 0],
    ]);

    expect(seedStr).toEqual(seed);
});
