import { expect, test } from 'vitest'
import * as fs from 'fs'
import * as path from 'path'

test('example input', () => {
    expect(example1()).toBe(2)
});

test('part1', () => {
    expect(part1()).toBe(510)
});

test('example2', () => {
    expect(example2()).toBe(4)
});

test('part2', () => {
    expect(part2()).toBe(553)
});

function part1(): number {
    const input = fs.readFileSync(path.join(__dirname, 'day2_input.txt'), 'utf8')
    const reports = parseInput(input)
    const result = findSafeReports(reports)
    return result
}

function example1(): any {
    const input = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`
    const reports = parseInput(input)
    const result = findSafeReports(reports)
    return result
}

function example2(): number {
    const input = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`
    const reports = parseInput(input)
    const result = findSafeReportsDampener(reports)
    return result
}

function part2(): number {
    const input = fs.readFileSync(path.join(__dirname, 'day2_input.txt'), 'utf8')
    const reports = parseInput(input)
    const result = findSafeReportsDampener(reports)
    return result
}

function findSafeReportsDampener(input: Report[]): number {
    const reports = input.map(x => x.levels)
    const result = reports.reduce(
        (count, report) =>
            count +
            (isSafe2(report) ||
                report.some((_, i) =>
                    isSafe2([...report.slice(0, i), ...report.slice(i + 1)])
                )
                ? 1
                : 0),
        0
    )
    return result;
}

function isSafe2(levels: number[]): boolean {
    const positiveReport = levels[0] - levels[1] > 0
    let prev = levels[0]
    for (let i = 1; i < levels.length; i++) {
        const diff = prev - levels[i]
        const isPositive = diff > 0
        if (isPositive !== positiveReport) {
            return false
        }
        const distance = Math.abs(diff)
        if (distance < 1 || distance > 3) {
            return false
        }
        prev = levels[i]
    }
    return true
}

interface Report {
    levels: number[]
}
function parseInput(input: string): Report[] {
    const lines = input.split('\n')
    const reports: Report[] = []
    for (let i = 0; i < lines.length; i++) {
        const line = lines[i]
        const parts = line.split(' ')
        if (parts.length < 2) continue
        const levels = parts.map(x => parseInt(x))
        reports.push({ levels })
    }

    return reports
}

function findSafeReports(reports: Report[]): number {
    const result = reports.filter(isSafe).length
    return result;
}

function isSafe(report: Report): boolean {
    const positiveReport = report.levels[0] - report.levels[1] > 0
    let prev = report.levels[0]
    for (let i = 1; i < report.levels.length; i++) {
        const diff = prev - report.levels[i]
        const isPositive = diff > 0
        if (isPositive !== positiveReport) {
            return false
        }
        const distance = Math.abs(diff)
        if (distance < 1 || distance > 3) {
            return false
        }
        prev = report.levels[i]
    }
    return true
}
