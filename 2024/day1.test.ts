import { expect, test } from 'vitest'
import * as fs from 'fs'
import * as path from 'path'


test('example input', () => {
    expect(example1()).toBe(11)
});

test('part1', () => {
    expect(part1()).toBe(2057374)
});

test('example2', () => {
    expect(example2()).toBe(31)
});

test('part2', () => {
    expect(part2()).toBe(23177084)
});

interface Input {
    a: number[],
    b: number[],
    length: number
}

function example2(): number {
    const input = `3   4
4   3
2   5
1   3
3   9
3   3`
    const parsed = parseInput(input)
    const result = calcultateSimilarity(parsed.a, parsed.b)
    return result;
}

function part2(): number {
    const filestream = fs.readFileSync(path.resolve(__dirname, 'day1_input.txt'), 'utf-8')
    const input = parseInput(filestream)
    const result = calcultateSimilarity(input.a, input.b)
    return result;
}

function calcultateSimilarity(a: number[], b: number[]): number {
    a.sort()
    b.sort()
    let occurences: [number, number][] = []
    for (let i = 0; i < a.length; i++) {
        const x = a[i]
        let occurence = 0;
        for (let j = 0; j < b.length; j++) {
            const y = b[j]
            if (x === y) {
                occurence++
            }
            if (y > x) break;
        }
        occurences.push([x, occurence])
    }
    const result = occurences.reduce((acc, cur) => acc + cur[1] * cur[0], 0)

    return result;
}

function parseInput(input: string): Input {
    const lines = input.split('\n')
    const a: number[] = []
    const b: number[] = []
    let i = 0;
    for (let line of lines) {
        const [x, y] = line.split('   ')
        const parsedX = parseInt(x)
        const parsedY = parseInt(y)
        if (isNaN(parsedX) || isNaN(parsedY)) continue
        a.push(parseInt(x))
        b.push(parseInt(y))
        i++;
    }

    return {
        a,
        b,
        length: i
    }
}

function calcultateDistance(a: number[], b: number[]): number {
    a.sort()
    b.sort()
    let distance = 0;
    for (let i = 0; i < a.length; i++) {
        distance += Math.abs(a[i] - b[i])
    }

    return distance
}


function example1(): number {
    const input = `3   4
4   3
2   5
1   3
3   9
3   3`

    const parsed = parseInput(input)
    const result = calcultateDistance(parsed.a, parsed.b)

    return result;
}

function part1(): number {
    const filestream = fs.readFileSync(path.resolve(__dirname, 'day1_input.txt'), 'utf-8')
    const input = parseInput(filestream)
    const result = calcultateDistance(input.a, input.b)
    return result;
}
