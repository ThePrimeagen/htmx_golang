const spots = [
    [-1, -1],
    [-1, 0],
    [-1, 1],
    [0, -1],
    //[0, 0], do not use yourself!
    [0, 1],
    [1, -1],
    [1, 0],
    [1, 1],
]

export function generatePreview(container: HTMLElement, seed: string, columns: number): void {
    const map = seedToDataMap(seed, columns);
}

function seedToString(data: number[][]): string {
    const colSize = data[0].length;
    const length = data.length * colSize;

    let str = "";
    for (let i = 0; i < length + 7; i += 8) {
        let value = 0;
        for (let bit = 7; bit >= 0; --bit) {
            const row = data[Math.floor(i / colSize)];
            if (!row) {
                break;
            }
            value |= row[i % colSize] << bit;
        }

        str += value.toString(16);
    }
    return str;
}

function seedToDataMap(data: string, rowSize: number, colSize: number): number[][] {
    const out = new Array(rowSize).fill(0).map(() => new Array(colSize).fill(0));

    for (let i = 0; i < data.length; ++i) {
        const s = data[i];
        const value = parseInt(s, 16);

        for (let bit = 7; bit >= 0; --bit) {
            const b = value & (1 << bit);
            out[Math.floor(i / colSize)][i % colSize] = b;
        }
    }

    return out;
}

function sum(data: number[][], row: number, col: number) {
    return spots.reduce((acc, spot) => {
        const r = data[row + spot[0]];
        if (!r) {
            return acc;
        }

        const value = r[col + spot[1]];
        return !value ? acc : acc + value;
    }, 0);
}

export class Conway {
    width!: number;
    height!: number;
    data!: number[][];

    private _tickCount!: number;
    private _running!: boolean;

    get tickCount(): number {
        return this._tickCount;
    }

    get running(): boolean {
        return this._running;
    }

    constructor(size: number) {
        this.init(size, new Array<number>(size).fill(0).map(() => new Array(size).fill(0)));
    }

    private init(size: number, data: number[][]) {
        this.width = size;
        this.height = size;
        this._tickCount = 0;
        this._running = false;

        this.data = data;
    }

    getSeedString(): string {
        return seedToString(this.data);
    }

    setSeed(seed: string, size: number): void {
        this.init(size, seedToDataMap(seed, size, size));
    }

    setData(data: number[][]): void {
        this.init(data.length, data);
    }

    tick() {
        this._tickCount++;

        const next = new Array(this.height).fill(0).map(() => new Array(this.width).fill(0));

        for (let row = 0; row < this.height; ++row) {
            for (let col = 0; col < this.width; ++col) {
                const cellValue = this.data[row][col];
                const value = sum(this.data, row, col);

                if (value === 3) {
                    next[row][col] = 1;
                } else if (value < 2 || value > 3) {
                    next[row][col] = 0;
                } else {
                    next[row][col] = cellValue;
                }
            }
        }

        this.data = next;
    }

}

