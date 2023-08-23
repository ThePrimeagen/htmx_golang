function generateSeed(seed) {
    let seedValue = 0;
    for (let i = 0; i < seed.length; ++i) {
        seedValue += seed.charCodeAt(i) * (i + 1);
    }

    return seedValue;
}

const conwaysChildren = [];
let conwaysValues = [];

function generate() {
    const rows = +document.getElementById("rows").value;
    const cols = +document.getElementById("cols").value;
    const seed = document.getElementById("seed").value;
    const seedValue = generateSeed(seed);

    if (isNaN(rows) || isNaN(cols)) {
        return;
    }

    const grid = document.getElementById("conway");
    grid.style.gridTemplateColumns = `repeat(${cols}, 1fr)`;

    // GET THE HELL OUT OF THE STARTUP
    grid.innerHTML = "";
    conwaysChildren.length = 0;

    for (let y = 0; y < rows; ++y) {
        for (let x = 0; x < cols; ++x) {
            const cell = document.createElement("span");
            grid.appendChild(cell);

            if (conwaysChildren[y] === undefined) {
                conwaysChildren[y] = [];
            }

            conwaysChildren[y][x] = cell;

            if (conwaysValues[y] === undefined) {
                conwaysValues[y] = [];
            }

            conwaysValues[y][x] = 0;
        }
    }

    seedMeDaddy(seedValue, conwaysValues);
    runConway(seedValue, 0, 500);
}

// write me mulberry functor
function mulberry(seed) {
    return function () {
        let t = seed += 0x6D2B79F5;
        t = Math.imul(t ^ t >>> 15, t | 1);
        t ^= t + Math.imul(t ^ t >>> 7, t | 61);
        return ((t ^ t >>> 14) >>> 0) / 4294967296;
    }
}


function seedMeDaddy(seedValue, array) {
    const rand = mulberry(seedValue);

    for (let y = 0; y < array.length; ++y) {
        for (let x = 0; x < array[y].length; ++x) {
            if (rand() < 0.33) {
                array[y][x] = 1;
            }
        }
    }
}

const spots = [
    [-1, -1],
    [-1, 0],
    [-1, 1],
    [0, -1],
    //[0, 0],
    [0, 1],
    [1, -1],
    [1, 0],
    [1, 1],
]

function sum(array, x, y) {
    return spots.reduce((acc, spot) => {
        const value = array[y + spot[0]][x + spot[1]];
        return isNaN(value) ? acc : acc + value;
    }, 0);
}

function runConway(seedValue, tick, pace) {

    const next = [];
    for (let y = 0; y < conwaysValues.length; ++y) {

    }

    setTimeout(() => {
        runConway(seedValue, tick + 1, pace);
    }, pace);
}

