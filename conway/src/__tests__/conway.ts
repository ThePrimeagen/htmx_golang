import { Conway } from '../conway';

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


