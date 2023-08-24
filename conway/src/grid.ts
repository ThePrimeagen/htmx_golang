let conwayGrid: HTMLElement | undefined = undefined;

let columnCount: number = 0;
const gridData: number[] = [];
const gridElements: HTMLElement[] = [];
const callbacks: (() => void)[] = [];

function initGridData(columns: number) {
    gridData.length = 0;
    gridData.length = columns * columns;
    for (let i = 0; i < columns * columns; i++) {
        gridData[i] = 0;
    }
}

function toggleCell(cell: HTMLElement) {
    if (cell.classList.contains("dead")) {
        cell.classList.remove("dead");
        cell.classList.add("alive");
    } else {
        cell.classList.add("dead");
        cell.classList.remove("alive");
    }
}

export function onUpdate(callback: () => void) {
    if (!callbacks.includes(callback)) {
        callbacks.push(callback);
    }
}

export function initGrid(element: HTMLElement, columns: number) {
    columnCount = columns;
    const foundConway = element.querySelector("#conway");
    if (foundConway === conwayGrid && foundConway !== undefined) {
        return;
    }

    if (gridData.length !== columns * columns) {
        initGridData(columns);
    }

    const div = document.createElement("div");
    div.id = "conway";
    div.style.display = "grid";
    div.style.gridTemplateColumns = `repeat(${columns}, 1fr)`;

    conwayGrid = div;

    gridElements.length = 0;

    // create the grid
    for (let i = 0; i < columns * columns; i++) {
        const cell = document.createElement("span");
        cell.classList.add("cell", gridData[i] ? "alive" : "dead");

        const cellIdx = i;
        cell.onclick = () => {
            gridData[cellIdx] = gridData[cellIdx] === 1 ? 0 : 1;
            toggleCell(cell);
            callbacks.forEach(cb => cb());
        }
        div.appendChild(cell);
        gridElements.push(cell);
    }

    element.appendChild(div);
}

export function updateGrid(data: number[][]) {
    if (conwayGrid === undefined) {
        return;
    }

    for (let r = 0; r < data.length; r++) {
        const row = data[r];
        for (let col = 0; col < row.length; col++) {
            const cell = row[col];
            const idx = r * data.length + col;

            if (cell != gridData[idx]) {
                toggleCell(gridElements[idx]);
                gridData[idx] = cell;
            }
        }
    }
}

export function getData(): number[][] {
    const out: number[][] = [];
    for (let i = 0; i < gridData.length; ++i) {
        const row = Math.floor(i / columnCount);
        if (!out[row]) {
            out[row] = [];
        }
        out[row].push(gridData[i]);
    }

    return out;
}
