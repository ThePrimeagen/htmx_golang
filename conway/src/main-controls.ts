import { Conway } from "./conway";

type Grid = {
    updateGrid(data: number[][]): void;
    getData(): number[][];
    onUpdate(cb: () => void): void;
}

let paceMS = 250;
let running = false;
let grid: Grid | undefined = undefined;
let seedInput: HTMLInputElement | undefined = undefined;
let seedInputAt: HTMLInputElement | undefined = undefined;
let frameCount: HTMLElement | undefined = undefined;
let frameCountInput: HTMLInputElement | undefined = undefined;
let pauseElement: HTMLButtonElement | undefined = undefined;
let paceElement: HTMLInputElement | undefined = undefined;
let conway: Conway | undefined = undefined;

export function initControls(menu: HTMLElement, columns: number, g: Grid) {
    seedInput = menu.querySelector(".seed-input") as HTMLInputElement;
    seedInputAt = menu.querySelector(".seed-input-at") as HTMLInputElement;
    frameCount = menu.querySelector(".frame-count") as HTMLElement;
    frameCountInput = menu.querySelector(".frame-count-input") as HTMLInputElement;
    pauseElement = menu.querySelector(".pause") as HTMLButtonElement;
    paceElement = menu.querySelector(".pace") as HTMLInputElement;
    paceElement.onchange = () => pace(paceElement);
    paceElement.value = paceMS.toString();

    if (!conway) {
        conway = new Conway(columns);
    }
    grid = g;
    grid.onUpdate(() => {
        if (!conway || !seedInput || !seedInputAt) {
            return;
        }

        const seedString = conway.getSeedString();
        seedInput.value = seedString;
        seedInputAt.value = seedString;
    });
}

export function seed(seed: string, columns: number) {
    if (!conway) {
        throw new Error("must call initControls first");
    }

    conway.setSeed(seed, columns);
}

export function pace(input?: HTMLInputElement) {
    if (!input) {
        return;
    }

    const nextPace = parseInt(input.value);
    if (isNaN(nextPace)) {
        return;
    }

    paceMS = nextPace;
}

function runLoop() {
    if (!running) {
        return;
    }
    tick();
    setTimeout(runLoop, paceMS);
}

export function run(): void {
    if (running || !conway || !grid || !seedInput || !seedInputAt) {
        return;
    }

    pauseElement?.removeAttribute("disabled");
    running = true;

    conway.data = grid.getData();

    const seedString = conway.getSeedString();
    seedInput.value = seedString;
    seedInputAt.value = seedString;

    runLoop();
}

export function pause(): void {
    running = false;
    pauseElement?.toggleAttribute("disabled");
}

function tick() {
    if (!conway || !frameCount || !frameCountInput || !grid || !seedInputAt) {
        return;
    }

    conway.tick();
    frameCount.innerText = conway.tickCount.toString();
    frameCountInput.value = conway.tickCount.toString();

    grid.updateGrid(conway.data);
    seedInputAt.value = conway.getSeedString();
}
