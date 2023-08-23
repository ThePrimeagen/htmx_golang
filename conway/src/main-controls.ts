import { Conway } from "./conway";

let controls: MainControls | undefined = undefined;

let running = false;
let paceMS = 250;

function run() {
    if (!running) {
        return;
    }
    tick();
    setTimeout(run, paceMS);
}

export function start(menu: HTMLElement, cells: HTMLElement) {
    if (running) {
        return;
    }

    if (!controls) {
        controls = new MainControls(
            50,
            menu, cells);
    } else {
        running = true;
    }

    controls.init();
    run();
}

export function tick() {
    if (controls) {
        controls.tick();
    }
}

export function pause() {
    running = false;
    controls?.init();
}

export function pace(input: HTMLInputElement) {
    paceMS = parseInt(input.value);
}


export function toggleCell(cell: HTMLElement) {
    if (controls) {
        controls.toggleCell(cell);
    }
}

class MainControls {
    conway: Conway;

    private seed: HTMLInputElement;
    private frameCount: HTMLInputElement;
    private pause: HTMLButtonElement;

    constructor(size: number, menu: HTMLElement, private cells: HTMLElement) {
        this.conway = new Conway(size);
        this.seed = menu.querySelector(".seed") as HTMLInputElement;
        this.frameCount = menu.querySelector(".frame-count") as HTMLInputElement;
        this.pause = menu.querySelector(".pause") as HTMLButtonElement;
    }

    updateSeed(seedStr: string) {
        this.seed.value = seedStr;
    }

    init() {
        const data: number[][] = new Array(this.conway.height).fill(0).map(() => new Array(this.conway.width).fill(0));
        for (let r = 0; r < this.conway.data.length; ++r) {
            const row = this.conway.data[r];
            for (let col = 0; col < row.length; ++col) {
                const cell = this.cells.querySelector(`.cell-${r}-${col}`) as HTMLElement;

                if (cell) {
                    data[r][col] = cell.classList.contains("alive") ? 1 : 0;
                }
            }
        }

        this.conway.setData(data);
        this.seed.value = this.conway.getSeedString();
    }

    tick() {
        if (this.pause.hasAttribute("disabled")) {
            this.pause.removeAttribute("disabled");
        }

        this.conway.tick();
        this.frameCount.value = this.conway.tickCount.toString();

        for (let r = 0; r < this.conway.data.length; ++r) {
            const row = this.conway.data[r];
            for (let col = 0; col < row.length; ++col) {
                const cell = this.cells.querySelector(`.cell-${r}-${col}`) as HTMLElement;

                if (cell) {
                    this.setCell(cell, row[col]);
                }
            }
        }
    }

    toggleCell(cell: HTMLElement) {
        if (cell.classList.contains("dead")) {
            cell.classList.remove("dead");
            cell.classList.add("alive");
        } else {
            cell.classList.add("dead");
            cell.classList.remove("alive");
        }
    }

    setCell(cell: HTMLElement, value: number) {
        const isDead = cell.classList.contains("dead");
        const toAlive = value === 1;

        if (isDead && toAlive) {
            cell.classList.remove("dead");
            cell.classList.add("alive");
        } else if (!isDead && !toAlive) {
            cell.classList.add("dead");
            cell.classList.remove("alive");
        }
    }

}
